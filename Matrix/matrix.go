package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	rows int
	cols int
	matx []int64
}

func New(s string) (*Matrix, error) {
	mem := []int64{}
	words := strings.Split(s, "\n")
	rows := len(words)
	cols := -1
	for _, row := range words {
		cls := strings.Fields(row)
		if cols == -1 {
			cols = len(cls)
		} else if len(cls) != cols {
			return nil, errors.New("broken matrix")
		}
		for _, c := range cls {
			if n, err := strconv.ParseInt(c, 10, 64); err == nil {
				mem = append(mem, n)
			} else {
				return nil, errors.New("broken matrix")
			}
		}
	}
	return &Matrix{rows: rows, cols: cols, matx: mem}, nil
}

func (m *Matrix) matrixTravel(x int, y int, l func(i, j int) int) [][]int {
	rws := make([][]int, x)
	for i := 0; i < x; i++ {
		if rws[i] == nil {
			rws[i] = make([]int, y)
		}
		for j := 0; j < y; j++ {
			rws[i][j] = int(m.matx[l(i, j)])
		}
	}
	return rws
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	return m.matrixTravel(m.cols, m.rows, func(i, j int) int {
		return i + j*m.cols
	})
}

func (m *Matrix) Rows() [][]int {
	return m.matrixTravel(m.rows, m.cols, func(i, j int) int {
		return m.cols*i + j
	})
}

func (m *Matrix) Set(row, col, val int) bool {
	if row >= m.rows || row < 0 || col >= m.cols || col < 0 {
		return false
	}
	m.matx[m.cols*row+col] = int64(val)
	return true
}
