package discord

import (
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/youssefhabri/z2bot-go/plugins"
	"github.com/youssefhabri/z2bot-go/scripting"
	"github.com/youssefhabri/z2bot-go/utils"
)

var startTime time.Time
var session *discordgo.Session
var discordErr error

func Init() {
	startTime = time.Now()

	utils.LogInfo("Logging in...")
	utils.LogInfo("Logging in with bot account token...")
	session, discordErr = discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	setupHandlers(session)

	scripting.Init(session)
	utils.LogError(discordErr)

	utils.LogInfo("Opening session...")
	discordErr = session.Open()
	utils.LogError(discordErr)

	utils.LogInfo("Sleeping...")
	// <-make(chan struct{})
}

func GetSession() *discordgo.Session {
	return session
}

func GetPrimaryChannel() string {
	return utils.FetchPrimaryTextChannelID(session)
}

func setupHandlers(session *discordgo.Session) {
	utils.LogInfo("Setting up event handlers...")

	session.AddHandler(system)
	session.AddHandler(testEmbedMsg)

	plugins.Register(session)

	session.AddHandler(func(sess *discordgo.Session, evt *discordgo.PresenceUpdate) {
		utils.LogDebug(fmt.Sprintf("PRESENSE UPDATE fired for user-ID: %s & username: %s", evt.User.ID, evt.User.Username))
		fmt.Printf("%+v\n", evt)
		fmt.Printf("%#v\n", evt.User)
		self := utils.FetchUser(sess, "@me")
		u := utils.FetchUser(sess, evt.User.ID)
		// Ignore self
		if u.ID == self.ID || u.Bot {
			return
		}
		// Handle online/offline notifications
		if evt.Status == "offline" {
			if _, ok := utils.GetOnlineUser(u.ID); ok {
				utils.DeleteOnlineUser(u.ID)
				// utils.SendMessage(sess, fmt.Sprintf(`**%s** went offline`, u.Username))
			}
		} else {
			if _, ok := utils.GetOnlineUser(u.ID); !ok {
				utils.SetOnlineUser(u.ID, u)
				// utils.SendMessage(sess, fmt.Sprintf(`**%s** is now online`, u.Username))
			}
		}
	})

	session.AddHandler(func(session *discordgo.Session, evt *discordgo.GuildCreate) {
		utils.LogInfo("GUILD_CREATE event fired")
		for _, presence := range evt.Presences {
			user := presence.User
			utils.LogInfo(fmt.Sprintf("Marked User online - ID: %s, Username: %s", user.ID, user.Username))
			utils.SetOnlineUser(user.ID, user)
		}
	})

	session.AddHandler(func(session *discordgo.Session, evt *discordgo.MessageReactionAdd) {
		if evt.MessageReaction.Emoji.Name == "❌" {
			utils.LogInfo("REACTION_CREATE event fired: " + evt.Emoji.Name)
			session.MessageReactionRemove(evt.ChannelID, evt.MessageID, "⬅", evt.UserID)
			session.MessageReactionRemove(evt.ChannelID, evt.MessageID, "❌", evt.UserID)
			session.MessageReactionRemove(evt.ChannelID, evt.MessageID, "➡", evt.UserID)
		}
	})
}
