package test

import (
	toolkit "github.com/byyam/go-leetcode-toolkit"
	"testing"
)

var grid = [][]int{
	{1, 2, 2},
	{2, 3, 4},
}

func Test_PrintGrid(t *testing.T) {
	toolkit.PrintGrid(grid)
}

func Test_CompareGrid(t *testing.T) {
	rst := toolkit.CompareGrid(grid, grid)
	t.Logf("compare grid:%v", rst)
}
