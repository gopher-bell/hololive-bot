package discord

import (
	"github.com/gopher-bell/hololive-bot/log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	botToken string
	guildID  string
)

func init() {
	botToken = os.Getenv("DISCORD_BOT_TOKEN")
	guildID = os.Getenv("DISCORD_GUILD_ID")
}

func New() (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.ZapLog.Errorw("invalid bot parameters", "bot_token", botToken, "error", err.Error())
		return nil, err
	}

	log.ZapLog.Infoln(botToken)
	log.ZapLog.Infoln(guildID)

	return session, nil
}
