package main

import (
	"fmt"
	"testing"
)

type TestItem struct {
	input  [][2]int
	result [][2]int
}

func TestMerge(t *testing.T) {
	testData := []TestItem{{
		[][2]int{{3, 6}, {3, 9}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {2, 7}},
		[][2]int{},
	}}
	for _, data := range testData {
		fmt.Println(merge(data.input))
	}
}
