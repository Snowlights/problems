package main

import "sort"

func minimumAddedInteger(nums1, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i := 2; ; i-- {
		diff := nums2[0] - nums1[i]
		j := 0
		for _, v := range nums1[i:] {
			if j < len(nums2) && nums2[j]-v == diff {
				j++
				if j == len(nums2) {
					return diff
				}
			}
		}
	}
}
