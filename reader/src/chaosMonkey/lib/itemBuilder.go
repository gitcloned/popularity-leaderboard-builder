package reader

import (
	"fmt"
	objects "liquide/re/popularity-leaderboard-builder/objects"
	"time"

	uuid "github.com/google/uuid"
)

type ItemBuilder struct {
	Store *ItemStore

	ItemTypes []string
	Topics    []string
	Stocks    []string
	Authors   []string
	Channels  []string
}

func (b *ItemBuilder) build() *objects.Item {

	itemType := pickRandom(b.ItemTypes)
	topic := pickRandom(b.Topics)
	stock := pickRandom(b.Stocks)
	author := pickRandom(b.Authors)
	channel := pickRandom(b.Channels)

	item := objects.Item{
		Id:       uuid.NewString(),
		Topic:    topic,
		Stock:    stock,
		Author:   author,
		ItemType: itemType,
		Channel:  channel,
		Content:  fmt.Sprintf("%s: %s %s", itemType, stock, topic),
	}

	return &item
}

func (b *ItemBuilder) Start() {

	go func() {

		for {

			item := b.build()

			b.Store.add(item)

			time.Sleep(time.Second)
		}
	}()
}
