package service

import (
	"OOP-Practice-GO/game/DTO"
	"container/list"
	"fmt"
	"testing"
)

func TestInitialStructure(t *testing.T) {
	cars := &Cars{list.New()}
	fmt.Println(cars.CarList)
	car := &DTO.Car{}
	fmt.Println(car.Position)
	fmt.Println(car.Name)
}

func TestListValueCheck(t *testing.T) {
	cars := &Cars{list.New()}
	cars.CarList.PushBack(&DTO.Car{0, "phobi"})
	cars.CarList.PushBack(&DTO.Car{2, "elena"})
	cars.CarList.PushBack(&DTO.Car{3, "honux"})

}
