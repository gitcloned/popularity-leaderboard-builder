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

func (r ChaosMonkeyReader) Read(d interfaces.EventDispatcher) {

	run := true

	log.Info("Chaos Monkey starting..")

	for run == true {

		log.Info("Emitting an item..")

		d.Queue <- objects.UserAction{
			Item: "Item 1",
		}

		log.Info("Sleeping now..")
		time.Sleep(8 * time.Second)
	}
}

func (m ChaosMonkeyReader) start() {

}
