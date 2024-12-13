package days

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"rayjseth.io/aoc-24/model"
)

type Stones []int

func Eleven() model.Result {
	day := uint8(11)
	stones := parseDay11Input("./inputs/day11")

	sCopy2 := make(Stones, len(stones))
	copy(sCopy2, stones)

	return model.Result{Day: &day, Part1: sCopy2.calcPart1()}
}

func (stones Stones) calcPart1() *int {
	tot := 0

	for range 25 {
		var newStones Stones

		for i := 0; i < len(stones); i++ {
			stone := stones[i]
			if stone == 0 {
				newStones = append(newStones, 1)
			} else {
				stoneStr := strconv.Itoa(stone)
				stoneStrLen := len(stoneStr)

				if stoneStrLen%2 == 0 {
					left, _ := strconv.Atoi(stoneStr[:stoneStrLen/2])
					right, _ := strconv.Atoi(stoneStr[stoneStrLen/2:])
					newStones = append(newStones, left, right)
				} else {
					newStones = append(newStones, stone*2024)
				}
			}
		}

		stones = newStones
	}

	tot = len(stones)
	return &tot
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
			stones = append(stones, fI)
		}
	}

	return stones
}
