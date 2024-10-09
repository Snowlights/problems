package a

import "sort"

func runeReserve(runes []int) int {
	sort.Ints(runes)
	ans, cnt := 1, 1
	for i, n := 1, len(runes); i < n; i++ {
		if runes[i]-runes[i-1] > 1 {
			cnt = 1 // 重新统计
		} else if cnt++; cnt > ans {
			ans = cnt
		}
	}
	return ans
}
