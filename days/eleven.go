package days

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	for rIter := range 45 {
		var newStones Stones

		for i := 0; i < len(stones); i++ {
			stone := stones[i]
			if stone == 0 {
				newStones = append(newStones, 1)
			} else {
				numDigits := int(math.Floor(math.Log10(float64(stone)))) + 1

				if numDigits%2 == 0 {
					divisor := int(math.Pow(10, float64(numDigits/2)))
					left := stone / divisor
					right := stone % divisor
					newStones = append(newStones, left, right)
				} else {
					newStones = append(newStones, stone*2024)
				}
			}
		}

		stones = newStones
		fmt.Println(rIter)
		fmt.Println(len(stones))
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
