package main

import "fmt"

func main() {
	// v := []int{1, 2, 2}
	v := []int{3, 2, 1, 2, 1, 7}
	fmt.Printf("%d\n", minIncrementForUnique(v))
}

func minIncrementForUnique(a []int) int {
	incrs := 0
	dups := make(map[int]bool)
	for i := range a {
		v := a[i]
		for dups[v] {
			incrs++
			v++
		}
		dups[v] = true
	}
	return incrs
}
