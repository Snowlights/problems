package double_97

import (
	"strconv"
)

// 1
func separateDigits(nums []int) []int {
	ans := make([]int, 0)
	for _, v := range nums {
		for _, val := range strconv.Itoa(v) {
			vv, _ := strconv.Atoi(string(val))
			ans = append(ans, vv)
		}
	}
	return ans
}

// 2
func maxCount(banned []int, n int, maxSum int) int {
	h := make(map[int]bool)
	for _, v := range banned {
		h[v] = true
	}
	sum, cnt := 0, 0
	for i := 1; i <= n; i++ {
		if h[i] {
			continue
		}
		sum += i
		if sum > maxSum {
			break
		}
		cnt++
	}
	return cnt
}

// 3
func maximizeWin(prizePositions []int, k int) int {
	pre := make([]int, len(prizePositions)+1)
	ans, left := 0, 0
	for right, p := range prizePositions {
		for p-prizePositions[left] > k {
			left++
		}
		ans = max(ans, right-left+1+pre[left])
		pre[right+1] = max(pre[right], right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 4
func isPossibleToCutPath(g [][]int) bool {
	m, n := len(g), len(g[0])
	var dfs func(int, int) bool
	dfs = func(x, y int) bool {
		if x == m-1 && y == n-1 {
			return true
		}
		g[x][y] = 0
		return x < m-1 && g[x+1][y] > 0 && dfs(x+1, y) ||
			y < n-1 && g[x][y+1] > 0 && dfs(x, y+1)
	}
	return !dfs(0, 0) || !dfs(0, 0)
}
