package dayseventeen

import (
	"fmt"
	"math"
)

// 7,1,5,2,4,0,7,6,1
// 2,4,1,2,7,5,1,3,4,4,5,5,0,3,3,0
// 35184372088832

func Solve2() {
	registers, instructions := readInputFile("./dayseventeen/input.txt")
	fmt.Println(registers, instructions)
	for i := 0; ; i++ {
		registers.A = i
		output := processInstructions2(registers, instructions)
		if len(output) == len(instructions) {
			fmt.Println("A:", registers.A)
			fmt.Println("Output:", output)
			fmt.Println("----")
		}
	}
}

func processInstructions2(registers Registers, instructions []int) []int {
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
			value := operandValue % 8
			if value == 2 && len(output) == 0 {
				return output
			}
			if instructions[len(output)] != value {
				return output
			}
			output = append(output, value)
			if len(output) >= len(instructions) {
				return output
			}
		}
		i += 2
	}
	return output
}
