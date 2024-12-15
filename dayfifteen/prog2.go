package dayfifteen

import (
	"adventofcode/utils"
	"fmt"
	"log"
)

const (
	DoubleWall       = "##"
	DoubleBox        = "[]"
	DoubleEmptySpace = ".."
	DoubleRobot      = "@."
	BoxLeftSide      = '['
	BoxRightSide     = ']'
)

func Solve2() {
	layout1, moves := readInputFile("./dayfifteen/input.txt")
	newLayout := [][]rune{}
	for i := range layout1 {
		layoutStr := ""
		for j := range layout1[i] {
			char := layout1[i][j]

			switch char {
			case Wall:
				layoutStr += DoubleWall
			case Robot:
				layoutStr += DoubleRobot
			case Box:
				layoutStr += DoubleBox
			case EmptySpace:
				layoutStr += DoubleEmptySpace
			}
		}
		newLayout = append(newLayout, []rune(layoutStr))
	}
	robotPos := [2]int{0, 0}
	for i := range newLayout {
		for j := range newLayout[i] {
			char := newLayout[i][j]

			if char == Robot {
				robotPos[0] = i
				robotPos[1] = j
			}
		}
	}
	clearOutputFile(newLayout, robotPos)
	lastMove := UpMove
	for index, move := range moves {
		robotPos = moveRobot2(newLayout, robotPos, move)
		lastMove = move
		fmt.Printf("Moved %d of %d\n", index+1, len(moves))
	}

	boxTotal := 0
	for i := range newLayout {
		for j := range newLayout[i] {
			char := newLayout[i][j]
			switch char {
			case BoxLeftSide:
				boxTotal += 100*i + j
			}
		}
	}
	log.Println("Total GPS: ", boxTotal)
	printLayout(newLayout, robotPos, lastMove)
}

func moveRobot2(layout [][]rune, robotPos [2]int, move rune) [2]int {
	moveDirection := moveToDirectionMap[move]
	nextX, nextY := getNextCoordsInDirection(robotPos[0], robotPos[1], moveDirection[0], moveDirection[1])

	if layout[nextX][nextY] == Wall {
		return robotPos
	}

	if layout[nextX][nextY] == EmptySpace {
		layout[nextX][nextY] = Robot
		layout[robotPos[0]][robotPos[1]] = EmptySpace
		robotPos[0], robotPos[1] = nextX, nextY
	}

	if layout[nextX][nextY] == BoxLeftSide || layout[nextX][nextY] == BoxRightSide {
		if move == LeftMove || move == RightMove {
			emptySpaceX, emptySpaceY := -1, -1
			lookbackX, lookbackY := nextX, nextY
			for {
				lookbackX, lookbackY = getNextCoordsInDirection(lookbackX, lookbackY, moveDirection[0], moveDirection[1])
				if layout[lookbackX][lookbackY] == EmptySpace {
					emptySpaceX, emptySpaceY = lookbackX, lookbackY
					break
				}

				if layout[lookbackX][lookbackY] == Wall {
					break
				}
			}

			if emptySpaceX != -1 && emptySpaceY != -1 {
				moveBoxWithinRow(layout, emptySpaceX, emptySpaceY, nextY)
				layout[nextX][nextY] = Robot
				layout[robotPos[0]][robotPos[1]] = EmptySpace
				robotPos[0], robotPos[1] = nextX, nextY
			}
		} else {
			boxOtherHalfCoordsX, boxOtherHalfCoordsY := -1, -1
			if layout[nextX][nextY] == BoxLeftSide {
				boxOtherHalfCoordsX, boxOtherHalfCoordsY = nextX, nextY+1
			} else if layout[nextX][nextY] == BoxRightSide {
				boxOtherHalfCoordsX, boxOtherHalfCoordsY = nextX, nextY-1
			}
			boxCoordinates := [][2]int{{nextX, nextY}, {boxOtherHalfCoordsX, boxOtherHalfCoordsY}}
			boxAddedMap := make(map[string]bool)
			boxAddedMap[utils.CoordsToString(nextX, nextY)] = true
			boxAddedMap[utils.CoordsToString(boxOtherHalfCoordsX, boxOtherHalfCoordsY)] = true
			index := 0
			for {
				coords := boxCoordinates[index]
				nextBoxX, nextBoxY := getNextCoordsInDirection(coords[0], coords[1], moveDirection[0], moveDirection[1])

				if layout[nextBoxX][nextBoxY] == Wall {
					break
				}

				if !boxAddedMap[utils.CoordsToString(nextBoxX, nextBoxY)] {
					if layout[nextBoxX][nextBoxY] == BoxLeftSide {
						nextBoxOtherHalfCoordsX, nextBoxOtherHalfCoordsY := nextBoxX, nextBoxY+1
						boxCoordinates = append(boxCoordinates, [2]int{nextBoxX, nextBoxY}, [2]int{nextBoxOtherHalfCoordsX, nextBoxOtherHalfCoordsY})
						boxAddedMap[utils.CoordsToString(nextBoxX, nextBoxY)] = true
						boxAddedMap[utils.CoordsToString(nextBoxOtherHalfCoordsX, nextBoxOtherHalfCoordsY)] = true
					} else if layout[nextBoxX][nextBoxY] == BoxRightSide {
						nextBoxOtherHalfCoordsX, nextBoxOtherHalfCoordsY := nextBoxX, nextBoxY-1
						boxCoordinates = append(boxCoordinates, [2]int{nextBoxX, nextBoxY}, [2]int{nextBoxOtherHalfCoordsX, nextBoxOtherHalfCoordsY})
						boxAddedMap[utils.CoordsToString(nextBoxX, nextBoxY)] = true
						boxAddedMap[utils.CoordsToString(nextBoxOtherHalfCoordsX, nextBoxOtherHalfCoordsY)] = true
					}
				}

				if index == len(boxCoordinates)-1 {
					for ; index >= 0; index-- {
						coords := boxCoordinates[index]
						nextSpaceX, nextSpaceY := getNextCoordsInDirection(coords[0], coords[1], moveDirection[0], moveDirection[1])
						layout[nextSpaceX][nextSpaceY] = layout[coords[0]][coords[1]]
						layout[coords[0]][coords[1]] = EmptySpace
					}
					layout[nextX][nextY] = Robot
					layout[robotPos[0]][robotPos[1]] = EmptySpace
					robotPos[0], robotPos[1] = nextX, nextY
					break
				}
				index++
			}

		}
	}

	return robotPos
}

func moveBoxWithinRow(layout [][]rune, newX, newY, oldY int) {
	row := layout[newX]

	row = append(row[:newY], row[newY+1:]...)
	newSlice := make([]rune, oldY+1)
	copy(newSlice, row[:oldY])
	newSlice[oldY] = EmptySpace

	row = append(newSlice, row[oldY:]...)
	layout[newX] = []rune(row)
}
