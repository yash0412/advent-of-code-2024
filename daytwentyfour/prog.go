package daytwentyfour

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Operation struct {
	Left     string
	Right    string
	Operator string
	Output   string
}

func Solve() {
	initValues, operations := readInputFile("daytwentyfour/input.txt")
	knownValuesMap := make(map[string]int)
	for _, line := range initValues {
		lineSplit := strings.Split(line, ": ")
		knownValuesMap[lineSplit[0]] = utils.StringToInt(lineSplit[1])
	}
	operationsList := make([]Operation, 0)
	for _, line := range operations {
		lineSplit := strings.Split(line, " -> ")
		operationsSplit := strings.Split(lineSplit[0], " ")
		newOp := Operation{Left: operationsSplit[0], Right: operationsSplit[2], Operator: operationsSplit[1], Output: lineSplit[1]}
		operationsList = append(operationsList, newOp)
	}
	performAllOperations(operationsList, knownValuesMap)
	outputFields := []string{}
	for outKey := range knownValuesMap {
		if strings.HasPrefix(outKey, "z") {
			outputFields = append(outputFields, outKey)
		}
	}
	sort.Slice(outputFields, func(i, j int) bool { return strings.ToLower(outputFields[i]) > strings.ToLower(outputFields[j]) })
	outputVals := []int{}
	for _, val := range outputFields {

		outputVals = append(outputVals, knownValuesMap[val])

	}
	fmt.Println("Output:", utils.IntArrayToString(outputVals, ""))
}

func performAllOperations(operationsList []Operation, knownValuesMap map[string]int) {
	operatedMap := map[int]bool{}

	for len(operatedMap) < len(operationsList) {
		for i, operation := range operationsList {
			if operatedMap[i] || !checkIfInputsAvailable(operation, knownValuesMap) {
				continue
			}
			operatedMap[i] = true
			output := 0
			switch operation.Operator {
			case "AND":
				output = andOp(knownValuesMap[operation.Left], knownValuesMap[operation.Right])
			case "OR":
				output = orOp(knownValuesMap[operation.Left], knownValuesMap[operation.Right])
			case "XOR":
				output = xorOp(knownValuesMap[operation.Left], knownValuesMap[operation.Right])
			}
			knownValuesMap[operation.Output] = output
		}
	}

}

func andOp(val1, val2 int) int {
	if val1 == 1 && val2 == 1 {
		return 1
	}
	return 0
}

func orOp(val1, val2 int) int {
	if val1 == 0 && val2 == 0 {
		return 0
	}
	return 1
}

func xorOp(val1, val2 int) int {
	if val1 == val2 {
		return 0
	}
	return 1
}

func checkIfInputsAvailable(operation Operation, knownValuesMap map[string]int) bool {
	_, leftExists := knownValuesMap[operation.Left]
	_, rightExists := knownValuesMap[operation.Right]
	return leftExists && rightExists
}

func readInputFile(fileName string) ([]string, []string) {
	inp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	initValues := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		initValues = append(initValues, line)
	}
	operations := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		operations = append(operations, line)
	}
	return initValues, operations
}
