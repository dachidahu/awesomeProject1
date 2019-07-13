package main

import (

	"sort"
	"strings"
)

type WordRes []interface{}

type Item struct {
	priority int
	item []interface{}
	index int
}

type PriorityQueue struct {
	wc WordCount
}


const K =  2

func (p PriorityQueue) Len() int {
	return len(p.wc.frequences)
}

func (p PriorityQueue) Swap(i ,j int)  {
	p.wc.frequences[i], p.wc.frequences[j] = p.wc.frequences[j], p.wc.frequences[i]
}

func (p PriorityQueue) Less(i, j int) bool {
	if p.wc.frequences[i][1] == p.wc.frequences[j][1] {
		return p.wc.wordIndex[p.wc.frequences[i][0].(string)] > p.wc.wordIndex[p.wc.frequences[i][0].(string)]
	}
	return  p.wc.frequences[i][1].(int) < p.wc.frequences[j][1].(int)
}





type  WordCount struct {
	frequences [][] interface{}
	wordIndex map[string]int
}

func (wc WordCount) Len() int {
	return len(wc.frequences)
}

func (wc WordCount) Swap(i,j int) {
	wc.frequences[i], wc.frequences[j] = wc.frequences[j], wc.frequences[i]
}

func (wc WordCount) Less(i, j int) bool {
	if wc.frequences[i][1] == wc.frequences[j][1] {
		return wc.wordIndex[wc.frequences[i][0].(string) ] < wc.wordIndex[wc.frequences[j][0].(string)]
	}
	return wc.frequences[i][1].(int) > wc.frequences[j][1].(int)
}

type A struct {
	a int
	k map[string]int
}

func test (a A) {
	a.a = 5
	a.k["test"] = a.a
}




func WordProcessing(s string) [][]interface{} {
	var words  = strings.Split(s, " ")
	var wordIdx = map[string]int{}
	var wordCount = map[string]int {}
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
	var wordFreq = [][]interface{} {
	}


	for key, value := range wordCount {
		wordFreq = append(wordFreq,[]interface{}{key,value})
	}
	wc := WordCount{
		wordFreq,
		wordIdx,
	}
	sort.Sort(wc)
	return wordFreq
}





func main() {
	//var word_freq = WordProcessing("Practice makes perfect. you'll only get Perfect by practice. just practice!")
	//fmt.Print(word_freq)
	testPQ()
}

