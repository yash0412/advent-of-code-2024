package daytwo

import "log"

func Solve2() {
	dataInput := readInputFile("./daytwo/test.txt")
	safeReports := 0
	for _, v := range dataInput {
		if checkIfValid(v, -1) {
			safeReports++
		}
	}
	log.Println("Safe reports: ", safeReports)
}

func checkIfValid(list []int, index int) bool {
	if index == len(list) {
		return false
	}
	newList := list
	if index > -1 {
		newList = removeElementFromList(list, index)
	}
	if len(newList) == 0 {
		return false
	}
	isIncreasing, largestDiffInc := checkIfIncreasing(newList)
	isDecreasing, largestDiffDes := checkIfDecreasing(newList)
	if (isIncreasing && (largestDiffInc > 0 && largestDiffInc < 4)) || (isDecreasing && (largestDiffDes > 0 && largestDiffDes < 4)) {
		return true
	}
	return checkIfValid(list, index+1)
}

func removeElementFromList(list []int, index int) []int {
	return append(list[:index], list[index+1:]...)
}
