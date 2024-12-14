package dayseven

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strings"
)

func Solve2() {
	results, values := readInputFile("./dayseven/input.txt")
	counter := int64(0)
	for i := range results {
		res := results[i]
		vals := values[i]
		result := findSum2(res, vals)
		if result == 0 {
			counter += int64(res)
		}
	}
	log.Println("Counter: ", counter)
}

func findSum2(res int, vals []int) int {
	if len(vals) == 0 {
		return 1
	}
	result3, success := applyOperations2(res, vals[len(vals)-1], "||")
	if result3 == 0 && success && len(vals) == 1 {
		return result3
	}
	if success {
		if findSum2(result3, vals[:len(vals)-1]) == 0 {
			return 0
		}
	}
	result1, success := applyOperations2(res, vals[len(vals)-1], "+")
	if result1 == 0 && success && len(vals) == 1 {
		return result1
	}
	if !success {
		return 1
	}
	if findSum2(result1, vals[:len(vals)-1]) == 0 {
		return 0
	}
	result2, success := applyOperations2(res, vals[len(vals)-1], "*")
	if result2 == 0 && success && len(vals) == 1 {
		return result2
	}
	if !success {
		return 1
	}
	if findSum2(result2, vals[:len(vals)-1]) == 0 {
		return 0
	}
	return 1
}

func applyOperations2(op1, op2 int, operator string) (int, bool) {
	switch operator {
	case "+":
		return op1 - op2, op1 >= op2
	case "*":
		return op1 / op2, op1%op2 == 0
	case "||":
		op1Str := fmt.Sprintf("%d", op1)
		op2Str := fmt.Sprintf("%d", op2)
		val := utils.StringToInt(strings.TrimSuffix(op1Str, op2Str))
		return val, strings.HasSuffix(op1Str, op2Str)
	}
	return 0, false
}
