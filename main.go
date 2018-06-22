package main

import (
	"github.com/youssefhabri/anitrend-bot/discord"
)

// Parameters from flag.
func main() {
	discord.Init()

	<-make(chan struct{})
}
