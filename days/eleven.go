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

type Stone int
type Stones []Stone

type BlinkRes struct {
	Stone Stone
	Depth int
}

func Eleven() model.Result {
	day := uint8(11)
	stones := parseDay11Input("./inputs/day11")

	memo := &map[BlinkRes]int{}
	return model.Result{Day: &day, Part1: stones.calcPart1(memo), Part2: stones.calcPart2(memo)}
}

func (stones Stones) calcPart1(memo *map[BlinkRes]int) *int {
	tot := 0
	for _, stone := range stones {
		tot += stone.blink(25, memo)
	}
	return &tot
}

func (stones Stones) calcPart2(memo *map[BlinkRes]int) *int {
	tot := 0
	for _, stone := range stones {
		tot += stone.blink(75, memo)
	}
	return &tot
}

func (stone Stone) blink(nBlinks int, memo *map[BlinkRes]int) int {
	result := 0
	if nBlinks > 0 {
		mRes, mExists := (*memo)[BlinkRes{Stone: stone, Depth: nBlinks}]
		if mExists {
			result += mRes
		} else {
			numDigits := int(math.Floor(math.Log10(float64(stone)))) + 1
			if stone == 0 {
				result += Stone(1).blink(nBlinks-1, memo)
			} else if numDigits%2 == 0 {
				divisor := int(math.Pow(10, float64(numDigits/2)))
				left := Stone(int(stone) / divisor)
				right := Stone(int(stone) % divisor)
				result += left.blink(nBlinks-1, memo)
				result += right.blink(nBlinks-1, memo)
			} else {
				result += (stone * 2024).blink(nBlinks-1, memo)
			}

			(*memo)[BlinkRes{Stone: stone, Depth: nBlinks}] = result
		}
	} else {
		return 1
	}
	return result
}

func parseDay11Input(filePath string) Stones {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file at %s", filePath)
	}
	defer file.Close()

	stones := Stones{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, f := range strings.Fields(scanner.Text()) {
			fI, err := strconv.Atoi(f)
			if err != nil {
				log.Panicf("Malformed input, expected only integers and whitespace - encountered %s", f)
			}
			stones = append(stones, Stone(fI))
		}
	}

	return stones
}
