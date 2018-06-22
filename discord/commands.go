package discord

import (
"fmt"
"os"
"strings"
"time"

"github.com/bwmarrin/discordgo"
"github.com/youssefhabri/anitrend-bot/utils"
	)

func system(sess *discordgo.Session, evt *discordgo.MessageCreate) {
	params := strings.Split(evt.Message.Content, " ")

	switch strings.ToLower(strings.TrimSpace(params[0])) {
	case utils.PREFIX + "uptime":
		hostname, err := os.Hostname()
		utils.PanicOnErr(err)
		duration := time.Now().Sub(startTime)
		utils.SendMessage(sess, fmt.Sprintf(
			"Uptime is: **%02d:%02d:%02d** (since **%s**) on **%s**",
			int(duration.Hours()),
			int(duration.Minutes())%60,
			int(duration.Seconds())%60,
			startTime.Format(time.Stamp),
			hostname))
		break
	}
}

func testEmbedMsg(sess *discordgo.Session, evt *discordgo.MessageCreate) {
	params := strings.Split(evt.Message.Content, " ")
	channelID := utils.FetchPrimaryTextChannelID(sess)

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
