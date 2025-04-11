package main

import (
	"slices"
)

func makeStringGood(s string) int {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	m := slices.Max(cnt[:])

	ans := len(s)
	f := [27]int{}
	for target := 0; target <= m; target++ {
		f[25] = min(cnt[25], abs(cnt[25]-target))
		for i := 24; i >= 0; i-- {
			x, y := cnt[i], cnt[i+1]
			// 单独操作 x（变成 target 或 0）
			f[i] = f[i+1] + min(x, abs(x-target))
			// x 变成 target 或 0，y 变成 target
			if y < target { // 只有当 y 需要变大时，才去执行第三种操作
				if x > target { // x 变成 target
					f[i] = min(f[i], f[i+2]+max(x-target, target-y))
				} else { // x 变成 0
					f[i] = min(f[i], f[i+2]+max(x, target-y))
				}
			}
		}
		ans = min(ans, f[0])
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
