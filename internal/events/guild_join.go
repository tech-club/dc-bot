package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tech-club/dc-bot/internal/config"
	"github.com/tech-club/dc-bot/pkg/log"
)

type GuildJoinHandler struct {
	log    log.Logger
	config *config.Config
}

func NewGuildJoinHandler(log log.Logger, config *config.Config) *GuildJoinHandler {
	return &GuildJoinHandler{
		log:    log,
		config: config,
	}
}

func (h *GuildJoinHandler) Handler(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	guild, err := s.Guild(e.GuildID)
	if err != nil {
		h.log.Warnln("failed getting guild object: ", err)
	}

	if e.Member.User.Bot {
		h.log.Debugf("bot '%s' joined guild '%s'", e.Member.User.String(), guild.Name)
		return
	}

	h.log.Debugf("member '%s' joined guild '%s'", e.Member.User.String(), guild.Name)

	_, err = s.ChannelMessageSendEmbed(h.config.Bot.GuildJoin.WelcomeChannelID, &discordgo.MessageEmbed{
		Title:       e.Member.User.String(),
		Description: "Welcome " + e.Member.Mention() + " on the " + guild.Name + " Discord Server",
		Color:       220582,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: e.Member.User.AvatarURL("256"),
		},
	})
	if err != nil {
		h.log.Warnln("failed to send welcome message: ", err)
	}

	userChannel, err := s.UserChannelCreate(e.Member.User.ID)
	if err != nil {
		h.log.Warnln("failed to create user channel: ", err)
	} else {
		_, err = s.ChannelMessageSendEmbed(userChannel.ID, &discordgo.MessageEmbed{
			Title:       guild.Name,
			Description: "Welcome " + e.Member.Mention() + " on the " + guild.Name + " Discord Server",
			Color:       220582,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: guild.IconURL(),
			},
		})
		if err != nil {
			h.log.Warnln("failed to send welcome message to user: ", err)
		}
	}
}
