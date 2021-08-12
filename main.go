package main

import (
	"OOP-Practice-GO/game/UI"
	"fmt"
)

func main() {
	ui := &UI.UI{}
	ui.MakeUI()
	err := ui.GetInputs()
	if err != nil {
		fmt.Println(err)
		return
	}

	ui.PresentGameStatus()
}
