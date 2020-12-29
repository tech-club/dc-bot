package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO starting the bot")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}