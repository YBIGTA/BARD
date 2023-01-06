package replicate

import "fmt"

type ListResponse[T interface{}] struct {
	Previous string `json:"previous"`
	Next     string `json:"next"`
	Results  []T    `json:"results"`
}

// APIError represents an error that occured on an API
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Type       string `json:"type"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("[%d:%s] %s", e.StatusCode, e.Type, e.Message)
}

// APIErrorResponse is the full error respnose that has been returned by an API.
type APIErrorResponse struct {
	Error APIError `json:"error"`
}

// EngineObject contained in an engine reponse
type EngineObject struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Owner  string `json:"owner"`
	Ready  bool   `json:"ready"`
}
