package anilist

const SEARCH_MEDIA_QUERY = `
query($search: String, $type: MediaType) {
  Media(search: $search, type: $type) {
    id
    idMal
    type
    title {
      english
      userPreferred
    }
    nextAiringEpisode {
      airingAt
    }
    status
    meanScore
    episodes
    chapters
	siteUrl
    externalLinks {
      site
      url
    }
    coverImage {
      medium
    }
	bannerImage
    description
  }
}
`

const SEARCH_USER_QUERY = `
query ($id: Int, $search: String) {
  User(id: $id, search: $search) {
    id
    name
    siteUrl
    avatar {
      large
    }
    about(asHtml: true)
    stats {
      watchedTime
      chaptersRead
    }
    favourites {
      manga {
        nodes {
          id
		  siteUrl
          title {
            romaji
            english
            native
            userPreferred
          }
        }
      }
      characters {
        nodes {
          id
		  siteUrl
          name {
            first
            last
            native
          }
        }
      }
      anime {
        nodes {
          id
		  siteUrl
          title {
            romaji
            english
            native
            userPreferred
          }
        }
      }
    }
  }
}
`

const SEARCH_CHARACTER_QUERY = `
query ($id: Int, $search: String) {
  Character(id: $id, search: $search) {
    id
	siteUrl
    description(asHtml: true)
    name {
      first
      last
      native
    }
    image {
      large
    }
    media {
      nodes {
        id
        type
		siteUrl
        title {
          romaji
          english
          native
          userPreferred
        }
      }
    }
  }
}
`