package daynineteen

import (
	"fmt"
)

func Solve2() {
	availableTowels, expectedDesigns := readInputFile("./daynineteen/input.txt")
	availableTowelsMap := make(map[string]bool)
	savedSubstrMap := make(map[string]int)
	for _, availableTowel := range availableTowels {
		availableTowelsMap[availableTowel] = true
	}
	possibleCount := 0
	for i := range expectedDesigns {
		isPossible := validateDesignAndCount(availableTowelsMap, savedSubstrMap, expectedDesigns[i])
		possibleCount += isPossible
	}

	fmt.Println("Possible Design total count:", possibleCount)
}

func validateDesignAndCount(availableTowelsMap map[string]bool, savedSubstrMap map[string]int, expectedDesign string) int {
	if len(expectedDesign) == 0 {
		return 1
	}
	if val, exists := savedSubstrMap[expectedDesign]; exists {
		return val
	}
	savedSubstrMap[expectedDesign] = 0
	windowSize := 0
	for windowSize < len(expectedDesign) {
		windowSize++
		substr := expectedDesign[:windowSize]
		if availableTowelsMap[substr] {
			possibleCount := validateDesignAndCount(availableTowelsMap, savedSubstrMap, expectedDesign[windowSize:])
			savedSubstrMap[expectedDesign] += possibleCount
		}
	}
	return savedSubstrMap[expectedDesign]
}
