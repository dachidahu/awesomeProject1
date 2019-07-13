package main

import (
	"sort"
)

/*
Given a collection of intervals, merge all overlapping intervals.
Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
*/

type Interval []int
type SortItems [][2]int

func (p SortItems) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}

func (p SortItems) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p SortItems) Len() int {
	return len(p)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func merge(intervals [][2]int) SortItems {
	var p SortItems = intervals
	sort.Sort(p)
	left := p[0]
	ans := SortItems{}
	for _, interval := range p {
		if interval[0] > left[1] {
			ans = append(ans, left)
			left = interval
		} else {
			left[1] = max(left[1], interval[1])
		}
	}
	if left != p[len(p)-1] {
		ans = append(ans, left)
	}
	return ans
}
