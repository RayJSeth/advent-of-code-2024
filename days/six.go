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
type Cell struct {
	Icon      rune
	Footprint Footprint
}
type Cells [][]Cell

// but why is there only one set of footprints in the sand?
// because of a bug on line 73, He replied.
type Footprint struct {
	Orientations []Heading
}
type FloorMap struct {
	Cells Cells
	Guard Guard
}

func Six() model.Result {
	day := uint8(6)
	fm := parseDay6Input("./inputs/day6")
	fm2 := fm.deepCopy()

	return model.Result{Day: &day, Part1: fm.calcPart1(), Part2: fm2.calcPart2(&fm)}
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

func (fm *FloorMap) calcPart2(originalRun *FloorMap) *int {
	loops := 0

	ogCells := fm.Cells.deepCopy()
	ogGuard := fm.Guard

	for row := 0; row < len(fm.Cells); row++ {
		for col := 0; col < len(fm.Cells[row]); col++ {

			if originalRun.Cells[row][col].Icon == 'x' {
				fm.Cells[row][col].Icon = '#'

				fm.Guard = ogGuard
				steps := 0
				for {
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
						break
					}

					steps += fm.handleMovement(newRow, newCol)

					if fm.isRetracingSteps() {
						loops++
						break
					}
				}

				fm.Cells[row][col].Icon = '.'
				fm.Cells = ogCells.deepCopy()
			}
		}
	}

	return &loops
}

func (original FloorMap) deepCopy() FloorMap {
	cellsCopy := original.Cells.deepCopy()
	guardCopy := original.Guard
	return FloorMap{Cells: cellsCopy, Guard: guardCopy}
}

func (original Cells) deepCopy() Cells {
	copied := make(Cells, len(original))
	for i := range original {
		copied[i] = append([]Cell(nil), original[i]...)
	}
	return copied
}

func (fm *FloorMap) handleMovement(newRow, newCol int) int {
	newSteps := 0
	if fm.Cells[newRow][newCol].Icon == '#' {
		fm.Guard.Heading = (fm.Guard.Heading + 1) % 4
		return newSteps
	}

	fm.Cells[newRow][newCol].Footprint.Orientations = append(fm.Cells[newRow][newCol].Footprint.Orientations, fm.Guard.Heading)

	if fm.Cells[newRow][newCol].Icon == '.' {
		newSteps++
	}

	fm.Cells[fm.Guard.Row][fm.Guard.Col].Icon = 'x'
	fm.Guard.Row = newRow
	fm.Guard.Col = newCol

	return newSteps
}

func (fm FloorMap) isRetracingSteps() bool {
	switch fm.Guard.Heading {
	case 0:
		chkPos1 := fm.Guard.Row - 1
		if chkPos1 >= 0 {
			for _, orientation := range fm.Cells[chkPos1][fm.Guard.Col].Footprint.Orientations {
				if orientation == fm.Guard.Heading {
					return true
				}
			}
		}
	case 1:
		chkPos1 := fm.Guard.Col + 1
		if chkPos1 < len(fm.Cells[fm.Guard.Row]) {
			for _, orientaion := range fm.Cells[fm.Guard.Row][chkPos1].Footprint.Orientations {
				if orientaion == fm.Guard.Heading {
					return true
				}
			}
		}
	case 2:
		chkPos1 := fm.Guard.Row + 1
		if chkPos1 < len(fm.Cells) {
			for _, orientation := range fm.Cells[chkPos1][fm.Guard.Col].Footprint.Orientations {
				if orientation == fm.Guard.Heading {
					return true
				}
			}
		}
	case 3:
		chkPos1 := fm.Guard.Col - 1
		if chkPos1 >= 0 {
			for _, orientation := range fm.Cells[fm.Guard.Row][chkPos1].Footprint.Orientations {
				if orientation == fm.Guard.Heading {
					return true
				}
			}
		}
	}
	return false
}

// func (fm FloorMap) printState(steps int) {
// 	fmt.Print("\033[H")

// 	for _, row := range fm.Cells {
// 		for _, cell := range row {
// 			fmt.Print(string(cell.Icon) + " ")
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

	cells := [][]Cell{}
	scanner := bufio.NewScanner(file)
	rowNum := 0
	guard := Guard{}
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		rowCells := []Cell{}
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
			rowCells = append(rowCells, Cell{Icon: []rune(cell)[0]})
		}
		cells = append(cells, rowCells)
		rowNum++
	}
	return FloorMap{Cells: cells, Guard: guard}
}
