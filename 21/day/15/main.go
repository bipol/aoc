package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getAdjacent(row, col int, om [][]int) [][]int {
	atTop := row == 0
	atBottom := row == len(om)-1
	atLeft := col == 0
	atRight := col == len(om[0])-1
	top := []int{}
	bottom := []int{}
	left := []int{}
	right := []int{}

	if !atTop {
		top = []int{row - 1, col}
	}
	if !atBottom {
		bottom = []int{row + 1, col}
	}
	if !atLeft {
		left = []int{row, col - 1}
	}
	if !atRight {
		right = []int{row, col + 1}
	}
	return [][]int{top, bottom, left, right}
}

type Coords struct {
	X        int
	Y        int
	Weight   int
	Distance int
}

type PQueue struct {
	queue []Coords
}

func (p *PQueue) Push(coord Coords) {

	p.queue = append(p.queue, coord)
	sort.SliceStable(p.queue, func(i, j int) bool {
		return p.queue[i].Distance < p.queue[j].Distance
	})
}

func (p *PQueue) Pop() Coords {
	node := p.queue[0]
	p.queue = p.queue[1:]
	return node
}

func (p *PQueue) Len() int {
	return len(p.queue)
}

func traverseDji(row, col int, maze [][]int) int {
	explored := make(map[[2]int]bool)
	prev := make(map[[2]int][2]int)
	frontier := &PQueue{}
	distance := make(map[[2]int]int)
	for x, row := range maze {
		for y := range row {
			distance[[2]int{x, y}] = math.MaxInt
		}
	}
	distance[[2]int{0, 0}] = 0
	frontier.Push(Coords{0, 0, 0, 0})
	for frontier.Len() != 0 {
		node := frontier.Pop()
		nodeKey := [2]int{node.X, node.Y}
		if node.X == len(maze)-1 && node.Y == len(maze[0])-1 {
			break
		}
		explored[nodeKey] = true
		adj := getAdjacent(node.X, node.Y, maze)
		for _, cell := range adj {
			if len(cell) != 0 {
				x, y := cell[0], cell[1]
				weight := maze[cell[0]][cell[1]]
				coordKey := [2]int{x, y}
				coord := Coords{x, y, weight, distance[coordKey]}
				_, found := explored[coordKey]
				if !found {
					if distance[nodeKey]+weight < distance[coordKey] {
						distance[coordKey] = distance[nodeKey] + weight
						prev[coordKey] = nodeKey
						coord = Coords{x, y, weight, distance[coordKey]}
						frontier.Push(coord)
					}
				}
			}
		}
	}

	risk := 0
	node := [2]int{len(maze) - 1, len(maze[0]) - 1}
	for true {
		risk += maze[node[0]][node[1]]
		node = prev[node]
		if node[0] == 0 && node[1] == 0 {
			break
		}

	}
	return risk
}

func main() {
	dat, err := os.ReadFile("./puzzle.txt")

	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(strings.TrimSpace(string(dat)), "\n")

	maze := make([][]int, len(rows))
	for idx, y := range rows {
		ays := strings.Split(y, "")
		iys := []int{}
		for _, ay := range ays {
			iy, _ := strconv.Atoi(ay)
			iys = append(iys, iy)
		}
		maze[idx] = iys
	}

	vMaze := make([][]int, len(rows)*5)
	for x := range vMaze {
		vMaze[x] = make([]int, len(maze[0])*5)
	}
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			for x, row := range maze {
				for y, cell := range row {
					if cell+(i+j) > 9 {
						vMaze[x+(len(maze)*j)][y+(len(maze[0])*i)] = cell + (i + j) - 9
					} else {
						vMaze[x+(len(maze)*j)][y+(len(maze[0])*i)] = cell + (i + j)
					}
				}
			}
		}
	}

	//for i, row := range vMaze {
	//	fmt.Println(i, ":", row)
	//}
	//path := make(map[Coords]bool)
	//traverse(0, 0, maze, path, 0, risk)
	fmt.Println(traverseDji(0, 0, vMaze))

}
