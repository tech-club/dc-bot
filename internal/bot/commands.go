package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/tech-club/dc-bot/internal/commands"
	"github.com/tech-club/dc-bot/internal/config"
	"github.com/tech-club/dc-bot/pkg/cmdHandler"
	"github.com/tech-club/dc-bot/pkg/log"
)

func registerCommands(s *discordgo.Session, log log.Logger, config *config.Config) {
	commandHandler := cmdHandler.NewCommandHandler(config.Bot.Prefix)
	commandHandler.OnError = func(err error, ctx *cmdHandler.Context) {
		fmt.Println("err while exec command: ", err)
	}

	commandHandler.RegisterCommand(commands.NewPingCommand(log.WithPrefix("cmd_ping")))
	commandHandler.RegisterCommand(commands.NewHelpCommand(log.WithPrefix("cmd_help")))

	s.AddHandler(commandHandler.MessageHandler)
}
