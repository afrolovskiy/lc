package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%t\n", repeatedSubstringPattern("abab"))
	fmt.Printf("%t\n", repeatedSubstringPattern("aba"))
	fmt.Printf("%t\n", repeatedSubstringPattern("abcabcabcabc"))
}

func repeatedSubstringPattern(s string) bool {
	for i := 1; i < len(s)/2+1; i++ {
		if len(s)%i != 0 {
			continue
		}
		p := s[:i]
		if strings.Repeat(p, len(s)/i) == s {
			return true
		}
	}
	return false
}
