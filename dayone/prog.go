package dayone

import (
	"adventofcode/utils"
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func Solve() {
	leftList, rightList := readInputFile()
	sort.Ints(leftList)
	sort.Ints(rightList)
	rightListCountMap := convertListToCountMap(rightList)
	sumOfDifference := 0
	sumOfSimilarity := 0
	for i := 0; i < len(leftList); i++ {
		diff := (rightList[i] - leftList[i])
		if diff < 0 {
			diff = -diff
		}
		sumOfDifference += diff
		sumOfSimilarity += leftList[i] * rightListCountMap[leftList[i]]
	}
	log.Println("Sum of differences: ", sumOfDifference)
	log.Println("Sum of similarities: ", sumOfSimilarity)
}

func readInputFile() (leftList []int, rightList []int) {
	inputFileName := "./dayone/input.txt"
	inp, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "   ")
		left := utils.StringToInt(lineSplit[0])
		right := utils.StringToInt(lineSplit[1])
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}
	return
}

func convertListToCountMap(list []int) map[int]int {
	m := make(map[int]int)
	for _, v := range list {
		if val := m[v]; val != 0 {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	return m
}
