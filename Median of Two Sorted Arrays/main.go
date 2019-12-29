package main

import "fmt"

func main() {
	// nums1 := []int{1, 2}
	// nums2 := []int{3, 4}
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l%2 == 0 {
		v1 := find(nums1, nums2, l/2)
		v2 := find(nums1, nums2, l/2+1)
		return float64(v1+v2) / 2
	}
	return float64(find(nums1, nums2, l/2+1))
}

func find(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}

	if k == 1 {
		n1, n2 := nums1[0], nums2[0]
		if n1 < n2 {
			return n1
		}
		return n2
	}

	if len(nums1) > len(nums2) {
		return find(nums2, nums1, k)
	}

	i1 := k / 2
	i2 := i1
	if i1 > len(nums1) {
		i1 = len(nums1)
	}
	if i2 > len(nums2) {
		i2 = len(nums2)
	}

	if nums1[i1-1] > nums2[i2-1] {
		return find(nums1, nums2[i2:], k-i2)
	}
	return find(nums1[i1:], nums2, k-i1)
}
