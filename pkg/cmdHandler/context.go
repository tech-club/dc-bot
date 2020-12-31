package cmdHandler

import "github.com/bwmarrin/discordgo"

type Context struct {
	Session *discordgo.Session
	Message *discordgo.Message
	args    []string
	Handler *CommandHandler
}
