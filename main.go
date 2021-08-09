package main

import (
	"OOP-Practice-GO/game/UI"
	"fmt"
)

func main() {
	ui:=&UI.UI{}
	err:=ui.GetInputs()
	if err!=nil{
		fmt.Println(err)
		return
	}
	ui.PresentGameStatus()
}
