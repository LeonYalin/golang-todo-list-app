package types

type EventType string

const (
	WELCOME     EventType = "welcome"
	ADD_TODO    EventType = "add_todo"
	REMOVE_TODO EventType = "remove_todo"
)

// type TodoMsg struct {
// 	Text string
// 	Type EventType
// }
