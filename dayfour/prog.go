package dayfour

import (
	"bufio"
	"log"
	"os"
)

func Solve() {
	input := readInputFile("./dayfour/input.txt")
	total := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'X' {
				res := findMAS(i, j, input)
				total += res
			}
		}
	}
	log.Println("Total: ", total)
}

func findMAS(x int, y int, input [][]rune) int {
	ops := getOperations(x, y, input)
	res := 0
	for _, v := range ops {
		char := input[v[0]][v[1]]
		if char == 'M' {
			if findAS(v[0], v[1], v[0]-x, v[1]-y, input) {
				res++
			}
		}
	}
	return res
}

func findAS(x int, y int, dx int, dy int, input [][]rune) bool {
	if !areCoordinatesValid(x+dx, y+dy, input) {
		return false
	}
	char := input[x+dx][y+dy]
	if char == 'A' {
		return findS(x+dx+dx, y+dy+dy, input)
	}
	return false
}

func findS(x int, y int, input [][]rune) bool {
	if !areCoordinatesValid(x, y, input) {
		return false
	}
	char := input[x][y]
	return char == 'S'
}

func getOperations(x int, y int, input [][]rune) [][]int {
	ops := make([][]int, 0)
	possibleOps := [][]int{
		{x - 1, y - 1}, {x - 1, y}, {x - 1, y + 1},
		{x, y - 1}, {x, y + 1},
		{x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1},
	}
	for _, v := range possibleOps {
		if areCoordinatesValid(v[0], v[1], input) {
			ops = append(ops, v)
		}
	}
	return ops
}

func areCoordinatesValid(x int, y int, input [][]rune) bool {
	xsize := len(input)
	if x < 0 || x >= xsize {
		return false
	}
	ysize := len(input[x])
	return y >= 0 && y < ysize
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
