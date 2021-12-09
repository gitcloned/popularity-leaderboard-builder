package main

import (
	"fmt"
	"liquide/re/popularity-leaderboard-builder/objects"
	reader "liquide/re/popularity-leaderboard-builder/reader/src/chaosMonkey"
	store "liquide/re/popularity-leaderboard-builder/store/src/memory"
	topology "liquide/re/popularity-leaderboard-builder/topology"
)

func main() {

	fmt.Println("Hello World")

	// TODO: use dig, https://blog.drewolson.org/dependency-injection-in-go

	// create topology tree
	tree := topology.Tree{}

	// create leader board store
	lbStore := store.InMemoryLeaderboardStore{}

	// channel which receives user action
	var channel chan objects.UserAction

	// get reader, and pass him the channel
	reader := reader.ChaosMonkeyReader{}
	reader.Read(channel)

	run := true
	for run == true {

		userAction := <-channel

		go func(u objects.UserAction) {
			tree.ProcessAction(u)
		}(userAction)
	}
}
