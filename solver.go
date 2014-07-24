package main

import "fmt"

const puzzleSize = 5

type group struct {
	value int
	op    string
	cells [][2]int
}

func validate(sets []group) bool {
	var field = make([][]bool, puzzleSize)
	for i := 0; i < puzzleSize; i++ {
		field[i] = make([]bool, puzzleSize)
	}

	for _, group := range sets {
		for _, pos := range group.cells {
			x, y := pos[0], pos[1]
			if x >= puzzleSize || y >= puzzleSize {
				return false
			}
			if field[x][y] {
				return false
			}
			field[x][y] = true
		}
	}

	sum := 0

	for i := 0; i < puzzleSize; i++ {
		for j := 0; j < puzzleSize; j++ {
			if field[i][j] {
				sum++
			}
		}
	}

	return sum == puzzleSize*puzzleSize

}

func main() {
	var sets []group

	sets = append(sets, group{2, "-", [][2]int{{0, 0}, {1, 0}}})
	sets = append(sets, group{4, "+", [][2]int{{0, 1}}})
	sets = append(sets, group{40, "*", [][2]int{{0, 2}, {1, 1}, {1, 2}}})
	sets = append(sets, group{2, "/", [][2]int{{0, 3}, {0, 4}}})

	sets = append(sets, group{7, "+", [][2]int{{1, 3}, {1, 4}, {2, 4}}})

	sets = append(sets, group{2, "/", [][2]int{{2, 0}, {3, 0}}})
	sets = append(sets, group{15, "*", [][2]int{{2, 1}, {2, 2}, {3, 1}}})
	sets = append(sets, group{9, "+", [][2]int{{2, 3}, {3, 3}}})

	sets = append(sets, group{2, "+", [][2]int{{3, 2}}})
	sets = append(sets, group{1, "-", [][2]int{{3, 4}, {4, 4}}})

	sets = append(sets, group{5, "+", [][2]int{{4, 0}, {4, 1}}})
	sets = append(sets, group{1, "-", [][2]int{{4, 2}, {4, 3}}})

	if validate(sets) {
		fmt.Println("valid configuration")
	} else {
		fmt.Println("invalid configuration")
	}

}
