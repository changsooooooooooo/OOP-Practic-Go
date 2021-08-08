package DTO

import (
	"fmt"
)

type DTO interface {
	MakeCandidate(string) error
	WhetherMoveOrNot(int)
}

type Car struct {
	Position int
	Name     string
}

func (c *Car) MakeCandidate(candidate string) error {
	if 5 < len(candidate) {
		err := fmt.Errorf("Name(%s) length must be under 5!", candidate)
		return err
	}
	c = &Car{0, candidate}
	return nil
}

func (c *Car) WhetherMoveOrNot(diceNum int) {
	if 3 < diceNum {
		c.Position += 1
	}
}
