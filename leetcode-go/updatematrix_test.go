package leetcode

import (
	"fmt"
	"testing"
)

func TestUpdatematrix(t *testing.T) {
	var matrix = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	t.Log(updateMatrix(matrix))
}

func updateMatrix(matrix [][]int) [][]int {
	fmt.Println(len(matrix))
	fmt.Println(len(matrix[0]))
	// for i, j := 0, 0; i < len(matrix) || j < len(matrix[0]); i, j = i+1, j+1 {

	// }
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {

		}
	}
	return matrix
}

func updateMatrix1(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])
	var helper func(i, j int)
	level := 0
	helper = func(i, j int) {
		if i < 0 || i > m-1 || j < 0 || j > n-1 {
			return
		}
		if matrix[i][j] == 1 {
			matrix[i][j] = level - 1
		}
	}
	// 遍历直到hasLv为false
	hasLv := true
	for hasLv {
		hasLv = false
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if matrix[i][j] == level {
					hasLv = true
					helper(i-1, j)
					helper(i+1, j)
					helper(i, j+1)
					helper(i, j-1)
				}
			}
		}
		level--
	}
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = -matrix[i][j]
		}
	}

	return matrix
}
