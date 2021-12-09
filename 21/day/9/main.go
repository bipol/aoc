package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type heightMap [][]int

func getAdjacent(row, col int, hm heightMap) []int {
	atTop := row == 0
	atBottom := row == len(hm)-1
	atLeft := col == 0
	atRight := col == len(hm[0])-1
	top := -1
	bottom := -1
	left := -1
	right := -1

	if !atTop {
		top = hm[row-1][col]
	}
	if !atBottom {
		bottom = hm[row+1][col]
	}
	if !atLeft {
		left = hm[row][col-1]
	}
	if !atRight {
		right = hm[row][col+1]
	}
	return []int{top, bottom, left, right}
}

func traverseBasin(row, col int, hm heightMap, basin []int, coords map[string]int) []int {
	coords[fmt.Sprint(row, col)] = 1
	atTop := row == 0
	atBottom := row == len(hm)-1
	atLeft := col == 0
	atRight := col == len(hm[0])-1

	point := hm[row][col]
	if point == 9 {
		return basin
	}
	basin = append(basin, point)

	//[5 8 7 8 8 6 7 8]
	if !atTop && coords[fmt.Sprint(row-1, col)] != 1 {
		basin = append(basin, traverseBasin(row-1, col, hm, []int{}, coords)...)
	}
	if !atBottom && coords[fmt.Sprint(row+1, col)] != 1 {
		basin = append(basin, traverseBasin(row+1, col, hm, []int{}, coords)...)
	}
	if !atLeft && coords[fmt.Sprint(row, col-1)] != 1 {
		basin = append(basin, traverseBasin(row, col-1, hm, []int{}, coords)...)
	}
	if !atRight && coords[fmt.Sprint(row, col+1)] != 1 {
		basin = append(basin, traverseBasin(row, col+1, hm, []int{}, coords)...)
	}
	return basin
}

func main() {
	dat, err := os.ReadFile("./puzzle.txt")

	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(strings.TrimSpace(string(dat)), "\n")
	heightMatrix := make(heightMap, len(rows))
	for idx, y := range rows {
		ays := strings.Split(y, "")
		iys := []int{}
		for _, ay := range ays {
			iy, _ := strconv.Atoi(ay)
			iys = append(iys, iy)
		}
		heightMatrix[idx] = iys
	}

	lowPoints := []int{}
	riskLevel := 0
	basins := []int{}
	for x, rows := range heightMatrix {
		for y, point := range rows {
			adj := getAdjacent(x, y, heightMatrix)
			foundSmaller := false
			for _, cell := range adj {
				if cell != -1 {
					if cell <= point {
						foundSmaller = true
						break
					}
				}
			}
			if !foundSmaller {
				lowPoints = append(lowPoints, point)
				riskLevel += point + 1
				coords := make(map[string]int)
				basin := traverseBasin(x, y, heightMatrix, []int{}, coords)
				basins = append(basins, len(basin))
			}

		}
	}
	fmt.Println("low points", lowPoints)
	fmt.Println("riskLevel", riskLevel)
	sort.Ints(basins)
	fmt.Println("three highest basin sizes", basins[len(basins)-3:])
	fmt.Println("product", basins[len(basins)-3:][0]*basins[len(basins)-3:][1]*basins[len(basins)-3:][2])

}
