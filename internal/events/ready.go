package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tech-club/dc-bot/pkg/log"
)

type ReadyHandler struct {
	log log.Logger
}

func NewReadyHandler(log log.Logger) *ReadyHandler {
	return &ReadyHandler{log: log.WithPrefix("ready")}
}

func (h *ReadyHandler) Handler(_ *discordgo.Session, e *discordgo.Ready) {
	h.log.Println("bot session is ready")
	h.log.Printf("bot is logged in as '%s'\n", e.User.String())
}
