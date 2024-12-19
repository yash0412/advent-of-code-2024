package daynineteen

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Solve() {
	availableTowels, expectedDesigns := readInputFile("./daynineteen/input.txt")
	availableTowelsMap := make(map[string]bool)
	savedSubstrMap := make(map[string]bool)
	for _, availableTowel := range availableTowels {
		availableTowelsMap[availableTowel] = true
	}
	possibleCount := 0
	for i := range expectedDesigns {
		isPossible := validateDesign(availableTowelsMap, savedSubstrMap, expectedDesigns[i])
		if isPossible {
			possibleCount++
		}
	}

	fmt.Println("Possible Design count:", possibleCount)
}

func validateDesign(availableTowelsMap, savedSubstrMap map[string]bool, expectedDesign string) bool {
	if len(expectedDesign) == 0 {
		return true
	}
	if val, exists := savedSubstrMap[expectedDesign]; exists {
		return val
	}
	savedSubstrMap[expectedDesign] = false
	windowSize := 0
	for windowSize < len(expectedDesign) {
		windowSize++
		substr := expectedDesign[:windowSize]
		if availableTowelsMap[substr] {
			if validateDesign(availableTowelsMap, savedSubstrMap, expectedDesign[windowSize:]) {
				savedSubstrMap[expectedDesign] = true
			}
		}
	}
	return savedSubstrMap[expectedDesign]
}

func readInputFile(fileName string) ([]string, []string) {
	inp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	availableTowels := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		availableTowels = append(availableTowels, strings.Split(line, ", ")...)
	}
	expectedDesigns := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		expectedDesigns = append(expectedDesigns, line)
	}
	return availableTowels, expectedDesigns
}
