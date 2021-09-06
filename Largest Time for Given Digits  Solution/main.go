package main

import "fmt"

func main() {
	fmt.Printf("%s\n", largestTimeFromDigits([]int{1, 2, 3, 4}))
	fmt.Printf("%s\n", largestTimeFromDigits([]int{5, 5, 5, 5}))
	fmt.Printf("%s\n", largestTimeFromDigits([]int{0, 0, 0, 0}))
	fmt.Printf("%s\n", largestTimeFromDigits([]int{0, 0, 1, 0}))
}

func largestTimeFromDigits(arr []int) string {
	res := ""
	for _, p := range permutations(arr) {
		hours := p[0]*10 + p[1]
		minutes := p[2]*10 + p[3]
		if hours <= 23 && minutes <= 59 {
			t := fmt.Sprintf("%d%d:%d%d", p[0], p[1], p[2], p[3])
			if t > res {
				res = t
			}
		}
	}
	return res
}

func permutations(arr []int) [][]int {
	if len(arr) <= 1 {
		return [][]int{arr}
	}

	res := make([][]int, 0, 0)
	for i := 0; i < len(arr); i++ {
		part := make([]int, 0, len(arr)-1)
		part = append(part, arr[0:i]...)
		part = append(part, arr[i+1:]...)
		for _, p := range permutations(part) {
			res = append(res, append([]int{arr[i]}, p...))
		}
	}
	return res
}
