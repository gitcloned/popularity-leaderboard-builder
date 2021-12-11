package objects

import (
	"fmt"
	"time"
)

type UserAction struct {
	Id         string
	Item       Item
	Channel    string
	User       string
	UserCohert string
	Timestamp  time.Time
	ActionType string
	Points     float64
}

func (u *UserAction) String() string {

	return fmt.Sprintf("%s - %s", u.Channel, u.Item.String())
}
