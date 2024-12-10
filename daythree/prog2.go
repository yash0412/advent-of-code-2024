package daythree

import (
	"log"
	"regexp"
	"strconv"
)

func Solve2() {
	input := readInputFile("./daythree/input.txt")

	regex := regexp.MustCompile(`(mul\(\d{0,3},\d{0,3}\))|(do\(\))|(don't\(\))`)

	matches := regex.FindAllString(input, -1)
	isEnabled := true
	total := 0
	for _, match := range matches {
		if match == "do()" {
			isEnabled = true
		} else if match == "don't()" {
			isEnabled = false
		} else if isEnabled {
			regexNum := regexp.MustCompile(`\d{1,3}`)
			nums := regexNum.FindAllString(match, -1)
			i, _ := strconv.Atoi(nums[0])
			j, _ := strconv.Atoi(nums[1])

			total += i * j
		}
	}
	log.Println("Total:", total)
}
