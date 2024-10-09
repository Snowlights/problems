package main

func numberOfRightTriangles(grid [][]int) int64 {

	m, n := len(grid), len(grid[0])
	r, c := make([]int, m), make([]int, n)
	for i, v := range grid {
		for j, vv := range v {
			if vv == 0 {
				continue
			}
			r[i]++
			c[j]++
		}
	}

	ans := 0
	for i, v := range grid {
		for j, vv := range v {
			if vv == 0 {
				continue
			}
			ans += (r[i] - 1) * (c[j] - 1)
		}
	}
	return int64(ans)
}
