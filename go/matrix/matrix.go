package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix interface {
	Rows() [][]int
	Cols() [][]int
	Set(int, int, int) bool
}

type matrix struct {
	rowCount int
	colCount int
	m        []int
}

func (m matrix) Rows() [][]int {
	res := make([][]int, m.rowCount)
	for i := range res {
		res[i] = []int{}
		for j := 0; j < m.colCount; j++ {
			res[i] = append(res[i], m.m[i*m.colCount+j])
		}
	}
	return res
}

func (m matrix) Cols() [][]int {
	res := make([][]int, m.colCount)
	for i := range res {
		res[i] = []int{}
		for j := 0; j < m.rowCount; j++ {
			res[i] = append(res[i], m.m[i+j*m.colCount])
		}
	}
	return res
}

func (m *matrix) Set(r, c, val int) bool {
	m.m[r*m.colCount+c] = val
	return true
}

func New(s string) (*matrix, error) {
	mat := &matrix{
		rowCount: 0,
		colCount: 0,
		m:        make([]int, 0),
	}
	rows := strings.Split(s, "\n")
	mat.rowCount = len(rows)
	tmpColCount := -1
	for _, row := range rows {
		row = strings.TrimSpace(row)
		colsInRow := strings.Split(row, " ")
		if tmpColCount == -1 {
			tmpColCount = len(colsInRow)
		} else if tmpColCount != len(colsInRow) {
			return nil, errors.New("invalid")
		}
		for _, col := range colsInRow {
			intVal, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			mat.m = append(mat.m, intVal)
		}
	}
	mat.colCount = tmpColCount
	return mat, nil
}
