package dayeleven

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	NUM_ITERATIONS = 25
)

type Stone struct {
	NextStone *Stone
	Number    int
}

func newStone(number int, nexStone *Stone) *Stone {
	return &Stone{Number: number, NextStone: nexStone}
}

func Solve() {
	inputNumbers := readInputFile("./dayeleven/input.txt")

	rootStone := createListOfStones(inputNumbers)
	for i := 0; i < NUM_ITERATIONS; i++ {
		rootStone = applyRulesOnStones(rootStone)
	}
	log.Println("Total Stones: ", printListOfStones(rootStone))

}

func createListOfStones(inputNumbers []int) *Stone {
	if len(inputNumbers) == 0 {
		return nil
	}
	var lastCreatedStone *Stone
	for i := len(inputNumbers) - 1; i >= 0; i-- {
		lastCreatedStone = newStone(inputNumbers[i], lastCreatedStone)
	}
	return lastCreatedStone
}

func printListOfStones(rootStone *Stone) int {
	stoneCount := 0
	currentStone := rootStone
	for {
		if currentStone == nil {
			return stoneCount
		}
		stoneCount++
		// log.Printf("Stone Index: %d, Number: %d", stoneCount, currentStone.Number)
		currentStone = currentStone.NextStone
	}
}

func applyRulesOnStones(rootStone *Stone) *Stone {
	currentStone := rootStone
	var previousStone *Stone
	for {
		if currentStone == nil {
			return rootStone
		}

		if currentStone.Number == 0 {
			currentStone.Number = 1
		} else if numStr := fmt.Sprintf("%d", currentStone.Number); len(numStr)%2 == 0 {
			firstNum, secondNum := utils.StringToInt(numStr[:len(numStr)/2]), utils.StringToInt(numStr[len(numStr)/2:])
			secondStone := newStone(secondNum, currentStone.NextStone)
			firstStone := newStone(firstNum, secondStone)
			if previousStone == nil {
				rootStone = firstStone
			} else {
				previousStone.NextStone = firstStone
			}
			previousStone = secondStone
			currentStone = secondStone.NextStone
			continue
		} else {
			currentStone.Number = currentStone.Number * 2024
		}
		previousStone = currentStone
		currentStone = currentStone.NextStone
	}
}

func readInputFile(fileName string) []int {
	inp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	input := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		for _, v := range lineSplit {
			val := utils.StringToInt(v)
			input = append(input, val)
		}
	}
	return input
}
