package DTO

import (
	"fmt"
)

type DTO interface {
	initDTO(string)
	MakeCandidate(string) error
	WhetherMoveOrNot(int)
}

type Car struct {
	Position int
	Name     string
}

func (c *Car) initDTO(candidate string) {
	c.Position = 0
	c.Name = candidate
}

func (c *Car) MakeCandidate(candidate string) error {
	if 5 < len(candidate) {
		err := fmt.Errorf("Name(%s) length must be under 5!", candidate)
		return err
	}
	c.initDTO(candidate)
	return nil
}

func (c *Car) WhetherMoveOrNot(diceNum int) {
	if 3 < diceNum {
		c.Position += 1
	}
}
