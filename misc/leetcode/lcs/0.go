package lcs

import "sort"

// lcs 01
func leastMinutes(n int) int {

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		a, b := dp[i-1]+1, dp[i/2]
		if i%2 == 1 {
			b = dp[i/2+1]
		}
		dp[i] = min(a, b+1)
	}
	return dp[n]
}

// lcs 02
func halfQuestions(questions []int) int {
	type pair struct {
		h, v int
	}
	n := len(questions) / 2
	pMap, pList := make(map[int]*pair), make([]*pair, 0)
	for i := 0; i < len(questions); i++ {
		p, ok := pMap[questions[i]]
		if !ok {
			p = &pair{
				h: questions[i],
				v: 0,
			}
			pList = append(pList, p)
			pMap[questions[i]] = p
		}
		p.v++
	}
	sort.Slice(pList, func(i, j int) bool {
		return pList[i].v > pList[j].v
	})
	ans := 0
	for _, v := range pList {
		if n <= 0 {
			break
		}
		ans++
		n -= v.v
	}
	return ans
}

// lcs 03
func largestArea(grid []string) int {
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	ans := 0
	for i, v := range grid {
		for j := range v {
			if v[j] == '0' || vis[i][j] {
				continue
			}
			vis[i][j] = true
			base := v[j]
			q, flag, res := [][]int{{i, j}}, false, 1
			for len(q) > 0 {
				tmp := q
				q = nil
				for _, t := range tmp {
					for _, d := range dir {
						x, y := t[0]+d[0], t[1]+d[1]
						if x < 0 || x >= m || y < 0 || y >= n {
							flag = true
						} else {
							if grid[x][y] == '0' {
								flag = true
							}
							if !vis[x][y] && grid[x][y] == base {
								q = append(q, []int{x, y})
								vis[x][y] = true
								res++
							}
						}
					}
				}
			}
			if !flag && res > ans {
				ans = res
			}
		}
	}

	return ans
}
