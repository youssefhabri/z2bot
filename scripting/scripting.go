package scripting

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/youssefhabri/zero2-go/utils"
)

var session *discordgo.Session

func Init(sess *discordgo.Session) {
	session = sess

	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("discord", discordLoader)
	if err := L.DoFile("scripts/main.lua"); err != nil {
		panic(err)
	}
}

func discordLoader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)
	// register other stuff
	L.SetField(mod, "name", lua.LString("value"))

	// returns the module
	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	"message": lMessage,
}

func lMessage(L *lua.LState) int {
	command := L.Get(1).String()
	arg2 := L.Get(2)
	var messages []string

	if arg2.Type() == lua.LTString {
		messages = append(messages, arg2.String())
	} else if arg2.Type() == lua.LTTable {
		L.ToTable(2).ForEach(func(key lua.LValue, value lua.LValue) {
			messages = append(messages, value.String())
		})
	}

	session.AddHandler(func(session *discordgo.Session, evt *discordgo.MessageCreate) {
		params := strings.Split(evt.Message.Content, " ")

		message := messages[0]

		if len(messages) > 1 {
			n := utils.Random(0, len(messages))
			message = messages[n]
		}

		switch strings.ToLower(strings.TrimSpace(params[0])) {
		case utils.PREFIX + command:
			utils.SendMessage(session, evt.ChannelID, message)
			break
		}
	})
	return 0
}
