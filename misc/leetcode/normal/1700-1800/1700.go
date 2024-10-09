package _700_1800

import "sort"

// 1706
func findBall(grid [][]int) []int {
	n := len(grid[0])
	ans := make([]int, n)
	for j := range ans {
		col := j // 球的初始列
		for _, row := range grid {
			dir := row[col]
			col += dir                                  // 移动球
			if col < 0 || col == n || row[col] != dir { // 到达侧边或 V 形
				col = -1
				break
			}
		}
		ans[j] = col // col >= 0 为成功到达底部
	}
	return ans
}

// 1712
func waysToSplit(a []int) (ans int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	for r := 2; r < n && 3*sum[r] <= 2*sum[n]; r++ {
		l1 := sort.SearchInts(sum[1:r], 2*sum[r]-sum[n]) + 1
		ans += sort.SearchInts(sum[l1:r], sum[r]/2+1)
	}
	return ans % (1e9 + 7)
}

// 1722
func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(from int) int {
		if fa[from] != from {
			fa[from] = find(fa[from])
		}
		return fa[from]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
	}

	for _, v := range allowedSwaps {
		merge(v[0], v[1])
	}
	sMap, tMap := make(map[int][]int), make([]map[int]int, n)
	for i := range source {
		to := find(i)
		sMap[to] = append(sMap[to], source[i])
		tm := tMap[to]
		if tm == nil {
			tm = make(map[int]int)
			tMap[to] = tm
		}
		tm[target[i]]++
	}
	ans := 0
	for i, s := range sMap {
		for _, v := range s {
			to := find(i)
			tm := tMap[to]
			if tm != nil && tm[v] > 0 {
				tm[v]--
				continue
			}
			ans++
		}
	}

	return ans
}
