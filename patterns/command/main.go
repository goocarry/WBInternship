package main

import (
	"fmt"

	"github.com/goocarry/wb-internship/patterns/command/command"
)

func main() {
	fmt.Println("command example")


	tv := &command.Tv{}
	on := &command.OnCommand{
		Device: tv,
	}
	off := &command.OffCommand{
		Device: tv,
	}

	onButton := &command.Button{Command: on} 
	offButton := &command.Button{Command: off} 

	onButton.Press()
	offButton.Press()
}