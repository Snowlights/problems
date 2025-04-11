package main

func comb2(num int) int {
	return num * (num - 1) / 2
}

func subsequencesWithMiddleMode(nums []int) int {
	n := len(nums)
	ans := n * (n - 1) * (n - 2) * (n - 3) * (n - 4) / 120 // 所有方案数
	suf := map[int]int{}
	for _, num := range nums {
		suf[num]++
	}
	pre := make(map[int]int, len(suf)) // 预分配空间
	// 枚举 x，作为子序列正中间的数
	for left, x := range nums[:n-2] {
		suf[x]--
		if left > 1 {
			right := n - 1 - left
			preX, sufX := pre[x], suf[x]
			// 不合法：只有一个 x
			ans -= comb2(left-preX) * comb2(right-sufX)
			// 不合法：只有两个 x，且至少有两个 y（y != x）
			for y, sufY := range suf {
				if y == x {
					continue
				}
				preY := pre[y]
				// 左边有两个 y，右边有一个 x，即 yy x xz（z 可以等于 y）
				ans -= comb2(preY) * sufX * (right - sufX)
				// 右边有两个 y，左边有一个 x，即 zx x yy（z 可以等于 y）
				ans -= comb2(sufY) * preX * (left - preX)
				// 左右各有一个 y，另一个 x 在左边，即 xy x yz（z != y）
				ans -= preY * sufY * preX * (right - sufX - sufY)
				// 左右各有一个 y，另一个 x 在右边，即 zy x xy（z != y）
				ans -= preY * sufY * sufX * (left - preX - preY)
			}
		}
		pre[x]++
	}
	return ans % 1_000_000_007
}
