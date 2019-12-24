package main

import "fmt"

func main() {
	v := []int{1, 2, 2}
	// v := []int{3, 2, 1, 2, 1, 7}
	fmt.Printf("%d\n", minIncrementForUnique(v))
}

func minIncrementForUnique(a []int) int {
	incrs := 0
	cnts := make(map[int]int)
	for i := range a {
		v := a[i]
		c := cnts[v]
		cnts[v] = c + 1
		for c > 0 {
			incrs += c
			v += c
			c = cnts[v]
			cnts[v] = c + 1
		}

	}
	return incrs
}
