package middlewares

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tech-club/dc-bot/pkg/cmdHandler"
	"github.com/tech-club/dc-bot/pkg/log"
)

type PermissionMiddleware struct {
	log log.Logger
}

func NewPermissionMiddleware(log log.Logger) *PermissionMiddleware {
	return &PermissionMiddleware{
		log: log,
	}
}

func (p PermissionMiddleware) Exec(ctx *cmdHandler.Context, cmd cmdHandler.Command) (bool, error) {
	if !cmd.AdminRequired() {
		return true, nil
	}

	guild, err := ctx.Session.Guild(ctx.Message.GuildID)
	if err != nil {
		return false, err
	}

	if guild.OwnerID == ctx.Message.Author.ID {
		return true, nil
	}

	roles := make(map[string]*discordgo.Role)
	for _, role := range guild.Roles {
		roles[role.ID] = role
	}

	for _, roleID := range ctx.Message.Member.Roles {
		if role, ok := roles[roleID]; ok && role.Permissions&discordgo.PermissionAdministrator > 0 {
			return false, nil
		}
	}

	_, err = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, &discordgo.MessageEmbed{
		Description: "You dont have the permission to execute this command!",
		Color:       16711680,
	})

	return false, err
}
