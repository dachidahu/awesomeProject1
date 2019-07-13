package main

func solve(str string) string {
	left := 0
	dict := map[rune]bool{}
	runes := []rune(str)
	ansStr := ""
	maxLength := 1 - 1<<31
	for i, v := range runes {
		for {
			if _, ok := dict[v]; !ok {
				break
			}
			delete(dict, runes[left])
			left += 1
		}
		if (i-left)+1 > maxLength {
			ansStr = string(runes[left : i+1])
			maxLength = i - left + 1
		}
		dict[v] = true
	}
	return ansStr
}
