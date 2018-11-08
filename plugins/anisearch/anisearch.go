package anisearch

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/youssefhabri/z2bot/plugins/anisearch/anilist"
	"github.com/youssefhabri/z2bot/utils"
)

func Register(session *discordgo.Session) {
	session.AddHandler(anisearch)
}

func anisearch(session *discordgo.Session, evt *discordgo.MessageCreate) {
	params := strings.Split(evt.Message.Content, " ")
	channelID := evt.ChannelID

	switch strings.ToLower(strings.TrimSpace(params[0])) {
	case utils.PREFIX + "anime":
		session.ChannelTyping(channelID)
		if len(params) > 1 {
			searchAnime(session, channelID, params[1:])
		} else {
			utils.SendMessage(session,  channelID,"`usage: "+utils.PREFIX+"anime <anime_name>`")
		}
		break
	case utils.PREFIX + "manga":
		session.ChannelTyping(channelID)
		if len(params) > 1 {
			searchManga(session, channelID, params[1:])
		} else {
			utils.SendMessage(session, channelID, "`usage: "+utils.PREFIX+"manga <manga_name>`")
		}
		break
	case utils.PREFIX + "user":
		session.ChannelTyping(channelID)
		if len(params) > 1 {
			searchUser(session, channelID, params[1:])
		} else {
			utils.SendMessage(session, channelID, "`usage: "+utils.PREFIX+"user <username>`")
		}
		break
	case utils.PREFIX + "character":
		session.ChannelTyping(channelID)
		if len(params) > 1 {
			searchCharacter(session, channelID, params[1:])
		} else {
			utils.SendMessage(session, channelID, "`usage: "+utils.PREFIX+"character <character_name>`")
		}
		break
	}
}

func searchAnime(session *discordgo.Session, channelID string, query []string) {
	var keyword = strings.Join(query, " ")
	anime, _ := anilist.SearchMedia(keyword, anilist.ANIME_T)
	if anime.ID != 0 {
		messageEmbed := utils.NewEmbed().
			SetColor(3447003).
			SetTitle(anime.Title.UserPreferred).
			SetURL(anime.SiteUrl).
			SetDescription(anime.Synopses(350)).
			SetImage(anime.BannerImage).
			SetThumbnail(anime.CoverImage.Medium).
			SetFooter("Status: " + anime.Status + ", Next episode: " + anime.NextEpisode() + " | Powered by AniList").
			AddField("Score", fmt.Sprintf("%d", anime.MeanScore)).
			AddField("Episodes", fmt.Sprintf("%d", anime.Episodes)).
			AddField("Streaming Services", anime.StreamingServices()).
			AddField("More info", anime.TrackingSites()).
			InlineAllFields().MessageEmbed

		utils.SendMessageEmbed(session, channelID, messageEmbed)
	} else {
		utils.SendMessage(session, channelID, "No anime was found or there was an error in the process")
	}
}

func searchManga(sess *discordgo.Session, channelID string, query []string) {
	var keyword = strings.Join(query, " ")
	manga, _ := anilist.SearchMedia(keyword, anilist.MANGA_T)
	if manga.ID != 0 {
		messageEmbed := utils.NewEmbed().
			SetColor(3447003).
			SetTitle(manga.Title.UserPreferred).
			SetURL(manga.SiteUrl).
			SetDescription(manga.Synopses(350)).
			SetThumbnail(manga.CoverImage.Medium).
			SetFooter("Status: " + manga.Status + " | Powered by AniList").
			AddField("Score", fmt.Sprintf("%d", manga.MeanScore)).
			AddField("Chapters", fmt.Sprintf("%d", manga.Chapters)).
			AddField("More info", manga.TrackingSites()).
			InlineAllFields().MessageEmbed

		utils.SendMessageEmbed(sess, channelID, messageEmbed)
	} else {
		utils.SendMessage(sess, channelID, "No manga was found or there was an error in the process")
	}
}

func searchUser(sess *discordgo.Session, channelID string, query []string) {
	var username = strings.Join(query, " ")
	username = strings.TrimSpace(username)

	user, _ := anilist.SearchUser(username)
	if user.ID != 0 {
		messageEmbed := utils.NewEmbed().
			SetColor(3447003).
			SetTitle(user.Name).
			SetURL(user.SiteUrl).
			SetDescription(user.AboutText(350)).
			SetThumbnail(user.Avatar.Large).
			SetFooter("Powered by AniList").
			AddField("Watched time", user.WatchedTime()).
			AddField("Chapters read", user.ChaptersRead()).
			AddField("Favorite Anime", user.GetFavoriteAnime()).
			AddField("Favorite Manga", user.GetFavoriteAnime()).
			AddField("Favorite Characters", user.GetFavoriteCharacters()).
			InlineAllFields().MessageEmbed

		utils.SendMessageEmbed(sess, channelID, messageEmbed)
	} else {
		utils.SendMessage(sess, channelID, "No user was found or there was an error in the process")
	}
}

func searchCharacter(session *discordgo.Session, channelID string, query []string) {
	var characterName = strings.Join(query, " ")
	characterName = strings.TrimSpace(characterName)

	character, _ := anilist.SearchCharacter(characterName)
	if character.ID != 0 {
		messageEmbed := utils.NewEmbed().
			SetColor(3447003).
			SetTitle(character.Name.First + " " + character.Name.Last).
			SetURL(character.SiteUrl).
			SetDescription(character.About(350)).
			SetThumbnail(character.Image.Large).
			SetFooter("Powered by AniList").
			AddField("Anime", character.GetMediaList(anilist.ANIME_T)).
			AddField("Manga", character.GetMediaList(anilist.MANGA_T)).
			InlineAllFields().MessageEmbed

		session.ChannelMessageSendEmbed(channelID, messageEmbed)
	} else {
		utils.SendMessage(session, channelID, "No character was found or there was an error in the process")
	}
}
