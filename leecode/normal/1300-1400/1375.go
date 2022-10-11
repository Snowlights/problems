package _300_1400

import "math"

// 1376
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {

	g, dp := make(map[int][]int), make([]int, n)
	for i, v := range manager {
		g[v] = append(g[v], i)
	}
	q := []int{headID}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, vv := range g[v] {
				dp[vv] = dp[v] + informTime[v]
				q = append(q, vv)
			}
		}
	}
	ans := dp[0]
	for _, v := range dp {
		ans = max(ans, v)
	}

	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 1385
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var res int
	for i := 0; i < len(arr1); i++ {
		isDistanceValue := true
		for j := 0; j < len(arr2); j++ {
			if int(math.Abs(float64(arr1[i]-arr2[j]))) <= d {
				isDistanceValue = false
				break
			}
		}
		if isDistanceValue {
			res++
		}
	}
	return res
}

// 1392
func longestPrefix(s string) string {
	ans := ""
	for i := 1; i < len(s); i++ {
		if s[:i] == s[len(s)-i:] {
			ans = s[:i]
		}
	}
	return ans
}
