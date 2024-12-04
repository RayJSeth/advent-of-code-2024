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

var iPre, iDo, iDont = "mul(", "do()", "don't()"
var iPost, iSep = ')', ','

func Three() model.Result {
	day := uint8(3)
	program := parseDay3Input("./inputs/day3")
	return model.Result{Day: &day, Part1: program.calcPart1(), Part2: program.calcPart2()}
}

func (p Program) calcPart1() *int {
	res := 0
pLoop:
	for i := 0; i < len(p)-8; i++ {
		multEnd := i + len(iPre)
		multSub := p[i:multEnd]
		if string(multSub) == iPre {
			iEnd := -1
			for j := multEnd + 1; j < len(p); j++ {
				r := rune(p[j])
				if r == iPost {
					iEnd = j
					break
				}
				if !(unicode.IsDigit(r) || r == iSep) {
					continue pLoop
				}
			}
			if iEnd-multEnd > 2 {
				iOps := strings.SplitN(string(p[multEnd:iEnd]), string(iSep), 2)
				multiplicand, _ := strconv.Atoi(iOps[0])
				multiplier, _ := strconv.Atoi(iOps[1])

				res += (multiplicand * multiplier)
			}
		}
	}
	return &res
}

func (p Program) calcPart2() *int {
	res, shouldDo := 0, true
pLoop:
	for i := 0; i < len(p)-8; i++ {
		multOrDoEnd := i + len(iPre)
		dontInstEnd := i + len(iDont)
		// luckily "mult(" and "do()" are same len so can just check one sub for match
		multOrDoSub := p[i:multOrDoEnd]
		// unluckily, "don't()" is a diff len, but still less than 8 so can lookahead
		// without risk of OOB
		dontSub := p[i:dontInstEnd]

		if string(dontSub) == iDont {
			shouldDo = false
		}
		if string(multOrDoSub) == iDo {
			shouldDo = true
		} else if shouldDo && string(multOrDoSub) == iPre {
			iEnd := -1
			for j := multOrDoEnd + 1; j < len(p); j++ {
				r := rune(p[j])
				if r == iPost {
					iEnd = j
					break
				}
				if !(unicode.IsDigit(r) || r == iSep) {
					continue pLoop
				}
			}
			if iEnd-multOrDoEnd > 2 {
				iOps := strings.SplitN(string(p[multOrDoEnd:iEnd]), string(iSep), 2)
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
