package days

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"rayjseth.io/aoc-24/model"
)

type Topography [][]int

func Ten() model.Result {
	day := uint8(10)
	t := parseDay10Input("./inputs/day10")
	return model.Result{Day: &day, Part1: t.calcPart1(), Part2: t.calcPart2()}
}

var validMoves = []Coord{
	{X: 0, Y: -1}, // up
	{X: 0, Y: 1},  // down
	{X: -1, Y: 0}, // left
	{X: 1, Y: 0},  // right
}

func (t Topography) calcPart1() *int {
	tScore := 0
	coords := []Coord{}
	for i, r := range t {
		for j, c := range r {
			if c == 0 {
				coords = append(coords, Coord{X: j, Y: i})
			}
		}
	}

	var wg sync.WaitGroup
	sChan := make(chan int, len(coords))

	for _, hiker := range coords {
		wg.Add(1)
		go func(c Coord) {
			defer wg.Done()
			sChan <- c.hike(t, true)
		}(hiker)
	}

	wg.Wait()
	close(sChan)

	for s := range sChan {
		tScore += s
	}

	return &tScore
}

func (t Topography) calcPart2() *int {
	tScore := 0
	coords := []Coord{}
	for i, r := range t {
		for j, c := range r {
			if c == 0 {
				// make a new hiker to send off on the trails!
				coords = append(coords, Coord{X: j, Y: i})
			}
		}
	}

	var wg sync.WaitGroup
	sChan := make(chan int, len(coords))

	for _, hiker := range coords {
		wg.Add(1)
		go func(c Coord) {
			defer wg.Done()
			sChan <- c.hike(t, false)
		}(hiker)
	}

	wg.Wait()
	close(sChan)

	for s := range sChan {
		tScore += s
	}

	return &tScore
}

func (c Coord) hike(t Topography, distinctPaths bool) int {
	pScore := 0
	hiked := make(map[Coord]bool)
	hiked[c] = true
	// superposition of all possible hikers given the state
	// of the topography
	pHikers := []Coord{c}
	for len(pHikers) > 0 {
		currHiker := pHikers[0]
		// remove the prev hiker as their superposition has either
		// been observed into a new position, or a possible path has been removed,
		// breaking the loop when all possibilities are depleted.
		// Don't @ me pysicists, this is how I picture it.
		pHikers = pHikers[1:]

		if t[currHiker.Y][currHiker.X] == 9 {
			pScore++
		}

		for _, dir := range validMoves {
			nHiker := Coord{X: currHiker.X + dir.X, Y: currHiker.Y + dir.Y}
			nInBounds := nHiker.X >= 0 && nHiker.X < len(t[0]) &&
				nHiker.Y >= 0 && nHiker.Y < len(t)

			if nInBounds {
				nGradual := t[nHiker.Y][nHiker.X] == t[currHiker.Y][currHiker.X]+1
				if nGradual {
					if distinctPaths {
						nHiked := hiked[nHiker]
						if !nHiked {
							// wow sheesh, this memo for paths was what had me banging my head
							// for hours on part 1. Then I hit part two and I'm like... there's no
							// heckin' way that I was solving part2 by mistake. So now I'm tired
							// and want to just make this conditional to switch it off for part2.
							if distinctPaths {
								hiked[nHiker] = true
								pHikers = append(pHikers, nHiker)
							}
						}
					} else {
						pHikers = append(pHikers, nHiker)
					}
				}
			}
		}
	}

	return pScore
}

func parseDay10Input(filePath string) Topography {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file at %s", filePath)

	}
	defer file.Close()

	t := Topography{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tLine := []int{}
		els := strings.Split(scanner.Text(), "")
		for _, el := range els {
			elI, err := strconv.Atoi(el)
			if err != nil {
				log.Panicf("Expected all numbers, received %s", el)
			}
			tLine = append(tLine, elI)
		}
		t = append(t, tLine)
	}
	return t
}
