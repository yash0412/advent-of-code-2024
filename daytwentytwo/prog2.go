package daytwentytwo

import (
	"adventofcode/utils"
	"fmt"
)

func Solve2() {
	input := readInputFile("daytwentytwo/input.txt")
	secretSum := 0
	patternMap := make(map[string]map[int]int)
	for _, val := range input {
		newSecret := findNSecretNumberAndMap(val, 2000, patternMap)
		secretSum += newSecret
	}
	fmt.Println("Sum:", secretSum)
	maxSum := 0
	for pattern := range patternMap {
		patternSum := 0
		for input := range patternMap[pattern] {
			patternSum += patternMap[pattern][input]
		}
		if patternSum > maxSum {
			maxSum = patternSum
		}
	}
	fmt.Println("Max Sum:", maxSum)
}

func findNSecretNumberAndMap(input int, iterations int, patternMap map[string]map[int]int) int {
	secret := input
	prevSecretLastChar := secret % 10
	operationsList := []int{}
	for i := 0; i < iterations; i++ {
		newSecret := secret * 64
		newSecret = newSecret ^ secret
		secret = newSecret % 16777216
		newSecret = secret / 32
		newSecret = newSecret ^ secret
		secret = newSecret % 16777216
		newSecret = secret * 2048
		newSecret = newSecret ^ secret
		secret = newSecret % 16777216
		secretLastChar := secret % 10
		diff := secretLastChar - prevSecretLastChar
		operationsList = insertNewElementAtTheEnd(operationsList, diff)

		if len(operationsList) == 4 {
			operationsListStr := utils.IntArrayToString(operationsList, ",")
			if _, ok := patternMap[operationsListStr]; ok {
				if _, ok := patternMap[operationsListStr][input]; !ok {
					patternMap[operationsListStr][input] = secretLastChar
				}
			} else {
				patternMap[operationsListStr] = make(map[int]int)
				patternMap[operationsListStr][input] = secretLastChar
			}
		}
		prevSecretLastChar = secretLastChar
	}
	return secret
}

func insertNewElementAtTheEnd(input []int, newElement int) []int {
	if len(input) < 4 {
		input = append(input, newElement)
		return input
	}
	input = append(input[1:], newElement)
	return input
}
