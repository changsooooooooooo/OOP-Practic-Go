package game

import (
	"fmt"
	"testing"
)

func TestInputGo(t *testing.T) {
	testSlice := make([]string, 0)
	testSlice = append(testSlice, "123")
	testSlice = append(testSlice, "123")
	testSlice = append(testSlice, "123")
	testSlice = append(testSlice, "123")
	testSlice = append(testSlice, "46")

	for _, v := range testSlice{
		fmt.Println(v)
	}
}

func TestMakeCandidate(t *testing.T) {
	name:="thomasss"
	if 5<len(name){
		err := fmt.Errorf("Name(%s) length must be under 5!", name)
		panic(err)
	}
}
