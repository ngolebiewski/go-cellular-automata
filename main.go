package main

import (
	"fmt"
	"time"
)

func life(cells []int) []int {
	newCells := make([]int, len(cells))
	// Rules, if adjacent to a 1 it becomes alive
	// If next to two 1s it dies
	newCells[0] = 0
	newCells[len(cells)-1] = 0
	for i := 1; i < len(cells)-1; i++ {
		switch {
		case (cells[i] == 1 && cells[i-1] == 0 && cells[i+1] == 0):
			newCells[i] = 1
		case (cells[i] == 1 && cells[i-1] == 1 && cells[i+1] == 0):
			newCells[i] = 1
		case (cells[i] == 1 && cells[i-1] == 0 && cells[i+1] == 1):
			newCells[i] = 1
		case (cells[i] == 1 && cells[i-1] == 1 && cells[i+1] == 1):
			newCells[i] = 0

		case (cells[i] == 0 && cells[i-1] == 0 && cells[i+1] == 0):
			newCells[i] = 0
		case (cells[i] == 0 && cells[i-1] == 1 && cells[i+1] == 0):
			newCells[i] = 1
		case (cells[i] == 0 && cells[i-1] == 0 && cells[i+1] == 1):
			newCells[i] = 1
		case (cells[i] == 0 && cells[i-1] == 1 && cells[i+1] == 1):
			newCells[i] = 1
		}
	}
	return newCells
}

func printCells(cells []int) {
	for _, val := range cells {
		switch val {
		case 0:
			fmt.Print(" ")
		case 1:
			fmt.Print("#")
		default:
			fmt.Print(" ")
		}
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Print("\n")
}

func main() {
	cells := make([]int, 151)
	cells[len(cells)/2+1] = 1

	printCells(cells)

	newIteration := cells
	for i := 0; i < 101; i++ {
		newIteration = life(newIteration)
		printCells(newIteration)
	}

}
