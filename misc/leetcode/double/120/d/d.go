package main

import (
	"sort"
)

func placedCoins(edges [][]int, cost []int) []int64 {
	g := make(map[int][]int)
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], v[1])
		g[v[1]] = append(g[v[1]], v[0])
	}

	ans := make([]int64, len(cost))
	for i := range ans {
		ans[i] = 1
	}
	var dfs func(x, fa int) []int
	dfs = func(x, fa int) []int {
		a := []int{cost[x]}
		for _, to := range g[x] {
			if to == fa {
				continue
			}
			a = append(a, dfs(to, x)...)
		}
		sort.Ints(a)
		n := len(a)
		if len(a) < 3 {
			ans[x] = 1
		} else {
			ans[x] = int64(max(a[0]*a[1]*a[n-1], a[n-3]*a[n-2]*a[n-1], 0))
		}
		if len(a) > 5 {
			a = append(a[:2], a[n-3:]...)
		}
		return a
	}
	dfs(0, -1)
	return ans
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
