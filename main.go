package main

import (
	"github.com/youssefhabri/z2bot/discord"
)

// Parameters from flag.
func main() {
	discord.Init()

	<-make(chan struct{})
}
