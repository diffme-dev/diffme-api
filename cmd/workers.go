package cmd

import (
	Infra "diffme.dev/diffme-api/internal/infra"
)

func StartWorkers() {
	taskserver, err := Infra.NewMachinery()

	if err != nil {
		println(err)
	}

	// register the tasks
	taskserver.RegisterTasks(map[string]interface{}{
		"send_email": func() {},
	})

	// start the worker
	worker := taskserver.NewWorker("machinery_worker", 10)

	if err := worker.Launch(); err != nil {
		println(err)
	}
}
