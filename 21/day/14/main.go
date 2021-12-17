package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./puzzle.txt")

	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(strings.TrimSpace(string(dat)), "\n")
	template := []rune(rows[0])

	pairInsertions := rows[2:]

	rules := make(map[string]rune)
	for _, k := range pairInsertions {
		split := strings.Fields(k)
		rules[split[0]] = rune(split[2][0])
	}

	fmt.Println(template)
	fmt.Println(rules)

	elementCount := make(map[string]int)
	for _, k := range template {
		elementCount[string(k)]++
	}

	//AB
	//ACB
	// AC CB
	pairs := make(map[string]int)
	//create original template pairs
	for y := 1; y < len(template); y++ {
		pair := fmt.Sprintf("%s%s", string(template[y-1]), string(template[y]))
		pairs[pair]++
	}

	for x := 1; x <= 40; x++ {
		fmt.Println(x)
		newPairs := make(map[string]int)
		for k, v := range pairs {
			rule := rules[k]
			newPair := fmt.Sprintf("%s%s", string(k[0]), string(rule))
			newPair2 := fmt.Sprintf("%s%s", string(rule), string(k[1]))
			elementCount[string(rule)] += v
			newPairs[newPair] += v
			newPairs[newPair2] += v
		}
		pairs = newPairs
	}

	sorted := []int{}
	for _, v := range elementCount {
		sorted = append(sorted, v)
	}
	sort.Ints(sorted)

	fmt.Println("template is size", len(template))
	fmt.Println("mc", sorted[len(sorted)-1])
	fmt.Println("lc", sorted[0])
	fmt.Println("mc - lc", sorted[len(sorted)-1]-sorted[0])

}
