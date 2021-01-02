package cmdHandler

type Command interface {
	Name() string
	Invokes() []string
	Description() string
	AdminRequired() bool
	Exec(ctx *Context) error
}
