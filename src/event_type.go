package src

type EventType string

const (
	WELCOME     EventType = "welcome"
	ADD_TODO    EventType = "add_todo"
	REMOVE_TODO EventType = "remove_todo"
)

// type eventTypes struct {
// 	WELCOME     EventType
// 	ADD_TODO    EventType
// 	DELETE_TODO EventType
// }

// var EventTypes = &eventTypes{
// 	WELCOME:     "welcome",
// 	ADD_TODO:    "add_todo",
// 	DELETE_TODO: "delete_todo",
// }

// type TodoMsg struct {
// 	Text string
// 	Type EventType
// }
