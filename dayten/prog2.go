package dayten

import (
	"adventofcode/utils"
	"log"
)

func Solve2() {
	input := readInputFile("./dayten/input.txt")
	totalTrailScore := 0
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 0 {
				reachedHeightMap := make(map[string]bool)
				trailScore := findNextStepRating(input, reachedHeightMap, i, j)
				totalTrailScore += trailScore
			}
		}
	}
	log.Println("Total Trail Score: ", totalTrailScore)
}

func findNextStepRating(input [][]int, reachedHeightMap map[string]bool, i, j int) int {
	currentHeight := input[i][j]
	// if reachedHeightMap[utils.CoordsToString(i, j)] {
	// 	return 0
	// }
	if currentHeight == 9 {
		// reachedHeightMap[utils.CoordsToString(i, j)] = true
		return 1
	}
	totalTrails := 0
	possibleCoords := [][]int{
		{i, j - 1}, {i - 1, j}, {i, j + 1}, {i + 1, j},
	}
	for _, coords := range possibleCoords {
		x, y := coords[0], coords[1]
		if utils.IsCoordinatesValid(x, y, len(input), len(input[i])) {
			nextHeight := input[x][y]
			if nextHeight-currentHeight == 1 {
				totalTrails += findNextStepRating(input, reachedHeightMap, x, y)
			}
		}
	}
	return totalTrails
}
