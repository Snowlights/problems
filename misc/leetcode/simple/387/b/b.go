package main

func countSubmatrices(grid [][]int, k int) int {
	m, n, ans := len(grid), len(grid[0]), 0
	a := make([][]int, m+1)
	for i := range a {
		a[i] = make([]int, n+1)
	}
	for i, v := range grid {
		for j, vv := range v {
			a[i+1][j+1] = a[i+1][j] + a[i][j+1] - a[i][j] + vv
			if a[i+1][j+1] <= k {
				ans++
			}
		}
	}
	return ans
}
