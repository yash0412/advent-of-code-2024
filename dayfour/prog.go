package dayfour

import (
	"adventofcode/utils"
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
	if !utils.IsCoordinatesValid(x+dx, y+dy, len(input), len(input[0])) {
		return false
	}
	char := input[x+dx][y+dy]
	if char == 'A' {
		return findS(x+dx+dx, y+dy+dy, input)
	}
	return false
}

func findS(x int, y int, input [][]rune) bool {
	if !utils.IsCoordinatesValid(x, y, len(input), len(input[0])) {
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
		if utils.IsCoordinatesValid(v[0], v[1], len(input), len(input[0])) {
			ops = append(ops, v)
		}
	}
	return ops
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
