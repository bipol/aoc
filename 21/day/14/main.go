package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./test.txt")

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
	for x := 1; x <= 10; x++ {
		new := []rune{template[0]}
		for y := 1; y < len(template); y++ {
			pair := fmt.Sprintf("%s%s", string(template[y-1]), string(template[y]))
			rule := rules[pair]
			elementCount[string(rule)]++
			new = append(new, rule)
			new = append(new, template[y])
		}
		template = new
		fmt.Println(x, "tempate size", len(template))
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
