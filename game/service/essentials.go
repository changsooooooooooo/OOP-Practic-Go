package service

import (
	"OOP-Practice-GO/game/DTO"
	"container/list"
	"fmt"
	"strings"
)

type Input interface {
	GetUserInputGameTurn()
	GetUserInputCandidate() error
}

type Output interface {
	RemainGameTurn() int
	TopRankingCandidates() Result
}

type Cars struct {
	CarList *list.List
}

type Turn struct {
	Turns int
}

type Result struct {
	ResultList *list.List
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
		c.CarList.PushBack(car)
	}
	return nil
}

func (c *Cars) TopRankingCandidates() *Result {
	return &Result{}
}

func (t *Turn) GetUserInputGameTurn() {
	var inputTurn int
	_, _ = fmt.Scanln(&inputTurn)
	t.Turns = inputTurn
}

func (t *Turn) RemainGameTurn() int {
	t.Turns -= 1
	return t.Turns
}
