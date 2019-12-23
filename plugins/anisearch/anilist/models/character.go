package models

import (
	"strings"

	"github.com/youssefhabri/zero2-go/utils"
)

type Character struct {
	ID          int    `json:"id"`
	SiteUrl     string `json:"siteUrl"`
	Description string `json:"description"`
	Name        struct {
		First  string `json:"first"`
		Last   string `json:"last"`
		Native string `json:"native"`
	} `json:"name"`
	Image struct {
		Large string `json:"large"`
	} `json:"image"`
	Media struct {
		Nodes []struct {
			ID      int    `json:"id"`
			Type    string `json:"type"`
			SiteUrl string `json:"siteUrl"`
			Title   struct {
				Romaji        string `json:"romaji"`
				English       string `json:"english"`
				Native        string `json:"native"`
				UserPreferred string `json:"userPreferred"`
			} `json:"title"`
		} `json:"nodes"`
	} `json:"media"`
}

func (c *Character) About(params ...int) string {
	textLength := 400
	if len(params) > 0 {
		textLength = params[0]
	}
	text := utils.StripHTML(c.Description)
	text = strings.Replace(text, "\n\n", "\n", -1)
	text = strings.Join(strings.Split(text, "\n"), "\n")
	if len(text) > textLength {
		return text[:textLength] + " ..."
	}
	return text
}

func (c *Character) GetMediaList(mType string) string {
	var list string
	//num := len(c.Media.Nodes)
	count := 0
	for _, media := range c.Media.Nodes {
		if media.Type == mType {
			list = list + "[" + media.Title.UserPreferred + "](" + media.SiteUrl + ")\n"
			count++
		}
		if count >= 5 {
			break
		}
	}
	//if num > 5 {
	//	list = list + fmt.Sprintf("+ %d more", len(c.Media.Nodes)-5)
	//}
	return list
}
