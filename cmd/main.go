package main

import (
	"diffme.dev/diffme-api/cmd/server"
	"diffme.dev/diffme-api/cmd/workers"
	"sync"
)

//var (
//	app = cli.NewApp()
//)

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		workers.StartWorkers()
		wg.Done()
	}()

	go func() {
		server.StartServer()
		wg.Done()
	}()

	wg.Wait()
}
