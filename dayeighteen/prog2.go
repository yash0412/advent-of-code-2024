package dayeighteen

import (
	"adventofcode/utils"
	"fmt"
)

func Solve2() {
	byteCoords := readInputFile("./dayeighteen/input.txt")
	corruptedPOS := make(map[string]bool)
	for i, coords := range byteCoords {
		if i == NUMBER_OF_BYTES {
			break
		}
		corruptedPOS[utils.CoordsToString(coords.X, coords.Y)] = true
	}
	for i := 0; i < len(byteCoords); i++ {
		index := NUMBER_OF_BYTES + i
		coords := byteCoords[index]
		corruptedPOS[utils.CoordsToString(coords.X, coords.Y)] = true
		cost, _ := (findShortestPathUsingBFS(corruptedPOS))
		if cost == -1 {
			fmt.Println("Index:", index, "Coords:", coords)
			break
		}
	}

}
