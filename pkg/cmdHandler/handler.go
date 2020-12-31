package cmdHandler

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type CommandHandler struct {
	prefix       string
	cmdInstances []Command
	cmdMap       map[string]Command
	middlewares  []Middleware
	OnError      func(err error, ctx *Context)
}

func NewCommandHandler(prefix string) *CommandHandler {
	return &CommandHandler{
		prefix:       prefix,
		cmdInstances: make([]Command, 0),
		cmdMap:       make(map[string]Command),
		middlewares:  make([]Middleware, 0),
		OnError:      func(error, *Context) {},
	}
}

func (c *CommandHandler) RegisterCommand(cmd Command) {
	c.cmdInstances = append(c.cmdInstances, cmd)
	for _, invoke := range cmd.Invokes() {
		c.cmdMap[invoke] = cmd
	}
}

func (c *CommandHandler) RegisterMiddleware(mw Middleware) {
	c.middlewares = append(c.middlewares, mw)
}

func (c *CommandHandler) MessageHandler(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author.Bot || e.Author.ID == s.State.User.ID || !strings.HasPrefix(e.Content, c.prefix) {
		return
	}

	parts := strings.Split(e.Content[len(c.prefix):], "")
	if len(parts) < 1 {
		return
	}

	invoke := strings.ToLower(parts[0])
	args := parts[1:]

	cmd, ok := c.cmdMap[invoke]
	if !ok || cmd == nil {
		return
	}

	ctx := &Context{
		Session: s,
		Message: e.Message,
		args:    args,
		Handler: c,
	}

	for _, mw := range c.middlewares {
		next, err := mw.Exec(ctx, cmd)
		if err != nil {
			c.OnError(err, ctx)
			return
		}

		if !next {
			return
		}
	}

	if err := cmd.Exec(ctx); err != nil {
		c.OnError(err, ctx)
	}
}
