package daytwelve

import (
	"adventofcode/utils"
	"fmt"
	"log"
)

func Solve2() {
	input := readInputFile("./daytwelve/input.txt")
	areas := make([]int, 0)
	sides := make([]int, 0)
	visitedMap := make(map[string]bool)
	sidesMap := make(map[string]bool)
	// sideVisitedMap := make(map[string]bool)
	for i := range input {
		for j := range input[i] {
			area := 0
			side := 0
			DFS2(visitedMap, sidesMap, input, i, j, &area, &side)
			// side = 0
			// countSides(sidesMap, sideVisitedMap, input, i, j, len(input), len(input[i]), &side)
			areas = append(areas, area)
			sides = append(sides, side)
		}
	}

	totalCost := 0

	for i := range areas {
		totalCost += areas[i] * sides[i]
	}

	log.Println("Total sides: ", sides)
	log.Println("Total aread: ", areas)
	log.Println("Total fence cost: ", totalCost)
	log.Println("Total sides: ", len(sidesMap))
}

func DFS2(visitedMap, sidesMap map[string]bool, input [][]rune, i, j int, area, side *int) {
	if visitedMap[utils.CoordsToString(i, j)] {
		return
	}
	visitedMap[utils.CoordsToString(i, j)] = true
	*area++

	possibleCoords := [][]int{
		{i, j - 1}, {i - 1, j}, {i, j + 1}, {i + 1, j},
	}

	for _, coords := range possibleCoords {
		x, y := coords[0], coords[1]
		dx, dy := x-i, y-j
		char := input[i][j]
		if utils.IsCoordinatesValid(x, y, len(input), len(input[i])) && input[x][y] == char {
			DFS2(visitedMap, sidesMap, input, x, y, area, side)
		} else {
			sidesMap[getSidesMapKey(i, j, dx, dy)] = true
			if dx == 0 {
				topAndBottomElemCoords := [][]int{
					{i - 1, j}, {i + 1, j},
				}
				shouldAddSide := true
				for _, tbc := range topAndBottomElemCoords {
					if utils.IsCoordinatesValid(tbc[0], tbc[1], len(input), len(input[i])) &&
						input[tbc[0]][tbc[1]] == char && sidesMap[getSidesMapKey(tbc[0], tbc[1], dx, dy)] {
						shouldAddSide = false
					}
				}

				if shouldAddSide {
					fmt.Println(char, i, j, dx, dy)
					*side++
				}
			}
			if dy == 0 {
				topAndBottomElemCoords := [][]int{
					{i, j - 1}, {i, j + 1},
				}

				shouldAddSide := true
				for _, tbc := range topAndBottomElemCoords {
					if utils.IsCoordinatesValid(tbc[0], tbc[1], len(input), len(input[i])) &&
						input[tbc[0]][tbc[1]] == char && sidesMap[getSidesMapKey(tbc[0], tbc[1], dx, dy)] {
						shouldAddSide = false
					}
				}

				if shouldAddSide {
					fmt.Println(char, i, j, dx, dy)
					*side++
				}
			}
		}
	}
}

func countSides(sidesMap, sideVisitedMap map[string]bool, input [][]rune, i, j, xsize, ysize int, side *int) {
	possibleSides := [][]int{
		{0, -1}, {-1, 0}, {0, 1}, {1, 0},
	}

	for _, possibleSide := range possibleSides {
		dx, dy := possibleSide[0], possibleSide[1]
		sideKey := getSidesMapKey(i, j, dx, dy)
		if sideVisitedMap[sideKey] || !sidesMap[sideKey] {
			continue
		}
		sideVisitedMap[sideKey] = true
		*side++
		x, y := i, j
		currentChar := input[i][j]
		for {
			x, y = x+dx, y+dy
			if !utils.IsCoordinatesValid(x, y, xsize, ysize) {
				break
			}
			nextChar := input[x][y]
			nextCharSideKey := getSidesMapKey(x, y, dx, dy)
			if nextChar == currentChar && sidesMap[nextCharSideKey] {
				sideVisitedMap[nextCharSideKey] = true

			} else {
				break
			}
		}
	}
	for _, possibleSide := range possibleSides {
		dx, dy := possibleSide[0], possibleSide[1]
		x, y := i+dx, j+dy

		if !utils.IsCoordinatesValid(x, y, xsize, ysize) {
			continue
		}
		countSides(sidesMap, sideVisitedMap, input, x, y, xsize, ysize, side)
	}

}

func getSidesMapKey(x, y, dx, dy int) string {
	return fmt.Sprintf("%s-%s", utils.CoordsToString(x, y), utils.CoordsToString(dx, dy))
}
