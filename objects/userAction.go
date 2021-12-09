package objects

import "time"

type ActionType int

const (
	Viewed = iota
	Replied
	Shared
	Rated
	Consumed
)

type UserAction struct {
	Id         string
	Item       string
	Channel    string
	User       string
	UserCohert string
	Timestamp  time.Time
	ActionType ActionType
	Points     float64
}
