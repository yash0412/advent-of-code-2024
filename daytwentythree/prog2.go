package daytwentythree

import (
	"adventofcode/utils"
	"fmt"
	"sort"
	"strings"
)

func Solve2() {
	networkMap := readInputFile("./daytwentythree/input.txt")
	allVertices := make(map[string]bool)
	for node := range networkMap {
		allVertices[node] = true
	}
	allCliques := bronKerbosch(networkMap, map[string]bool{}, allVertices, map[string]bool{})
	maxClique := map[string]bool{}

	for _, clique := range allCliques {
		if len(clique) > len(maxClique) {
			maxClique = clique
		}
	}
	comps := []string{}
	for comp := range maxClique {
		comps = append(comps, comp)
	}

	sort.Slice(comps, func(i, j int) bool { return strings.ToLower(comps[i]) < strings.ToLower(comps[j]) })

	fmt.Println(utils.StringArrayToString(comps, ","))
}

func bronKerbosch(networkMap map[string]map[string]bool,
	R map[string]bool, P map[string]bool, X map[string]bool) []map[string]bool {

	cliques := make([]map[string]bool, 0)
	if len(P) == 0 && len(X) == 0 {
		cliques = append(cliques, R)
	}

	for vertice := range P {
		newClique := map[string]bool{}
		for key := range R {
			newClique[key] = true
		}
		newClique[vertice] = true
		// 	let newP = new Set([...P].filter((x) => graph.get(v).has(x)));
		// let newX = new Set([...X].filter((x) => graph.get(v).has(x)));
		// cliques = new Set([...cliques, ...bronKerbosch(newR, newP, newX, graph)]);
		newP := make(map[string]bool)
		for newVertice := range P {
			if networkMap[vertice][newVertice] {
				newP[newVertice] = true
			}
		}

		newX := make(map[string]bool)
		for newVertice := range X {
			if networkMap[vertice][newVertice] {
				newX[newVertice] = true
			}
		}
		// fmt.Println(newClique, newP, newX)
		cliques = append(cliques, bronKerbosch(networkMap, newClique, newP, newX)...)
		delete(P, vertice)
		X[vertice] = true
	}
	return cliques
}
