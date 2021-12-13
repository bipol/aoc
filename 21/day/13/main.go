package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./puzzle.txt")

	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(strings.TrimSpace(string(dat)), "\n")
	dots := make(map[int][]int)
	instrStart := 0
	for i, k := range rows {
		if k == "" {
			instrStart = i
			break
		}
		coords := strings.Split(k, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		if v, ok := dots[y]; ok {
			dots[y] = append(v, x)
		} else {
			dots[y] = []int{x}
		}
	}
	folds := rows[instrStart+1:]
	instr := [][]string{}
	for _, k := range folds {
		fields := strings.Fields(k)
		instr = append(instr, strings.Split(fields[len(fields)-1], "="))
	}
	for _, instruction := range instr {
		fmt.Println("starting instruction", instruction)
		foldPlace, _ := strconv.Atoi(instruction[1])
		switch instruction[0] {
		case "y":
			keys := []int{}
			for k, _ := range dots {
				keys = append(keys, k)
			}
			sort.Ints(keys)
			//fmt.Println("y before", instr, dots)
			for x := foldPlace; x <= keys[len(keys)-1]; x++ {
				dots[foldPlace-(x-foldPlace)] = append(dots[foldPlace-(x-foldPlace)], dots[x]...)
				delete(dots, x)
			}
		case "x":
			//fmt.Println("x before", instr, dots)
			for k, v := range dots {
				newPlace := []int{}
				for _, x := range v {
					if x > foldPlace {
						newPlace = append(newPlace, foldPlace-(x-foldPlace))
					} else {
						newPlace = append(newPlace, x)
					}
					dots[k] = newPlace
				}
			}
			//fmt.Println("x after", instr, dots)
		}
		fmt.Println("finishing instruction", instruction)
	}

	numDots := 0
	for _, v := range dots {
		f := make(map[int]bool)
		for _, k := range v {
			if _, ok := f[k]; !ok {
				numDots++
				f[k] = true
			}
		}
	}

	fmt.Println("total number of visible dots", numDots)
	fmt.Println(dots)
	printDots(dots)
}

func printDots(dots map[int][]int) {
	keys := []int{}
	for k, _ := range dots {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	maxX := 0
	// for all the columns
	for x := 0; x <= keys[len(keys)-1]; x++ {
		// go over all of our dots
		if ints, ok := dots[x]; ok {
			if len(ints) > 0 {
				sort.Ints(ints)
				if ints[len(ints)-1] > maxX {
					maxX = ints[len(ints)-1]
				}
				for y := 0; y <= ints[len(ints)-1]; y++ {
					found := false
					for _, k := range ints {
						if y == k {
							found = true
							break
						}
					}
					if found {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
			} else {
				for x := 0; x <= maxX; x++ {
					fmt.Print(".")
				}
			}
			fmt.Println("")
		}

	}
}
