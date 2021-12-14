package main

import (
	provider "liquide/re/popularity-leaderboard-builder/reader/providers"

	"liquide/re/popularity-leaderboard-builder/topology"
	"sync"

	log "github.com/sirupsen/logrus"
)

// Options represent options for EventDispatcher.
type Options struct {
	MaxWorkers   int // Number of workers to spawn.
	MaxQueueSize int // Maximum length for the queue to hold events.
}

func main() {

	log.Info("Starting..")

	// TODO: use dig, https://blog.drewolson.org/dependency-injection-in-go
	// TODO: use viper: https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d
	// TODO: use google wire DI
	// TODO: file based logging using rotating logs
	// TODO: test cases
	// TODO: metrics API, or simple GUI - can add trees/branches from GUI?
	// TODO: Redis sorted set to have TTL
	// TODO: Handle redis connection break, and restarts
	// TODO: memory profiling

	garden := &topology.Garden{
		Trees: []topology.Tree{
			{
				Name:          "item",
				ItemFieldName: "Id",
				Branches: []topology.Branch{
					{
						Name:  "Channel",
						Field: "Channel",
						Branches: []topology.Branch{
							{
								Name:  "Cohert",
								Field: "UserCohert",
							},
						},
					},
				},
			},
			{
				Name:          "stock",
				ItemFieldName: "Stock",
				Branches: []topology.Branch{
					{
						Name:  "Cohert",
						Field: "UserCohert",
					},
					{
						Name:  "Channel",
						Field: "Channel",
					},
				},
			},
		},
	}

	opts := Options{
		MaxWorkers:   100,
		MaxQueueSize: 100000,
	}

	// get reader, and pass him the channel
	reader, err := provider.ReaderProvider()

	if err != nil {
		panic(err)
	}

	d := reader.Read(opts.MaxQueueSize)

	log.Info("Would start the workers .. ")
	wg := sync.WaitGroup{}
	for i := 0; i < opts.MaxWorkers; i++ {
		wg.Add(1) // Add a wait group for each worker
		// Spawn a worker
		go func(garden *topology.Garden) {
			for {
				select {
				case userAction := <-d.Queue:
					garden.ProcessAction(&userAction)
				}
			}
		}(garden)
	}
	wg.Wait()
}
