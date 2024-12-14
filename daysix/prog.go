package daysix

import (
	"adventofcode/utils"
	"bufio"
	"log"
	"os"
	"time"
)

func Solve() {
	startTime := time.Now()
	inputMaze := readInputFile("./daysix/input.txt")
	caratPos := []int{0, 0}
	obstaclePos := make(map[string]struct{})
	movesMap := make(map[string]struct{})
	for i := range inputMaze {
		for j := range inputMaze[i] {
			if inputMaze[i][j] == '#' {
				obstaclePos[utils.CoordsToString(i, j)] = struct{}{}
			}
			if inputMaze[i][j] == '^' {
				caratPos[0], caratPos[1] = i, j
				movesMap[utils.CoordsToString(i, j)] = struct{}{}
			}
		}
	}
	mazeLength, mazeWidth := len(inputMaze), len(inputMaze[0])

	counterDx, counterDy := -1, 0
	dx, dy := -1, 0
	for {
		if moveCaratAndCountSteps(caratPos, dx, dy, mazeLength, mazeWidth, movesMap, obstaclePos) {
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
	}
	log.Println("Total Steps: ", len(movesMap))
	log.Println("Time Taken: ", time.Since(startTime))
}

func moveCaratAndCountSteps(caratPos []int, dx, dy, mazeLength, mazeWidth int, movesMap, obstaclePos map[string]struct{}) bool {
	for {
		x, y := caratPos[0], caratPos[1]
		nextX, nextY := x+dx, y+dy
		if nextX >= mazeLength || nextX < 0 || nextY >= mazeWidth || nextY < 0 {
			return true
		}
		if _, found := obstaclePos[utils.CoordsToString(nextX, nextY)]; found {
			return false
		}
		movesMap[utils.CoordsToString(nextX, nextY)] = struct{}{}
		caratPos[0], caratPos[1] = nextX, nextY
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
