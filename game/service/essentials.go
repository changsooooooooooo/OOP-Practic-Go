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

type Process interface {
	initCandidates(int)
	ReflectControllerNumber()
}

type Output interface {
	RemainGameTurn()
	TopRankingCandidates()
	topRankingScore() int
}

type ControllerAction interface {
	rollController(c *Cars)
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

func (c *Cars) initCandidates(length int) {
	c.CarList = make([]*DTO.Car, length)
	for i := 0; i < length; i++ {
		c.CarList[i] = &DTO.Car{}
	}
}

func (c *Cars) GetUserInputCandidate() error {
	var inputSeries string
	_, _ = fmt.Scanln(&inputSeries)
	candidates := strings.Split(inputSeries, ",")
	c.initCandidates(len(candidates))
	for i, v := range candidates {
		err := c.CarList[i].MakeCandidate(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Cars) topRankingScore() int {
	topScore := 0
	for _, v := range c.CarList {
		if topScore < v.Position {
			topScore = v.Position
		}
	}
	return topScore
}

//값을 추후에 아래에서 복사해줘야 하는 단점이 있지만, 한 번만 계산하면 된다는 장점이 있음
//즉, 늘 꾸준히 struct 에서 topScore 를 관리할 필요는 없지

func (c *Cars) ReflectControllerNumber(cl *Controller) {
	cl.rollController(c)
	for i, v := range cl.ControllerList {
		c.CarList[i].WhetherMoveOrNot(v)
	}
}

func (r *Result) TopRankingCandidates(c *Cars) {
	topScore := c.topRankingScore()
	r.ResultList = make([]*DTO.Car, 0)
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

func (cl *Controller) rollController(c *Cars) {
	length := len(c.CarList)
	cl.ControllerList = make([]int, length)
	for i := 0; i < length; i++ {
		randNum := rand.Intn(10)
		cl.ControllerList[i] = randNum
	}
}
