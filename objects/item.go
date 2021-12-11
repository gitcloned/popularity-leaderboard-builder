package objects

type Item struct {
	Id       string
	ItemType string
	Stock    string
	Topic    string
	Content  string
	Author   string
	Channel  string
}

func (i *Item) String() string {

	return i.Content
}
