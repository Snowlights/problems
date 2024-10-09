package main

func minimumOperationsToWriteY(grid [][]int) int {
	y, noty, m := [3]int{}, [3]int{}, len(grid)
	for i, v := range grid[:m/2] {
		for j, vv := range v {
			if i == j || i == m-1-j {
				y[vv]++
			} else {
				noty[vv]++
			}
		}
	}
	for _, v := range grid[m/2:] {
		for j, vv := range v {
			if m/2 == j {
				y[vv]++
			} else {
				noty[vv]++
			}
		}
	}
	mx := 0
	for i, v := range y {
		for j, vv := range noty {
			if i != j {
				if v+vv > mx {
					mx = v + vv
				}
			}
		}
	}
	return m*m - mx
}
