package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./puzzle.txt")

	if err != nil {
		log.Fatal(err)
	}

	pos := strings.Split(strings.TrimSpace(string(dat)), ",")
	numPos := []int{}

	for _, y := range pos {
		num, _ := strconv.Atoi(y)
		numPos = append(numPos, num)
	}

	m := 0
	for x, y := range numPos {
		if x == 0 || y > m {
			m = y
		}
	}

	mFuel := 0
	fuelPos := 0
	for i := 0; i < m; i++ {
		fuel := 0
		for _, y := range numPos {
			if i > y {
				steps := (i - y)
				// triangle function
				fuel += (steps * (steps + 1)) / 2
			} else {
				steps := (y - i)
				fuel += (steps * (steps + 1)) / 2
			}
		}
		if i == 0 || fuel < mFuel {
			mFuel = fuel
			fuelPos = i
		}
	}

	fmt.Println("minimal fuel position", fuelPos)
	fmt.Println("minimal fuel", mFuel)
}
