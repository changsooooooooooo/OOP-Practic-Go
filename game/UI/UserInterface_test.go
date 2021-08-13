package UI

import (
	"OOP-Practice-GO/game/service"
	"fmt"
	"strings"
	"testing"
)

func TestInitUI(t *testing.T) {
	ui := &UI{}
	ui.Game = service.InitGame()
	fmt.Printf("%p", ui.Game)
}

func TestStringBuilderWriteString(t *testing.T) {
	rs := strings.Builder{}
	rs.WriteString("Hello")
	fmt.Printf("%x\n", rs.String())
	rs.WriteString("HE")
	fmt.Printf("%x\n", rs.String())
}

func TestStringBuilderWriteRune(t *testing.T) {
	rs := strings.Builder{}
	rs.WriteRune('H')
	fmt.Printf("%x\n", rs.String())
	rs.WriteRune('E')
	fmt.Printf("%x\n", rs.String())
}
