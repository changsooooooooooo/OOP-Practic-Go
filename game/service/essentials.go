package service

import (
	"OOP-Practice-GO/game/DTO"
	"fmt"
	"math/rand"
	"strings"
)

type Input interface {
	GetUserInputGameTurn()
	GetUserInputCandidate() error
}

type Output interface {
	RemainGameTurn()
	TopRankingCandidates()
	TopRankingScore() int
}

type ControllerAction interface {
	RollController(c *Cars)
}

type Cars struct {
	CarList []*DTO.Car
}

type Turn struct {
	Turns int
}

type Controller struct {
	ControllerList []int
}

type Result struct {
	ResultList []*DTO.Car
}

func (c *Cars) GetUserInputCandidate() error {
	var inputSeries string
	_, _ = fmt.Scanln(&inputSeries)
	candidates := strings.Split(inputSeries, ",")
	for _, v := range candidates {
		car := &DTO.Car{}
		err := car.MakeCandidate(v)
		if err != nil {
			return err
		}
		c.CarList = append(c.CarList, car)
	}
	return nil
}

func (c *Cars) TopRankingScore() int {
	topScore := 0
	for _, v := range c.CarList {
		if topScore < v.Position {
			topScore = v.Position
		}
	}
	return topScore
}

func (r *Result) TopRankingCandidates(c *Cars) {
	topScore := c.TopRankingScore()
	for _, v := range c.CarList {
		if v.Position == topScore {
			r.ResultList = append(r.ResultList, v)
		}
	}
}

func (t *Turn) GetUserInputGameTurn() {
	var inputTurn int
	_, _ = fmt.Scanln(&inputTurn)
	t.Turns = inputTurn
}

func (t *Turn) RemainGameTurn() {
	t.Turns -= 1
}

func (cl *Controller) RollController(c *Cars) {
	length := len(c.CarList)
	cl.ControllerList = make([]int, length)
	for i := 0; i < length; i++ {
		randNum := rand.Intn(10)
		cl.ControllerList = append(cl.ControllerList, randNum)
	}
}
