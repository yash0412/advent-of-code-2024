package daysixteen

import (
	"adventofcode/models"
	"adventofcode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
)

type QueueElement struct {
	Reindeer Reindeer
	Cost     int
}

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
	cost := solveMazeWithBFS(wallPOSMap, currentPOS, endPOS)
	log.Println("Min Cost:", cost)
}

func solveMazeWithBFS(wallPOSMap map[string]bool, currentPOS Reindeer, endPOS [2]int) int {
	visitedMap := make(map[string]bool)
	queue := make([]QueueElement, 0)
	queue = append(queue, QueueElement{Reindeer: currentPOS, Cost: 0})
	visitedMap[getSidesMapKey(currentPOS.Position.X, currentPOS.Position.Y, currentPOS.Direction.X, currentPOS.Direction.Y)] = true
	lowestCost := -1
	for len(queue) > 0 {
		currentElement := queue[0]
		queue = queue[1:]
		currentPOS = currentElement.Reindeer
		if currentPOS.isAtPOS(endPOS) {
			return currentElement.Cost
		}

		possibleSides := [][]int{
			{0, -1}, {-1, 0}, {0, 1}, {1, 0},
		}
		for sideIndex := range possibleSides {
			side := possibleSides[sideIndex]
			newPos := currentPOS.moveReindeer(side[0], side[1])
			if newPos.isOnAWall(wallPOSMap) || visitedMap[getSidesMapKey(newPos.Position.X, newPos.Position.Y, newPos.Direction.X, newPos.Direction.Y)] {
				continue
			}
			visitedMap[getSidesMapKey(currentPOS.Position.X, currentPOS.Position.Y, currentPOS.Direction.X, currentPOS.Direction.Y)] = true
			moveCost := currentPOS.calculateMoveCost(newPos)
			if moveCost >= 2000 {
				continue
			}
			queue = append(queue, QueueElement{Reindeer: newPos, Cost: currentElement.Cost + moveCost})
		}
		queue = sortQueueElements(queue)
	}
	return lowestCost
}

func getSidesMapKey(x, y, dx, dy int) string {
	return fmt.Sprintf("%s-%s", utils.CoordsToString(x, y), utils.CoordsToString(dx, dy))
}

func sortQueueElements(queue []QueueElement) []QueueElement {
	// use bubble sort
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(queue)-1; i++ {
			if queue[i].Cost > queue[i+1].Cost {
				queue[i], queue[i+1] = queue[i+1], queue[i]
				sorted = false

			}
		}
	}
	return queue
}

func solveMaze(wallPOSMap, visitedMap map[string]bool, currentPOS Reindeer, endPOS [2]int) (bool, int) {
	if currentPOS.isAtPOS(endPOS) {
		return true, 0
	}
	visitedMap[utils.CoordsToString(currentPOS.Position.X, currentPOS.Position.Y)] = true
	possibleSides := [][]int{
		{0, -1}, {-1, 0}, {0, 1}, {1, 0},
	}
	minCost := 0
	solved := false
	for sideIndex := range possibleSides {
		side := possibleSides[sideIndex]
		newPos := currentPOS.moveReindeer(side[0], side[1])
		if newPos.isOnAWall(wallPOSMap) || visitedMap[utils.CoordsToString(newPos.Position.X, newPos.Position.Y)] {
			continue
		}
		moveCost := currentPOS.calculateMoveCost(newPos)
		branchSolved, cost := solveMaze(wallPOSMap, (visitedMap), newPos, endPOS)
		if branchSolved {
			solved = true
			if minCost == 0 || cost+moveCost < minCost {
				minCost = cost + moveCost
			}
		}
	}
	visitedMap[utils.CoordsToString(currentPOS.Position.X, currentPOS.Position.Y)] = false
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
