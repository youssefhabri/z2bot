package xkcd

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/youssefhabri/zero2-go/utils"
)

func Register(session *discordgo.Session) {
	session.AddHandler(fetchComics)
}

func fetchComics(session *discordgo.Session, evt *discordgo.MessageCreate) {
	params := strings.Split(evt.Message.Content, " ")
	channelID := evt.ChannelID

	xkcd := NewClient()
	var comic Comic

	switch strings.ToLower(strings.TrimSpace(params[0])) {
	case utils.PREFIX + "xkcd":
		session.ChannelTyping(channelID)
		if len(params) > 1 {
			switch params[1] {
			case "latest":
				comic = xkcd.LatestComic()
				showComic(session, channelID, &comic)
				break
			case "random":
				comic = xkcd.RandomComic()
				showComic(session, channelID, &comic)
				break
			default:
				comic = xkcd.Comic(params[1])
				showComic(session, channelID, &comic)
				break
			}
		} else {
			comic = xkcd.RandomComic()
			showComic(session, channelID, &comic)
		}
		break
	}
}

func showComic(session *discordgo.Session, channelID string, comic *Comic) {
	messageEmbed := utils.NewEmbed().
		SetColor(3447003).
		SetTitle(comic.GetTitle()).
		SetDescription(comic.ImgALT).
		SetURL(comic.GetLink()).
		SetImage(comic.ImgURL).
		SetFooter("Powered by xkcd").
		MessageEmbed

	utils.SendMessageEmbed(session, channelID, messageEmbed)
}