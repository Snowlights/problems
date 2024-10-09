package main

import (
	_ "math/bits"
)

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	dis := make([][]int, 26)
	for i := range dis {
		dis[i] = make([]int, 26)
		for j := range dis[i] {
			if j != i {
				dis[i][j] = 1e12
			}
		}
	}

	for i := range original {
		c, d := int(original[i]-'a'), int(changed[i]-'a')
		dis[c][d] = min(dis[c][d], cost[i])
	}

	for k := range dis {
		for i := range dis {
			for j := range dis {
				// 不选 k，选 k
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}
	ans := 0
	for i := range source {
		if source[i] != target[i] {
			ans += dis[int(source[i]-'a')][(target[i] - 'a')]
		}
	}
	if ans >= 1e12 {
		return -1
	}
	return int64(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
