package asynq

import (
	"context"
	domain "diffme.dev/diffme-api/internal/modules/snapshots"
	"encoding/json"
	"fmt"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/hibiken/asynq"
	"log"
)

type CreateChangePayload struct {
	previous domain.Snapshot
	next     domain.Snapshot
}

type Data struct{ name string }

func CreateChangeHandler(ctx context.Context, t *asynq.Task) error {
	var payload CreateChangePayload

	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	payload.next.Data = &Data{name: "hi"}

	log.Printf("Previous: %+v", payload.previous.Data)
	log.Printf("Next: %+v", payload.next.Data)

	//prevBytes, err := encoders.EncodeJSON(payload.previous.Data)
	//nextBytes, err := encoders.EncodeJSON(payload.next.Data)

	original := []byte(`{"name": "John", "age": 24, "height": 3.21}`)
	target := []byte(`{"name": "Jane", "age": 24}`)

	patch, err := jsonpatch.CreateMergePatch(original, target)

	if err != nil {
		println(err)
	}

	log.Printf("Patch %s", patch)

	return nil
}
