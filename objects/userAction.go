package objects

import "time"

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
