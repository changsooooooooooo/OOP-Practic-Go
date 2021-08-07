package game

import (
	"container/list"
	"fmt"
	"strings"
)

type Input interface {
	GetUserInputGameTurn() int
	GetUserInputCandidate() *list.List
}

type DTO interface{
	MakeCandidate(string) interface{}
}

type Car struct{
	position int
	name string
}

type Cars struct{
	carList *list.List
}

type Turn struct{
	turns int
}

func (c *Car) MakeCandidate(candidate string) (*Car, error) {

	if 5<len(candidate){
		err := fmt.Errorf("Name(%s) length must be under 5!", candidate)
		return &Car{}, err
	}

	car := &Car{0, candidate}
	return car, nil
}

func (c *Cars) GetUserInputCandidate() *list.List{
	var inputSeries string
	_, _ = fmt.Scanln(&inputSeries)
	candidates := strings.Split(inputSeries, ",")

	car := &Car{}
	carList := list.New()

	for _, v := range candidates{
		candidate, err:=car.MakeCandidate(v)
		if err != nil{
			return list.New()
		}
		carList.PushBack(candidate)
	}
	return carList
}

func (gt *Turn) GetUserInputGameTurn() int{
	var inputTurn int
	_, _ = fmt.Scanln(&inputTurn)
	return inputTurn
}
