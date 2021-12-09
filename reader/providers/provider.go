package reader

import (
	interfaces "liquide/re/popularity-leaderboard-builder/reader/interfaces"
	chaos "liquide/re/popularity-leaderboard-builder/reader/src/chaosMonkey"
)

func ReaderProvider() (interfaces.UserActionReader, error) {
	return chaos.ChaosMonkeyReader{}, nil
}
