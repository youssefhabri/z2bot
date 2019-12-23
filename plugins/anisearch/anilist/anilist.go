package anilist

import (
	"github.com/animenotifier/anilist"
	"github.com/youssefhabri/zero2-go/plugins/anisearch/anilist/models"
)

const (
	ANIME_T string = "ANIME"
	MANGA_T string = "MANGA"
)

type Variables struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}

func SearchMedia(keyword string, mediaType string) (models.Media, error) {
	type Variables struct {
		Search string `json:"search"`
		Type   string `json:"type"`
	}

	body := struct {
		Query     string    `json:"query"`
		Variables Variables `json:"variables"`
	}{
		Query: SEARCH_MEDIA_QUERY,
		Variables: Variables{
			Search: keyword,
			Type:   mediaType,
		},
	}

	// Query response
	response := new(struct {
		Data struct {
			Media models.Media `json:"Media"`
		} `json:"data"`
	})

	err := anilist.Query(body, response)

	if err != nil {
		return models.Media{}, err
	}

	return response.Data.Media, nil
}

func SearchUser(username string) (models.User, error) {
	type Variables struct {
		Search string `json:"search"`
	}

	body := struct {
		Query     string    `json:"query"`
		Variables Variables `json:"variables"`
	}{
		Query: SEARCH_USER_QUERY,
		Variables: Variables{
			Search: username,
		},
	}

	// Query response
	response := new(struct {
		Data struct {
			User models.User `json:"User"`
		} `json:"data"`
	})

	err := anilist.Query(body, response)
	if err != nil {
		return models.User{}, err
	}

	return response.Data.User, nil
}

func SearchCharacter(character string) (models.Character, error) {
	type Variables struct {
		Search string `json:"search"`
	}

	body := struct {
		Query     string    `json:"query"`
		Variables Variables `json:"variables"`
	}{
		Query: SEARCH_CHARACTER_QUERY,
		Variables: Variables{
			Search: character,
		},
	}

	// Query response
	response := new(struct {
		Data struct {
			Character models.Character `json:"Character"`
		} `json:"data"`
	})

	err := anilist.Query(body, response)
	if err != nil {
		return models.Character{}, err
	}

	return response.Data.Character, nil
}
