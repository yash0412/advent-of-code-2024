package daytwentythree

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type QueueElement struct {
	Comp       string
	ParentComp *QueueElement
	Cost       int
}

func Solve() {
	networkMap := readInputFile("./daytwentythree/input.txt")
	foundNetworks := make(map[string]bool)
	for computer := range networkMap {
		findPathWithCost3(networkMap, computer, foundNetworks)
	}
	santaCount := 0
	for connectedNetworks := range foundNetworks {
		if strings.HasPrefix(connectedNetworks, "t") || strings.Contains(connectedNetworks, ",t") {
			santaCount++
		}
	}
	fmt.Println("Count:", santaCount)
}

func findPathWithCost3(networkMap map[string]map[string]struct{}, startingComp string, foundNetworks map[string]bool) {
	queue := []QueueElement{{Comp: startingComp, Cost: 0, ParentComp: nil}}
	for len(queue) > 0 {
		currentComp := queue[0]
		queue = queue[1:]
		if currentComp.Comp == startingComp && currentComp.Cost == 3 {
			foundNetworks[findAllCompsInPath(currentComp)] = true
		}
		if currentComp.Cost == 3 {
			continue
		}
		adjacentComps := networkMap[currentComp.Comp]
		for adjacentComp := range adjacentComps {
			if adjacentComp == startingComp {
				if currentComp.Cost == 2 {
					queue = append(queue, QueueElement{Comp: adjacentComp, Cost: currentComp.Cost + 1, ParentComp: &currentComp})
				}
			} else {
				queue = append(queue, QueueElement{Comp: adjacentComp, Cost: currentComp.Cost + 1, ParentComp: &currentComp})
			}
		}
	}
}

func findAllCompsInPath(currentComp QueueElement) string {
	compStr := []string{}
	currentElem := currentComp.ParentComp
	for currentElem != nil {
		compStr = append(compStr, currentElem.Comp)
		currentElem = currentElem.ParentComp
	}
	sort.Slice(compStr, func(i, j int) bool { return strings.ToLower(compStr[i]) < strings.ToLower(compStr[j]) })
	return utils.StringArrayToString(compStr, ",")
}

func readInputFile(fileName string) map[string]map[string]struct{} {
	inp, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()
	scanner := bufio.NewScanner(inp)
	input := make(map[string]map[string]struct{})
	for scanner.Scan() {
		line := scanner.Text()
		comps := strings.Split(line, "-")
		if _, ok := input[comps[0]]; !ok {
			input[comps[0]] = make(map[string]struct{})
		}
		if _, ok := input[comps[1]]; !ok {
			input[comps[1]] = make(map[string]struct{})
		}
		input[comps[0]][comps[1]] = struct{}{}
		input[comps[1]][comps[0]] = struct{}{}
	}
	return input
}
