package main

func findMissingAndRepeatedValues(grid [][]int) []int {
	h, ans, n := make(map[int]int), make([]int, 2), len(grid)
	for _, v := range grid {
		for _, vv := range v {
			h[vv]++
			if h[vv] == 2 {
				ans[0] = vv
			}
		}
	}
	for i := 1; i <= n*n; i++ {
		if h[i] == 0 {
			ans[1] = i
			break
		}
	}
	return ans
}
