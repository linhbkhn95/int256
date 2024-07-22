package main

import (
	"fmt"

	"github.com/linhbkhn95/int256"
)

type PoolState struct {
	Price *int256.Int
	Tick  int64
}

func main() {
	state1 := &PoolState{
		Price: fromString("100000000000000000000000"),
		Tick:  1000,
	}
	state2 := &PoolState{
		Price: fromString("-100000000000000000000000"),
		Tick:  1000,
	}

	fmt.Println("state1", state1)
	fmt.Println("state2", state2)

}

func fromString(str string) *int256.Int {
	i, _ := new(int256.Int).SetString(str)
	return i
}
