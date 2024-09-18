package sensors

type Item interface {
	GetItemID() string
	GetItemType() string
	Event
}
