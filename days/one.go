package days

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"rayjseth.io/aoc-24/model"
)

type lists struct {
	List1 []int
	List2 []int
}

func One() model.Result {
	day := uint8(1)
	lists := parseDay1Input("./inputs/day1").sortLists()
	return model.Result{Day: &day, Part1: lists.calcPart1(), Part2: lists.calcPart2()}
}

func (l lists) calcPart1() *int {
	dSum := 0
	for i, ls1Val := range l.List1 {
		ls2Val := l.List2[i]
		dSum += int(math.Abs(float64(ls2Val) - float64(ls1Val)))
	}
	return &dSum
}

func (l lists) calcPart2() *int {
	dSim := 0

	for i := 0; i < len(l.List1); i++ {
		j, streak := 0, 0
		// move forward to the first similarity
		for j < len(l.List2) && l.List2[j] < l.List1[i] {
			j++
		}

		// move through groups and count multiplier
		for j < len(l.List2) && l.List2[j] == l.List1[i] {
			streak++
			j++
		}

		// if no similarity, streak is 0 so add doesn't do anything and moves to next list1 item
		dSim += l.List1[i] * streak
	}

	return &dSim
}

func (l lists) sortLists() lists {
	slices.Sort(l.List1)
	slices.Sort(l.List2)
	return l
}

func parseDay1Input(filePath string) lists {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file at %s", filePath)

	}
	defer file.Close()

	parsed := lists{List1: []int{}, List2: []int{}}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		locs := strings.Fields(scanner.Text())
		if len(locs) != 2 {
			log.Fatal("Expected input should be pairs of strings separated by whitespace")
		}
		l1Int, err := strconv.ParseInt(locs[0], 10, 32)
		if err != nil {
			log.Fatalf("Expected %s to be int", locs[0])
		}
		l2Int, err := strconv.ParseInt(locs[1], 10, 32)
		if err != nil {
			log.Fatalf("Expected %s to be int", locs[1])
		}
		parsed.List1 = append(parsed.List1, int(l1Int))
		parsed.List2 = append(parsed.List2, int(l2Int))
	}
	return parsed
}
