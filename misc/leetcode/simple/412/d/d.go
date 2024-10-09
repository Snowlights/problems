package main

import "slices"

var pow10 = [...]int{1, 10, 100, 1000, 10000, 100000, 1000000}

func countPairs(nums []int) int {
	slices.Sort(nums)
	ans := 0
	cnt := make(map[int]int)
	a := [len(pow10)]int{}
	for _, x := range nums {
		st := map[int]struct{}{x: {}} // 不交换
		m := 0
		for v := x; v > 0; v /= 10 {
			a[m] = v % 10
			m++
		}
		for i := 0; i < m; i++ {
			for j := i + 1; j < m; j++ {
				if a[i] == a[j] { // 小优化
					continue
				}
				y := x + (a[j]-a[i])*(pow10[i]-pow10[j])
				st[y] = struct{}{} // 交换一次
				a[i], a[j] = a[j], a[i]
				for p := i + 1; p < m; p++ {
					for q := p + 1; q < m; q++ {
						st[y+(a[q]-a[p])*(pow10[p]-pow10[q])] = struct{}{} // 交换两次
					}
				}
				a[i], a[j] = a[j], a[i]
			}
		}
		for x := range st {
			ans += cnt[x]
		}
		cnt[x]++
	}
	return ans
}
