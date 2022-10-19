package main

import (
	"fmt"

	"github.com/goocarry/wb-internship/patterns/state/player"
)

func main() {
	fmt.Println("state example")

	songs := []string{"1.mp3", "2.mp3", "3.mp3"}
	player := player.NewPlayer(songs)

	player.ClickPlay()
	player.ClickLock()
	player.ClickNext()
}