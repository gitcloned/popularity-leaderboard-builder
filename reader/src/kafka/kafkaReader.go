package reader

import (
	"liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/reader/interfaces"
)

type KafkaUserActionReader struct {
	interfaces.UserActionReader
}

func (r KafkaUserActionReader) Read(chan objects.UserAction) {
	// TODO: https://docs.confluent.io/clients-confluent-kafka-go/current/overview.html#id1
}
