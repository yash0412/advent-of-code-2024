package dayfour

import "log"

func Solve2() {
	input := readInputFile("./dayfour/input.txt")
	total := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'A' {
				res := findX_MAS(i, j, input)
				total += res
			}
		}
	}
	log.Println("Total: ", total)
}

func findX_MAS(x int, y int, input [][]rune) int {
	ops := getXOperations(x, y, input)
	if len(ops) < 4 {
		return 0
	}
	topLeftCorner := ops[0]
	bottomLeftCorner := ops[1]
	topRightCorner := ops[2]
	bottomRightCorner := ops[3]

	if (input[topLeftCorner[0]][topLeftCorner[1]] == 'M' &&
		input[bottomRightCorner[0]][bottomRightCorner[1]] == 'S') ||
		(input[topLeftCorner[0]][topLeftCorner[1]] == 'S' &&
			input[bottomRightCorner[0]][bottomRightCorner[1]] == 'M') {
		if (input[bottomLeftCorner[0]][bottomLeftCorner[1]] == 'M' &&
			input[topRightCorner[0]][topRightCorner[1]] == 'S') ||
			(input[bottomLeftCorner[0]][bottomLeftCorner[1]] == 'S' &&
				input[topRightCorner[0]][topRightCorner[1]] == 'M') {
			return 1
		}
	}
	return 0
}

func getXOperations(x int, y int, input [][]rune) [][]int {
	ops := make([][]int, 0)
	possibleOps := [][]int{
		{x - 1, y - 1}, {x - 1, y + 1},
		{x + 1, y - 1}, {x + 1, y + 1},
	}
	for _, v := range possibleOps {
		if areCoordinatesValid(v[0], v[1], input) {
			ops = append(ops, v)
		}
	}
	return ops
}
