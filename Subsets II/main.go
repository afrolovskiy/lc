package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	v := []int{1, 2, 2}
	// v := []int{1, 2, 2, 3}
	sort.Ints(v)
	fmt.Printf("%v\n", subsets(v))
}

func subsets(nums []int) [][]int {
	var res [][]int
	dups := make(map[string]struct{})
	for i := 0; i < 1<<uint(len(nums)); i++ {
		ss := []int{}
		for j := 0; j < len(nums); j++ {
			if i&(1<<uint(j)) > 0 {
				ss = append(ss, nums[j])
			}
		}

		k := key(ss)
		if _, ok := dups[k]; !ok {
			res = append(res, ss)
			dups[k] = struct{}{}
		}
	}
	return res
}

func key(nums []int) string {
	strs := make([]string, len(nums))
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	return strings.Join(strs, ",")
}
