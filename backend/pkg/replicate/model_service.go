package replicate

import (
	"context"
	"fmt"
)

type ModelResponse struct {
	Url           string               `json:"url"`
	Owner         string               `json:"owner"`
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	Visibility    string               `json:"visibility"`
	GithubURL     string               `json:"github_url"`
	PaperURL      string               `json:"paper_url"`
	LicenseURL    string               `json:"license_url"`
	LatestVersion ModelVersionResponse `json:"latest_version"`
}

type ModelVersionResponse struct {
	Id            string      `json:"id"`
	CreatedAt     string      `json:"created_at"`
	CogVersion    string      `json:"cog_version"`
	OpenapiSchema interface{} `json:"openapi_schema"`
}

type ModelVersionListResponse = ListResponse[ModelVersionResponse]

type ModelCollectionsResponse struct {
	Name        string          `json:"name"`
	Slug        string          `json:"slug"`
	Description string          `json:"description"`
	Models      []ModelResponse `json:"models"`
}

type GetModelReqParams struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

func (c *client) GetModel(ctx context.Context, params GetModelReqParams) (*ModelResponse, error) {

	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/models/%s/%s", params.Owner, params.Name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	output := new(ModelResponse)

	if err := getResponseObject(resp, &output); err != nil {
		return nil, err
	}

	return output, nil

}

func (c *client) GetModelVersionList(ctx context.Context, params GetModelReqParams) (*ModelVersionListResponse, error) {

	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/models/%s/%s/versions", params.Owner, params.Name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	output := new(ModelVersionListResponse)

	if err := getResponseObject(resp, &output); err != nil {
		return nil, err
	}

	return output, nil
}

type GetModelVersionReqParams struct {
	GetModelReqParams
	Id string
}

func (c *client) GetModelVersion(ctx context.Context, params GetModelVersionReqParams) (*ModelVersionResponse, error) {

	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/models/%s/%s/versions/%s", params.Owner, params.Name, params.Id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	output := new(ModelVersionResponse)

	if err := getResponseObject(resp, &output); err != nil {
		return nil, err
	}

	return output, nil
}

type GetModelCollectionsReqParams struct {
	CollectionSlug string `json:"collection_slug"`
}

func (c *client) GetModelCollections(ctx context.Context, params GetModelCollectionsReqParams) (*ModelCollectionsResponse, error) {

	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/collections/%s", params.CollectionSlug), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.performRequest(req)
	if err != nil {
		return nil, err
	}

	output := new(ModelCollectionsResponse)

	if err := getResponseObject(resp, &output); err != nil {
		return nil, err
	}

	return output, nil

}
