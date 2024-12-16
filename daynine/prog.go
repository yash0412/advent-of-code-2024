package daynine

import (
	"adventofcode/utils"
	"bufio"
	"log"
	"os"
)

func Solve() {
	inputString := readInputFile("./daynine/input.txt")
	newInput := defragmentDisk(createDiskMap(inputString))
	log.Println("Checksum: ", calculateChecksum(newInput))
	clearOutputFile("output.txt")
	printLayout("output.txt", newInput)
}

func clearOutputFile(fileName string) {
	if err := os.Truncate("./daynine/"+fileName, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
}

func printLayout(fileName string, input []int64) {
	f, err := os.OpenFile("./daynine/"+fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(utils.Int64ArrayToString((input), "|")); err != nil {
		panic(err)
	}
	if _, err = f.WriteString("\n"); err != nil {
		panic(err)
	}
}

func calculateChecksum(input []int64) int64 {
	checksum := int64(0)
	for i, v := range input {
		if v == -1 {
			continue
		}
		checksum += int64(int64(i) * v)
	}
	return checksum
}

func defragmentDisk(input []int64) []int64 {
	firstElement, lastElement := 0, len(input)-1
	for {
		if firstElement > lastElement {
			log.Println(firstElement, lastElement)
			break
		}
		for {
			if input[firstElement] == -1 {
				break
			}
			firstElement++
		}
		for {
			if input[lastElement] != -1 {
				break
			}
			lastElement--
		}
		if firstElement > lastElement {
			log.Println(firstElement, lastElement)
			break
		}
		input[firstElement], input[lastElement] = input[lastElement], input[firstElement]
	}
	return input
}

func createDiskMap(input string) []int64 {
	fileId := int64(0)
	diskMap := make([]int64, 0)
	for i := range input {
		char := string(input[i])
		charNum := utils.StringToInt(char)
		if i%2 != 0 {
			for k := 0; k < charNum; k++ {
				diskMap = append(diskMap, -1)
			}
		} else {
			for k := 0; k < charNum; k++ {
				diskMap = append(diskMap, fileId)
			}
			fileId++
		}
	}
	return diskMap
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
