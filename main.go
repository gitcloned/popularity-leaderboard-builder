package main

import (
	"liquide/re/popularity-leaderboard-builder/objects"

	reader "liquide/re/popularity-leaderboard-builder/reader/interfaces"
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

	// create topology tree
	tree := topology.Tree{
		Branches: []topology.Branch{
			{
				Path:  "Channel",
				Field: "Channel",
				Branches: []topology.Branch{
					{
						Path:  "Cohert",
						Field: "UserCohert",
					},
				},
			},
		},
	}

	opts := Options{
		MaxWorkers:   100,
		MaxQueueSize: 100000,
	}

	d := reader.EventDispatcher{
		Queue:    make(chan objects.UserAction, opts.MaxQueueSize),
		Finished: false,
	}
	// get reader, and pass him the channel
	reader, err := provider.ReaderProvider()

	if err != nil {
		panic(err)
	}

	go func() {
		reader.Read(d)
	}()

	log.Info("Would start the workers .. ")
	wg := sync.WaitGroup{}
	for i := 0; i < opts.MaxWorkers; i++ {
		wg.Add(1) // Add a wait group for each worker
		// Spawn a worker
		go func() {
			for {
				select {
				case userAction := <-d.Queue:
					log.Info("Processing an item..")
					tree.ProcessAction(&userAction)
				}
			}
		}()
	}
	wg.Wait()
}
