package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tech-club/dc-bot/internal/bot"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		bot, err := bot.New()
		if err != nil {
			panic(err)
		}

		bot.Run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
