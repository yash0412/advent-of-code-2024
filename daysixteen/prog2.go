package daysixteen

import (
	"adventofcode/utils"
	"fmt"
)

func Solve2() {
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
	CheckPoints(maze, wallPOSMap, currentPOS, endPOS)
}

func CheckPoints(maze [][]rune, wallPOSMap map[string]bool, currentPOS Reindeer, endPOS [2]int) {
	counter := 0
	fullCost, _ := solveMazeWithBFS(wallPOSMap, currentPOS, endPOS)
	fmt.Println("Full Cost", fullCost)
	for i := len(maze) - 1; i >= 0; i-- {
		for j := range maze[i] {
			char := maze[i][j]
			if char != WALL && char != START && char != END {
				fmt.Println("i,j9", i, j)
				cost1, lastPos := solveMazeWithBFS(wallPOSMap, currentPOS, [2]int{i, j})
				if cost1 == -1 {
					continue
				}
				cost2, _ := solveMazeWithBFS(wallPOSMap, lastPos, endPOS)
				if cost2 == -1 {
					continue
				}
				if cost1+cost2 == fullCost {
					fmt.Println("i,j", i, j)
					counter++
				}
			}
		}
	}
	fmt.Println("Counter", counter+2)
}
