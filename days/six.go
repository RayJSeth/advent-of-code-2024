package days

import (
	"bufio"
	"log"
	"os"
	"strings"

	"rayjseth.io/aoc-24/model"
)

type Heading int

const (
	North Heading = iota
	East
	South
	West
)

type Guard struct {
	Heading
	Row int
	Col int
}

type Cells [][]string

type FloorMap struct {
	Cells Cells
	Guard Guard
}

func Six() model.Result {
	day := uint8(6)
	fm := parseDay6Input("./inputs/day6")
	return model.Result{Day: &day, Part1: fm.calcPart1()}
}

func (fm *FloorMap) calcPart1() *int {
	// a journey of a thousand miles...
	steps := 1
	for {
		// optional cool visualization mode!
		// fm.printState(steps)
		var newRow int
		var newCol int
		switch fm.Guard.Heading {
		case 0:
			newRow = fm.Guard.Row - 1
			newCol = fm.Guard.Col
		case 1:
			newRow = fm.Guard.Row
			newCol = fm.Guard.Col + 1
		case 2:
			newRow = fm.Guard.Row + 1
			newCol = fm.Guard.Col
		case 3:
			newRow = fm.Guard.Row
			newCol = fm.Guard.Col - 1
		}
		if newRow < 0 || newRow >= len(fm.Cells) || newCol < 0 || newCol >= len(fm.Cells[0]) {
			return &steps
		}

		steps += fm.handleMovement(newRow, newCol)

	}
}

func (fm *FloorMap) handleMovement(newRow, newCol int) int {
	newSteps := 0
	if fm.Cells[newRow][newCol] == "#" {
		fm.Guard.Heading = (fm.Guard.Heading + 1) % 4
		fm.updateDirection()
		return newSteps
	}

	fm.updateDirection()

	if fm.Cells[newRow][newCol] == "." {
		newSteps++
	}

	fm.Guard.Row = newRow
	fm.Guard.Col = newCol

	return newSteps
}

func (fm *FloorMap) updateDirection() {
	switch fm.Guard.Heading {
	case 0:
		fm.Cells[fm.Guard.Row][fm.Guard.Col] = "^"
	case 1:
		fm.Cells[fm.Guard.Row][fm.Guard.Col] = ">"
	case 2:
		fm.Cells[fm.Guard.Row][fm.Guard.Col] = "v"
	case 3:
		fm.Cells[fm.Guard.Row][fm.Guard.Col] = "<"
	}
}

// func (fm FloorMap) printState(steps int) {
// 	fmt.Print("\033[H")

// 	for _, row := range fm.Cells {
// 		for _, cell := range row {
// 			fmt.Print(cell + " ")
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println("Steps:", steps)
// }

func parseDay6Input(filePath string) FloorMap {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file at %s", filePath)

	}
	defer file.Close()

	cells := [][]string{}
	scanner := bufio.NewScanner(file)
	rowNum := 0
	guard := Guard{}
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		for colNum, cell := range row {
			if cell == "^" {
				guard.Heading = 0
				guard.Row = rowNum
				guard.Col = colNum
			}
			if cell == ">" {
				guard.Heading = 1
				guard.Row = rowNum
				guard.Col = colNum
			} else if cell == "v" {
				guard.Heading = 2
				guard.Row = rowNum
				guard.Col = colNum
			} else if cell == "<" {
				guard.Heading = 3
				guard.Row = rowNum
				guard.Col = colNum
			}
		}
		cells = append(cells, row)
		rowNum++
	}
	return FloorMap{Cells: cells, Guard: guard}
}
