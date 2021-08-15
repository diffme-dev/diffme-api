package workers

import (
	"context"
	Infra "diffme.dev/diffme-api/internal/infra"
	"github.com/RichardKnop/machinery/v1/tasks"
)

var (
	SnapshotCreated = "SnapshotCreated"
	ChangeCreated   = "ChangeCreated"
	Worker          = "diffme_worker"
)

func StartWorkers() {
	println("[starting workers]")

	taskserver, err := Infra.NewMachinery()

	if err != nil {
		println(err)
	}

	// register the tasks
	taskserver.RegisterTasks(map[string]interface{}{
		SnapshotCreated: OnSnapshotCreated,
		ChangeCreated:   OnChangeCreated,
	})

	// start the worker
	worker := taskserver.NewWorker(Worker, 10)

	if err := worker.Launch(); err != nil {
		println(err)
	}
}

func OnSnapshotCreated(ctx context.Context, arg tasks.Arg) error {
	println(arg.Value)
	return nil
}

func OnChangeCreated(ctx context.Context, arg tasks.Arg) error {
	println(arg.Value)
	return nil
}
