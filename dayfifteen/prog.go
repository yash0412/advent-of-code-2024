package dayfifteen

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Wall       = '#'
	WallStr    = "#"
	Robot      = '@'
	Box        = 'O'
	EmptySpace = '.'
	UpMove     = '^'
	DownMove   = 'v'
	LeftMove   = '<'
	RightMove  = '>'
)

var moveToDirectionMap = map[rune][2]int{
	UpMove:    {-1, 0},
	DownMove:  {1, 0},
	LeftMove:  {0, -1},
	RightMove: {0, 1},
}

func Solve() {
	layout, moves := readInputFile("./dayfifteen/input.txt")
	wallPOSMap := make(map[string]bool)
	boxPOSMap := make(map[string]bool)
	robotPos := [2]int{0, 0}

	for i := range layout {
		for j := range layout[i] {
			char := layout[i][j]

			switch char {
			case Wall:
				wallPOSMap[utils.CoordsToString(i, j)] = true
			case Robot:
				robotPos[0] = i
				robotPos[1] = j
			case Box:
				boxPOSMap[utils.CoordsToString(i, j)] = true
			}
		}
	}

	for _, move := range moves {
		robotPos = moveRobot(layout, wallPOSMap, boxPOSMap, robotPos, move)
		// printLayout(layout, move)
		// fmt.Println(robotPos)
	}

	boxTotal := 0
	for i := range layout {
		for j := range layout[i] {
			char := layout[i][j]
			switch char {
			case Box:
				boxTotal += 100*i + j
			}
		}
	}
	log.Println("Total GPS: ", boxTotal)
}

func moveRobot(layout [][]rune, wallPOSMap, boxPOSMap map[string]bool, robotPos [2]int, move rune) [2]int {
	moveDirection := moveToDirectionMap[move]
	nextX, nextY := getNextCoordsInDirection(robotPos[0], robotPos[1], moveDirection[0], moveDirection[1])

	if layout[nextX][nextY] == EmptySpace {
		layout[nextX][nextY] = Robot
		layout[robotPos[0]][robotPos[1]] = EmptySpace
		robotPos[0], robotPos[1] = nextX, nextY
	}

	if wallPOSMap[utils.CoordsToString(nextX, nextY)] {
		return robotPos
	}

	if boxPOSMap[utils.CoordsToString(nextX, nextY)] {
		emptySpaceX, emptySpaceY := -1, -1
		lookbackX, lookbackY := nextX, nextY
		for {
			lookbackX, lookbackY = getNextCoordsInDirection(lookbackX, lookbackY, moveDirection[0], moveDirection[1])
			if layout[lookbackX][lookbackY] == EmptySpace {
				emptySpaceX, emptySpaceY = lookbackX, lookbackY
				break
			}

			if wallPOSMap[utils.CoordsToString(lookbackX, lookbackY)] {
				break
			}
		}

		if emptySpaceX != -1 && emptySpaceY != -1 {
			layout[emptySpaceX][emptySpaceY] = Box
			boxPOSMap[utils.CoordsToString(emptySpaceX, emptySpaceY)] = true
			delete(boxPOSMap, utils.CoordsToString(nextX, nextY))
			layout[nextX][nextY] = Robot
			layout[robotPos[0]][robotPos[1]] = EmptySpace
			robotPos[0], robotPos[1] = nextX, nextY
		}
	}
	return robotPos
}

func getNextCoordsInDirection(x, y, dx, dy int) (int, int) {
	return x + dx, y + dy
}

func printLayout(layout [][]rune, move rune) {
	layoutStr := ""
	for i := range layout {
		for j := range layout[i] {
			char := layout[i][j]
			layoutStr += (string(char))
		}
		layoutStr += ("\n")
	}
	fmt.Println("Move made: ", string(move))
	fmt.Println(layoutStr)
}

func readInputFile(fileName string) ([][]rune, string) {
	inp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	layout := make([][]rune, 0)
	moves := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, WallStr) {
			layout = append(layout, []rune(line))
		} else {
			moves += line
		}
	}
	return layout, moves
}
