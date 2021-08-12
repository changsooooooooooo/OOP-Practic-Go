package service

import "fmt"

type GameProcess interface {
	InitGame()
	DoGame()
	ReturnResult()
	ValidateInputCheck() error
}

type Game struct {
	Candidates *Cars
	GameTurn   *Turn
	Controller *Controller
	Winner     *Result
	IsFinish   bool
}

func (g *Game) InitGame() {
	g.Candidates = &Cars{}
	g.GameTurn = &Turn{}
	g.Controller = &Controller{}
	g.Winner = &Result{}
	g.IsFinish = false
}

func (g *Game) ValidateInputCheck() error {
	fmt.Println(1)
	err := g.Candidates.GetUserInputCandidate()
	if err != nil {
		return err
	}
	g.GameTurn.GetUserInputGameTurn()
	return nil
}

func (g *Game) DoGame() {
	cl := g.Controller
	for -1 < g.GameTurn.Turns {
		g.Candidates.ReflectControllerNumber(cl)
	}
	g.IsFinish = true
}

func (g *Game) ReturnResult() {
	cars := g.Candidates
	g.Winner.TopRankingCandidates(cars)
}
