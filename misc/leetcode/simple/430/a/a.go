package main

func minimumOperations(grid [][]int) (ans int) {
	for j := range grid[0] {
		pre := -1
		for _, row := range grid {
			x := row[j]
			ans += max(pre+1-x, 0)
			pre = max(pre+1, x)
		}
	}
	return
}
