package bot

import "github.com/tech-club/dc-bot/internal/middlewares"

func (b *Bot) registerMiddlewares() {
	b.CmdHandler.RegisterMiddleware(middlewares.NewPermissionMiddleware(b.Log.WithPrefix("mw_permission")))
}
