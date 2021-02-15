package leetcode

import (
	"fmt"
	"testing"
)

func TestAddToArrayFormOfInteger(t *testing.T) {
	var (
		a = []int{0}
		k = 34
	)
	t.Log(addToArrayForm1(a, k))
}

// 翻转
func addToArrayForm1(A []int, K int) []int {
	var (
		lens = len(A)
		add  int
	)
	A = reverse1(A)
	for i := 0; i < lens || K != 0; i, K = i+1, K/10 {
		fmt.Println(i)
		if i >= lens {
			A = append(A, 0)
		}
		sum := K%10 + A[i] + add
		add = sum / 10
		A[i] = sum % 10
	}
	if add > 0 {
		A = append(A, add)
	}
	A = reverse1(A)
	return A
}

func reverse1(sli []int) []int {
	for i := 0; i < len(sli)/2; i++ {
		sli[i], sli[len(sli)-1-i] = sli[len(sli)-1-i], sli[i]
	}
	return sli
}
