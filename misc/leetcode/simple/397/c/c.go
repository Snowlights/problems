package main

import "math"

func maxScore(grid [][]int) int {
	ans := math.MinInt
	colMin := make([]int, len(grid[0]))
	for i := range colMin {
		colMin[i] = math.MaxInt
	}
	for _, row := range grid {
		preMin := math.MaxInt
		for j, x := range row {
			ans = max(ans, x-min(preMin, colMin[j]))
			colMin[j] = min(colMin[j], x)
			preMin = min(preMin, colMin[j])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
