package model

import (
	"fmt"
	"log"
)

var sp1 = "Part1"
var sp2 = "Part2"
var sNI = "not implemented"

type Result struct {
	Day   *uint8
	Part1 *int
	Part2 *int
}

func (r Result) Print() {
	if r.Day == nil {
		log.Fatal("Invalid Result struct (Day must not be nil)")
	}

	fmt.Printf("\nDay %d\n", *r.Day)

	if r.Part1 != nil {
		fmt.Printf("%s %d\n", sp1, *r.Part1)
	} else {
		fmt.Printf("%s %s\n", sp1, sNI)
	}

	if r.Part2 != nil {
		fmt.Printf("%s %d\n", sp2, *r.Part2)
	} else {
		fmt.Printf("%s %s\n", sp2, sNI)
	}
}
