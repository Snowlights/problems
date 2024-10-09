package _800_1900

import "sort"

// 1838
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	pre := make([]int, len(nums)+1)
	for i, v := range nums {
		pre[i+1] = pre[i] + v
	}
	ans := 0

	for r, v := range nums {
		l := sort.Search(r, func(i int) bool {
			return (r-i)*v-pre[r]+pre[i] <= k
		})
		ans = max(ans, r-l+1)
	}

	return ans
}
