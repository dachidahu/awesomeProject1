package main

import (
	"fmt"
	"sort"
)

func TwoSum1(arr []int, target int) (*[2]int, int) {
	l, r := 0, len(arr)-1
	sort.Ints(arr)
	for l <= r {
		sum := arr[l] + arr[r]
		if sum == target {
			return &[2]int{l, r}, 0
		} else if sum > target {
			r -= 1
		} else {
			l += 1
		}
	}
	return nil, -1
}

func TwoSum(arr []int, target int) ([2]int, int) {
	dict := map[int]int{}
	var ans [2]int
	for i, v := range arr {
		rem := target - v
		if _, ok := dict[rem]; ok {
			return [2]int{dict[rem], i}, 0
		}

		if _, ok := dict[v]; !ok {
			dict[v] = i
		}
		dict[v] = i
	}
	return ans, -1
}

func TestSum() {
	testCase := []int{0, 1, 1, 4}
	if _, error := TwoSum(testCase, 3); error == 0 {
		fmt.Print("error!!!")
	}
	if arr, error := TwoSum(testCase, 2); error != 0 || arr != [2]int{1, 2} {
		fmt.Print("error!!")
	}
	fmt.Println(TwoSum(testCase, 6))
}
