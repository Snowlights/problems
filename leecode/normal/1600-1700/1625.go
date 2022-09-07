package _600_1700

import "sort"

// 1630
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var n = len(l)
	var res = make([]bool, n)
	for i := 0; i < n; i++ {
		s, e := l[i], r[i]
		if s >= e {
			res[i] = false
			continue
		}
		tempArr := make([]int, e-s+1)
		copy(tempArr, nums[s:e+1])
		sort.Ints(tempArr)
		ans := true
		temp := tempArr[1] - tempArr[0]
		for j := 1; j < e-s+1; j++ {
			if tempArr[j]-tempArr[j-1] != temp {
				ans = false
				break
			}
		}
		res[i] = ans
	}
	return res
}
