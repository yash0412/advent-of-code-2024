package dayfourteen

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Width              = 101
	Length             = 103
	NumberOfSeconds    = 100
	MinNumberOfSeconds = 10000
)

type Coords struct {
	X int
	Y int
}

type Robot struct {
	Position Coords
	Velocity Coords
}

func (robot *Robot) MoveRobot() {
	newX, newY := robot.Position.X+robot.Velocity.X, robot.Position.Y+robot.Velocity.Y
	if newX < 0 {
		newX = Width + newX
	}
	if newX >= Width {
		newX = newX - Width
	}
	if newY < 0 {
		newY = Length + newY
	}
	if newY >= Length {
		newY = newY - Length
	}
	newCoords := Coords{X: newX, Y: newY}
	robot.Position = newCoords
}

func Solve() {
	input := readInputFile("./dayfourteen/input.txt")
	robots := make([]Robot, 0)
	for _, val := range input {
		robot := Robot{}
		valSplit := strings.Split(val, " ")
		posSplit := strings.Split(valSplit[0], "=")
		robot.Position.X, robot.Position.Y = utils.StringToCoords(posSplit[1])
		velSplit := strings.Split(valSplit[1], "=")
		robot.Velocity.X, robot.Velocity.Y = utils.StringToCoords(velSplit[1])
		robots = append(robots, robot)
	}

	for i := 0; i < NumberOfSeconds; i++ {
		for index := range robots {
			robots[index].MoveRobot()
		}
	}

	quandrants := make([]int, 4)
	for index := range robots {
		if quad := getQuadrantIndexForCoord(robots[index].Position); quad > -1 {
			quandrants[quad]++
		}
	}
	log.Println("Quads: ", quandrants[0]*quandrants[1]*quandrants[2]*quandrants[3])
}

func Solve2() {
	input := readInputFile("./dayfourteen/input.txt")
	robots := make([]Robot, 0)
	for _, val := range input {
		robot := Robot{}
		valSplit := strings.Split(val, " ")
		posSplit := strings.Split(valSplit[0], "=")
		robot.Position.X, robot.Position.Y = utils.StringToCoords(posSplit[1])
		velSplit := strings.Split(valSplit[1], "=")
		robot.Velocity.X, robot.Velocity.Y = utils.StringToCoords(velSplit[1])
		robots = append(robots, robot)
	}

	for i := 0; i < MinNumberOfSeconds; i++ {
		robotPosMap := make(map[Coords]bool)
		for index := range robots {
			robots[index].MoveRobot()
			robotPosMap[robots[index].Position] = true
		}
		generateOutputFile(robotPosMap, i)
	}

}

func generateOutputFile(robotPosMap map[Coords]bool, iteration int) {
	outputStr := ""
	for i := 0; i < Width; i++ {
		for j := 0; j < Length; j++ {
			if robotPosMap[Coords{X: i, Y: j}] {
				outputStr += "1"
			} else {
				outputStr += "."
			}
		}
		outputStr += "\n"
	}
	if iteration == 6511 {
		os.WriteFile(fmt.Sprintf("./dayfourteen/renders/output-%d.txt", iteration+1), []byte(outputStr), os.ModeAppend)
	}
}

func getQuadrantIndexForCoord(coords Coords) int {
	midX := Width / 2
	midY := Length / 2
	if coords.X < midX {
		if coords.Y < midY {
			return 0
		}
		if coords.Y > midY {
			return 2
		}
	}
	if coords.X > midX {
		if coords.Y < midY {
			return 1
		}
		if coords.Y > midY {
			return 3
		}
	}
	return -1
}

func readInputFile(fileName string) []string {
	inp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	input := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	return input
}
