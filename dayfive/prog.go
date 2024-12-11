package dayfive

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// rule 33|55
// files 33,55,77

func Solve() {
	rules, files := readInputFile("./dayfive/input.txt")
	rulesMap := make(map[string]map[string]struct{})
	for i := 0; i < len(rules); i++ {
		rule := rules[i]
		ruleParts := strings.Split(rule, "|")
		less := (ruleParts[0])
		more := (ruleParts[1])
		if _, found := rulesMap[less]; found {
			rulesMap[less][more] = struct{}{}
		} else {
			rulesMap[less] = map[string]struct{}{more: {}}
		}
	}
	validTotal := 0
	invalidTotal := 0
	for _, fileStr := range files {
		fileData := strings.Split(fileStr, ",")
		if isFileValid(fileData, rulesMap) {
			intVal, _ := strconv.Atoi(fileData[len(fileData)/2])
			validTotal += intVal
		} else {
			validFile := sortInvalidFiles(fileData, rulesMap)
			intVal, _ := strconv.Atoi(validFile[len(validFile)/2])
			invalidTotal += intVal
		}
	}
	log.Println("Valid Total: ", validTotal)
	log.Println("Invalid Total: ", invalidTotal)
}

func isFileValid(fileData []string, rules map[string]map[string]struct{}) bool {
	for i := 0; i < len(fileData)-1; i++ {
		less := fileData[i]
		more := fileData[i+1]
		if _, found := rules[less][more]; !found {
			return false
		}
	}
	return true
}

func sortInvalidFiles(fileData []string, rules map[string]map[string]struct{}) []string {
	sort.Slice(fileData, func(i, j int) bool {
		a := fileData[i]
		b := fileData[j]
		_, isLess := rules[a][b]
		return isLess
	})
	return fileData
}

func readInputFile(fileName string) ([]string, []string) {
	inp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	rules := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rules = append(rules, line)
	}
	files := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		files = append(files, line)
	}
	return rules, files
}
