package _00_400

import "math"

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
