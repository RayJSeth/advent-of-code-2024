package days

import (
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"rayjseth.io/aoc-24/model"
)

type Program string

var iPre, iPost, iSep = "mul(", ')', ','

func Three() model.Result {
	day := uint8(3)
	program := parseDay3Input("./inputs/day3")
	return model.Result{Day: &day, Part1: program.calcPart1()}
}

func (p Program) calcPart1() *int {
	res := 0
pLoop:
	for i := 0; i < len(p)-8; i++ {
		iStart := i + len(iPre)
		sub := p[i:iStart]
		if string(sub) == iPre {
			// mul(
			iEnd := -1
			for j := iStart + 1; j < len(p); j++ {
				r := rune(p[j])
				if r == iPost {
					iEnd = j
					break
				}
				if !(unicode.IsDigit(r) || r == iSep) {
					continue pLoop
				}
			}
			if iEnd-iStart > 2 {
				iOps := strings.SplitN(string(p[iStart:iEnd]), string(iSep), 2)
				// can't have err here since already checked if string is
				// only digits or comma, and removed the comma already
				multiplicand, _ := strconv.Atoi(iOps[0])
				multiplier, _ := strconv.Atoi(iOps[1])

				res += (multiplicand * multiplier)
			}
		}
	}
	return &res
}

func parseDay3Input(filePath string) Program {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error opening file at %s", filePath)

	}
	return Program(file)
}
