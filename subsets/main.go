package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// v := []int{1, 2, 2}
	v := []int{1, 2, 2, 3}
	sort.Ints(v)
	fmt.Printf("%v\n", subsets(v))
}

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	var res [][]int
	dups := make(map[string]bool)

	add := func(x []int) {
		k := key(x)
		if !dups[k] {
			res = append(res, x)
			dups[k] = true
		}
	}

	for i := range nums {
		add([]int{nums[i]})

		ss := subsets(nums[i+1:])
		for j := range ss {
			add(ss[j])
			add(append([]int{nums[i]}, ss[j]...))
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
