package replicate

import (
	"context"
	"time"
)

const (
	defaultPollingInterval = 1
)

type Model interface {
	Predict(ctx context.Context, payload interface{}) (interface{}, error)
}

type model struct {
	version           string
	webhook_completed string
	client            Client
}

type ModelConfig struct {
	Version          string
	WebhookCompleted string
}

func (c *client) NewModel(config ModelConfig) Model {
	model := &model{
		client:            c,
		version:           config.Version,
		webhook_completed: config.WebhookCompleted,
	}

	return model
}

func (m *model) Predict(ctx context.Context, payload interface{}) (interface{}, error) {

	res, err := m.client.CreatePrediction(ctx, CreatePredictionPayload{
		Version:          m.version,
		WebhookCompleted: m.webhook_completed,
		Input:            payload,
	})

	if err != nil {
		return nil, err
	}

	// TODO : make this configurable
	interval := time.Duration(defaultPollingInterval * time.Second)

	prediction_id := res.Id
	m.client.GetPrediction(ctx, GetPredictionParams{Id: prediction_id})

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		prediction, err := m.client.GetPrediction(ctx, GetPredictionParams{Id: prediction_id})
		if err != nil {
			return nil, err
		}

		if prediction.Status == PredictionStatusSucceeded {
			return prediction.Output, nil
		}
		if prediction.Status == PredictionStatusFailed {
			return prediction, err
		}
		if prediction.Status == PredictionStatusCanceled {
			return prediction, err
		}
	}

	panic("unreachable")

}
