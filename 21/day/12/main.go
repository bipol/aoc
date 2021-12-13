package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func traversePath(node string, caveGraph map[string][]string, path []string) [][]string {
	paths := [][]string{}
	np := make([]string, len(path))
	copy(np, path)
	np = append(np, node)
	if node == "end" {
		paths = append(paths, np)
		return paths
	}
	// for each of these we have a new possible path
	for _, next := range caveGraph[node] {
		if next == "start" {
			continue
		}
		//fmt.Println("moving from", node, "to", next)
		visited := make(map[string]int)
		skip := false
		//a b a => b
		for _, n := range np {
			if unicode.IsLower(rune(n[0])) {
				visited[n]++
				if n == next {
					if visited[n] > 1 {
						skip = true
					}
				}
			}
		}
		multis := []string{}
		for k, v := range visited {
			if v >= 2 {
				multis = append(multis, k)
			}

		}
		if skip || len(multis) > 1 {
			//fmt.Println("skipped", np)
			continue
		}
		//fmt.Println("accepted", np, "multis", multis)
		paths = append(paths, traversePath(next, caveGraph, np)...)
	}
	return paths
}

func main() {
	dat, err := os.ReadFile("./puzzle.txt")

	if err != nil {
		log.Fatal(err)
	}

	caveGraph := make(map[string][]string)
	rows := strings.Split(strings.TrimSpace(string(dat)), "\n")

	// you can go forward or backward
	for _, row := range rows {
		nodes := strings.Split(row, "-")
		caveGraph[nodes[0]] = append(caveGraph[nodes[0]], nodes[1])
		found := false
		for _, node := range caveGraph[nodes[1]] {
			if node == nodes[0] {
				found = true
			}
		}
		if !found {
			caveGraph[nodes[1]] = append(caveGraph[nodes[1]], nodes[0])
		}
	}

	fmt.Println(caveGraph)
	totalPaths := [][]string{}
	for _, node := range caveGraph["start"] {
		totalPaths = append(totalPaths, traversePath(node, caveGraph, []string{"start"})...)
	}

	fmt.Println("found", len(totalPaths), "paths to end")
	onlyOnce := 0
	for _, path := range totalPaths {
		found := false
		foundTwice := false
		for _, node := range path {
			if unicode.IsLower(rune(node[0])) && node != "start" && node != "end" {
				if !found {
					found = true
				} else {
					foundTwice = true
				}
			}
		}
		if !foundTwice {
			onlyOnce++
		}
		//fmt.Println(strings.Join(path, ","))
	}
	fmt.Println(len(totalPaths), "paths visit small caves at most once")
	fmt.Println(onlyOnce, "paths visit small caves at most once")
}
