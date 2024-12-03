package days

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"rayjseth.io/aoc-24/model"
)

type reports [][]int

func Two() model.Result {
	day := uint8(2)
	reports := parseDay2Input("./inputs/day2")
	return model.Result{Day: &day, Part1: reports.calcPart1()}
}

func (r reports) calcPart1() *int {
	nSafe := 0

reportsLoop:
	for i, report := range r {
		if len(report) < 2 {
			log.Fatalf("Error line %d - each line must have at least two numbers", i)
		}

		iDiff := report[1] - report[0]
		if isSevereDelta(iDiff) {
			continue
		}
		isAsc := iDiff > 0

		// start at 2 since this is a lookback and first pair already processed above
		for j := 2; j < len(report); j++ {
			diff := report[j] - report[j-1]
			if isSevereDelta(diff) || isAsc != (diff > 0) {
				continue reportsLoop
			}
		}
		nSafe++
	}

	return &nSafe
}

func isSevereDelta(diff int) bool {
	return diff == 0 || math.Abs((float64(diff))) > 3
}

func parseDay2Input(filePath string) reports {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file at %s", filePath)

	}
	defer file.Close()

	var parsed reports
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var levels []int
		for _, s := range strings.Fields(scanner.Text()) {
			level, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("file contains nonnumeric string %s", s)
			}
			levels = append(levels, level)
		}
		parsed = append(parsed, levels)
	}

	return parsed
}
