package _00_800

import "math"

// 733
var (
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	currColor := image[sr][sc]
	if currColor != newColor {
		dfs(image, sr, sc, currColor, newColor)
	}
	return image
}

func dfs(image [][]int, x, y, color, newColor int) {
	if image[x][y] == color {
		image[x][y] = newColor
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			if mx >= 0 && mx < len(image) && my >= 0 && my < len(image[0]) {
				dfs(image, mx, my, color, newColor)
			}
		}
	}
}

// 740
func deleteAndEarn(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sum := make([]int, 1e4+5)
	for _, v := range nums {
		sum[v] += v
	}
	dp := make([]int64, 1e4+5)
	dp[1] = int64(sum[1])
	dp[2] = max(int64(dp[1]), int64(sum[2]))
	for i := 3; i < 1e4+5; i++ {
		dp[i] = max(dp[i-2]+int64(sum[i]), dp[i-1])
	}
	return int(dp[1e4+4])
}

// 746
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i, v := range cost {
		if i > 0 {
			dp[i] = min(dp[i], dp[i-1]+v)
		}
		if i > 1 {
			dp[i] = min(dp[i], dp[i-2]+v)
		}
	}

	return min(dp[len(cost)-1], dp[len(cost)-2])
}
