package _00_1000

import (
	"math"
	"sort"
)

// 926
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	type pair struct {
		diff, pro int
	}
	pairList := make([]*pair, 0)
	for i, v := range difficulty {
		pairList = append(pairList, &pair{
			diff: v,
			pro:  profit[i],
		})
	}
	sort.Slice(pairList, func(i, j int) bool {
		return pairList[i].diff < pairList[j].diff
	})

	ans := 0
	for _, w := range worker {
		tmp, idx := 0, 0
		for idx < len(pairList) && pairList[idx].diff <= w {
			tmp = max(tmp, pairList[idx].pro)
			idx++
		}
		ans += tmp
	}
	return ans
}

// 931
func minFallingPathSum(matrix [][]int) int {

	n := len(matrix)
	dp := make([][]int, n)
	dir := []int{-1, 0, 1}
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range dp {
		dp[0][i] = matrix[0][i]
	}

	for i := 1; i < n; i++ {
		for j := range dp[i] {
			m := math.MaxInt32
			for _, d := range dir {
				y := j + d
				if 0 <= y && y < n {
					m = min(m, dp[i-1][y])
				}
			}
			dp[i][j] = m + matrix[i][j]
		}
	}

	ans := math.MaxInt32
	for _, v := range dp[n-1] {
		ans = min(ans, v)
	}

	return ans
}

// 934
func shortestBridge(grid [][]int) int {

	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	var dfs func(i, j, val int)
	dfs = func(i, j, val int) {
		grid[i][j] = val
		for _, d := range dir {
			x, y := d[0]+i, d[1]+j
			if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) && grid[x][y] == 1 {
				dfs(x, y, val)
			}
		}
	}

	tmp := 2
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				dfs(i, j, tmp)
				tmp++
			}
		}
	}

	n := len(grid)
	ans := math.MaxInt32
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				q, length := [][]int{{i, j}}, 0
				vis := make([][]bool, n)
				for k := range vis {
					vis[k] = make([]bool, n)
				}
				vis[i][j] = true
				for len(q) > 0 {
					t := q
					q = nil
					for _, v := range t {
						x, y := v[0], v[1]
						if grid[x][y] == 3 {
							ans = min(ans, length)
							q = nil
							break
						}
						for _, d := range dir {
							xx, yy := x+d[0], y+d[1]
							if 0 <= xx && xx < n && 0 <= yy && yy < n && grid[xx][yy] != 2 && !vis[xx][yy] {
								q = append(q, []int{xx, yy})
								vis[xx][yy] = true
							}
						}
					}
					length++
				}
			}
		}
	}
	return ans - 1
}

// 940
func distinctSubseqII(s string) int {
	dp := make([][26]int64, len(s))
	mod := int64(1e9 + 7)
	dp[0][s[0]-'a'] = 1

	for i := 1; i < len(s); i++ {
		total := int64(0)
		for j := 0; j < 26; j++ {
			total += dp[i-1][j]
		}
		dp[i] = dp[i-1]
		dp[i][s[i]-'a'] = total%mod + 1
	}

	total := int64(0)
	for _, v := range dp[len(s)-1] {
		total += v
	}
	return int(total % mod)
}

// 946
func validateStackSequences(pushed []int, popped []int) bool {
	s, idx := []int{}, 0
	for _, v := range pushed {
		for len(s) > 0 && s[len(s)-1] == popped[idx] {
			idx++
			s = s[:len(s)-1]
		}
		s = append(s, v)
	}
	for len(s) > 0 && s[len(s)-1] == popped[idx] {
		idx++
		s = s[:len(s)-1]
	}
	return idx == len(popped)
}

// 947
func removeStones(stones [][]int) int {
	g := make(map[int][]int)

	for i := 0; i < len(stones)-1; i++ {
		a := stones[i]
		for j := i + 1; j < len(stones); j++ {
			b := stones[j]
			if a[0] == b[0] || a[1] == b[1] {
				g[i] = append(g[i], j)
				g[j] = append(g[j], i)
			}
		}
	}

	vis, cnt := make(map[int]bool), 0
	for i := 0; i < len(stones); i++ {
		if !vis[i] {
			cnt++
			vis[i] = true
			q := []int{i}
			for len(q) > 0 {
				tmp := q
				q = nil
				for _, v := range tmp {
					for _, to := range g[v] {
						if !vis[to] {
							q = append(q, to)
							vis[to] = true
						}
					}
				}
			}
		}
	}
	return len(stones) - cnt
}
