package game

import (
	"container/list"
	"fmt"
	"strings"
)

type Input interface {
	GetUserInputGameTurn()
	GetUserInputCandidate() error
}

type Cars struct{
	carList *list.List
}

type Turn struct{
	turns int
}

func (c *Cars) GetUserInputCandidate() error{
	var inputSeries string
	_, _ = fmt.Scanln(&inputSeries)
	candidates := strings.Split(inputSeries, ",")

	c.carList = list.New()

	for _, v := range candidates{
		car := &Car{}
		err:=car.MakeCandidate(v)
		if err != nil{
			return err
		}
		c.carList.PushBack(car)
	}
	return nil
}

func (t *Turn) GetUserInputGameTurn(){
	var inputTurn int
	_, _ = fmt.Scanln(&inputTurn)
	t.turns=inputTurn
}

