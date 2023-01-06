package replicate

import (
	"context"
	"fmt"
)

type CreatePredictionPayload struct {
	/*
	 * The ID of the model version to use for the prediction
	 */
	Version string `json:"version"`
	/*
	 * The Model's Input as a JSON object
	 */
	Input interface{} `json:"input"`
	/*
	 * HTTPS URL to receive a webhook when the prediction is completed
	 */
	WebhookCompleted string `json:"webhook_completed"`
}

type PredictionRelatedEndpoints struct {
	Get    string `json:"get"`
	Cancel string `json:"cancel"`
}

type PredictionStatus string

const (
	PredictionStatusStarting   PredictionStatus = "starting"
	PredictionStatusProcessing PredictionStatus = "processing"
	PredictionStatusSucceeded  PredictionStatus = "succeeded"
	PredictionStatusFailed     PredictionStatus = "failed"
	PredictionStatusCanceled   PredictionStatus = "canceled"
)

type CreatePredictionResponse struct {
	Id          string                     `json:"id"`
	Version     string                     `json:"version"`
	Urls        PredictionRelatedEndpoints `json:"urls"`
	CreatedAt   string                     `json:"created_at"`
	StartedAt   string                     `json:"started_at"`
	CompletedAt string                     `json:"completed_at"`
	Status      PredictionStatus           `json:"status"`
	Input       interface{}                `json:"input"`
	Output      interface{}                `json:"output"`
	Error       interface{}                `json:"error"`
	Logs        interface{}                `json:"logs"`
	Metrics     interface{}                `json:"metrics"`
}

type GetPredictionResponse struct {
	CreatePredictionResponse
	Source PredictionSource `json:"source"`
}

type PredictionSource string

const (
	PredictionSourceAPI PredictionSource = "api"
	PredictionSourceWeb PredictionSource = "web"
)

type PredictionListResult struct {
	Id          string                     `json:"id"`
	Version     string                     `json:"version"`
	Urls        PredictionRelatedEndpoints `json:"urls"`
	CreatedAt   string                     `json:"created_at"`
	StartedAt   string                     `json:"started_at"`
	CompletedAt string                     `json:"completed_at"`
	Source      PredictionSource           `json:"source"`
	Status      PredictionStatus           `json:"status"`
}

type PredictionListResponse = ListResponse[PredictionListResult]

func (c *client) CreatePrediction(ctx context.Context, payload CreatePredictionPayload) (*CreatePredictionResponse, error) {

	req, err := c.newRequest(ctx, "POST", "/predictions", payload)

	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	output := new(CreatePredictionResponse)

	if err := getResponseObject(resp, &output); err != nil {
		return nil, err
	}

	return output, nil
}

type GetPredictionParams struct {
	Id string
}

func (c *client) GetPrediction(ctx context.Context, params GetPredictionParams) (*GetPredictionResponse, error) {

	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/predictions/%s", params.Id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	output := new(GetPredictionResponse)

	if err := getResponseObject(resp, &output); err != nil {
		return nil, err
	}

	return output, nil

}

func (c *client) GetPredictionList(ctx context.Context) (*PredictionListResponse, error) {

	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/predictions"), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	output := new(PredictionListResponse)

	if err := getResponseObject(resp, &output); err != nil {
		return nil, err
	}

	return output, nil

}

type CancelPredictionParams struct {
	Id string
}

func (c *client) CancelPrediction(ctx context.Context, params CancelPredictionParams) (interface{}, error) {

	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/predictions/%s/cancel", params.Id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	output := new(interface{})

	if err := getResponseObject(resp, &output); err != nil {
		return nil, err
	}

	return output, nil

}
