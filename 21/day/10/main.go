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

	score := make(map[string]int)
	validOpenChars := make(map[string]int)
	validClosedChars := make(map[string]string)
	validOpenChars["("] = 1
	validOpenChars["{"] = 1
	validOpenChars["<"] = 1
	validOpenChars["["] = 1
	validClosedChars[")"] = "("
	validClosedChars["}"] = "{"
	validClosedChars[">"] = "<"
	validClosedChars["]"] = "["
	score[")"] = 3
	score["}"] = 1197
	score[">"] = 25137
	score["]"] = 57
	totalScore := 0

	closingScore := make(map[string]int)
	closingScore["("] = 1
	closingScore["{"] = 3
	closingScore["<"] = 4
	closingScore["["] = 2

	scores := []int{}
	for _, row := range rows {
		openChunk := []string{}
		totalFixScore := 0
		symbols := strings.Split(row, "")
		corrupt := false
		for _, char := range symbols {
			if _, ok := validOpenChars[char]; ok {
				openChunk = append(openChunk, char)
				continue
			}
			if openChar, ok := validClosedChars[char]; ok {
				if openChunk[len(openChunk)-1] == openChar {
					openChunk = openChunk[:len(openChunk)-1]
				} else {
					totalScore += score[char]
					corrupt = true
					break
				}
			}
		}
		if !corrupt {
			openChunk = reverseStringSlice(openChunk)
			for _, opened := range openChunk {
				totalFixScore = totalFixScore * 5
				totalFixScore += closingScore[opened]
			}
			scores = append(scores, totalFixScore)
		}
	}

	sort.Ints(scores)
	// 1 2 3
	mid := scores[(len(scores)+1)/2-1]
	fmt.Println("total score of invalid closing chars", totalScore)
	fmt.Println("scores", scores)
	fmt.Println("mid", mid)
}

func reverseStringSlice(s []string) []string {
	reversed := []string{}
	for x := 1; x <= len(s); x++ {
		reversed = append(reversed, s[len(s)-x])
	}
	return reversed
}
