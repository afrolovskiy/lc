package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%v\n", partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Printf("%v\n", partitionLabels("eccbbbbdec"))
	fmt.Printf("%v\n", partitionLabels("vhaagbqkaq"))
}

func partitionLabels(s string) []int {
	var res []int

	bi := -1
	pi := -1
	seen := make(map[rune]struct{})

	for i, r := range s {
		if _, ok := seen[r]; !ok {
			li := strings.LastIndexByte(s, byte(r))
			seen[r] = struct{}{}
			if li > bi {
				bi = li
			}
		}
		if i == bi {
			res = append(res, bi-pi)
			pi = bi
		}
	}
	return res
}
