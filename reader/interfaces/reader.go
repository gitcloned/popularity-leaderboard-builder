package reader

import "liquide/re/popularity-leaderboard-builder/objects"

type EventDispatcher struct {
	Queue    chan objects.UserAction
	Finished bool
}

type UserActionReader interface {
	Read(maxQueueSize int) EventDispatcher
}
