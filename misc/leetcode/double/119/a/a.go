package main

func findIntersectionValues(nums1, nums2 []int) []int {
	h, c := make(map[int]bool), make(map[int]bool)
	for _, v := range nums1 {
		h[v] = true
	}
	ans := make([]int, 2)
	for _, v := range nums2 {
		if h[v] {
			ans[1]++
		}
		c[v] = true
	}
	for _, v := range nums1 {
		if c[v] {
			ans[0]++
		}
	}
	return ans
}
