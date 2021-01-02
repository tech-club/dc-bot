package bot

import (
	"fmt"
	"github.com/tech-club/dc-bot/internal/config"
	"github.com/tech-club/dc-bot/pkg/cmdHandler"
	"github.com/tech-club/dc-bot/pkg/log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session    *discordgo.Session
	Log        log.Logger
	Config     *config.Config
	CmdHandler *cmdHandler.CommandHandler
}

func New() (*Bot, error) {
	var err error
	var bot Bot

	bot.Config = config.LoadConfig()
	bot.Log = log.New(os.Stderr, bot.Config.Log.Level, bot.Config.Log.Dir)
	bot.Session, err = discordgo.New("Bot " + bot.Config.Bot.Token)
	if err != nil {
		return nil, err
	}

	bot.Session.Identify.Intents = discordgo.MakeIntent(
		discordgo.IntentsGuildMembers | discordgo.IntentsGuildMessages)

	bot.CmdHandler = cmdHandler.NewCommandHandler(bot.Config.Bot.Prefix)
	bot.CmdHandler.OnError = func(err error, ctx *cmdHandler.Context) {
		fmt.Println("err while exec command: ", err)
	}

	return &bot, nil
}

func (b *Bot) Run() {
	b.registerEvents()
	b.registerCommands()
	b.registerMiddlewares()

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
