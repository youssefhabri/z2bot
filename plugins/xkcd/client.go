package xkcd

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/youssefhabri/z2bot-go/utils"
)

type Comic struct {
	Number     int    `json:"num"`
	Title      string `json:"title"`
	SafeTitle  string `json:"safe_title"`
	ImgURL     string `json:"img"`
	ImgALT     string `json:"alt"`
	Transcript string `json:"transcript"`
	News       string `json:"news"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
}

func (c *Comic) GetTitle() string {
	return c.Title + " ("+ c.Day +"/"+ c.Month +"/"+ c.Year +")"
}

func (c *Comic) GetLink() string {
	if c.Link == "" {
		return "https://xkcd.com/" + strconv.Itoa(c.Number)
	}
	return c.Link
}

type Client struct {
	HTTPClient *http.Client
	UseHTTPS   bool
}

func NewClient() *Client {
	client := &Client{
		HTTPClient: http.DefaultClient,
		UseHTTPS:   true,
	}
	return client
}

func (c *Client) baseURL() string {
	protocol := "http://"
	if c.UseHTTPS {
		protocol = "https://"
	}
	return protocol + "xkcd.com"
}

func (c *Client) comicRequest(path string) (Comic, error) {
	var comic Comic

	res, err := c.HTTPClient.Get(c.baseURL() + path)
	if err != nil {
		utils.LogError(err)
		return comic, err
	}

	err = json.NewDecoder(res.Body).Decode(&comic)
	if err != nil {
		utils.LogError(err)
		return comic, err
	}

	return comic, nil
}

func (c *Client) LatestComic() Comic {
	comic, _ := c.comicRequest("/info.0.json")
	return comic
}

func (c *Client) RandomComic() Comic {
	numStr := strconv.Itoa(rand.Intn(c.LatestComic().Number))
	comic, _ := c.comicRequest("/" + numStr + "/info.0.json")
	return comic
}

func (c *Client) Comic(number string) Comic {
	comic, _ := c.comicRequest("/" + number + "/info.0.json")
	return comic
}
