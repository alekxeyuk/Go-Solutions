package matrix

import (
	"strconv"
	"strings"
)

// Define the Matrix and Pair types here.
type Matrix struct {
	rows, cols int
	array      [][]int
	rowsMax    []int
	colsMin    []int
}

type Pair struct {
	x, y int
}

func New(s string) (*Matrix, error) {
	a := [][]int{}
	rows := []int{}
	cols := []int{}

	for rI, r := range strings.Split(s, "\n") {
		t := strings.Split(r, " ")
		tI := make([]int, 0, len(t))
		for cI, n := range t {
			i, r := strconv.ParseInt(n, 10, 32)
			if r != nil {
				continue
			}
			tI = append(tI, int(i))

			if rI >= len(rows) {
				rows = append(rows, int(i))
			} else if int(i) > rows[rI] {
				rows[rI] = int(i)
			}
			if cI >= len(cols) {
				cols = append(cols, int(i))
			} else if int(i) < cols[cI] {
				cols[cI] = int(i)
			}
		}
		a = append(a, tI)
	}

	return &Matrix{rows: len(a), cols: len(a[0]), array: a, rowsMax: rows, colsMin: cols}, nil
}

func (m *Matrix) Saddle() []Pair {
	pairs := []Pair{}
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			if m.array[r][c] >= m.rowsMax[r] && m.array[r][c] <= m.colsMin[c] {
				pairs = append(pairs, Pair{x: r + 1, y: c + 1})
			}
		}
	}
	return pairs
}
