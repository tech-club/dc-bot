package bot

import (
	"github.com/tech-club/dc-bot/internal/events"
)

func (b *Bot) registerEvents() {
	b.Session.AddHandler(events.NewReadyHandler(b.Log.WithPrefix("ready")).Handler)
	b.Session.AddHandler(events.NewGuildJoinHandler(b.Log.WithPrefix("guild_join"), b.Config).Handler)
}
