package main

import "strconv"

func findColumnWidth(grid [][]int) []int {
	ans := make([]int, len(grid[0]))
	for j := range grid[0] {
		for _, row := range grid {
			ans[j] = max(ans[j], len(strconv.Itoa(row[j])))
		}
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
