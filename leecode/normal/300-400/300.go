package _00_400

import (
	"math"
	"sort"
)

// 300
func lengthOfLIS(nums []int) int {

	//使得f按照升序排列
	f := []int{}
	for _, e := range nums {
		//h := e
		if i := sort.SearchInts(f, e); i < len(f) {
			f[i] = e
		} else {
			f = append(f, e)
		}
	}
	return len(f)
}

// 309
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	f0, f1, f2 := -prices[0], 0, 0
	for i := 1; i < n; i++ {
		newf0 := max(f0, f2-prices[i])
		newf1 := f0 + prices[i]
		newf2 := max(f1, f2)
		f0, f1, f2 = newf0, newf1, newf2
	}
	return max(f1, f2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 310
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		deg[x]++
		deg[y]++
	}

	q := []int{}
	for i, d := range deg {
		if d == 1 {
			q = append(q, i)
		}
	}

	remainNodes := n
	for remainNodes > 2 {
		remainNodes -= len(q)
		tmp := q
		q = nil
		for _, x := range tmp {
			for _, y := range g[x] {
				deg[y]--
				if deg[y] == 1 {
					q = append(q, y)
				}
			}
		}
	}
	return q
}

// 322
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	for _, v := range coins {
		if v > amount {
			continue
		}
		dp[v] = 1
	}
	for i := 0; i <= amount; i++ {
		for _, c := range coins {
			if c > i {
				continue
			}
			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
