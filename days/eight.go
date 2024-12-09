package days

import (
	"bufio"
	"log"
	"os"

	"rayjseth.io/aoc-24/model"
)

type Coord struct {
	X int
	Y int
}

type Antenna struct {
	Kind   rune
	Coords []Coord
}

type Antinode struct {
	Coord Coord
}

type Antinodes []Antinode

type Antennas []Antenna

type City struct {
	Antennas Antennas
	Limits   Coord
}

func Eight() model.Result {
	day := uint8(8)
	city := parseDay8Input("./inputs/day8")
	return model.Result{Day: &day, Part1: city.calcPart1(), Part2: city.calcPart2()}
}

func (c City) calcPart1() *int {
	uAntinodes := Antinodes{}
	for _, a := range c.Antennas {
		if len(a.Coords) > 1 {
			for i, coord := range a.Coords {
				for j := i + 1; j < len(a.Coords); j++ {
					nCoord := a.Coords[j]
					dX := coord.X - nCoord.X
					dY := coord.Y - nCoord.Y

					a1X := coord.X + dX
					a1Y := coord.Y + dY
					a2X := nCoord.X - dX
					a2Y := nCoord.Y - dY

					if a1X > -1 && a1X < c.Limits.X &&
						a1Y > -1 && a1Y < c.Limits.Y {
						aNode1 := Antinode{Coord: Coord{X: a1X, Y: a1Y}}
						if !uAntinodes.contains(aNode1) {
							uAntinodes = append(uAntinodes, aNode1)
						}
					}
					if a2X > -1 && a2X < c.Limits.X &&
						a2Y > -1 && a2Y < c.Limits.Y {
						aNode2 := Antinode{Coord: Coord{X: a2X, Y: a2Y}}
						if !uAntinodes.contains(aNode2) {
							uAntinodes = append(uAntinodes, aNode2)
						}
					}
				}
			}
		}
	}

	uANLen := len(uAntinodes)
	return &uANLen
}

func (c City) calcPart2() *int {
	uAntinodes := Antinodes{}
	for _, a := range c.Antennas {
		if len(a.Coords) > 1 {
			for i, coord := range a.Coords {
				for j := i + 1; j < len(a.Coords); j++ {
					nCoord := a.Coords[j]

					projections := 1
					for {
						dX := (coord.X - nCoord.X) * projections
						dY := (coord.Y - nCoord.Y) * projections

						a1X := coord.X + dX
						a1Y := coord.Y + dY
						a2X := nCoord.X - dX
						a2Y := nCoord.Y - dY

						a1OffGrid := a1X < 0 || a1X >= c.Limits.X || a1Y < 0 || a1Y >= c.Limits.Y
						a2OffGrid := a2X < 0 || a2X >= c.Limits.X || a2Y < 0 || a2Y >= c.Limits.Y

						if !a1OffGrid {
							aCoord1 := Coord{X: a1X, Y: a1Y}
							aNode1 := Antinode{Coord: aCoord1}
							if !c.Antennas.hasAntennaAtCoord(aCoord1) && !uAntinodes.contains(aNode1) {
								uAntinodes = append(uAntinodes, aNode1)
							}
						}
						if !a2OffGrid {
							aCoord2 := Coord{X: a2X, Y: a2Y}
							aNode2 := Antinode{Coord: aCoord2}
							if !c.Antennas.hasAntennaAtCoord(aCoord2) && !uAntinodes.contains(aNode2) {
								uAntinodes = append(uAntinodes, aNode2)
							}
						}

						if a1OffGrid && a2OffGrid {
							break
						}
						projections++
					}
				}
			}
		}
	}

	uANLen := len(uAntinodes)
	for _, a := range c.Antennas {
		if len(a.Coords) > 1 {
			uANLen += len(a.Coords)
		}
	}

	return &uANLen
}

func (an Antinodes) contains(match Antinode) bool {
	for _, a := range an {
		if match.Coord.X == a.Coord.X && match.Coord.Y == a.Coord.Y {
			return true
		}
	}
	return false
}

func (as Antennas) hasAntennaAtCoord(c Coord) bool {
	for _, a := range as {
		for _, ac := range a.Coords {
			if c.X == ac.X && c.Y == ac.Y {
				return true
			}
		}
	}

	return false
}

func parseDay8Input(filePath string) City {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file at %s", filePath)

	}
	defer file.Close()

	antennas := Antennas{}
	scanner := bufio.NewScanner(file)
	i, limitX := 0, -1
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		if limitX > -1 {
			if limitX != len(runes) {
				log.Panic("Only rectangular towns allowed in this puzzle!")
			}
		} else {
			limitX = len(runes)
		}
		for j, r := range runes {
			if (r > 47 && r < 58) ||
				(r > 64 && r < 91) ||
				(r > 96 && r < 123) {
				newCoord := Coord{j, i}
				aIdx := -1
				for k, antenna := range antennas {
					if antenna.Kind == r {
						aIdx = k
						break
					}
				}
				if aIdx == -1 {
					newAntenna := Antenna{Kind: r}
					newAntenna.Coords = append(newAntenna.Coords, newCoord)
					antennas = append(antennas, newAntenna)
				} else {
					antennas[aIdx].Coords = append(antennas[aIdx].Coords, newCoord)
				}
			}
		}
		i++
	}

	return City{Antennas: antennas, Limits: Coord{X: limitX, Y: i}}
}
