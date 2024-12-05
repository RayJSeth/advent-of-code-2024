package days

import (
	"bufio"
	"log"
	"os"
	"strings"

	"rayjseth.io/aoc-24/model"
)

type WordSearch [][]string

var match = "XMAS"
var matchLen = len(match)

func Four() model.Result {
	day := uint8(4)
	wordSearch := parseDay4Input("./inputs/day4")
	return model.Result{Day: &day, Part1: wordSearch.calcPart1()}
}

func (w WordSearch) calcPart1() *int {
	hits := 0

	for i, l := range w {
		for j := 0; j < len(l); j++ {
			// 8 cardinal directions (n,s,e,w,ne,nw,se,sw) - skipping neutral (0, 0)
			// this works like a clock where the "big hand" xDir sweeps -1, -1, -1,  0,  0,  0,  1,  1,  1
			// while yDir using modulo acts as the "small hand" and ticks -1,  0,  1, -1,  0,  1, -1,  0   1
			// generating all combinations :)
			for dir := 0; dir < 9; dir++ {
				if dir == 4 {
					continue
				}
				xDir := (dir/3 - 1)
				yDir := (dir%3 - 1)
				if w.checkDir(i, j, xDir, yDir) { // right *
					hits++
				}
			}
		}
	}
	return &hits
}

func (w WordSearch) checkDir(rowIdx int, colIdx int, xDir int, yDir int) bool {
	rows := len(w)
	cols := len(w[rowIdx])

	// skip if not enough room left in a dir for a match
	if rowIdx+xDir*(matchLen-1) < 0 || rowIdx+xDir*(matchLen-1) >= rows ||
		colIdx+yDir*(matchLen-1) < 0 || colIdx+yDir*(matchLen-1) >= cols {
		return false
	}

	var substring string
	for i := 0; i < len(match); i++ {
		nr := rowIdx + xDir*i
		nc := colIdx + yDir*i
		substring += w[nr][nc]
	}

	return substring == match
}

func parseDay4Input(filePath string) WordSearch {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file at %s", filePath)

	}
	defer file.Close()

	parsed := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsed = append(parsed, strings.Split(scanner.Text(), ""))
	}
	return parsed
}
