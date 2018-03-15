package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type matrix [][]int

func (m matrix) Rows() [][]int {
	rowCount := len(m)
	colCount := len(m[0])
	rows := make([][]int, rowCount)
	for r := 0; r < rowCount; r++ {
		rows[r] = make([]int, colCount)
		for c := 0; c < colCount; c++ {
			rows[r][c] = m[r][c]
		}
	}
	return rows
}
func (m matrix) Cols() [][]int {
	rowCount := len(m)
	colCount := len(m[0])
	cols := make([][]int, colCount)
	for c := 0; c < colCount; c++ {
		cols[c] = make([]int, rowCount)
		for r := 0; r < rowCount; r++ {
			cols[c][r] = m[r][c]
		}
	}
	return cols
}
func (m matrix) Set(r, c, val int) bool {
	if r >= len(m) || c >= len(m[0]) || r < 0 || c < 0 {
		return false
	}
	m[r][c] = val
	return true
}

func New(s string) (m matrix, err error) {
	lines := strings.Split(s, "\n")
	m = make(matrix, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		if i >= 1 && len(m[i-1]) != len(fields) {
			return nil, fmt.Errorf("inconsistent field length")
		}
		m[i] = make([]int, len(fields))
		for j, field := range fields {
			m[i][j], err = strconv.Atoi(string(field))
			if err != nil {
				return
			}
		}
	}
	return
}
