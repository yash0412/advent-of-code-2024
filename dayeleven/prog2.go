package dayeleven

import (
	"adventofcode/utils"
	"fmt"
	"log"
)

func Solve2() {
	inputNumbers := readInputFile("./dayeleven/input.txt")
	stoneNumbersCountMap := createStoneNumbersCountMap(inputNumbers)
	for i := 0; i < NUM_ITERATIONS*3; i++ {
		stoneNumbersCountMap = applyRulesOnStonesNew(stoneNumbersCountMap)
	}
	log.Println("Total Stones: ", getSumOfUniqueCounts(stoneNumbersCountMap))
}

func applyRulesOnStonesNew(stoneNumbersCountMap map[int]int) map[int]int {
	stoneNumbersCountMapNew := make(map[int]int)
	for k, v := range stoneNumbersCountMap {
		if k == 0 {
			stoneNumbersCountMapNew[1] += v
		} else if numStr := fmt.Sprintf("%d", k); len(numStr)%2 == 0 {
			firstNum, secondNum := utils.StringToInt(numStr[:len(numStr)/2]), utils.StringToInt(numStr[len(numStr)/2:])
			stoneNumbersCountMapNew[firstNum] += v
			stoneNumbersCountMapNew[secondNum] += v
		} else {
			stoneNumbersCountMapNew[k*2024] += v
		}
	}
	return stoneNumbersCountMapNew
}

func printMap(stoneNumbersCountMap map[int]int) {
	for k, v := range stoneNumbersCountMap {
		log.Println(k, v)
	}
}

func getSumOfUniqueCounts(stoneNumbersCountMap map[int]int) int {
	count := 0
	for k := range stoneNumbersCountMap {
		count += stoneNumbersCountMap[k]
	}
	return count
}

func createStoneNumbersCountMap(numbers []int) map[int]int {
	stoneNumbersCountMap := make(map[int]int)
	for _, v := range numbers {
		stoneNumbersCountMap[v]++
	}

	return stoneNumbersCountMap
}
