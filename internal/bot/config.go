package bot

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Bot botConfig
	Log logConfig
}

type botConfig struct {
	Prefix string
	Token  string
}

type logConfig struct {
	Level string
	Dir   string
}

func LoadConfig() *Config {
	setDefaults()
	loadConfig()

	return &Config{
		Bot: botConfig{
			Prefix: viper.GetString("bot.prefix"),
			Token:  viper.GetString("bot.token"),
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
