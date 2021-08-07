package game

import (
	"container/list"
	"fmt"
	"testing"
)

func TestInitialStructure(t *testing.T) {
	cars:=&Cars{list.New()}
	fmt.Println(cars.carList)
	car:=&Car{}
	fmt.Println(car.position)
	fmt.Println(car.name)
}
