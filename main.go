package main

import (
	"fmt"
	"liquide/re/popularity-leaderboard-builder/objects"
	reader "liquide/re/popularity-leaderboard-builder/reader/src/chaosMonkey"
	"liquide/re/popularity-leaderboard-builder/topology"
)

func main() {

	fmt.Println("Hello World")

	// TODO: use dig, https://blog.drewolson.org/dependency-injection-in-go

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

	// channel which receives user action
	var channel chan objects.UserAction

	// get reader, and pass him the channel
	reader := reader.ChaosMonkeyReader{}
	reader.Read(channel)

	run := true
	for run == true {

		userAction := <-channel

		go func(u *objects.UserAction) {

			tree.ProcessAction(u)
		}(&userAction)
	}
}
