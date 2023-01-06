package replicate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client interface {
	GetModel(ctx context.Context, params GetModelReqParams) (*ModelResponse, error)
	GetModelVersionList(ctx context.Context, params GetModelReqParams) (*ModelVersionListResponse, error)
	GetModelVersion(ctx context.Context, params GetModelVersionReqParams) (*ModelVersionResponse, error)
	GetModelCollections(ctx context.Context, params GetModelCollectionsReqParams) (*ModelCollectionsResponse, error)

	CreatePrediction(ctx context.Context, payload CreatePredictionPayload) (*CreatePredictionResponse, error)
	GetPrediction(ctx context.Context, params GetPredictionParams) (*GetPredictionResponse, error)
	GetPredictionList(ctx context.Context) (*PredictionListResponse, error)
	CancelPrediction(ctx context.Context, params CancelPredictionParams) (interface{}, error)

	NewModel(config ModelConfig) Model
}

type client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
	userAgent  string
}

const (
	defaultTimeOutSeconds = 10
	defaultBaseURL        = "https://api.replicate.com/v1"
	defaultUserAgent      = "replicate-go-client"
)

func NewClient(apiKey string) Client {

	httpClient := &http.Client{
		Timeout: time.Duration(time.Second * defaultTimeOutSeconds),
	}

	c := &client{
		httpClient: httpClient,
		baseURL:    defaultBaseURL,
		apiKey:     apiKey,
		userAgent:  defaultUserAgent,
	}

	return c
}

func jsonBodyReader(body interface{}) (io.Reader, error) {
	if body == nil {
		return bytes.NewBuffer(nil), nil
	}
	raw, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed encoding json: %w", err)
	}
	return bytes.NewBuffer(raw), nil
}

func (c *client) newRequest(ctx context.Context, method, path string, payload interface{}) (*http.Request, error) {
	bodyReader, err := jsonBodyReader(payload)
	if err != nil {
		return nil, err
	}
	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.apiKey))
	return req, nil
}

func (c *client) performRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if err := checkForSuccess(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func getResponseObject(rsp *http.Response, v interface{}) error {
	defer rsp.Body.Close()
	if err := json.NewDecoder(rsp.Body).Decode(v); err != nil {
		return fmt.Errorf("invalid json response: %w", err)
	}
	return nil
}

// returns an error if this response includes an error.
func checkForSuccess(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read from body: %w", err)
	}
	var result APIErrorResponse
	if err := json.Unmarshal(data, &result); err != nil {
		// if we can't decode the json error then create an unexpected error
		apiError := APIError{
			StatusCode: resp.StatusCode,
			Type:       "Unexpected",
			Message:    string(data),
		}
		return apiError
	}
	result.Error.StatusCode = resp.StatusCode
	return result.Error
}
