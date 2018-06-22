package plugins

import (
	"github.com/bwmarrin/discordgo"
	"github.com/youssefhabri/anitrend-bot/plugins/admin"
	"github.com/youssefhabri/anitrend-bot/plugins/anisearch"
	"github.com/youssefhabri/anitrend-bot/plugins/xkcd"
)

func Register(session *discordgo.Session) {

	// Register plugins
	admin.Register(session)
	anisearch.Register(session)
	xkcd.Register(session)
}
