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

type Calibration struct {
	Target   int
	Operands []int
}

type Calibrations []Calibration

func Seven() model.Result {
	day := uint8(7)
	cs := parseDay7Input("./inputs/day7")
	return model.Result{Day: &day, Part1: cs.calcPart1()}
}

func (cs Calibrations) calcPart1() *int {
	var cTot int
	var wg sync.WaitGroup
	rc := make(chan int, len(cs))

	for _, c := range cs {
		wg.Add(1)
		go func(c Calibration) {
			defer wg.Done()
			if calcCombinations(c.Operands, c.Operands[0], c.Target) {
				rc <- c.Target
			}
		}(c)
	}

	wg.Wait()
	close(rc)

	for r := range rc {
		cTot += r
	}

	return &cTot
}

func calcCombinations(operands []int, currentValue int, target int) bool {
	var recFunc func(index int, currentValue int) bool
	recFunc = func(index int, currentValue int) bool {
		if index == len(operands) {
			return currentValue == target
		}

		addResult := recFunc(index+1, currentValue+operands[index])
		if addResult {
			return true
		}

		return recFunc(index+1, currentValue*operands[index])
	}

	return recFunc(1, currentValue)
}

func parseDay7Input(filePath string) Calibrations {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file at %s", filePath)

	}
	defer file.Close()

	var calibrations Calibrations

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		partSplit := strings.Split(text, ": ")
		tInt, err := strconv.Atoi(partSplit[0])
		if err != nil {
			log.Panicf("Malformed input, expected int, received %s", partSplit[0])
		}
		oSArr := strings.Split(partSplit[1], " ")
		var oInts = make([]int, len(oSArr))
		for i, oStr := range oSArr {
			oInt, err := strconv.Atoi(oStr)
			if err != nil {
				log.Panicf("Malformed input, expected int, received %s", oStr)
			}
			oInts[i] = oInt
		}
		calibrations = append(calibrations, Calibration{Target: tInt, Operands: oInts})
	}
	return calibrations
}
