package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/tech-club/dc-bot/pkg/cmdHandler"
	"github.com/tech-club/dc-bot/pkg/log"
	"strings"
)

type HelpCommand struct {
	log log.Logger
}

func NewHelpCommand(log log.Logger) *HelpCommand {
	return &HelpCommand{
		log: log,
	}
}

func (c HelpCommand) Name() string {
	return "help"
}

func (c HelpCommand) Invokes() []string {
	return []string{"help", "h"}
}

func (c HelpCommand) Description() string {
	return "shows help for all commands"
}

func (c HelpCommand) AdminRequired() bool {
	return false
}

func (c HelpCommand) Exec(ctx *cmdHandler.Context) error {
	c.log.Debugf("user '%s' is executing help command", ctx.Message.Author.String())

	fields := make([]*discordgo.MessageEmbedField, 0)

	for _, cmd := range ctx.Handler.CmdInstances {
		var name string

		for _, cmdName := range cmd.Invokes() {
			name += fmt.Sprintf("`%s`, ", cmdName)
		}

		name = strings.TrimSuffix(name, ", ")

		field := &discordgo.MessageEmbedField{
			Name:   name,
			Value:  cmd.Description(),
			Inline: false,
		}
		fields = append(fields, field)
	}

	helpMessage := &discordgo.MessageEmbed{
		Title:       "Help",
		Description: fmt.Sprintf("help for all commands\nPrefix: `%s`", ctx.Handler.Prefix),
		Fields:      fields,
	}

	_, err := ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, helpMessage)
	return err
}
