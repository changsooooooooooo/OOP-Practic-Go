package service

import (
	"OOP-Practice-GO/game/DTO"
	"fmt"
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
