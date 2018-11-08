package admin

import (
"strconv"
"strings"



"github.com/bwmarrin/discordgo"
"github.com/youssefhabri/z2bot/utils"
)

func Register(session *discordgo.Session) {
	session.AddHandler(system)
}

func system(sess *discordgo.Session, evt *discordgo.MessageCreate) {
	params := strings.Split(evt.Message.Content, " ")
	channelID := evt.ChannelID

	switch strings.ToLower(strings.TrimSpace(params[0])) {
	case utils.PREFIX + "cleanup":
		limit := 10
		if len(params) > 1 && len(params[1]) > 0 {
			limit, _ = strconv.Atoi(params[1])
		}
		messages, _ := sess.ChannelMessages(channelID, limit, "", "", "")
		var messagesIDs []string
		for _, message := range messages {
			messagesIDs = append(messagesIDs, message.ID)
		}
		err := sess.ChannelMessagesBulkDelete(channelID, messagesIDs)
		if err != nil {
			utils.LogDebug(err)
		}
		break

	}
}
