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

type Operation int

const (
	Add Operation = iota
	Mult
	Concat
)

type Calibration struct {
	Target   int
	Operands []int
}

type Calibrations []Calibration

func Seven() model.Result {
	day := uint8(7)
	cs := parseDay7Input("./inputs/day7")
	return model.Result{Day: &day, Part1: cs.calcPart1(), Part2: cs.calcPart2()}
}

func (cs Calibrations) calcPart1() *int {
	var cTot int
	var wg sync.WaitGroup
	rc := make(chan int, len(cs))

	for _, c := range cs {
		wg.Add(1)
		go func(c Calibration) {
			defer wg.Done()
			if calcCombinations(c.Operands, c.Operands[0], c.Target, []Operation{Add, Mult}) {
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

func (cs Calibrations) calcPart2() *int {
	var cTot int
	var wg sync.WaitGroup
	rc := make(chan int, len(cs))

	for _, c := range cs {
		wg.Add(1)
		go func(c Calibration) {
			defer wg.Done()
			if calcCombinations(c.Operands, c.Operands[0], c.Target, []Operation{Add, Mult, Concat}) {
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

func calcCombinations(operands []int, currentValue int, target int, ops []Operation) bool {
	shouldAdd, shouldMult, shouldConcat := false, false, false
	for _, op := range ops {
		switch op {
		case Add:
			shouldAdd = true
		case Mult:
			shouldMult = true
		case Concat:
			shouldConcat = true
		}
	}

	var recFunc func(index int, currentValue int) bool
	recFunc = func(index int, currentValue int) bool {
		if index == len(operands) {
			return currentValue == target
		}

		if shouldAdd && recFunc(index+1, currentValue+operands[index]) {
			return true
		}

		if shouldMult && recFunc(index+1, currentValue*operands[index]) {
			return true
		}

		if shouldConcat {
			concat, _ := strconv.Atoi(strconv.Itoa(currentValue) + strconv.Itoa(operands[index]))
			return recFunc(index+1, concat)
		}

		return false
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
