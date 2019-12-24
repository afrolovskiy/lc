package main

import (
	"fmt"
)

func main() {
	// v := []int{1, 2, 3}
	v := []int{1, 2, 3, 4}
	fmt.Printf("%v\n", subsets(v))
}

func subsets(nums []int) [][]int {
	res := make([][]int, 1<<len(nums))
	for i := range res {
		ss := []int{}
		for j := 0; j < len(nums); j++ {
			if i&(1<<j) > 0 {
				ss = append(ss, nums[j])
			}
		}
		res[i] = ss
	}
	return res
}
