package bot

import (
	"github.com/tech-club/dc-bot/pkg/log"
	"os"
	"os/signal"
	"syscall"

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

func (b *Bot) Run() {
	registerEvents(b.Session, b.Log)
	err := b.Session.Open()
	if err != nil {
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	b.Log.Println("closing bot session")
	err = b.Session.Close()
	if err != nil {
		panic(err)
	}
}
