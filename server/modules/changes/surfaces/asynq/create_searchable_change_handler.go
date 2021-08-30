package asynq

import (
	"context"
	"diffme.dev/diffme-api/server/modules/changes"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
)

type ChangeCreatedPayload struct {
	change domain.Change
}

func (e *ChangeAsynqSurface) CreateSearchableChangeHandler(ctx context.Context, t *asynq.Task) error {
	var payload ChangeCreatedPayload

	err := json.Unmarshal(t.Payload(), &payload)

	if err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	_, err = e.changeUseCases.IndexSearchableChange(payload.change)

	return err
}
