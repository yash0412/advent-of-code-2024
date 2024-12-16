package daysixteen

import (
	"adventofcode/models"
	"adventofcode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
)

var counter = 0

type Reindeer struct {
	Position  models.Coords
	Direction models.Coords
}

func (r Reindeer) isAtPOS(pos [2]int) bool {
	return r.Position.X == pos[0] && r.Position.Y == pos[1]
}

func (r Reindeer) isOnAWall(wallPOSMap map[string]bool) bool {
	return wallPOSMap[utils.CoordsToString(r.Position.X, r.Position.Y)]
}

func (r Reindeer) moveReindeer(dx, dy int) Reindeer {
	return Reindeer{
		Position: models.Coords{
			X: r.Position.X + dx,
			Y: r.Position.Y + dy,
		},
		Direction: models.Coords{
			X: dx, Y: dy,
		},
	}
}

func (r Reindeer) isEqual(reindeer Reindeer) bool {
	return r.Position.X == reindeer.Position.X && r.Position.Y == reindeer.Position.Y &&
		r.Direction.X == reindeer.Direction.X && r.Direction.Y == reindeer.Direction.Y
}

func (r Reindeer) calculateMoveCost(targetPos Reindeer) int {
	if r.isEqual(targetPos) {
		return 0
	}
	cost := 0
	pdx, pdy := targetPos.Position.X-r.Position.X, targetPos.Position.Y-r.Position.Y
	if absoluteSum(pdx, pdy) == 1 {
		cost += 1
	}
	ddx, ddy := targetPos.Direction.X-r.Direction.X, targetPos.Direction.Y-r.Direction.Y
	if ddx == 0 && ddy == 0 {
	} else if ddx == 0 || ddy == 0 {
		cost += 2000
	} else {
		cost += 1000
	}
	return cost
}

func absoluteSum(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a + b
}

const (
	WALL  = '#'
	START = 'S'
	END   = 'E'
	PATH  = '.'
)

func Solve() {
	maze := readInputFile("./daysixteen/input.txt")
	currentPOS := Reindeer{}
	wallPOSMap := make(map[string]bool)
	endPOS := [2]int{}
	visitedMap := make(map[string]bool)
	for i := range maze {
		for j := range maze[i] {
			char := maze[i][j]
			switch char {
			case WALL:
				wallPOSMap[utils.CoordsToString(i, j)] = true
			case START:
				currentPOS.Position.X, currentPOS.Position.Y = i, j
				currentPOS.Direction.X, currentPOS.Direction.Y = 0, 1
			case END:
				endPOS[0] = i
				endPOS[1] = j
			}
		}
	}
	solved, cost := solveMaze(wallPOSMap, visitedMap, currentPOS, endPOS)
	if solved {
		log.Println("Min Cost:", cost)
	}
}

func solveMaze(wallPOSMap, visitedMap map[string]bool, currentPOS Reindeer, endPOS [2]int) (bool, int) {
	if currentPOS.isAtPOS(endPOS) {
		return true, 0
	}
	if currentPOS.isOnAWall(wallPOSMap) {
		return false, 0
	}
	if visitedMap[utils.CoordsToString(currentPOS.Position.X, currentPOS.Position.Y)] {
		return false, 0
	}
	visitedMap[utils.CoordsToString(currentPOS.Position.X, currentPOS.Position.Y)] = true
	counter++
	possibleSides := [][]int{
		{0, -1}, {-1, 0}, {0, 1}, {1, 0},
	}
	if counter%10000 == 0 {
		fmt.Println("Counter:", counter, "Current POS:", currentPOS.Position)
	}
	minCost := 0
	solved := false
	for sideIndex := range possibleSides {
		side := possibleSides[sideIndex]
		newPos := currentPOS.moveReindeer(side[0], side[1])
		moveCost := currentPOS.calculateMoveCost(newPos)
		branchSolved, cost := solveMaze(wallPOSMap, copyVisitedMap(visitedMap), newPos, endPOS)
		if branchSolved {
			solved = true
			if minCost == 0 || cost+moveCost < minCost {
				minCost = cost + moveCost
			}
		}
	}

	return solved, minCost
}

func copyVisitedMap(visitedMap map[string]bool) map[string]bool {
	newMap := make(map[string]bool)
	for key, value := range visitedMap {
		newMap[key] = value
	}
	return newMap
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
