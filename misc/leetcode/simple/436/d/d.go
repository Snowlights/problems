package main

import (
	"slices"
	"sort"
)

func maxScore(points []int, m int) int64 {
	right := (m + 1) / 2 * slices.Min(points)
	ans := sort.Search(right, func(low int) bool {
		// 二分最小的不满足要求的 low+1，即可得到最大的满足要求的 low
		low++
		left := m
		pre := 0
		for i, p := range points {
			k := (low-1)/p + 1 - pre          // 还需要操作的次数
			if i == len(points)-1 && k <= 0 { // 最后一个数已经满足要求
				break
			}
			k = max(k, 1)   // 至少要走 1 步
			left -= k*2 - 1 // 左右横跳
			if left < 0 {
				return true
			}
			pre = k - 1 // 右边那个数顺带操作了 k-1 次
		}
		return false
	})
	return int64(ans)
}
