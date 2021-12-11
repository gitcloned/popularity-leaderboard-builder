package reader

import (
	objects "liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/reader/interfaces"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func pickRandomUser(list []objects.User) objects.User {

	if len(list) == 0 {
		return objects.User{}
	}

	index := rand.Intn(len(list))
	return list[index]
}

type ActionBuilder struct {
	ActionTypes []string
	Users       []objects.User

	Store *ItemStore
}

func (b *ActionBuilder) build(item *objects.Item) *objects.UserAction {

	actionType := pickRandom(b.ActionTypes)
	user := pickRandomUser(b.Users)

	action := objects.UserAction{
		ActionType: actionType,
		Item:       *item,
		Id:         uuid.NewString(),
		User:       user.Name,
		UserCohert: user.Cohert,
		Timestamp:  time.Now(),
		Points:     rand.Float64() * 10,
	}

	return &action
}

func (b *ActionBuilder) Start(d *interfaces.EventDispatcher) {

	for {

		logrus.Info("Building action, picking an item")

		// pick item
		item := b.Store.pick()

		if item != nil {

			// build action
			action := b.build(item)

			logrus.Info("Dispatching action")

			if action != nil {

				// add to queue
				d.Queue <- *action
			}
		}

		time.Sleep(time.Second)
	}
}
