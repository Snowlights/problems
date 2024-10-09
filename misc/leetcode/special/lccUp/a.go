package lccUp

import "sort"

// 1
func temperatureTrend(temperatureA []int, temperatureB []int) int {
	ans, tmp := 0, 0
	for i := 1; i < len(temperatureA); i++ {
		if temperatureA[i] > temperatureA[i-1] {
			if temperatureB[i] > temperatureB[i-1] {
				tmp++
			} else {
				tmp = 0
			}
			ans = max(ans, tmp)
		} else if temperatureA[i] == temperatureA[i-1] {
			if temperatureB[i] == temperatureB[i-1] {
				tmp++
			} else {
				tmp = 0
			}
			ans = max(ans, tmp)
		} else if temperatureA[i] < temperatureA[i-1] {
			if temperatureB[i] < temperatureB[i-1] {
				tmp++
			} else {
				tmp = 0
			}
			ans = max(ans, tmp)
		}
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 2
func transportationHub(path [][]int) int {

	inDegree, outDegree := make(map[int]int), make(map[int]int)
	h := make(map[int]bool)
	for _, v := range path {
		outDegree[v[0]]++
		inDegree[v[1]]++
		h[v[0]] = true
		h[v[1]] = true
	}
	for i := 0; i <= 1000; i++ {
		if inDegree[i] == len(h)-1 && outDegree[i] == 0 {
			return i
		}
	}

	return -1
}

// 3
func ballGame(k int, g []string) (ans [][]int) {
	var dir4 = []struct{ x, y int }{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // 上右下左
	n, m := len(g), len(g[0])
	vis := make([][][4]bool, n)
	type pair struct{ x, y, di int }
	q := []pair{}
	for i := range vis {
		vis[i] = make([][4]bool, m)
	}
	for i, r := range g {
		for j, v := range r {
			if v == 'O' {
				for k := 0; k < 4; k++ {
					vis[i][j][k] = true
					q = append(q, pair{i, j, k})
				}
			}
		}
	}
	cs := [][2]int{{0, 0}, {0, m - 1}, {n - 1, 0}, {n - 1, m - 1}}
	inans := map[[2]int]bool{}
	for step := 1; len(q) > 0 && step <= k+1; step++ {
		tmp := q
		q = nil
	oo:
		for _, p := range tmp {
			di := p.di
			x, y := p.x, p.y
			b := g[x][y]
			if b == 'e' {
				di = (di + 3) % 4
			} else if b == 'W' {
				di = (di + 1) % 4
			}
			x += dir4[di].x
			y += dir4[di].y
			if 0 <= x && x < n && 0 <= y && y < m {
				if !vis[x][y][di] {
					vis[x][y][di] = true
					q = append(q, pair{x, y, di})
				}
			} else {
				pp := [2]int{p.x, p.y}
				for _, c := range cs {
					if c == pp {
						continue oo
					}
				}
				if g[p.x][p.y] == '.' && !inans[pp] {
					inans[pp] = true
					ans = append(ans, []int{p.x, p.y})
				}
			}
		}
	}
	sort.Slice(ans, func(i, j int) bool { a, b := ans[i], ans[j]; return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] })
	return
}

// 4
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closeLampInTree(root *TreeNode) (ans int) {
	type pair struct {
		o        *TreeNode
		tp2, tp3 bool
	}
	dp := map[pair]int{}
	var f func(*TreeNode, bool, bool) int
	f = func(o *TreeNode, all, t bool) (res int) {
		if o == nil {
			return
		}
		p := pair{o, all, t}
		if v, ok := dp[p]; ok {
			return v
		}
		if o.Val == 1 == (all == t) {
			r1 := f(o.Left, all, false) + f(o.Right, all, false) + 1
			r2 := f(o.Left, !all, false) + f(o.Right, !all, false) + 1
			r3 := f(o.Left, all, true) + f(o.Right, all, true) + 1
			r123 := f(o.Left, !all, true) + f(o.Right, !all, true) + 3
			res = min(min(r1, r2), min(r3, r123))
		} else {
			r0 := f(o.Left, all, false) + f(o.Right, all, false)
			r12 := f(o.Left, !all, false) + f(o.Right, !all, false) + 2
			r13 := f(o.Left, all, true) + f(o.Right, all, true) + 2
			r23 := f(o.Left, !all, true) + f(o.Right, !all, true) + 2
			res = min(min(r0, r12), min(r13, r23))
		}
		dp[p] = res
		return
	}
	ans = f(root, false, false)
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 5
func unSuitability(a []int) (ans int) {
	mx := int(-1e9)
	for _, v := range a {
		if v > mx {
			mx = v
		}
	}
	n := len(a)
	// mx*4
	ans = sort.Search(mx*4, func(lim int) bool {
		if a[0] > lim {
			return false
		}
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, lim+1)
			for j := range dp[i] {
				dp[i][j] = -1e9
			}
		}
		var f func(int, int) int
		f = func(i, pos int) int {
			if i == n {
				return 0
			}
			dv := &dp[i][pos]
			if *dv != -1e9 {
				return *dv
			}
			res := int(-1e8)

			if pos+a[i] <= lim {
				res2 := f(i+1, pos+a[i])
				res = min(res2+a[i], 0)
			}

			res2 := f(i+1, max(pos-a[i], 0))
			if res2-a[i] >= -lim {
				res = max(res, res2-a[i])
			}

			*dv = res
			return res
		}
		ans = f(0, 0)
		return ans >= -lim
	})
	return
}
