package service

import (
	"OOP-Practice-GO/game/DTO"
	"fmt"
	"math/rand"
	"testing"
)

func TestInitialStructure(t *testing.T) {
	cars := &Cars{}
	fmt.Println(cars.CarList)
	car := &DTO.Car{}
	fmt.Println(car.Position)
	fmt.Println(car.Name)
}

func TestListValueCheck(t *testing.T) {
	cars := &Cars{}
	cars.CarList = append(cars.CarList, &DTO.Car{0, "Phobi"})
	for _, v := range cars.CarList {
		fmt.Println(v)
	}
}

func rollController(cl *Controller) {
	length := 3
	cl.ControllerList = make([]int, length)
	//tdd 필요
	for i := 0; i < length; i++ {
		randNum := rand.Intn(10)
		fmt.Print(randNum)
		cl.ControllerList[i] = randNum
	}
}

func TestControllerPointer(t *testing.T) {
	cl := &Controller{}
	for i := 0; i < 3; i++ {
		rollController(cl)
		fmt.Println(cl.ControllerList)
	}
}
