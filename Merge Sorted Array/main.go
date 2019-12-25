package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3

	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
	fmt.Printf("res: %v\n", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2[:n])
	sort.Ints(nums1[:m+n])
}
