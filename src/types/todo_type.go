package types

type Handler func(args ...any)

type TodoOption struct {
	Key     string
	Text    string
	Handler Handler
}
