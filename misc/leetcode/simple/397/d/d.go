package main

import (
	"math"
)

func findPermutation(a []int) []int {
	n := len(a)
	u := 1<<n - 1
	f := make([][]int, 1<<n)
	g := make([][]int, 1<<n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = math.MaxInt
		}
		g[i] = make([]int, n)
	}
	for j := range f[u] {
		f[u][j] = abs(j - a[0])
		g[u][j] = -1
	}
	for s := u - 2; s > 0; s -= 2 { // 注意偶数不含 0，是无效状态
		for j := 0; j < n; j++ {
			if s>>j&1 == 0 { // 无效状态，因为 j 一定在 s 中
				continue
			}
			for k := 1; k < n; k++ {
				if s>>k&1 > 0 { // k 之前填过
					continue
				}
				v := f[s|1<<k][k] + abs(j-a[k])
				if v < f[s][j] {
					f[s][j] = v
					g[s][j] = k // 记录该状态下填了哪个数
				}
			}
		}
	}

	ans := make([]int, 0, n)
	for s, j := 0, 0; j >= 0; {
		ans = append(ans, j)
		s |= 1 << j
		j = g[s][j]
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
