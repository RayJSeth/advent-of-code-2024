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

type Report []int
type Reports []Report

func Two() model.Result {
	day := uint8(2)
	reports := parseDay2Input("./inputs/day2")
	return model.Result{Day: &day, Part1: reports.calcPart1(), Part2: reports.calcPart2()}
}

func (r Reports) calcPart1() *int {
	nSafe := 0
	for i, report := range r {
		if len(report) < 2 {
			log.Fatalf("Error line %d - each line must have at least two numbers", i)
		}
		if report.isSafe() {
			nSafe++
		}
	}

	return &nSafe
}

func (r Reports) calcPart2() *int {
	nSafe := 0

reportsLoop:
	for i, report := range r {
		if len(report) < 3 {
			log.Fatalf("Error line %d - each line must have at least three numbers", i)
		}

		for j := 0; j < len(report); j++ {
			modifiedReport := append(append(Report{}, report[:j]...), report[j+1:]...)
			// TODO, maybe return unsafe idx and target modifications rather than brute force
			// challenge would be discerning if violation is curr or lookbehind to know where
			// to cut
			if modifiedReport.isSafe() {
				nSafe++
				continue reportsLoop
			}
		}
	}
	return &nSafe
}

func (r Report) isSafe() bool {
	isAsc := r[1] > r[0]
	for i := 1; i < len(r); i++ {
		diff := r[i] - r[i-1]
		if isInvalidDelta(diff) || isAsc != (diff > 0) {
			return false
		}
	}
	return true
}

func isInvalidDelta(diff int) bool {
	return diff == 0 || math.Abs((float64(diff))) > 3
}

func parseDay2Input(filePath string) Reports {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file at %s", filePath)

	}
	defer file.Close()

	var parsed Reports
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
