package store

import "liquide/re/popularity-leaderboard-builder/objects"

type ItemStore interface {
	ReadItem() objects.Item
	PutItem(objects.Item)
}
