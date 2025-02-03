package main

import (
	"fmt"
	"time"
)

func life(cells []int) []int {
	newCells := make([]int, len(cells))
	/* Serpinski Triangle Rules: // 1,1,1,0 0,1,1,1
	if adjacent to a 1 it becomes alive
	If in between two 1s it dies */
	// 1,1,1,0 0,0,1,1 is the turing complete one

	// USER: update the 1 to a 0 or 0 to 1 for each case: newCells[i] = 1 to newCells[i] = 0
	newCells[0] = 0
	newCells[len(cells)-1] = 0
	for i := 1; i < len(cells)-1; i++ {
		switch {
		case (cells[i] == 1 && cells[i-1] == 0 && cells[i+1] == 0): // 010
			newCells[i] = 1 // Update these!
		case (cells[i] == 1 && cells[i-1] == 1 && cells[i+1] == 0): // 110
			newCells[i] = 1
		case (cells[i] == 1 && cells[i-1] == 0 && cells[i+1] == 1): // 011
			newCells[i] = 1
		case (cells[i] == 1 && cells[i-1] == 1 && cells[i+1] == 1): // 111
			newCells[i] = 0

		case (cells[i] == 0 && cells[i-1] == 0 && cells[i+1] == 0): // 000
			newCells[i] = 0
		case (cells[i] == 0 && cells[i-1] == 1 && cells[i+1] == 0): // 100
			newCells[i] = 1 // switch this from 1 to 0 for turing complete/serpinski triangle
		case (cells[i] == 0 && cells[i-1] == 0 && cells[i+1] == 1): // 001
			newCells[i] = 1
		case (cells[i] == 0 && cells[i-1] == 1 && cells[i+1] == 1): // 101
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
	time.Sleep(50 * time.Millisecond) //50 is good, pause between lines
	fmt.Print("\n")
}

func main() {
	// USER: update the length of the slice 'cells' below to another odd number
	cells := make([]int, 151)
	cells[len(cells)/2+1] = 1

	printCells(cells)

	newIteration := cells
	// USER: update the number of lines to be printed out in this for loop.
	// If you want to print 333 lines, change to: i < 333
	for i := 0; i < 201; i++ {
		newIteration = life(newIteration)
		printCells(newIteration)
	}

}
