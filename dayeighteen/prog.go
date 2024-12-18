package dayeighteen

import (
	"adventofcode/models"
	"adventofcode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	NUMBER_OF_BYTES = 1024
	GRID_WIDTH      = 71
	GRID_HEIGHT     = 71
)

type QueueElement struct {
	Postion models.Coords
	Cost    int
}

func Solve() {
	byteCoords := readInputFile("./dayeighteen/input.txt")
	corruptedPOS := make(map[string]bool)
	for i, coords := range byteCoords {
		if i == NUMBER_OF_BYTES {
			break
		}
		corruptedPOS[utils.CoordsToString(coords.X, coords.Y)] = true
	}
	cost, visitedNodes := (findShortestPathUsingBFS(corruptedPOS))
	fmt.Println("Cost:", cost)
	printGrid(corruptedPOS, visitedNodes)
}

func printGrid(corruptedPOS, visitedNodes map[string]bool) {

	for i := 0; i < GRID_HEIGHT; i++ {
		outStr := ""
		for j := 0; j < GRID_WIDTH; j++ {
			coordKey := utils.CoordsToString(i, j)
			if corruptedPOS[coordKey] {
				outStr += "#"
				continue
			}
			if visitedNodes[coordKey] {
				outStr += "*"
				continue
			}
			outStr += "."
		}
		fmt.Println(outStr)
	}
}

func findShortestPathUsingBFS(corruptedPOS map[string]bool) (int, map[string]bool) {
	visitedMap := make(map[string]bool)
	queue := make([]QueueElement, 0)
	currentPOS := models.Coords{X: 0, Y: 0}
	endPOS := models.Coords{X: GRID_HEIGHT - 1, Y: GRID_WIDTH - 1}
	queue = append(queue, QueueElement{Postion: currentPOS, Cost: 0})
	visitedMap[utils.CoordsToString(currentPOS.X, currentPOS.Y)] = true
	lowestCost := -1
	for len(queue) > 0 {
		currentElement := queue[0]
		queue = queue[1:]
		currentPOS = currentElement.Postion
		if currentPOS.IsAtPOS(endPOS) {
			return currentElement.Cost, visitedMap
		}

		possibleSides := [][]int{
			{0, -1}, {-1, 0}, {0, 1}, {1, 0},
		}
		for sideIndex := range possibleSides {
			side := possibleSides[sideIndex]
			newPos := currentPOS.MovePos(side[0], side[1])
			if !newPos.IsWithinBounds(GRID_WIDTH, GRID_HEIGHT) ||
				corruptedPOS[utils.CoordsToString(newPos.X, newPos.Y)] || visitedMap[utils.CoordsToString(newPos.X, newPos.Y)] {
				continue
			}
			visitedMap[utils.CoordsToString(newPos.X, newPos.Y)] = true
			queue = append(queue, QueueElement{Postion: newPos, Cost: currentElement.Cost + 1})
		}
	}
	return lowestCost, visitedMap
}

func readInputFile(fileName string) []models.Coords {
	inp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	input := make([]models.Coords, 0)
	for scanner.Scan() {
		line := scanner.Text()
		y, x := utils.StringToCoords(line)
		input = append(input, models.Coords{
			X: x,
			Y: y,
		})
	}
	return input
}
