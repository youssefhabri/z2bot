package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/youssefhabri/zero2-go/utils"
)

//id
//idMal
//title { english }
//nextAiringEpisode { airingAt }
//status
//meanScore
//episodes
//externalLinks { site, url }
//coverImage { medium }
//description

type Media struct {
	ID    int `json:"id"`
	IDMal int `json:"idMal"`
	Title struct {
		English       string `json:"english"`
		UserPreferred string `json:"userPreferred"`
	} `json:"title"`
	NextAiringEpisode struct {
		AiringAt        int64 `json:"airingAt"`
		TimeUntilAiring int64 `json:"timeUntilAiring"`
	} `json:"nextAiringEpisode"`
	Status        string `json:"status"`
	MeanScore     int    `json:"meanScore"`
	Episodes      int    `json:"episodes"`
	Chapters      int    `json:"chapters"`
	SiteUrl       string `json:"siteUrl"`
	ExternalLinks []struct {
		Site string `json:"site"`
		Url  string `json:"url"`
	} `json:"externalLinks"`
	CoverImage struct {
		Medium string `json:"medium"`
	} `json:"coverImage"`
	BannerImage string `json:"bannerImage"`
	Description string `json:"description"`
}

func (m *Media) Synopses(params ...int) string {
	textLength := 400
	if len(params) > 0 {
		textLength = params[0]
	}
	text := utils.StripHTML(m.Description)
	text = strings.Replace(text, "\n\n", "\n", -1)
	segs := strings.Split(text, "\n")
	if len(segs) > 5 {
		segs = strings.Split(text, "\n")[:5]
	}
	text = strings.Join(segs, "\n")
	if len(text) > textLength {
		return text[:textLength] + " ..."
	}
	return text
}

func (m *Media) StreamingServices() string {
	if len(m.ExternalLinks) > 0 {
		var list []string
		for _, service := range m.ExternalLinks {
			list = append(list, "["+service.Site+"]("+service.Url+")")
		}
		return strings.Join(list, ", ")
	}
	return "Not available."
}

func (m *Media) TrackingSites() string {
	var list []string
	list = append(list, fmt.Sprintf("[AniList](https://anilist.co/anime/%d)", m.ID))
	list = append(list, fmt.Sprintf("[MAL](https://myanimelist.com/anime/%d)", m.IDMal))

	return strings.Join(list, ", ")
}

func (m *Media) NextEpisode() string {
	if m.NextAiringEpisode.AiringAt != 0 {
		return time.Unix(m.NextAiringEpisode.AiringAt, 0).Format("Jan _2 15:04 MST")
	}
	return "Ended"
}
