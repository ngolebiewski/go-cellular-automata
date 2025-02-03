package main

import (
	"flag"
	"fmt"
	"time"
)

func getArgs() (int, int, int) {
	rule := flag.Int("rule", 126, "Elementary Automata Rule Number (Integer 0-255)")
	width := flag.Int("width", 151, "Width of design, Integer, i.e. 151")
	lines := flag.Int("lines", 200, "Number of lines/rows, Integer, i.e. 200")
	flag.Parse()
	// fmt.Println(*rule, *width, *lines) // debug
	return *rule, *width, *lines
}

func ruleNumToBinary(ruleNum int) map[string]int {
	ruleNum = ruleNum % 256                  // keep rule int within 8bit range of 0-255
	rule8bit := fmt.Sprintf("%08b", ruleNum) // converts to 8bit binary
	fmt.Println(rule8bit)
	letters := "abcdefgh"
	ruleMap := make(map[string]int)
	for i := 0; i < 8; i++ {
		ruleMap[string(letters[i])] = int(rule8bit[i] - '0')
	}
	return ruleMap
}

func life(cells []int, ruleMap map[string]int) []int {
	newCells := make([]int, len(cells))
	/* Serpinski Triangle Rules: // 1,1,1,0 0,1,1,1
	if adjacent to a 1 it becomes alive
	If in between two 1s it dies */
	// 1,1,1,0 0,0,1,1 is the turing complete one

	newCells[0] = 0
	newCells[len(cells)-1] = 0
	for i := 1; i < len(cells)-1; i++ {
		switch {
		case (cells[i] == 1 && cells[i-1] == 0 && cells[i+1] == 0): // 010
			newCells[i] = ruleMap["f"] // f
		case (cells[i] == 1 && cells[i-1] == 1 && cells[i+1] == 0): // 110
			newCells[i] = ruleMap["b"] // b
		case (cells[i] == 1 && cells[i-1] == 0 && cells[i+1] == 1): // 011
			newCells[i] = ruleMap["e"] // e
		case (cells[i] == 1 && cells[i-1] == 1 && cells[i+1] == 1): // 111
			newCells[i] = ruleMap["a"] // a

		case (cells[i] == 0 && cells[i-1] == 0 && cells[i+1] == 0): // 000
			newCells[i] = ruleMap["h"] // h
		case (cells[i] == 0 && cells[i-1] == 1 && cells[i+1] == 0): // 100
			newCells[i] = ruleMap["d"] // d
		case (cells[i] == 0 && cells[i-1] == 0 && cells[i+1] == 1): // 001
			newCells[i] = ruleMap["g"] // g
		case (cells[i] == 0 && cells[i-1] == 1 && cells[i+1] == 1): // 101
			newCells[i] = ruleMap["c"] // c
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
	rule, width, lines := getArgs()
	ruleMap := ruleNumToBinary(rule)

	//debug
	/*for key, value := range ruleMap {
		fmt.Printf("%s: %d\n", key, value)
	}*/

	// update the length of the slice 'cells' below to another odd number
	cells := make([]int, width)
	cells[len(cells)/2+1] = 1

	printCells(cells)

	newIteration := cells
	// update the number of lines to be printed out in this for loop.
	// If you want to print 333 lines, change to: i < 333
	for i := 0; i < lines; i++ {
		newIteration = life(newIteration, ruleMap)
		printCells(newIteration)
	}

}
