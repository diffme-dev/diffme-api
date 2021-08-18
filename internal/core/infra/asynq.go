package infra

import (
	"diffme.dev/diffme-api/internal/config"
	"github.com/hibiken/asynq"
)

var redisAddr = config.GetConfig().RedisUri

var (
	SnapshotCreated = "SnapshotCreated"
	ChangeCreated   = "ChangeCreated"
)

func NewAsynqServer() (*asynq.Server, *asynq.ServeMux) {

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 10,
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()

	return srv, mux
}

func NewAsynqClient() *asynq.Client {
	return asynq.NewClient(
		asynq.RedisClientOpt{Addr: redisAddr},
	)
}
