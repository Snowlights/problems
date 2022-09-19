package _800_1900

import "sort"

// 1818
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	vec := append(sort.IntSlice{}, nums1...)
	sort.Ints(vec)

	sum, maxn, n := 0, 0, len(nums1)
	for i, v := range nums2 {
		diff := abs(nums1[i] - v)

		sum += diff
		j := vec.Search(v)
		if j < n {
			maxn = max(maxn, diff-(vec[j]-v))
		}
		if j > 0 {
			maxn = max(maxn, diff-(v-vec[j-1]))
		}
	}

	return (sum - maxn) % (1e9 + 7)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
