package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getAdjacent(row, col int, om [][]int) [][]int {
	atTop := row == 0
	atBottom := row == len(om)-1
	atLeft := col == 0
	atRight := col == len(om[0])-1
	top := []int{}
	topLeft := []int{}
	topRight := []int{}
	bottom := []int{}
	bottomLeft := []int{}
	bottomRight := []int{}
	left := []int{}
	right := []int{}

	if !atTop {
		top = []int{row - 1, col}
		if !atLeft {
			topLeft = []int{row - 1, col - 1}
		}
		if !atRight {
			topRight = []int{row - 1, col + 1}
		}
	}
	if !atBottom {
		bottom = []int{row + 1, col}
		if !atLeft {
			bottomLeft = []int{row + 1, col - 1}
		}
		if !atRight {
			bottomRight = []int{row + 1, col + 1}
		}
	}
	if !atLeft {
		left = []int{row, col - 1}
	}
	if !atRight {
		right = []int{row, col + 1}
	}
	return [][]int{top, bottom, left, right, topLeft, topRight, bottomLeft, bottomRight}
}

func handleFlashes(octoMatrix [][]int, hasFlashed map[string]int, count int) int {
	// handle flashes
	for x, row := range octoMatrix {
		for y, octopus := range row {
			// part ii: any octopus over 9 flashes
			if octopus > 9 {
				// only can flash once
				if _, ok := hasFlashed[fmt.Sprint(x, y)]; ok {
					continue
				}
				hasFlashed[fmt.Sprint(x, y)] = 1
				count++
				// part ii: increment neighbors
				adj := getAdjacent(x, y, octoMatrix)
				for _, neighbor := range adj {
					if len(neighbor) > 0 {
						octoMatrix[neighbor[0]][neighbor[1]] += 1
					}
				}
				//part ii: handle those flashes
				count = handleFlashes(octoMatrix, hasFlashed, count)

				// part iii: set to 0
				octoMatrix[x][y] = 0
			}
		}
	}

	return count
}

func simulate(octoMatrix [][]int) int {
	hasFlashed := make(map[string]int)
	count := 0
	// part i: increase by 1
	for x, row := range octoMatrix {
		for y, octopus := range row {
			octoMatrix[x][y] = octopus + 1
		}
	}

	count = handleFlashes(octoMatrix, hasFlashed, count)

	return count

}

func main() {
	dat, err := os.ReadFile("./puzzle.txt")

	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(strings.TrimSpace(string(dat)), "\n")
	octoMatrix := make([][]int, len(rows))
	for idx, y := range rows {
		ays := strings.Split(y, "")
		iys := []int{}
		for _, ay := range ays {
			iy, _ := strconv.Atoi(ay)
			iys = append(iys, iy)
		}
		octoMatrix[idx] = iys
	}

	agg := 0
	max := 100
	for x := 0; x < max; x++ {
		count := simulate(octoMatrix)
		agg += count
		//fmt.Println(count, "flashes after step", x)
		foundNonFlash := false
		for _, row := range octoMatrix {
			for _, octopus := range row {
				if octopus != 0 {
					foundNonFlash = true
				}
			}
		}
		if !foundNonFlash {
			fmt.Println("step", x+1, "all flashed")
			return
		}
		if x == max-1 && foundNonFlash {
			max += 100
		}

	}
	fmt.Println(agg, "total flashes")
}
