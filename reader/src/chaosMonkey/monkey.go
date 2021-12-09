package reader

import (
	"liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/reader/interfaces"
)

type ChaosMonkeyReader struct {
	interfaces.UserActionReader
}

func (r ChaosMonkeyReader) Read(chan objects.UserAction) {

}

func (m ChaosMonkeyReader) start() {

}
