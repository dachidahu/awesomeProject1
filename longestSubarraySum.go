package main

//longest Longest Well-Performing Interval
func longestWPI(hours []int) int {
	maxLen := 0
	for i := 0; i < len(hours); i++ {
		s := 0
		for j := i; j < len(hours); j++ {
			if hours[j] < 9 {
				s -= 1
			} else {
				s += 1
			}
			if s > 0 && maxLen < j-i+1 {
				maxLen = j - i + 1
			}
		}
	}
	return maxLen
}

func longestWPI1(hours []int) int {
	dict := map[int]int{}
	res := 0
	s := 0
	//dict[0] = -1
	for i, v := range hours {
		if v > 8 {
			s += 1
		} else {
			s -= 1
		}
		if s > 0 {
			res = i + 1
		} else {
			if idx, ok := dict[s-1]; ok {
				res = max(res, i-idx)
			}
		}
		if _, ok := dict[s]; !ok {
			dict[s] = i
		}
	}
	return res
}
