package enums

type TaskStatus int64

const (
	NoTask    TaskStatus = 0
	OnHold    TaskStatus = 1
	InQueue   TaskStatus = 2
	PickedUp  TaskStatus = 3
	Delivered TaskStatus = 4
	Cancelled TaskStatus = 5
)
