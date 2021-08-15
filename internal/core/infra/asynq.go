package infra

import (
	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

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
