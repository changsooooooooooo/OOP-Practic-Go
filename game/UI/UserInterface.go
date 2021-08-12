package UI

import (
	"OOP-Practice-GO/game/service"
	"fmt"
	"strings"
)

type ShowInterface interface {
	GetInputs() error
	ShowStatus()
	ShowResult()
	ChangeStatusType(rPos int) string
}

type UI struct {
	Game *service.Game
}

func InitUI() *UI {
	ui := &UI{}
	ui.Game = service.InitGame()
	return ui
}

func (ui *UI) GetInputs() error {
	err := ui.Game.ValidateInputCheck()
	if err != nil {
		return err
	}
	return nil
}

func (ui *UI) ShowStatus() {
	status := ui.Game.Candidates
	for _, v := range status.CarList {
		tempStatus := ui.ChangeStatusType(v.Position)
		fmt.Printf("%s : %s\n", v.Name, tempStatus)
	}
	fmt.Println()
}

func (ui *UI) ShowResult() {
	ui.Game.ReturnResult()
	result := ui.Game.Winner
	length := len(result.ResultList)
	rs := strings.Builder{}
	for i, v := range result.ResultList {
		rs.WriteString(v.Name)
		if i < length-1 {
			rs.WriteString(",")
		}
	}
	fmt.Printf("%s가 우승하였습니다.\n", rs.String())
}

func (ui *UI) ChangeStatusType(rPos int) string {
	sb := strings.Builder{}
	for i := 0; i < rPos; i++ {
		sb.WriteString("_")
	}
	return sb.String()
}

func (ui *UI) PresentGameStatus() {
	for ui.Game.IsFinish != true {
		ui.Game.DoGame()
		ui.ShowStatus()
	}
	ui.ShowResult()
}
