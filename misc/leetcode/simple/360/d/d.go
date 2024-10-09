package main

import "math/bits"

func getMaxFunctionValue(receiver []int, K int64) int64 {
	type pair struct{ pa, sum int }
	n, m := len(receiver), bits.Len(uint(K))
	pa := make([][]pair, n)
	for i, p := range receiver {
		pa[i] = make([]pair, m)
		pa[i][0] = pair{p, p}
	}
	for i := 0; i+1 < m; i++ {
		for x := range pa {
			p := pa[x][i]
			pp := pa[p.pa][i]
			pa[x][i+1] = pair{pp.pa, p.sum + pp.sum} // 合并节点值之和
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		x := i
		sum := i // 节点值之和，初始为节点 i
		for k := uint(K); k > 0; k &= k - 1 {
			p := pa[x][bits.TrailingZeros(k)]
			sum += p.sum
			x = p.pa
		}
		ans = max(ans, sum)
	}
	return int64(ans)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
