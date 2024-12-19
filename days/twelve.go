package days

import (
	"bufio"
	"log"
	"os"

	"rayjseth.io/aoc-24/model"
)

type Garden [][]rune

type Region struct {
	Area      int
	Perimeter int
}
type Regions []Region
type RegionMap map[rune]Regions
type CoordPair [2]int
type Visited map[CoordPair]bool

func Twelve() model.Result {
	day := uint8(12)
	garden := parseDay12Input("./inputs/day12")

	return model.Result{Day: &day, Part1: garden.calcPart1()}
}

func (garden Garden) calcPart1() *int {
	cost := 0
	visited := Visited{}
	rMap := RegionMap{}

	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[i]); j++ {
			coord := CoordPair{i, j}
			if !visited[coord] {
				crop := garden[i][j]
				region := visited.surveyRegion(garden, coord, crop)
				rMap[crop] = append(rMap[crop], region)
			}
		}
	}

	for _, regions := range rMap {
		for _, region := range regions {
			cost += region.Area * region.Perimeter
		}
	}

	return &cost
}

func (visited *Visited) surveyRegion(garden Garden, coords CoordPair, crop rune) Region {
	i, j := coords[0], coords[1]

	(*visited)[CoordPair{i, j}] = true

	region := Region{}

	isLeftMost := j == 0 || garden[i][j-1] != crop
	isRightmost := j == len(garden[i])-1 || garden[i][j+1] != crop
	isTopmost := i == 0 || garden[i-1][j] != crop
	isBottommost := i == len(garden)-1 || garden[i+1][j] != crop

	region.Area++

	if isTopmost {
		region.Perimeter++
	} else {
		north := CoordPair{i - 1, j}
		if !(*visited)[north] {
			northRegion := visited.surveyRegion(garden, north, crop)
			region.Area += northRegion.Area
			region.Perimeter += northRegion.Perimeter
		}
	}
	if isBottommost {
		region.Perimeter++
	} else {
		south := CoordPair{i + 1, j}
		if !(*visited)[south] {
			southRegion := visited.surveyRegion(garden, south, crop)
			region.Area += southRegion.Area
			region.Perimeter += southRegion.Perimeter
		}
	}
	if isRightmost {
		region.Perimeter++
	} else {
		east := CoordPair{i, j + 1}
		if !(*visited)[east] {
			eastRegion := visited.surveyRegion(garden, east, crop)
			region.Area += eastRegion.Area
			region.Perimeter += eastRegion.Perimeter
		}
	}
	if isLeftMost {
		region.Perimeter++
	} else {
		west := CoordPair{i, j - 1}
		if !(*visited)[west] {
			westRegion := visited.surveyRegion(garden, west, crop)
			region.Area += westRegion.Area
			region.Perimeter += westRegion.Perimeter
		}
	}

	return region
}

func parseDay12Input(filePath string) Garden {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file at %s", filePath)
	}
	defer file.Close()

	garden := Garden{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		garden = append(garden, []rune(scanner.Text()))
	}

	return garden
}
