package days

import (
	"log"
	"os"
	"strconv"
	"strings"

	"rayjseth.io/aoc-24/model"
)

type Disk []string

type Block struct {
	ID     string
	IsFile bool
	Len    int
}

func Nine() model.Result {
	day := uint8(9)
	disk := parseDay9Input("./inputs/day9")
	return model.Result{Day: &day, Part1: disk.calcPart1(), Part2: disk.calcPart2()}
}

func (d Disk) calcPart1() *int {
	cSum := 0
	fragged := make(Disk, len(d))
	copy(fragged, d)

	left := 0
	right := len(fragged) - 1

	for {
		for left < len(fragged) && fragged[left] != "." {
			left++
		}

		for right >= 0 && fragged[right] == "." {
			right--
		}

		if left < right {
			fragged[left], fragged[right] = fragged[right], fragged[left]
			left++
			right--
		} else {
			break
		}
	}

	for i, f := range fragged {
		if f == "." {
			break
		}
		fI, _ := strconv.Atoi(f)
		cSum += i * fI
	}

	return &cSum
}

func (d Disk) calcPart2() *int {
	cSum := 0

	bs := d.toBlocks()
	defragged := make([]Block, len(bs))
	copy(defragged, bs)

	left := 0
	right := len(defragged) - 1

	for {
		for left < len(defragged) && defragged[left].ID != "." {
			left++
		}

		for right >= 0 && defragged[right].ID == "." {
			right--
		}

		if left < right {
			break
			// check if left is empty space "." and has Len >= right Len
			// if so, swap right into left, and swap the equivalent number of "." right
			// this may leave some "." leftover on the left side after the swapped values
		} else {
			break
		}

	}

	return &cSum
}

func (d Disk) toBlocks() []Block {
	bs := []Block{}
	if len(d) > 0 {
		currId := "0"
		blkLen := 0
		for _, f := range d {
			if currId == f {
				blkLen++
			} else {
				bs = append(bs, Block{ID: currId, IsFile: currId != ".", Len: blkLen})
				currId = f
				blkLen = 1
			}
		}
	}
	return bs
}

func parseDay9Input(filePath string) Disk {
	disk := Disk{}
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error opening file at %s", filePath)
	}
	fID := 0
	for i, s := range strings.Split(string(file), "") {
		isFile := i%2 == 0
		blockLen, err := strconv.Atoi(s)
		if err != nil {
			log.Panicf("input must be all digits or dots, received %s", s)
		}
		if isFile {
			for j := 0; j < blockLen; j++ {
				disk = append(disk, strconv.Itoa(fID))
			}
			fID++
		} else {
			for j := 0; j < blockLen; j++ {
				disk = append(disk, ".")
			}
		}
	}
	return Disk(disk)
}
