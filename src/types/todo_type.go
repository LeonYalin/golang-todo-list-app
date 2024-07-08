package types

type TodoOption struct {
	Key     string
	Text    string
	Handler func(args ...any)
}
