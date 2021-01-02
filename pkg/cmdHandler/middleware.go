package cmdHandler

type Middleware interface {
	Exec(ctx *Context, cmd Command) (bool, error)
}
