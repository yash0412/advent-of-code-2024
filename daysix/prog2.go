package daysix

import (
	"fmt"
	"log"
	"time"
)

func Solve2() {
	startTime := time.Now()
	inputMaze := readInputFile("./daysix/input.txt")
	caratPos := []int{0, 0}
	obstaclePos := make(map[string]struct{})

	for i := range inputMaze {
		for j := range inputMaze[i] {
			if inputMaze[i][j] == '#' {
				obstaclePos[coordsToString(i, j)] = struct{}{}
			}
			if inputMaze[i][j] == '^' {
				caratPos[0], caratPos[1] = i, j
			}
		}
	}
	counter := 0
	mazeLength, mazeWidth := len(inputMaze), len(inputMaze[0])
	for i := 0; i < mazeLength; i++ {
		for j := 0; j < mazeWidth; j++ {
			if i == caratPos[0] && j == caratPos[1] {
				continue
			}
			if _, exists := obstaclePos[coordsToString(i, j)]; exists {
				continue
			}
			encounteredObstacles := make(map[string]struct{})
			obstaclePos[coordsToString(i, j)] = struct{}{}
			movesMap := make(map[string]struct{})
			movesMap[coordsToString(caratPos[0], caratPos[1])] = struct{}{}
			counterDx, counterDy := -1, 0
			dx, dy := -1, 0
			oldMovesCount := len(movesMap)
			oldEOCount := len(encounteredObstacles)
			newCaratPos := []int{caratPos[0], caratPos[1]}
			for {
				if moveCaratAndCountStepsAndObs(newCaratPos, dx, dy, mazeLength, mazeWidth, movesMap, encounteredObstacles, obstaclePos) {
					break
				}
				counterDx, counterDy = (counterDx + 1), (counterDy + 1)
				if counterDx%2 == 0 {
					dx = 0
				} else if (counterDx/2)%2 == 0 {
					dx = 1
				} else {
					dx = -1
				}
				if counterDy%2 == 0 {
					dy = 0
				} else if (counterDy/2)%2 == 0 {
					dy = 1
				} else {
					dy = -1
				}

				if oldMovesCount == len(movesMap) && (oldEOCount == len(encounteredObstacles)) {
					counter++
					break
				}
				oldMovesCount = len(movesMap)
				oldEOCount = len(encounteredObstacles)
			}
			delete(obstaclePos, coordsToString(i, j))
		}
	}
	log.Println("Counted: ", counter)
	log.Println("Time Taken: ", time.Since(startTime))
}

func moveCaratAndCountStepsAndObs(caratPos []int, dx, dy, mazeLength, mazeWidth int, movesMap, encounteredObstacles, obstaclePos map[string]struct{}) bool {
	for {
		x, y := caratPos[0], caratPos[1]
		nextX, nextY := x+dx, y+dy
		if nextX >= mazeLength || nextX < 0 || nextY >= mazeWidth || nextY < 0 {
			return true
		}
		if _, found := obstaclePos[coordsToString(nextX, nextY)]; found {
			encounteredObstacles[fmt.Sprintf("%s,%d,%d", coordsToString(nextX, nextY), dx, dy)] = struct{}{}
			return false
		}
		movesMap[coordsToString(nextX, nextY)] = struct{}{}
		caratPos[0], caratPos[1] = nextX, nextY
	}
}
