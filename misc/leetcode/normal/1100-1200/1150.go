package _100_1200

import "math"

// 1162
func maxDistance(grid [][]int) int {
	ans := 0
	a, b := [][]int{}, [][]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				a = append(a, []int{i, j})
			} else {
				b = append(b, []int{i, j})
			}
		}
	}

	if len(a) == 0 || len(b) == 0 {
		return -1
	}

	for _, v := range b {
		tmp := math.MaxInt32
		for _, u := range a {
			val := abs(v[0]-u[0]) + abs(v[1]-u[1])
			tmp = min(tmp, val)
		}
		ans = max(ans, tmp)
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
