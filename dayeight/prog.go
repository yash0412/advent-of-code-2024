package dayeight

import (
	"adventofcode/utils"
	"bufio"
	"log"
	"os"
)

func Solve() {
	input := readInputFile("./dayeight/input.txt")
	antennaPos := make(map[rune]map[string]bool)
	for i := range input {
		for j := range input[i] {
			char := input[i][j]
			if char != '.' {
				if _, exists := antennaPos[char]; !exists {
					antennaPos[char] = make(map[string]bool)
				}
				antennaPos[char][utils.CoordsToString(i, j)] = true
			}
		}
	}
	antinodePos := make(map[string]bool)
	for antennaType := range antennaPos {
		for antenna1 := range antennaPos[antennaType] {
			for antenna2 := range antennaPos[antennaType] {
				if antenna1 == antenna2 {
					continue
				}
				antenna1x, antenna1y := utils.StringToCoords(antenna1)
				antenna2x, antenna2y := utils.StringToCoords(antenna2)
				dx, dy := getCoordsDiff(antenna1x, antenna1y, antenna2x, antenna2y)
				antinode1x, antinode1y := antenna1x-dx, antenna1y-dy
				antinode2x, antinode2y := antenna2x+dx, antenna2y+dy
				if utils.IsCoordinatesValid(antinode1x, antinode1y, len(input), len(input[0])) {
					antinodePos[utils.CoordsToString(antinode1x, antinode1y)] = true
				}
				if utils.IsCoordinatesValid(antinode2x, antinode2y, len(input), len(input[0])) {
					antinodePos[utils.CoordsToString(antinode2x, antinode2y)] = true
				}
			}
		}
	}
	log.Println("Antinodes: ", len(antinodePos))
}

func Solve2() {
	input := readInputFile("./dayeight/input.txt")
	antennaPos := make(map[rune]map[string]bool)
	for i := range input {
		for j := range input[i] {
			char := input[i][j]
			if char != '.' {
				if _, exists := antennaPos[char]; !exists {
					antennaPos[char] = make(map[string]bool)
				}
				antennaPos[char][utils.CoordsToString(i, j)] = true
			}
		}
	}
	antinodePos := make(map[string]bool)
	for antennaType := range antennaPos {
		for antenna1 := range antennaPos[antennaType] {
			for antenna2 := range antennaPos[antennaType] {
				if antenna1 == antenna2 {
					continue
				}
				antenna1x, antenna1y := utils.StringToCoords(antenna1)
				antenna2x, antenna2y := utils.StringToCoords(antenna2)
				antinodePos[utils.CoordsToString(antenna1x, antenna1y)] = true
				antinodePos[utils.CoordsToString(antenna2x, antenna2y)] = true
				dx, dy := getCoordsDiff(antenna1x, antenna1y, antenna2x, antenna2y)
				currAntenna1x, currAntenna1y := antenna1x, antenna1y
				for {
					antinode1x, antinode1y := currAntenna1x-dx, currAntenna1y-dy
					if utils.IsCoordinatesValid(antinode1x, antinode1y, len(input), len(input[0])) {
						antinodePos[utils.CoordsToString(antinode1x, antinode1y)] = true
						currAntenna1x, currAntenna1y = antinode1x, antinode1y
					} else {
						break
					}
				}
				currAntenna2x, currAntenna2y := antenna2x, antenna2y
				for {
					antinode2x, antinode2y := currAntenna2x-dx, currAntenna2y-dy
					if utils.IsCoordinatesValid(antinode2x, antinode2y, len(input), len(input[0])) {
						antinodePos[utils.CoordsToString(antinode2x, antinode2y)] = true
						currAntenna2x, currAntenna2y = antinode2x, antinode2y
					} else {
						break
					}
				}
			}
		}
	}
	log.Println("Antinodes: ", len(antinodePos))
}

func getCoordsDiff(coords1x, coords1y, coords2x, coords2y int) (int, int) {
	return coords2x - coords1x, coords2y - coords1y
}

func readInputFile(fileName string) [][]rune {
	inp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	input := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, []rune(line))
	}
	return input
}
