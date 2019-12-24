package main

import (
	"fmt"
	"unicode"
)

func main() {
	v := "a1b2"
	// v := "C"
	// v := "3z4"
	// v := "12345"
	fmt.Printf("%#v\n", letterCasePermutation(v))
}

func letterCasePermutation(s string) []string {
	var letters []int
	for i, ch := range s {
		if ch < '0' || ch > '9' {
			letters = append(letters, i)
		}
	}

	res := make([]string, 1<<uint(len(letters)))
	for i := range res {
		ss := []rune(s)
		for j, idx := range letters {
			if i&(1<<uint(j)) > 0 {
				r := ss[idx]
				if 'a' <= r && r <= 'z' {
					ss[idx] = unicode.ToUpper(ss[idx])
				} else {
					ss[idx] = unicode.ToLower(ss[idx])
				}
			}
		}
		res[i] = string(ss)
	}

	return res
}
