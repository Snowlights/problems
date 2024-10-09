package main

func maximumSetSize(nums1 []int, nums2 []int) int {
	s1, s2 := make(map[int]bool), make(map[int]bool)
	n, common := len(nums1), 0
	for _, v := range nums1 {
		s1[v] = true
	}
	for _, v := range nums2 {
		if s2[v] {
			continue
		}
		s2[v] = true
		if s1[v] {
			common++
		}
	}

	c1 := min(n/2, len(s1)-common)
	c2 := min(n/2, len(s2)-common)
	return min(n, c1+c2+common)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
