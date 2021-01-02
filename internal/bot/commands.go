package bot

import (
	"github.com/tech-club/dc-bot/internal/commands"
)

func (b *Bot) registerCommands() {
	b.CmdHandler.RegisterCommand(commands.NewPingCommand(b.Log.WithPrefix("cmd_ping")))
	b.CmdHandler.RegisterCommand(commands.NewHelpCommand(b.Log.WithPrefix("cmd_help")))

	b.Session.AddHandler(b.CmdHandler.MessageHandler)
}
