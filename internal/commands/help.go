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
		nameField := &discordgo.MessageEmbedField{
			Name:   "Name",
			Value:  fmt.Sprintf("**%s**", cmd.Name()),
			Inline: true,
		}

		var invokes string
		for _, invoke := range cmd.Invokes() {
			invokes += fmt.Sprintf("`%s%s`, ", ctx.Handler.Prefix, invoke)
		}

		invokes = strings.TrimSuffix(invokes, ", ")

		invokesField := &discordgo.MessageEmbedField{
			Name:   "Usage",
			Value:  invokes,
			Inline: true,
		}

		descriptionField := &discordgo.MessageEmbedField{
			Name:   "Description",
			Value:  cmd.Description(),
			Inline: true,
		}

		fields = append(fields, nameField, invokesField, descriptionField)
	}

	helpMessage := &discordgo.MessageEmbed{
		Title:       "Help",
		Description: "help for all commands",
		Fields:      fields,
	}

	_, err := ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, helpMessage)
	return err
}
