package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	wordArray := strings.Fields(s)
	var counts = make(map[string]int)
	for _, v := range wordArray {
		_, ok := counts[v]
		if ok {
			counts[v] += 1
		} else {
			counts[v] = 1
		}
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
