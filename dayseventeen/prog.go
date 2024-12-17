package dayseventeen

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Registers struct {
	A int
	B int
	C int
}

type Operand struct {
	IsCombo bool
}

var (
	operandMap = map[int]Operand{
		0: {
			IsCombo: true,
		},
		1: {
			IsCombo: false,
		},
		2: {
			IsCombo: true,
		},
		3: {
			IsCombo: false,
		},
		4: {
			IsCombo: false,
		},
		5: {
			IsCombo: true,
		},
		6: {
			IsCombo: true,
		},
		7: {
			IsCombo: true,
		},
	}
)

func Solve() {
	registers, instructions := readInputFile("dayseventeen/input.txt")
	fmt.Println(registers, instructions)
	output := processInstructions(registers, instructions)
	fmt.Println("Output", utils.IntArrayToString(output, ","))
}

func processInstructions(registers Registers, instructions []int) []int {
	output := []int{}
	for i := 0; i < len(instructions); {
		instruction := instructions[i]
		operand := instructions[i+1]
		operandValue := getOperandValue(registers, operand, operandMap[instruction].IsCombo)

		switch instruction {
		case 0:
			numerator := registers.A
			denominator := math.Pow(2, float64(operandValue))

			result := int(numerator / int(denominator))
			registers.A = result
		case 6:
			numerator := registers.A
			denominator := math.Pow(2, float64(operandValue))

			result := int(numerator / int(denominator))
			registers.B = result
		case 7:
			numerator := registers.A
			denominator := math.Pow(2, float64(operandValue))

			result := int(numerator / int(denominator))
			registers.C = result
		case 1:
			left := registers.B
			registers.B = left ^ operandValue
		case 2:
			registers.B = operandValue % 8
		case 3:
			if registers.A != 0 {
				i = operandValue
				continue
			}
		case 4:
			registers.B = registers.B ^ registers.C
		case 5:
			output = append(output, operandValue%8)
		}
		i += 2
	}
	return output
}

func getOperandValue(registers Registers, operand int, isCombo bool) int {
	if !isCombo {
		return operand
	}
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return registers.A
	case 5:
		return registers.B
	case 6:
		return registers.C
	}
	log.Fatalf("Invalid Operand %d", operand)
	return 0
}

func readInputFile(fileName string) (Registers, []int) {
	inp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	regitsers := Registers{}
	instructions := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "Register A") {
			dataStr := strings.TrimPrefix(line, "Register A: ")
			regitsers.A = utils.StringToInt(dataStr)
		}
		if strings.HasPrefix(line, "Register B") {
			coordsStr := strings.TrimPrefix(line, "Register B: ")
			regitsers.B = utils.StringToInt(coordsStr)
		}
		if strings.HasPrefix(line, "Register C") {
			coordsStr := strings.TrimPrefix(line, "Register B: ")
			regitsers.C = utils.StringToInt(coordsStr)
		}
		if strings.HasPrefix(line, "Program") {
			coordsStr := strings.TrimPrefix(line, "Program: ")
			instructions = utils.StringSliceToIntSlice(strings.Split(coordsStr, ","))
		}
	}
	return regitsers, instructions
}
