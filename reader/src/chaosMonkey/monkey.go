package reader

import (
	"liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/reader/interfaces"
	"time"

	log "github.com/sirupsen/logrus"
)

type ChaosMonkeyReader struct {
	interfaces.UserActionReader
}

func (r ChaosMonkeyReader) Read(maxQueueSize int) interfaces.EventDispatcher {

	run := true

	d := interfaces.EventDispatcher{
		Queue:    make(chan objects.UserAction, maxQueueSize),
		Finished: false,
	}

	log.Info("Chaos Monkey starting..")

	go func() {

		for run == true {

			log.Info("Emitting an item..")

			d.Queue <- objects.UserAction{
				Item:       "Item 1",
				Channel:    "Ch 1",
				UserCohert: "UC 1",
				Points:     1.0,
			}

			time.Sleep(8 * time.Second)
		}
	}()

	return d
}

func (m ChaosMonkeyReader) start() {

}
