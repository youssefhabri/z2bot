package plugins

import (
	"github.com/bwmarrin/discordgo"
	"github.com/youssefhabri/z2bot/plugins/admin"
	"github.com/youssefhabri/z2bot/plugins/anisearch"
	"github.com/youssefhabri/z2bot/plugins/xkcd"
)

func Register(session *discordgo.Session) {

	// Register plugins
	admin.Register(session)
	anisearch.Register(session)
	xkcd.Register(session)
}
