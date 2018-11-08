package discord

import (
"fmt"
	"github.com/youssefhabri/z2bot/utils/colors"
	"os"
"strings"
"time"

"github.com/bwmarrin/discordgo"
"github.com/youssefhabri/z2bot/utils"
	)

func system(session *discordgo.Session, evt *discordgo.MessageCreate) {
	params := strings.Split(evt.Message.Content, " ")

	switch strings.ToLower(strings.TrimSpace(params[0])) {
	case utils.PREFIX + "uptime":
		hostname, err := os.Hostname()
		utils.PanicOnErr(err)
		duration := time.Now().Sub(startTime)
		utils.SendMessage(session, evt.ChannelID, fmt.Sprintf(
			"Uptime is: **%02d:%02d:%02d** (since **%s**) on **%s**",
			int(duration.Hours()),
			int(duration.Minutes())%60,
			int(duration.Seconds())%60,
			startTime.Format(time.Stamp),
			hostname))
		break
	case utils.PREFIX + "help":
		animeCommands := []string{
			fmt.Sprintf("**%sanime**     **<anime>**        Search for an anime.", utils.PREFIX),
			fmt.Sprintf("**%smanga**     **<manga>**        Search for a manga.", utils.PREFIX),
			fmt.Sprintf("**%scharacter** **<character>**    Search for an anime character.", utils.PREFIX),
			fmt.Sprintf("**%suser**      **<user>**         Search for an anilist user.", utils.PREFIX),
		}
		xkcdCommands := []string{
			fmt.Sprintf("**%sxkcd** **[random]**     Search for a random xkcd comic.", utils.PREFIX),
			fmt.Sprintf("**%sxkcd** **<number>**     Search for a specific xkcd comic.", utils.PREFIX),
			fmt.Sprintf("**%sxkcd** **latest**       Search for a the latest xkcd comic.", utils.PREFIX),
		}
		messageEmbed := utils.NewEmbed().
			SetColor(colors.DEFAULT).
			SetTitle("Zero Two's Commands").
			SetDescription("The list of usable commands for the best waifu, Zero Two").
			SetFooter("Powered by the best waifu, Zero Two").
			AddField("AniList commands", strings.Join(animeCommands, "\n")).
			AddField("xkcd commands", strings.Join(xkcdCommands, "\n")).
			MessageEmbed

		utils.SendMessageEmbed(session, evt.ChannelID, messageEmbed)
		break
	}
}

func testEmbedMsg(sess *discordgo.Session, evt *discordgo.MessageCreate) {
	params := strings.Split(evt.Message.Content, " ")
	channelID := evt.ChannelID

	switch strings.ToLower(strings.TrimSpace(params[0])) {
	case utils.PREFIX + "w":
		tn := time.Now()
		date := fmt.Sprintf("%s %d, %d", tn.Month().String(), tn.Day(), tn.Year())
		messageFooter := discordgo.MessageEmbedFooter{
			Text: "Sent by " + evt.Author.Username + " | " + date,
		}
		messageImage := discordgo.MessageEmbedImage{
			URL: "https://cdn.discordapp.com/attachments/453320571262992395/453350317707362326/discordgo-bot.png",
		}
		messageEmbed := discordgo.MessageEmbed{
			Color:       3447003,
			Title:       "Embedded Messages",
			Description: "Testing Embedded messages in DiscordGo :smile:",
			Image:       &messageImage,
			Footer:      &messageFooter,
		}

		controlEmojis := []*discordgo.Emoji {
			{Name:"prev", ID:"⬅"},
			{Name:"exit", ID:"❌"},
			{Name:"next", ID:"➡"},
		}

		message, _ :=sess.ChannelMessageSendEmbed(channelID, &messageEmbed)
		for _, emoji := range controlEmojis {
			sess.MessageReactionAdd(channelID, message.ID, emoji.ID)
		}
		break
	}
}
