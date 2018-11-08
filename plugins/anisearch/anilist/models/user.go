package models

import (
		"fmt"
	"github.com/youssefhabri/z2bot/utils"
	"strings"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	SiteUrl string `json:"siteUrl"`
	Avatar struct {
		Large string `json:"large"`
	} `json:"avatar"`
	About string `json:"about"`
	Stats struct {
		WatchedTime  int64 `json:"watchedTime"`
		ChaptersRead int   `json:"chaptersRead"`
	} `json:"stats"`
	Favourites struct {
		Manga struct {
			Nodes []struct {
				ID      int    `json:"id"`
				SiteUrl string `json:"siteUrl"`
				Title struct {
					Romaji        string `json:"romaji"`
					English       string `json:"english"`
					Native        string `json:"native"`
					UserPreferred string `json:"userPreferred"`
				} `json:"title"`
			} `json:"nodes"`
		} `json:"manga"`
		Characters struct {
			Nodes []struct {
				ID      int    `json:"id"`
				SiteUrl string `json:"siteUrl"`
				Name struct {
					First  string `json:"first"`
					Last   string `json:"last"`
					Native string `json:"native"`
				} `json:"name"`
			} `json:"nodes"`
		} `json:"characters"`
		Anime struct {
			Nodes []struct {
				ID      int    `json:"id"`
				SiteUrl string `json:"siteUrl"`
				Title struct {
					Romaji        string `json:"romaji"`
					English       string `json:"english"`
					Native        string `json:"native"`
					UserPreferred string `json:"userPreferred"`
				} `json:"title"`
			} `json:"nodes"`
		} `json:"anime"`
	} `json:"favourites"`
}

func (u *User) UsernameLink() string {
	return "[" + u.Name + "](https://anilist.co/user/" + u.Name + ")"
}

func (u *User) AboutText(params ...int) string {
	textLength := 400
	if len(params) > 0 {
		textLength = params[0]
	}
	text := utils.StripHTML(u.About)
	text = strings.Replace(text, "\n\n", "\n", -1)
	text = strings.Join(strings.Split(text, "\n"), "\n")
	if len(text) > textLength {
		return text[:textLength] + " ..."
	}
	return text
}

func (u *User) WatchedTime() string {
	minutes := (u.Stats.WatchedTime) % 60
	hours := (u.Stats.WatchedTime / 60) % 24
	days := u.Stats.WatchedTime / 60 / 24

	return fmt.Sprintf("%d days, %d:%.2d", days, hours, minutes)
}

func (u *User) ChaptersRead() string {
	return fmt.Sprintf("%d", u.Stats.ChaptersRead)
}

func (u *User) GetFavoriteAnime() string {
	var fav string
	num := len(u.Favourites.Anime.Nodes)
	for i := 0; i < utils.Min(num, 5); i++ {
		anime := u.Favourites.Anime.Nodes[i]
		fav = fav + "[" + anime.Title.UserPreferred + "](" + anime.SiteUrl + ")\n"
	}
	if num > 5 {
		fav = fav + fmt.Sprintf("+ %d more", len(u.Favourites.Anime.Nodes)-5)
	}
	return fav
}

func (u *User) GetFavoriteManga() string {
	var fav string
	num := len(u.Favourites.Manga.Nodes)
	for i := 0; i < utils.Min(num, 5); i++ {
		manga := u.Favourites.Manga.Nodes[i]
		fav = fav + "[" + manga.Title.UserPreferred + "](" + manga.SiteUrl + ")\n"
	}
	if num > 5 {
		fav = fav + fmt.Sprintf("+ %d more", len(u.Favourites.Manga.Nodes)-5)
	}
	return fav
}

func (u *User) GetFavoriteCharacters() string {
	var fav string
	num := len(u.Favourites.Characters.Nodes)
	for i := 0; i < utils.Min(num, 5); i++ {
		character := u.Favourites.Characters.Nodes[i]
		fav = fav + "[" + character.Name.First+ " "+ character.Name.Last +"](" + character.SiteUrl + ")\n"
	}
	if num > 5 {
		fav = fav + fmt.Sprintf("+ %d more", len(u.Favourites.Characters.Nodes)-5)
	}
	return fav
}
