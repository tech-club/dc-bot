package commands

import (
	"github.com/tech-club/dc-bot/pkg/cmdHandler"
	"github.com/tech-club/dc-bot/pkg/log"
)

type PingCommand struct {
	log log.Logger
}

func NewPingCommand(log log.Logger) *PingCommand {
	return &PingCommand{
		log: log,
	}
}

func (c *PingCommand) Invokes() []string {
	return []string{"ping", "p"}
}

func (c *PingCommand) Description() string {
	return "ping command returns 'pong' and mention the member"
}

func (c *PingCommand) AdminRequired() bool {
	return false
}

func (c *PingCommand) Exec(ctx *cmdHandler.Context) error {
	c.log.Debugf("user '%s' is executing ping command", ctx.Message.Author.String())

	_, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "pong "+ctx.Message.Author.Mention())
	return err
}
