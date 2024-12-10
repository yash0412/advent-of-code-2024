package daythree

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Solve() {
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	regexNum := regexp.MustCompile(`\d{1,3}`)
	input := readInputFile("./daythree/input.txt")
	matches := regex.FindAllString(input, -1)
	total := 0
	for _, match := range matches {
		nums := regexNum.FindAllString(match, -1)
		i, _ := strconv.Atoi(nums[0])
		j, _ := strconv.Atoi(nums[1])

		total += i * j
	}
	log.Println("Total:", total)
}

func readInputFile(fileName string) string {
	inp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	input := ""
	for scanner.Scan() {
		line := scanner.Text()
		input += line
	}
	return input
}
