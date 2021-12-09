package store

type LeaderboardStore interface {
	IncrementScoreForAnItem(boardName string, item string, points float64)
}
