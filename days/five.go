package days

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"rayjseth.io/aoc-24/model"
)

type Rule struct {
	MustBefore int
	MustAfter  int
}

type PageSet []int
type PageSets []PageSet

type Update struct {
	Rules []Rule
	Sets  PageSets
}

func Five() model.Result {
	day := uint8(5)
	update := parseDay5Input("./inputs/day5")
	return model.Result{Day: &day, Part1: update.calcPart1(), Part2: update.calcPart2()}
}

func (u Update) calcPart1() *int {
	mPgSum := 0

	for _, set := range u.Sets {
		isCorrectOrder := true
		for _, rule := range u.Rules {
			isCorrectOrder = isCorrectOrder && rule.checkSet(set)
		}
		if isCorrectOrder {
			mPgSum += set.getMiddlePage()
		}
	}

	return &mPgSum
}

func (u Update) calcPart2() *int {
	mPgSum := 0

	for _, set := range u.Sets {
		if set.fixSet(u.Rules) {
			mPgSum += set.getMiddlePage()
		}
	}

	return &mPgSum
}

func (set PageSet) fixSet(rules []Rule) bool {
	neededFix := false
	// swap until no more swaps can be made
	// would be better to detect conflicting rules and panic but I'd need
	// some kind of graph, so do cheap exit case for now
	maxIters := 1000
	for {
		swapped := false
		for _, rule := range rules {
			mustBeforeIdx := indexOfPage(set, rule.MustBefore)
			mustAfterIdx := indexOfPage(set, rule.MustAfter)

			if mustBeforeIdx != -1 && mustAfterIdx != -1 && mustBeforeIdx > mustAfterIdx {
				set[mustBeforeIdx], set[mustAfterIdx] = set[mustAfterIdx], set[mustBeforeIdx]
				neededFix = true
				swapped = true
				maxIters--
			}
		}

		if maxIters == 0 {
			log.Panicf("exceeded %d shuffles, check rules for contradiction", maxIters)
		}

		if !swapped {
			break
		}
	}
	return neededFix
}

func indexOfPage(arr []int, match int) int {
	for i, val := range arr {
		if val == match {
			return i
		}
	}

	return -1
}

func (set PageSet) getMiddlePage() int {
	setLen := len(set)
	if setLen%2 != 1 {
		log.Panic("Page sets must be in odd numbered sets")
	}
	middleIdx := setLen / 2
	return set[middleIdx]
}

func (rule Rule) checkSet(set PageSet) bool {
	mustBeforeIdx := indexOfPage(set, rule.MustBefore)
	mustAfterIdx := indexOfPage(set, rule.MustAfter)

	return mustBeforeIdx == -1 || mustAfterIdx == -1 || mustBeforeIdx < mustAfterIdx
}

func parseDay5Input(filePath string) Update {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file at %s", filePath)

	}
	defer file.Close()

	rules := []Rule{}
	sets := PageSets{}
	scanner := bufio.NewScanner(file)
	rulesParsed := false
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			rulesParsed = true
			continue
		}
		if rulesParsed {
			sPgNums := strings.Split(text, ",")
			pgNums := []int{}
			for _, sPgNum := range sPgNums {
				pgNum, err := strconv.Atoi(sPgNum)
				if err != nil {
					log.Panicf("Invalid page, received %s which is not an integer", sPgNum)
				}
				pgNums = append(pgNums, pgNum)
			}
			sets = append(sets, pgNums)
		} else {
			sRule := strings.Split(text, "|")
			if len(sRule) != 2 {
				log.Panicf("Invalid rule, must be two strings separated by a | and no whitespace")
			}
			rule := Rule{}
			for i, sRulePt := range sRule {
				rulePt, err := strconv.Atoi(sRulePt)
				if err != nil {
					log.Panicf("Invalid rule, rule parts must be integers")
				}
				if i == 0 {
					rule.MustBefore = rulePt
				} else {
					rule.MustAfter = rulePt
				}
			}
			rules = append(rules, rule)
		}
	}
	return Update{Rules: rules, Sets: sets}
}
