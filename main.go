package main

import (
	"github.com/subosito/gotenv"
	"github.com/youssefhabri/zero2-go/discord"
)

// Parameters from flag.
func main() {
	gotenv.Load()

	discord.Init()

	<-make(chan struct{})
}
