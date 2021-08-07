package game

import (
	"fmt"
)

type DTO interface{
	MakeCandidate(string) interface{}
}

type Car struct{
	position int
	name string
}

func (c *Car) MakeCandidate(candidate string) error {

	if 5<len(candidate){
		err := fmt.Errorf("Name(%s) length must be under 5!", candidate)
		return err
	}

	c = &Car{0, candidate}
	return nil
}

