package infra

import (
	Machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

var cnf = &config.Config{
	Broker:        "redis://localhost:6379",
	ResultBackend: "redis://localhost:6379",
}

func NewMachinery() (*Machinery.Server, error) {

	taskserver, err := Machinery.NewServer(cnf)

	if err != nil {
		return nil, err
	}

	return taskserver, nil
}
