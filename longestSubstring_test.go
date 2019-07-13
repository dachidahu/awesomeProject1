package main

import (
	"testing"
)

type TestDataItem struct {
	inputs string
	result string
}

func TestSolve(t *testing.T) {
	dataItems := []TestDataItem{
		{"a", "a"},
		{"", ""},
		{"aaaaaaa", "a"},
		{"abcdefg", "abcdefg"},
		{"abcdbzcaabcd", "dbzca"},
		{"accdecx", "decx"},
	}

	for _, test_item := range dataItems {
		if actual := solve(test_item.inputs); actual != test_item.result {
			t.Error("acutal:", actual, ", expected:", test_item.result)
		}
	}
}
