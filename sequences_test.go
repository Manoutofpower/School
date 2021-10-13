package main

import (
	"fmt"
	"testing"
)


type cell struct {
	x, y int
}

func calculateAliveCells(world [][]byte) []cell {
	aliveCells := []cell{}
	for i, v := range world {
		for j, r := range v {
			if r == 255 {
				aliveCells = append(aliveCells, cell{i, j})
			}
		}
	}
	return aliveCells
}


func TestRange(t *testing.T) {

	v := [][]byte{
		{255, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 255, 4},
		{1, 2, 3, 255},
	}

	aliveCells := calculateAliveCells(v)

	if len(aliveCells) != 3 {
		t.Fatal("not right number")
	}

	fmt.Println(aliveCells);

}
