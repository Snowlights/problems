package main

import (
	_ "math/bits"
	"sort"
)

func minimumCost(nums []int) int64 {
	pal := make([]int, 0, 109999)
	// 哨兵。根据题目来定，也可以设置成 -2e9 等
	pal = append(pal, 0)
	initPalindromeNumber := func() {
		const mx int = 1e9
		// 严格按顺序从小到大生成所有回文数
	outer:
		for base := 1; ; base *= 10 {
			// 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
			for i := base; i < base*10; i++ {
				x := i
				for t := i / 10; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				if x > mx {
					break outer
				}
				pal = append(pal, x)
			}
			// 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
			for i := base; i < base*10; i++ {
				x := i
				for t := i; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				if x > mx {
					break outer
				}
				pal = append(pal, x)
			}
		}

		// 哨兵。根据 mx 调整，如果 mx 是 2e9 的话要写成 mx+2
		pal = append(pal, mx+1)
	}
	initPalindromeNumber()

	sort.Ints(nums)
	n := len(nums)
	i := sort.SearchInts(pal, nums[n/2])
	cost := func(i int) int {
		res := 0
		for _, v := range nums {
			res += abs(v - i)
		}
		return res
	}
	return int64(min(cost(pal[i-1]), cost(pal[i])))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
