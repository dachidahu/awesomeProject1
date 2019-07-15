package main

import "testing"

type TestItem2 struct {
	input  []int
	output int
}

var data = []TestItem2{
	{[]int{1, 2}, 0},
	{[]int{9, 9, 9, 9, 9, 9, 9}, 7},
	{[]int{9, 9, 1, 1, 1, 9, 9}, 7},
	{[]int{9, 9, 6, 0, 6, 6, 9}, 3},
	{[]int{6, 6, 6, 0, 9, 6, 9, 9}, 5},
}

func TestLongestSubarray(t *testing.T) {

	for _, tdata := range data {
		if k2 := longestWPI1(tdata.input); k2 != tdata.output {
			t.Error("fail", tdata.input, tdata.output, k2)
		}
		if k1 := longestWPI(tdata.input); k1 != tdata.output {
			t.Error("fail", tdata.input, tdata.output, k1)
		}
	}
}
