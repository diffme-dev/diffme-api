package asynq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
)

type EmailDeliveryPayloadd struct {
	UserID     int
	TemplateID string
}

func CreateSearchableChangeHandler(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayloadd

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Sending Email to User: user_id=%d, template_id=%s", p.UserID, p.TemplateID)

	return nil
}
