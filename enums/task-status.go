package enums

type TaskStatus int64

const (
	New       TaskStatus = 0
	InProcess TaskStatus = 1
	Done      TaskStatus = 2
	Error     TaskStatus = 3
)
