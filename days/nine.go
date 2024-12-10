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

type Blocks []Block

func Nine() model.Result {
	day := uint8(9)
	disk := parseDay9Input("./inputs/day9")
	return model.Result{Day: &day, Part1: disk.calcPart1(), Part2: disk.calcPart2()}
}

func (d Disk) calcPart1() *int {
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

	return fragged.calcChecksum()
}

func (d Disk) calcPart2() *int {
	blocks := d.toBlocks()

	for {
		moved := false

		for i := len(blocks) - 1; i >= 0; i-- {
			if blocks[i].IsFile {
				for j := 0; j < i; j++ {
					if !blocks[j].IsFile && blocks[j].Len >= blocks[i].Len {
						if blocks[j].Len == blocks[i].Len {
							blocks[i], blocks[j] = blocks[j], blocks[i]
						} else {
							blocks[j].Len -= blocks[i].Len
							blocks = append(blocks[:j], append(Blocks{blocks[i]}, blocks[j:]...)...)
							blocks[i+1].ID = blocks[j+1].ID
							// omfg I was banging my head against the wall for hours  because I forgot to flip this.
							// "Dont' worry about swapping" I said, "You can just modify the right side directly" I said.
							// SMH my head.
							blocks[i+1].IsFile = false
						}

						moved = true
						break
					}
				}
			}
			if moved {
				break
			}
		}
		if !moved {
			break
		}
	}

	// yeah I know I coulda calc'd based on blocks but after getting stuck on that silly
	// isFile flip for a long time, I honestly can't be arsed - and I already have a known
	// good calc for the disk format :shrug:
	return blocks.toDisk().calcChecksum()
}

func (d Disk) toBlocks() Blocks {
	bs := Blocks{}

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
		bs = append(bs, Block{ID: currId, IsFile: currId != ".", Len: blkLen})
	}

	return bs
}

func (bs Blocks) toDisk() Disk {
	disk := Disk{}

	for _, b := range bs {
		for i := 0; i < b.Len; i++ {
			disk = append(disk, b.ID)
		}
	}

	return disk
}

func (d Disk) calcChecksum() *int {
	cSum := 0

	for i, f := range d {
		if f == "." {
			continue
		}
		fI, _ := strconv.Atoi(f)
		cSum += i * fI
	}

	return &cSum
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
