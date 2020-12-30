package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tech-club/dc-bot/internal/events"
	"github.com/tech-club/dc-bot/pkg/log"
)

func registerEvents(s *discordgo.Session, log log.Logger) {
	s.AddHandler(events.NewReadyHandler(log).Handler)
}
