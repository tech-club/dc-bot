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

func (p PingCommand) Invokes() []string {
	return []string{"ping", "p"}
}

func (p PingCommand) Description() string {
	return "ping command returns 'pong' and mention the member"
}

func (p PingCommand) AdminRequired() bool {
	return false
}

func (p PingCommand) Exec(ctx *cmdHandler.Context) error {
	p.log.Debugf("user '%s' is executing ping command", ctx.Message.Author.String())

	_, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "pong "+ctx.Message.Author.Mention())
	return err
}
