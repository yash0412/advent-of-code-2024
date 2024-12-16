package daythirteen

import (
	"adventofcode/models"
	"adventofcode/utils"
	"bufio"
	"log"
	"os"
	"strings"
)

type Machine struct {
	A     models.Coords
	B     models.Coords
	Prize models.Coords
}

func Solve() {
	machines := readInputFile("./daythirteen/input.txt")
	totalWin := 0
	for index := range machines {
		A, B := findAandB(machines[index])
		if A == B && B == 0 || (A >= 100 || B >= 100) {
			continue
		}
		totalWin += A*3 + B*1
	}
	log.Println("Total Prize: ", totalWin)
}

func Solve2() {
	machines := readInputFile("./daythirteen/input.txt")
	totalWin := 0
	for index := range machines {
		machines[index].Prize.X += 10000000000000
		machines[index].Prize.Y += 10000000000000
		A, B := findAandB(machines[index])
		if A == B && B == 0 {
			continue
		}
		totalWin += A*3 + B*1
	}
	log.Println("Total Prize: ", totalWin)
}

func findAandB(machine Machine) (int, int) {
	X := machine.Prize.X
	Y := machine.Prize.Y

	x1 := machine.A.X
	y1 := machine.A.Y

	x2 := machine.B.X
	y2 := machine.B.Y

	base := (x2*y1 - x1*y2)
	Aval := (Y*x2 - X*y2)
	Bval := (X*y1 - Y*x1)

	if Aval%base == 0 && Bval%base == 0 {
		return Aval / base, Bval / base
	}

	return 0, 0
}

func readInputFile(fileName string) []Machine {
	inp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	machines := make([]Machine, 0)
	machine := Machine{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			machines = append(machines, machine)
			machine = Machine{}
			continue
		}
		if strings.HasPrefix(line, "Button A") {
			coordsStr := strings.TrimPrefix(line, "Button A: ")
			coordsStrSplit := strings.Split(coordsStr, ", ")
			coordsStrXSplit := strings.Split(coordsStrSplit[0], "+")
			coordsStrYSplit := strings.Split(coordsStrSplit[1], "+")
			machine.A.X = utils.StringToInt(coordsStrXSplit[1])
			machine.A.Y = utils.StringToInt(coordsStrYSplit[1])
		}
		if strings.HasPrefix(line, "Button B") {
			coordsStr := strings.TrimPrefix(line, "Button B: ")
			coordsStrSplit := strings.Split(coordsStr, ", ")
			coordsStrXSplit := strings.Split(coordsStrSplit[0], "+")
			coordsStrYSplit := strings.Split(coordsStrSplit[1], "+")
			machine.B.X = utils.StringToInt(coordsStrXSplit[1])
			machine.B.Y = utils.StringToInt(coordsStrYSplit[1])
		}
		if strings.HasPrefix(line, "Prize") {
			coordsStr := strings.TrimPrefix(line, "Prize: ")
			coordsStrSplit := strings.Split(coordsStr, ", ")
			coordsStrXSplit := strings.Split(coordsStrSplit[0], "=")
			coordsStrYSplit := strings.Split(coordsStrSplit[1], "=")
			machine.Prize.X = utils.StringToInt(coordsStrXSplit[1])
			machine.Prize.Y = utils.StringToInt(coordsStrYSplit[1])
		}
	}
	return machines
}
