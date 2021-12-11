package reader

import (
	objects "liquide/re/popularity-leaderboard-builder/objects"
	"math/rand"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type RankedItem struct {
	Item objects.Item
	Rank int
}

type ItemStore struct {
	list []RankedItem
	Size int
}

func (s *ItemStore) trim() {

	lock.Lock()
	defer lock.Unlock()

	length := len(s.list)

	if length > s.Size {
		s.list = s.list[0 : length-s.Size]
	}
}

func (s *ItemStore) add(item *objects.Item) {

	lock.Lock()
	defer lock.Unlock()

	s.list = append(s.list, RankedItem{
		Item: *item,
		Rank: 0,
	})

	// logrus.Info(fmt.Sprintf("Added an item to store, len: %d", len(s.list)))
}

func (s *ItemStore) sortFromIndex(index int) {

	for index > 0 {

		item := s.list[index]
		parent := s.list[index-1]

		if parent.Rank >= item.Rank {
			break
		}

		temp := parent
		parent = item
		item = temp

		index = index - 1
	}
}

func (s *ItemStore) pick() *objects.Item {

	lock.Lock()
	defer lock.Unlock()

	// logrus.Info(fmt.Sprintf("Picking an item from store, len: %d", len(s.list)))

	if len(s.list) == 0 {
		return nil
	}

	index := rand.Intn(len(s.list))

	rItem := s.list[index]
	rItem.Rank = rItem.Rank + 1

	s.sortFromIndex(index)

	return &rItem.Item
}

func (s *ItemStore) Start() {

	go func() {

		for {

			s.trim()

			time.Sleep(time.Second)
		}
	}()
}
