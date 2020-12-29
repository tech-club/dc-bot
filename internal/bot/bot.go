package bot

import (
	"github.com/tech-club/dc-bot/pkg/log"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session
	Log     log.Logger
	Config  *Config
}

func New() (*Bot, error) {
	var err error
	var bot Bot

	bot.Config = LoadConfig()
	bot.Log = log.New(os.Stderr, bot.Config.Log.Level, bot.Config.Log.Dir)
	bot.Session, err = discordgo.New("Bot " + bot.Config.Bot.Token)
	if err != nil {
		return nil, err
	}

	return &bot, nil
}
