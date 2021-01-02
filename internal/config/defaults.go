package config

import "github.com/spf13/viper"

func setDefaults() {
	viper.SetDefault("bot.prefix", "!")
	viper.SetDefault("bot.token", "your_discord_bot_token")

	viper.SetDefault("bot.guild_join.welcome_channel_id", "channel_id_for_welcome_messages")

	viper.SetDefault("database.username", "dbuser")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.driver", "mysql")
	viper.SetDefault("database.host", "0.0.0.0")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("database.name", "dc_bot")

	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.dir", ".")
}
