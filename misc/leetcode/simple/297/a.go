package simple_297

import (
	"math"
)

// 1
func calculateTax(brackets [][]int, income int) float64 {
	ans := float64(0)
	val := brackets[0][0]
	if val > income {
		val = income
	}
	ans += float64(val) * float64(brackets[0][1]) / 100
	for i := 1; i < len(brackets); i++ {
		val = brackets[i][0]
		if val > income {
			val = income
		}
		if val-brackets[i-1][0] < 0 {
			continue
		}
		ans += float64(val-brackets[i-1][0]) * float64(brackets[i][1]) / 100
	}
	return ans
}

// 2
func minPathCost(grid [][]int, moveCost [][]int) int {

	dp := make([][]int, len(grid)-1)
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	for i, g := range grid[0] {
		for k, c := range moveCost[g] {
			dp[0][k] = min(dp[0][k], c+grid[0][i])
		}
	}

	for i := 1; i < len(grid)-1; i++ {
		for j, g := range grid[i] {
			for k, c := range moveCost[g] {
				dp[i][k] = min(dp[i][k], dp[i-1][j]+c+grid[i][j])
			}
		}
	}

	ans := math.MaxInt32
	for i, v := range dp[len(grid)-2] {
		ans = min(ans, v+grid[len(grid)-1][i])
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 3 目前不会
func distributeCookies(cookies []int, k int) int {
	return 0
}

// 4
func distinctNames(ideas []string) int64 {

	headerMap := make(map[byte]bool)
	exitMap := make(map[string]bool)
	val := make([][]int, 26)
	for i := 0; i < 26; i++ {
		val[i] = make([]int, 26)
	}
	for _, v := range ideas {
		exitMap[v] = true
		headerMap[v[0]] = true
	}

	for _, idea := range ideas {
		for v, _ := range headerMap {
			s := string(v) + idea[1:]
			if exitMap[s] {
				continue
			}
			val[idea[0]-'a'][v-'a']++
		}
	}

	var ans int64 = 0
	for header, _ := range headerMap {
		for h, _ := range headerMap {
			ans += int64(val[header-'a'][h-'a'] * val[h-'a'][header-'a'])
		}
	}

	return ans
}
