package main

import (
	"diffme.dev/diffme-api/cmd/server"
	"diffme.dev/diffme-api/cmd/workers"
)

func main() {
	println("[starting everything]")

	server.StartServer()
	workers.StartWorkers()
}
