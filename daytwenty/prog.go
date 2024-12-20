package daytwenty

import (
	"adventofcode/models"
	"adventofcode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
)

type QueueElement struct {
	Postion models.Coords
	Cost    int
}

const (
	WALL  = '#'
	START = 'S'
	END   = 'E'
	PATH  = '.'
)

var (
	GRID_WIDTH  = 80
	GRID_HEIGHT = 80
)

func Solve() {
	maze := readInputFile("./daytwenty/input.txt")
	startPos := models.Coords{}
	endPOS := models.Coords{}
	wallPOSMap := make(map[string]bool)
	for i := range maze {
		for j := range maze[i] {
			char := maze[i][j]
			switch char {
			case WALL:
				wallPOSMap[utils.CoordsToString(i, j)] = true
			case START:
				startPos.X, startPos.Y = i, j
			case END:
				endPOS.X, endPOS.Y = i, j
			}
		}
	}
	GRID_WIDTH = len(maze)
	GRID_HEIGHT = len(maze[0])
	allMinCost, _ := findShortestPathUsingBFS(wallPOSMap, startPos, endPOS)
	fmt.Println("Min Cost", allMinCost, len(wallPOSMap))
	minCostMap := make(map[int]int)
	counter := 0
	for i := range wallPOSMap {
		delete(wallPOSMap, i)
		cost, _ := findShortestPathUsingBFS(wallPOSMap, startPos, endPOS)
		counter++
		minCostMap[allMinCost-cost] += 1
		wallPOSMap[i] = true
	}
	moreThan100SavingsCount := 0
	for k, v := range minCostMap {
		if k >= 100 {
			moreThan100SavingsCount += v
		}
	}
	log.Printf("The number of paths that save more than 100 steps is %d", moreThan100SavingsCount)
}

func findShortestPathUsingBFS(wallPOSMap map[string]bool, startPos, endPOS models.Coords) (int, map[string]bool) {
	visitedMap := make(map[string]bool)
	queue := make([]QueueElement, 0)
	queue = append(queue, QueueElement{Postion: startPos, Cost: 0})
	visitedMap[utils.CoordsToString(startPos.X, startPos.Y)] = true
	lowestCost := -1
	for len(queue) > 0 {
		currentElement := queue[0]
		queue = queue[1:]
		startPos = currentElement.Postion
		if startPos.IsAtPOS(endPOS) {
			return currentElement.Cost, visitedMap
		}

		possibleSides := [][]int{
			{0, -1}, {-1, 0}, {0, 1}, {1, 0},
		}
		for sideIndex := range possibleSides {
			side := possibleSides[sideIndex]
			newPos := startPos.MovePos(side[0], side[1])
			if !newPos.IsWithinBounds(GRID_WIDTH, GRID_HEIGHT) ||
				wallPOSMap[utils.CoordsToString(newPos.X, newPos.Y)] || visitedMap[utils.CoordsToString(newPos.X, newPos.Y)] {
				continue
			}
			visitedMap[utils.CoordsToString(newPos.X, newPos.Y)] = true
			queue = append(queue, QueueElement{Postion: newPos, Cost: currentElement.Cost + 1})
		}
	}
	return lowestCost, visitedMap
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
