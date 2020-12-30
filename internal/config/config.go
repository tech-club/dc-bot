package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Bot botConfig
	Log logConfig
}

type botConfig struct {
	Prefix    string
	Token     string
	GuildJoin guildJoinConfig
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

	return &Config{
		Bot: botConfig{
			Prefix: viper.GetString("bot.prefix"),
			Token:  viper.GetString("bot.token"),
			GuildJoin: guildJoinConfig{
				WelcomeChannelID: viper.GetString("bot.guild_join.welcome_channel_id"),
			},
		},
		Log: logConfig{
			Level: viper.GetString("log.level"),
			Dir:   viper.GetString("log.dir"),
		},
	}
}

func setDefaults() {
	viper.SetDefault("bot.prefix", "!")
	viper.SetDefault("bot.token", "your_discord_bot_token")

	viper.SetDefault("bot.guild_join.welcome_channel_id", "channel_id_for_welcome_messages")

	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.dir", ".")
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
