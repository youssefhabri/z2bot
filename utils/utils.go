package utils

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
	"errors"
	"log"
	"os"
	"bytes"
	"html/template"
	"html"
)

const PREFIX = "a2!"
var logger *log.Logger
var usersOnline map[string]*discordgo.User

func init() {
	logger = log.New(os.Stderr, "  ", log.Ldate|log.Ltime)
	usersOnline = make(map[string]*discordgo.User)
}

func GetOnlineUsers() map[string]*discordgo.User {
	return usersOnline
}

func GetOnlineUser(userId string) (*discordgo.User, bool) {
	user, ok := usersOnline[userId]
	return user, ok
}

func GetOnlineUserByName(username string) (*discordgo.User, bool) {
	for _, user := range usersOnline {
		if user.Username == username {
			return user, true
		}
	}
	return nil, false
}

func SetOnlineUser(userId string, user *discordgo.User) {
	usersOnline[userId] = user
}

func DeleteOnlineUser(userId string) {
	delete(usersOnline, userId)
}

func LogError(v ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(v...)
}

func LogDebug(v ...interface{}) {
	logger.SetPrefix("DEBUG ")
	logger.Println(v...)
}

func LogInfo(v ...interface{}) {
	logger.SetPrefix("INFO  ")
	logger.Println(v...)
}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

/* Tries to call a method and checking if the method returned an error, if it
did check to see if it's HTTP 502 from the Discord API and retry for
`attempts` number of times. */
func RetryOnBadGateway(f func() error) {
	var err error
	for i := 0; i < 3; i++ {
		err = f()
		if err != nil {
			if strings.HasPrefix(err.Error(), "HTTP 502") {
				// If the error is Bad Gateway, try again after 1 sec.
				time.Sleep(1 * time.Second)
				continue
			} else {
				// Otherwise panic !
				PanicOnErr(err)
			}
		} else {
			// In case of no error, return.
			return
		}
	}
}

func FetchUser(sess *discordgo.Session, userid string) *discordgo.User {
	var result *discordgo.User
	RetryOnBadGateway(func() error {
		var err error
		result, err = sess.User(userid)
		if err != nil {
			return err
		}
		return nil
	})
	return result
}

func FetchUserByName(sess *discordgo.Session, username string) *discordgo.User {
	var result *discordgo.User
	RetryOnBadGateway(func() error {
		var err error
		result, err = sess.User(username)
		if err != nil {
			return err
		}
		return nil
	})
	return result
}

func FetchPrimaryTextChannelID(sess *discordgo.Session) string {
	var channelid string
	RetryOnBadGateway(func() error {
		guilds, err := sess.UserGuilds(1, "", "")
		if err != nil {
			return err
		}
		guild, err := sess.Guild(guilds[0].ID)
		if err != nil {
			return err
		}
		channels, err := sess.GuildChannels(guild.ID)
		if err != nil {
			return err
		}
		for _, channel := range channels {
			channel, err = sess.Channel(channel.ID)
			if err != nil {
				return err
			}
			if channel.Type == discordgo.ChannelTypeGuildText {
				channelid = channel.ID
				return nil
			}
		}
		return errors.New("No primary channel found")
	})
	return channelid
}

func SendMessage(session *discordgo.Session, message string) {
	channelID := FetchPrimaryTextChannelID(session)
	session.ChannelTyping(channelID)
	LogInfo("SENDING MESSAGE:", message)
	RetryOnBadGateway(func() error {
		_, err := session.ChannelMessageSend(channelID, message)
		return err
	})
}

func SendMessageEmbed(session *discordgo.Session, messageEmbed *discordgo.MessageEmbed) {
	channelID := FetchPrimaryTextChannelID(session)
	// LogInfo("SENDING MESSAGE:", messageEmbed)
	RetryOnBadGateway(func() error {
		_, err := session.ChannelMessageSendEmbed(channelID, messageEmbed)
		return err
	})
}

func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// HTML strips html tags, replace common entities, and escapes <>&;'" in the result.
// Note the returned text may contain entities as it is escaped by HTMLEscapeString, and most entities are not translated.
// Original code from: https://github.com/kennygrant/sanitize
func StripHTML(s string) (output string) {

	// Shortcut strings with no tags in them
	if !strings.ContainsAny(s, "<>") {
		output = s
	} else {

		// First remove line breaks etc as these have no meaning outside html tags (except pre)
		// this means pre sections will lose formatting... but will result in less unintentional paras.
		//s = strings.Replace(s, "\n", "", -1)

		// Then replace line breaks with newlines, to preserve that formatting
		s = strings.Replace(s, "</p>", "\n", -1)
		s = strings.Replace(s, "<br>", "\n", -1)
		s = strings.Replace(s, "</br>", "\n", -1)
		s = strings.Replace(s, "<br/>", "\n", -1)
		s = strings.Replace(s, "<br />", "\n", -1)

		// Walk through the string removing all tags
		b := bytes.NewBufferString("")
		inTag := false
		for _, r := range s {
			switch r {
			case '<':
				inTag = true
			case '>':
				inTag = false
			default:
				if !inTag {
					b.WriteRune(r)
				}
			}
		}
		output = b.String()
	}

	// Remove a few common harmless entities, to arrive at something more like plain text
	output = strings.Replace(output, "&#8216;", "'", -1)
	output = strings.Replace(output, "&#8217;", "'", -1)
	output = strings.Replace(output, "&#8220;", "\"", -1)
	output = strings.Replace(output, "&#8221;", "\"", -1)
	output = strings.Replace(output, "&nbsp;", " ", -1)
	output = strings.Replace(output, "&quot;", "\"", -1)
	output = strings.Replace(output, "&apos;", "'", -1)

	// Translate some entities into their plain text equivalent (for example accents, if encoded as entities)
	output = html.UnescapeString(output)

	// In case we have missed any tags above, escape the text - removes <, >, &, ' and ".
	output = template.HTMLEscapeString(output)

	// After processing, remove some harmless entities &, ' and " which are encoded by HTMLEscapeString
	output = strings.Replace(output, "&#34;", "\"", -1)
	output = strings.Replace(output, "&#39;", "'", -1)
	output = strings.Replace(output, "&amp; ", "& ", -1)     // NB space after
	output = strings.Replace(output, "&amp;amp; ", "& ", -1) // NB space after

	return output
}