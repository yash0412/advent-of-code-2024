package daytwentytwo

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
)

func Solve() {
	input := readInputFile("daytwentytwo/input.txt")
	secretSum := 0
	for _, val := range input {
		newSecret := findNSecretNumber(val, 2000)
		fmt.Println("Secret:", val, newSecret)
		secretSum += newSecret
	}
	fmt.Println("Sum:", secretSum)
}

func findNSecretNumber(input int, iterations int) int {
	secret := input
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

	}
	return secret
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
		val := utils.StringToInt(line)
		input = append(input, val)
	}
	return input
}
