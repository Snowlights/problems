package _500_1600

import "sort"

// 1552
func maxDistance(position []int, m int) int {
	sort.Ints(position)
	n := len(position)
	check := func(num int) bool {
		pre := position[0]
		cnt := 1
		for i := 1; i < n; i++ {
			if position[i]-pre >= num {
				pre = position[i]
				cnt++
			}
		}
		return cnt < m
	}
	idx := sort.Search(position[len(position)-1], func(i int) bool {
		return check(i)
	})
	return idx - 1
}

// 1557
func findSmallestSetOfVertices(n int, edges [][]int) []int {

	indegree := make(map[int]int)
	g := make(map[int][]int)
	vis := make(map[int]bool)
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
		indegree[edge[1]]++
	}

	ans := []int{}
	for i := 0; i < n; i++ {
		if indegree[i] != 0 || vis[i] {
			continue
		}
		ans = append(ans, i)
		vis[i] = true
		q := []int{i}
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, v := range tmp {
				for _, v2 := range g[v] {
					if !vis[v2] {
						q = append(q, v2)
						vis[v2] = true
					}
				}
			}
		}

	}
	return ans
}

// 1559
func containsCycle(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	type pair struct {
		x, y, lx, ly int
	}

	for i, v := range grid {
		for j := range v {
			if !vis[i][j] {
				vis[i][j] = true
				q := []pair{{
					x:  i,
					y:  j,
					lx: i,
					ly: j,
				}}
				for len(q) > 0 {
					tmp := q
					q = nil
					for _, vv := range tmp {
						for _, d := range dir {
							x, y := vv.x+d[0], vv.y+d[1]
							if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == grid[vv.x][vv.y] {
								if x == vv.lx && y == vv.ly {
									continue
								}
								if vis[x][y] {
									return true
								}
								vis[x][y] = true
								q = append(q, pair{
									x:  x,
									y:  y,
									lx: vv.x,
									ly: vv.y,
								})
							}
						}
					}
				}
			}
		}
	}

	return false
}

// 1567
func getMaxLen(nums []int) (ans int) {
	pos, neg := 0, 0
	for _, num := range nums {
		if num > 0 {
			pos++
			if neg > 0 {
				neg++
			}
		} else if num == 0 {
			pos, neg = 0, 0
		} else {
			if neg > 0 {
				pos, neg = neg+1, pos+1
			} else {
				pos, neg = 0, pos+1
			}
		}
		ans = max(ans, pos)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 1574
func findLengthOfShortestSubarray(arr []int) int {

	n := len(arr)
	if n <= 1 {
		return 0
	}

	pre := 1
	for i := 1; i < n && arr[i] >= arr[i-1]; i++ {
		pre++
	}
	if pre == n {
		return 0
	}

	suf := 1
	for i := n - 1; i > 0 && arr[i] >= arr[i-1]; i-- {
		suf++
	}

	min := n - pre
	if n-suf < min {
		min = n - suf
	}

	j := n - suf
	for i := 0; i < pre; i++ {
		for ; j < n; j++ {
			if arr[j] >= arr[i] {
				if j-i-1 < min {
					min = j - i - 1
				}
				break
			}
		}
	}

	return min
}
