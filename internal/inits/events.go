package inits

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tech-club/dc-bot/internal/events"
)

func InitEventHandlers(s *discordgo.Session) {
	s.AddHandler(events.NewReadyHandler().Handler)
}
