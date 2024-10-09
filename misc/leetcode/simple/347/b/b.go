package main

func differenceOfDistinctValues(g [][]int) (ans [][]int) {
	n, m := len(g), len(g[0])

	top, button := func(i, j int) int {
		h := make(map[int]bool)
		i--
		j--
		for i >= 0 && j >= 0 {
			h[g[i][j]] = true
			i--
			j--
		}
		return len(h)
	}, func(i, j int) int {
		h := make(map[int]bool)
		i++
		j++
		for i < n && j < m {
			h[g[i][j]] = true
			i++
			j++
		}
		return len(h)
	}

	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
		for j := range ans[i] {
			ans[i][j] = abs(top(i, j) - button(i, j))
		}
	}
	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
