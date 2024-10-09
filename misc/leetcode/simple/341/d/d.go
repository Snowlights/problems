package main

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建树
	}

	cnt := make([]int, n)
	for _, t := range trips {
		end := t[1]
		var dfs func(int, int) bool
		dfs = func(x, fa int) bool {
			if x == end { // 到达终点（注意树只有唯一的一条简单路径）
				cnt[x]++    // 统计从 start 到 end 的路径上的点经过了多少次
				return true // 找到终点
			}
			for _, y := range g[x] {
				if y != fa && dfs(y, x) {
					cnt[x]++ // 统计从 start 到 end 的路径上的点经过了多少次
					return true
				}
			}
			return false // 未找到终点
		}
		dfs(t[0], -1)
	}

	// 类似 337. 打家劫舍 III https://leetcode.cn/problems/house-robber-iii/
	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		notHalve := price[x] * cnt[x] // x 不变
		halve := notHalve / 2         // x 减半
		for _, y := range g[x] {
			if y != fa {
				nh, h := dfs(y, x)     // 计算 y 不变/减半的最小价值总和
				notHalve += min(nh, h) // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
				halve += nh            // x 减半，那么 y 只能不变
			}
		}
		return notHalve, halve
	}
	nh, h := dfs(0, -1)
	return min(nh, h)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
