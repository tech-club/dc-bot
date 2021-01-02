package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tech-club/dc-bot/pkg/db"
)

type Config struct {
	Bot botConfig
	Db  dbConfig
	Log logConfig
}

type botConfig struct {
	Prefix    string
	Token     string
	GuildJoin guildJoinConfig
}

type dbConfig struct {
	Username string
	Password string
	Driver   db.Driver
	Host     string
	Port     int
	Name     string
}

type logConfig struct {
	Level string
	Dir   string
}

type guildJoinConfig struct {
	WelcomeChannelID string
}

func LoadConfig() *Config {
	setDefaults()
	loadConfig()

	dbDriver, err := db.StringToDriver(viper.GetString("database.driver"))
	if err != nil {
		panic(err)
	}

	return &Config{
		Bot: botConfig{
			Prefix: viper.GetString("bot.prefix"),
			Token:  viper.GetString("bot.token"),
			GuildJoin: guildJoinConfig{
				WelcomeChannelID: viper.GetString("bot.guild_join.welcome_channel_id"),
			},
		},
		Db: dbConfig{
			Username: viper.GetString("database.username"),
			Password: viper.GetString("database.password"),
			Driver:   dbDriver,
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			Name:     viper.GetString("database.name"),
		},
		Log: logConfig{
			Level: viper.GetString("log.level"),
			Dir:   viper.GetString("log.dir"),
		},
	}
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %s", err))
	}

	viper.WatchConfig()
}
