package cmdHandler

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type CommandHandler struct {
	Prefix       string
	CmdInstances []Command
	cmdMap       map[string]Command
	middlewares  []Middleware
	OnError      func(err error, ctx *Context)
}

func NewCommandHandler(prefix string) *CommandHandler {
	return &CommandHandler{
		Prefix:       prefix,
		CmdInstances: make([]Command, 0),
		cmdMap:       make(map[string]Command),
		middlewares:  make([]Middleware, 0),
		OnError:      func(error, *Context) {},
	}
}

func (c *CommandHandler) RegisterCommand(cmd Command) {
	c.CmdInstances = append(c.CmdInstances, cmd)
	for _, invoke := range cmd.Invokes() {
		c.cmdMap[invoke] = cmd
	}
}

func (c *CommandHandler) RegisterMiddleware(mw Middleware) {
	c.middlewares = append(c.middlewares, mw)
}

func (c *CommandHandler) MessageHandler(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author.Bot || e.Author.ID == s.State.User.ID || !strings.HasPrefix(e.Content, c.Prefix) {
		return
	}

	parts := strings.Split(e.Content[len(c.Prefix):], " ")
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
		Args:    args,
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
