package dayseven

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	results, values := readInputFile("./dayseven/input.txt")
	counter := int64(0)
	for i := range results {
		res := results[i]
		vals := values[i]
		result := findSum(res, vals)
		if result == 0 {
			counter += int64(res)
		}
	}
	log.Println("Counter: ", counter)
}

func findSum(res int, vals []int) int {
	if len(vals) == 0 {
		return 1
	}
	result1, success := applyOperations(res, vals[len(vals)-1], "+")
	if result1 == 0 && success && len(vals) == 1 {
		return result1
	}
	if !success {
		return 1
	}
	if findSum(result1, vals[:len(vals)-1]) == 0 {
		return 0
	}
	result2, success := applyOperations(res, vals[len(vals)-1], "*")
	if result2 == 0 && success && len(vals) == 1 {
		return result2
	}
	if !success {
		return 1
	}
	if findSum(result2, vals[:len(vals)-1]) == 0 {
		return 0
	}
	return 1
}

func applyOperations(op1, op2 int, operator string) (int, bool) {
	switch operator {
	case "+":
		return op1 - op2, op1 >= op2
	case "*":
		return op1 / op2, op1%op2 == 0
	}
	return 0, false
}

func readInputFile(inputFileName string) (result []int, values [][]int) {
	inp, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ": ")
		left, err := strconv.Atoi(lineSplit[0])
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, left)
		rightSplit := strings.Split(lineSplit[1], " ")
		rightList := []int{}
		for i := range rightSplit {
			x, _ := strconv.Atoi(rightSplit[i])
			rightList = append(rightList, x)
		}
		values = append(values, rightList)
	}
	return
}
