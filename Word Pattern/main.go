package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Printf("%t\n", wordPattern("abba", "dog cat cat dog"))
	fmt.Printf("%t\n", wordPattern("abba", "dog cat cat fish"))
	fmt.Printf("%t\n", wordPattern("aaaa", "dog cat cat dog"))
	fmt.Printf("%t\n", wordPattern("abba", "dog dog dog dog"))
}

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	wordMap := make(map[string]rune)
	runeMap := make(map[rune]string)
	for i, r := range pattern {
		w := words[i]

		wr, ok := wordMap[w]
		if ok {
			if r != wr {
				return false
			}
		} else {
			wordMap[w] = r
		}

		ww, ok := runeMap[r]
		if ok {
			if w != ww {
				return false
			}
		} else {
			runeMap[r] = w
		}
	}
	return true
}
