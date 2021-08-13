package UI

import (
	"OOP-Practice-GO/game/service"
	"fmt"
	"testing"
)

func TestInitUI(t *testing.T) {
	ui := &UI{}
	ui.Game = service.InitGame()
	fmt.Println(ui.Game)
}
