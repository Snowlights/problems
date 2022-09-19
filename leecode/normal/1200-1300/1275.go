package _200_1300

import "sort"

// 1281
func groupThePeople(groupSizes []int) [][]int {

	h := make(map[int][]int)
	for i, v := range groupSizes {
		h[v] = append(h[v], i)
	}
	ans := make([][]int, 0)
	for k, v := range h {
		for i := 0; i < len(v); i += k {
			ans = append(ans, v[i:i+k])
		}
	}
	return ans
}

// 1283
func smallestDivisor(nums []int, threshold int) int {
	return sort.Search(1e9, func(i int) bool {
		if i == 0 {
			return false
		}
		val := 0
		for _, v := range nums {
			val += v / i
			if v%i > 0 {
				val++
			}
		}
		return val <= threshold
	})
}
