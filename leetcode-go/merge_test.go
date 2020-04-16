package leetcode

import (
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	var in = [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	t.Log(merge(in))
}

func merge(intervals [][]int) [][]int {
	hl := len(intervals)
	res := [][]int{}
	if hl == 0 {
		return res
	}

	//1 go语言排序语法
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//2 申请空间res
	res = append(res, intervals[0])

	i := 1
	//3.开始遍历
	for i < hl {
		resEnd := res[len(res)-1]
		//2.1 大于，找到left和right，然后修改res最后一个元素
		if resEnd[len(resEnd)-1] >= intervals[i][0] {
			left := resEnd[0]
			right := maxa(resEnd[len(resEnd)-1], intervals[i][len(intervals[i])-1])
			res[len(res)-1] = []int{left, right}
		} else { //2.3 不大于直接添加
			res = append(res, intervals[i])
		}
		i++
	}
	return res
}

func maxa(a, b int) int {
	if a < b {
		return b
	}

	return a
}
