package discord

import (
	//"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/youssefhabri/z2bot/plugins"
	"github.com/youssefhabri/z2bot/utils"
	"github.com/yuin/gopher-lua"
)

var startTime time.Time
var session *discordgo.Session
var discordErr error

var accountToken = "NDUzNzczMDAxODA1MTM1ODgz.DfjwlA.CiarL8FdMZS-27VlOrWdrgqCOl8"

func Init() {
	startTime = time.Now()

	utils.LogInfo("Logging in...")

	utils.LogInfo("Logging in with bot account token...")
	session, discordErr = discordgo.New("Bot " + accountToken)
	setupHandlers(session)
	setupScripting(session)
	utils.PanicOnErr(discordErr)

	utils.LogInfo("Opening session...")
	discordErr = session.Open()
	utils.PanicOnErr(discordErr)

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
		utils.LogDebug("PRESENSE UPDATE fired for user-ID:", evt.User.ID)
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

	session.AddHandler(func(sess *discordgo.Session, evt *discordgo.GuildCreate) {
		utils.LogInfo("GUILD_CREATE event fired")
		for _, presence := range evt.Presences {
			user := presence.User
			utils.LogInfo("Marked user-ID online:", user.ID)
			utils.SetOnlineUser(user.ID, user)
		}
	})

	session.AddHandler(func(sess *discordgo.Session, evt *discordgo.MessageReactionAdd) {
		utils.LogInfo("REACTION_CREATE event fired: " + evt.Emoji.ID + " " + evt.Emoji.Name)
		if evt.MessageReaction.Emoji.ID == "❌" {
			sess.MessageReactionsRemoveAll(evt.ChannelID, evt.MessageID)
		}
	})
}

func setupScripting(session *discordgo.Session) {
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("discord", Loader)
	if err := L.DoFile("scripts/main.lua"); err != nil {
		panic(err)
	}
}

func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)
	// register other stuff
	L.SetField(mod, "name", lua.LString("value"))

	// returns the module
	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	"message": LMessage,
}

func LMessage(L *lua.LState) int {
	command := L.Get(1).String()
	message := L.Get(2).String()

	utils.LogDebug(command, message)

	session.AddHandler(func(session *discordgo.Session, evt *discordgo.MessageCreate) {
		params := strings.Split(evt.Message.Content, " ")

		switch strings.ToLower(strings.TrimSpace(params[0])) {
		case utils.PREFIX + command:
			utils.SendMessage(session, evt.ChannelID, message)
			break
		}
	})
	return 0
}
