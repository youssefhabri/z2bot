package plugins

import (
	"github.com/bwmarrin/discordgo"
	//"github.com/youssefhabri/zero2-go/plugins/admin"
	"github.com/youssefhabri/zero2-go/plugins/anisearch"
	"github.com/youssefhabri/zero2-go/plugins/xkcd"
)

func Register(session *discordgo.Session) {

	// Register plugins
	// admin.Register(session)
	anisearch.Register(session)
	xkcd.Register(session)
}
