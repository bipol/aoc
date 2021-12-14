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
	template := strings.Split(rows[0], "")

	pairInsertions := rows[2:]

	rules := make(map[string]string)
	for _, k := range pairInsertions {
		split := strings.Fields(k)
		rules[split[0]] = split[2]
	}

	fmt.Println(template)
	fmt.Println(rules)

	for x := 1; x <= 10; x++ {
		new := []string{template[0]}
		for y := 1; y < len(template); y++ {
			pair := fmt.Sprintf("%s%s", template[y-1], template[y])
			rule := rules[pair]
			//fmt.Println("insert", rule, "between", template[y-1], template[y])
			new = append(new, rule)
			new = append(new, template[y])
			//			fmt.Println(new)
		}
		template = new
		fmt.Println(x, "tempate size", len(template))
	}

	elementCount := make(map[string]int)
	for _, elem := range template {
		elementCount[elem]++
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
