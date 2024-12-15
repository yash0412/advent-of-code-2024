package daytwelve

import (
	"adventofcode/utils"
	"bufio"
	"log"
	"os"
)

func Solve() {
	input := readInputFile("./daytwelve/input.txt")
	areas := make([]int, 0)
	perimeters := make([]int, 0)
	visitedMap := make(map[string]bool)
	for i := range input {
		for j := range input[i] {
			area := 0
			perimeter := 0
			DFS(visitedMap, input, i, j, &area, &perimeter)
			areas = append(areas, area)
			perimeters = append(perimeters, perimeter)
		}
	}

	totalCost := 0
	perimeterSum := 0
	for i := range areas {
		totalCost += areas[i] * perimeters[i]
		perimeterSum += perimeters[i]
	}

	log.Println("Total fence cost: ", totalCost, perimeterSum)
}

func DFS(visitedMap map[string]bool, input [][]rune, i, j int, area, perimeter *int) {
	if visitedMap[utils.CoordsToString(i, j)] {
		return
	}
	visitedMap[utils.CoordsToString(i, j)] = true
	*area++
	if i == 0 {
		*perimeter++
	}
	if j == 0 {
		*perimeter++
	}
	if i == len(input)-1 {
		*perimeter++
	}
	if j == len(input[i])-1 {
		*perimeter++
	}

	possibleCoords := [][]int{
		{i, j - 1}, {i - 1, j}, {i, j + 1}, {i + 1, j},
	}

	for _, coords := range possibleCoords {
		x, y := coords[0], coords[1]
		if utils.IsCoordinatesValid(x, y, len(input), len(input[i])) {
			char := input[i][j]
			if input[x][y] == char {
				DFS(visitedMap, input, x, y, area, perimeter)
			} else {
				*perimeter++
			}
		}
	}
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
