package daytwentyfive

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Solve() {
	locks, keys := readInputFile("daytwentyfive/input.txt")
	compatibleCount := 0
	for _, lock := range locks {
		for _, key := range keys {
			if checkIfCompatible(lock, key) {
				compatibleCount++
			}
		}
	}

	fmt.Println("Compatible combinations:", compatibleCount)
}

func checkIfCompatible(lock [5]int, key [5]int) bool {
	for i := range lock {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}

func readInputFile(fileName string) ([][5]int, [][5]int) {
	inp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	locks := make([][5]int, 0)
	keys := make([][5]int, 0)
	isLock := false
	isKey := false
	newLock := [5]int{}
	newKey := [5]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if isLock {
				locks = append(locks, newLock)
			}
			if isKey {
				keys = append(keys, subtractValueFromEachEntry(newKey))
			}
			isLock = false
			isKey = false
			newLock = [5]int{}
			newKey = [5]int{}
		}
		if isLock {
			lineSplit := strings.Split(line, "")
			for i, char := range lineSplit {
				if char == "#" {
					newLock[i]++
				}
			}
			continue
		}
		if isKey {
			lineSplit := strings.Split(line, "")
			for i, char := range lineSplit {
				if char == "#" {
					newKey[i]++
				}
			}
			continue
		}
		if line == "#####" {
			isLock = true
		}

		if line == "....." {
			isKey = true
		}
	}

	return locks, keys
}

func subtractValueFromEachEntry(input [5]int) [5]int {
	for i := range input {
		input[i]--
	}
	return input
}
