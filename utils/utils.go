package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func CoordsToString(i, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func StringToCoords(coordsStr string) (int, int) {
	coordsStrSplit := strings.Split(coordsStr, ",")
	return StringToInt(coordsStrSplit[0]), StringToInt(coordsStrSplit[1])
}

func StringToInt(input string) int {
	val, _ := strconv.Atoi(input)
	return val
}

func StringSliceToIntSlice(input []string) []int {
	res := make([]int, 0)
	for i := range input {
		res = append(res, StringToInt(input[i]))
	}
	return res
}

func IsCoordinatesValid(x, y, xsize, ysize int) bool {
	if x < 0 || x >= xsize {
		return false
	}
	return y >= 0 && y < ysize
}

func IntArrayToString(input []int, sep string) string {
	res := ""
	for i := range input {
		res += strconv.Itoa(input[i]) + sep
	}
	return res
}
