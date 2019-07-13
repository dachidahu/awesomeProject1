package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type PQItem struct {
	word      string
	frequence int
	index     int
}

type PriorityQueue1 struct {
	cap       int
	items     []*PQItem
	wordIndex map[string]int
}

func (pq PriorityQueue1) Len() int {
	return len(pq.items)
}

func (pq PriorityQueue1) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = i
	pq.items[j].index = j
}

func (pq PriorityQueue1) Less(i, j int) bool {
	if pq.items[i].frequence == pq.items[j].frequence {
		return pq.wordIndex[pq.items[i].word] > pq.wordIndex[pq.items[j].word]
	}
	return pq.items[i].frequence < pq.items[j].frequence
}

func (pq *PriorityQueue1) Push(x interface{}) {
	n := len((*pq).items)
	item := x.(*PQItem)
	item.index = n
	(*pq).items = append((*pq).items, item)
}

func (pq *PriorityQueue1) Pop() interface{} {
	old := pq
	n := len((*old).items)
	item := old.items[n-1]

	item.index = -1
	(*pq).items = (*pq).items[0 : n-1]
	return item
}

func escape(word string) string {
	word = strings.ToLower(word)
	var res string
	for _, r := range word {
		if r < 'a' || r > 'z' {
			continue
		}
		res += string(r)
	}
	return res
}

func TopK(s string, k int) *PriorityQueue1 {
	var words = strings.Split(s, " ")
	var wordIdx = map[string]int{}
	var wordCount = map[string]int{}

	items := []*PQItem{}
	pq := PriorityQueue1{
		k,
		items,
		wordIdx,
	}

	for idx, word := range words {
		word = escape(word)
		if wordCount[word] > 0 {
			wordCount[word] += 1
		} else {
			wordCount[word] = 1
		}
		if _, ok := wordIdx[word]; !ok {
			wordIdx[word] = idx
		}
	}

	heap.Init(&pq)

	for key, value := range wordCount {
		heap.Push(&pq, &PQItem{key, value, 0})
		if pq.Len() > pq.cap {
			heap.Pop(&pq)
		}
	}
	fmt.Print(pq.items[0])
	return &pq

}

func main() {
	TopK("Practice makes perfect. you'll only get Perfect by practice. just practice!", 2)
}
