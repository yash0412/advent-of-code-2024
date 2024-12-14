package daytwo

import (
	"adventofcode/utils"
	"bufio"
	"log"
	"os"
	"strings"
)

func Solve() {
	dataInput := readInputFile("./daytwo/input.txt")
	safeReports := 0
	for _, v := range dataInput {
		isIncreasing, largestDiffInc := checkIfIncreasing(v)
		isDecreasing, largestDiffDes := checkIfDecreasing(v)
		if (isIncreasing && (largestDiffInc > 0 && largestDiffInc < 4)) || (isDecreasing && (largestDiffDes > 0 && largestDiffDes < 4)) {
			safeReports++
		}
	}
	log.Println("Safe reports: ", safeReports)
}

func readInputFile(fileName string) (dataInput [][]int) {
	inp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		var intList []int
		for _, v := range lineSplit {
			val := utils.StringToInt(v)
			intList = append(intList, val)
		}
		dataInput = append(dataInput, intList)
	}
	return
}

func checkIfIncreasing(list []int) (bool, int) {
	largestDiff := 0
	for i := 0; i < len(list)-1; i++ {
		if list[i] >= list[i+1] {
			return false, 0
		}
		if largestDiff < list[i+1]-list[i] {
			largestDiff = list[i+1] - list[i]
		}
	}
	return true, largestDiff
}

func checkIfDecreasing(list []int) (bool, int) {
	largestDiff := 0
	for i := 0; i < len(list)-1; i++ {
		if list[i] <= list[i+1] {
			return false, 0
		}
		if largestDiff < list[i]-list[i+1] {
			largestDiff = list[i] - list[i+1]
		}
	}
	return true, largestDiff
}
