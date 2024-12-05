package days

import (
	"bufio"
	"log"
	"os"
	"strings"

	"rayjseth.io/aoc-24/model"
)

type WordSearch [][]string

func Four() model.Result {
	day := uint8(4)
	wordSearch := parseDay4Input("./inputs/day4")
	return model.Result{Day: &day, Part1: wordSearch.calcPart1(), Part2: wordSearch.calcPart2()}
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
				if w.checkDir(i, j, xDir, yDir, "XMAS") { // right *
					hits++
				}
			}
		}
	}

	return &hits
}

func (w WordSearch) calcPart2() *int {
	hits := 0

	for i, l := range w {
		l = w[i]
		for j := 0; j < len(l); j++ {
			if w.checkCross(i, j, "MAS") {
				hits++
			}
		}
	}

	return &hits
}

func (w WordSearch) checkDir(rowIdx int, colIdx int, xDir int, yDir int, match string) bool {
	rows := len(w)
	cols := len(w[rowIdx])
	matchLen := len(match)

	// skip if not enough room left in a dir for a match
	if rowIdx+xDir*(matchLen-1) < 0 || rowIdx+xDir*(matchLen-1) >= rows ||
		colIdx+yDir*(matchLen-1) < 0 || colIdx+yDir*(matchLen-1) >= cols {
		return false
	}

	var substring string
	for i := 0; i < matchLen; i++ {
		nr := rowIdx + xDir*i
		nc := colIdx + yDir*i
		substring += w[nr][nc]
	}

	return substring == match
}

func (w WordSearch) checkCross(rowIdx int, colIdx int, match string) bool {
	rows := len(w)
	cols := len(w[rowIdx])
	// split in half flooring since centering on rune
	matchLen := len(match)
	if matchLen%2 != 1 {
		log.Panic("cannot check cross on even length strings")
	}

	matchExtn := int(float64(matchLen / 2))
	// skip if not enough room left in all dirs for a cross
	if rowIdx-matchExtn < 0 || rowIdx+matchExtn >= rows ||
		colIdx-matchExtn < 0 || colIdx+matchExtn >= cols {
		return false
	}

	// using go is like always being in an interview
	// where they ask you to do something basic without a std lib func
	var matchRev string
	matchRunes := []rune(match)
	for i, j := 0, len(matchRunes)-1; i < j; i, j = i+1, j-1 {
		matchRunes[i], matchRunes[j] = matchRunes[j], matchRunes[i]
	}
	matchRev = string(matchRunes)

	var NWSESubstring string
	var NESWSubstring string
	for i := 1; i <= matchExtn; i++ {
		currCenter := w[rowIdx][colIdx]
		NWSESubstring += w[rowIdx-i][colIdx-i] + currCenter + w[rowIdx+i][colIdx+i]
		NESWSubstring += w[rowIdx-i][colIdx+i] + currCenter + w[rowIdx+i][colIdx-i]
	}

	if (NWSESubstring == match || NWSESubstring == matchRev) &&
		(NESWSubstring == match || NESWSubstring == matchRev) {
		return true
	}

	return false
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
