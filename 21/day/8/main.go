package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// signal wires a thru g
// signal segments a thru g

func calculateNums(outputs []string, numSize map[int][]string) int {
	count := 0
	for _, y := range outputs {
		output := strings.Split(y, " ")
		for _, o := range output {
			possNumbers := numSize[len(o)]
			if len(possNumbers) == 1 {
				count++
			}
		}
	}
	return count
}

func uncrossWires(signal string, wireMapping map[string]string, signalTranslation map[string]string) string {
	uncrossedSignal := []string{}
	splitted := strings.Split(signal, "")
	for _, y := range splitted {
		uncrossedSignal = append(uncrossedSignal, wireMapping[y])
	}
	sort.Strings(uncrossedSignal)
	return signalTranslation[strings.Join(uncrossedSignal, "")]
}

func main() {
	dat, err := os.ReadFile("./puzzle.txt")

	if err != nil {
		log.Fatal(err)
	}

	// the correct wire config for each displayed number
	correctCombos := map[string]string{}
	correctCombos["abcefg"] = "0"
	correctCombos["cf"] = "1"
	correctCombos["acdeg"] = "2"
	correctCombos["acdfg"] = "3"
	correctCombos["bcdf"] = "4"
	correctCombos["abdfg"] = "5"
	correctCombos["abdefg"] = "6"
	correctCombos["acf"] = "7"
	correctCombos["abcdefg"] = "8"
	correctCombos["abcdfg"] = "9"

	// the size of each wire config, for each number
	numSize := map[int][]string{}

	// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
	numSize[2] = []string{"1"}
	numSize[3] = []string{"7"}
	numSize[4] = []string{"4"}
	numSize[5] = []string{"2", "3", "5"}
	numSize[6] = []string{"0", "6", "9"}
	numSize[7] = []string{"8"}

	ret := strings.Split(strings.TrimSpace(string(dat)), "\n")
	outputs := []string{}
	signals := []string{}
	for _, y := range ret {
		outputs = append(outputs, strings.TrimSpace(strings.Split(y, "|")[1]))
		signals = append(signals, strings.Split(y, "|")[0])
	}

	agg := 0
	//acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab
	for idx, y := range signals {
		fmt.Println("eval signal", idx)
		wireConfig := map[string]string{}
		key := map[string]int{}
		numToSignal := map[int]string{}
		possibleNumbers := map[int][]string{}
		for _, signal := range strings.Split(strings.TrimSpace(y), " ") {
			switch len(signal) {
			case 2:
				key[signal] = 1
				numToSignal[1] = signal
				break
			case 3:
				key[signal] = 7
				numToSignal[7] = signal
				break
			case 4:
				key[signal] = 4
				numToSignal[4] = signal
				break
			case 7:
				key[signal] = 8
				numToSignal[8] = signal
				break
			default:
				for _, y := range numSize[len(signal)] {
					idx, _ := strconv.Atoi(y)
					possibleNumbers[idx] = append(possibleNumbers[idx], signal)
				}
				break
			}
		}

		// acf -> cf should leave a behind
		a := compareStrings(strings.Split(numToSignal[7], ""), strings.Split(numToSignal[1], ""))
		wireConfig["a"] = a[0]
		zeroes := possibleNumbers[0] // could be zero, six, or nine
		leftovers := []string{}
		for _, y := range zeroes {
			acf := strings.Split(numToSignal[7], "")
			// 0 => beg
			// 6 => bdeg
			// 9 => bdg
			removedAcf := removeWires(acf, strings.Split(y, ""))
			if len(removedAcf) == 4 {
				key[y] = 6
				numToSignal[6] = y
			} else {
				leftovers = append(leftovers, y)
			}
		}

		abdefg := strings.Split(numToSignal[6], "")
		// remove abdefg from abcdfg or abcefg, get c
		removedAbdefg := removeWires(abdefg, strings.Split(leftovers[0], ""))
		wireConfig["c"] = removedAbdefg[0]

		acf := strings.Split(numToSignal[7], "")
		// remove a and c from acf, get f
		f := removeWires([]string{wireConfig["c"], wireConfig["a"]}, acf)
		wireConfig["f"] = f[0]

		for _, y := range leftovers { // 0 or 9
			bcdf := strings.Split(numToSignal[4], "")
			// 0 => aeg
			// 9 =>  ag
			removedBcdf := removeWires(bcdf, strings.Split(y, ""))
			if len(removedBcdf) == 2 {
				key[y] = 9
				numToSignal[9] = y
			} else {
				key[y] = 0
				numToSignal[0] = y
			}
		}

		bcdf := strings.Split(numToSignal[4], "")
		abcdfg := strings.Split(numToSignal[9], "")
		// bcdf removed from abcdfg is ag
		removedBcdf := removeWires(bcdf, abcdfg)
		// remove a from ag and you get g
		g := removeWires([]string{wireConfig["a"]}, removedBcdf)
		wireConfig["g"] = g[0]

		abcefg := strings.Split(numToSignal[0], "")
		// bcdf removed from abcefg is aeg
		removedBcdf = removeWires(bcdf, abcefg)
		// ag removed from aeg is e
		e := removeWires([]string{wireConfig["a"], wireConfig["g"]}, removedBcdf)
		wireConfig["e"] = e[0]
		// a c e f g from abcefg is b
		b := removeWires([]string{wireConfig["a"], wireConfig["c"], wireConfig["e"], wireConfig["f"], wireConfig["g"]}, abcefg)
		wireConfig["b"] = b[0]
		d := removeWires([]string{wireConfig["a"], wireConfig["b"], wireConfig["c"], wireConfig["f"], wireConfig["g"]}, abcdfg)
		wireConfig["d"] = d[0]

		//fdgacbe cefdb cefbgd gcbe
		translation := map[string]string{}

		//invert wireconfig
		for x, y := range wireConfig {
			translation[y] = x
		}

		totalNum := []string{}
		for _, y := range strings.Split(outputs[idx], " ") {
			combo := uncrossWires(y, translation, correctCombos)
			totalNum = append(totalNum, combo)
		}
		num, _ := strconv.Atoi(strings.Join(totalNum, ""))
		agg += num
	}

	fmt.Println("number of 1, 7, 4, or 8s", calculateNums(outputs, numSize))
	fmt.Println("total", agg)
}

func removeWires(a, b []string) []string {
	newString := []string{}
	for _, y := range b {
		found := false
		for _, x := range a {
			if x == y {
				found = true
			}
		}
		if !found {
			newString = append(newString, y)
		}
	}
	return newString
}

func removeWire(a string, b []string) []string {
	newString := []string{}
	for _, y := range b {
		if a != y {
			newString = append(newString, y)
		}
	}
	return newString
}

func compareStrings(a, b []string) []string {
	diff := []string{}
	for _, y := range a {
		found := false
		for _, k := range b {
			if y == k {
				found = true
			}
		}
		if !found {
			diff = append(diff, y)
		}
	}
	return diff

}
