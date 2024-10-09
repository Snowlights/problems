package _021spring

import (
	"math"
	"sort"
)

// 1
func storeWater(bucket []int, vat []int) int {
	finish := true
	for _, v := range vat {
		if v != 0 {
			finish = false
			break
		}
	}
	if finish {
		return 0
	}

	//需要额外补多少下才能满足能在time次内完成
	ans := int(2 * 1e4)

	for time := 1; time <= int(1e4); time++ {
		finish := true
		add := 0
		for i, _ := range bucket {
			n := int(math.Ceil(float64(vat[i]) / float64(time)))
			if n > bucket[i] {
				add += n - bucket[i]
				finish = false
			}
		}

		if finish {
			return min(ans, time)
		}

		ans = min(time+add, ans)
	}

	return ans

}

// 2
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxValue(root *TreeNode, k int) int {
	var dfs func(root *TreeNode, k int) []int
	dfs = func(root *TreeNode, k int) []int {
		if root == nil {
			return []int{0}
		}
		l, r := dfs(root.Left, k), dfs(root.Right, k)
		res := make([]int, k+1)
		for i := 0; i < len(l); i++ {
			for j := 0; j < len(r); j++ {
				res[0] = max(res[0], l[i]+r[j])
				if i+j+1 <= k {
					res[i+j+1] = max(res[i+j+1], l[i]+r[j]+root.Val)
				}
			}

		}
		return res
	}

	result, ans := dfs(root, k), 0
	for i := 0; i <= len(result); i++ {
		if i <= k {
			ans = max(ans, result[i])
		}
	}
	return ans
}

// 3
func electricCarPlan(paths [][]int, cnt, start, end int, charge []int) int {
	n, inf := len(charge), int(1e9)
	type pair struct {
		first, second int
	}
	const mx = 400
	type pt struct {
		u, crg int
	}
	var (
		f   [mx][mx]int
		vis [mx][mx]bool
		e   [mx][]*pair
		ans int = inf
	)

	for i := 0; i < n; i++ {
		for j := 0; j <= cnt; j++ {
			f[i][j] = inf
		}
	}

	for _, p := range paths {
		e[p[0]] = append(e[p[0]], &pair{p[1], p[2]})
		e[p[1]] = append(e[p[1]], &pair{p[0], p[2]})
	}
	f[start][0], vis[start][0] = 0, true
	q := []*pt{{start, 0}}
	for len(q) > 0 {
		nico := q[0]
		q = q[1:]
		u, crg := nico.u, nico.crg
		vis[u][crg] = false
		for i := crg + 1; i <= cnt; i++ {
			if f[u][i] > f[u][crg]+charge[u]*(i-crg) {
				f[u][i] = f[u][crg] + charge[u]*(i-crg)
				if !vis[u][i] {
					vis[u][i] = true
					q = append(q, &pt{u, i})
				}
			}
		}

		for _, x := range e[u] {
			y, rm := x.first, crg-x.second
			if rm >= 0 && f[y][rm] > f[u][crg]+x.second {
				f[y][rm] = f[u][crg] + x.second
				if !vis[y][rm] {
					vis[y][rm] = true
					q = append(q, &pt{y, rm})
				}
			}
		}
	}

	for i := 0; i <= cnt; i++ {
		ans = min(ans, f[end][i])
	}
	return ans
}

// 4
func maxGroupNumber(tiles []int) int {
	type pair struct {
		tile, freq int
	}
	last, cntMap, cnt := 0, make(map[int]*pair), make([]*pair, 0)
	for _, t := range tiles {
		p, ok := cntMap[t]
		if !ok {
			p = &pair{
				tile: t,
			}
			cntMap[t] = p
			cnt = append(cnt, p)
		}
		p.freq++
	}
	var dp [5][5]int
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i].tile < cnt[j].tile
	})
	for _, v := range cnt {
		tile, freq := v.tile, v.freq
		if last != tile-1 {
			hi := dp[0][0]
			dp = [5][5]int{}
			for i := range dp {
				for j := range dp[i] {
					dp[i][j] = -1
				}
			}
			dp[0][0] = hi
		}

		ndp := [5][5]int{}
		for i := range ndp {
			for j := range ndp[i] {
				ndp[i][j] = -1
			}
		}

		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if dp[i][j] < 0 {
					continue
				}

				for k := 0; k <= min(i, min(j, freq)); k++ {
					ni := j - k
					for nj := 0; nj <= min(4, freq-k); nj++ {
						ndp[ni][nj] = max(ndp[ni][nj], dp[i][j]+k+(freq-k-nj)/3)
					}
				}

			}
		}

		dp = ndp
		last = tile
	}

	ans := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			ans = max(ans, dp[i][j])
		}
	}

	return ans
}

// 5 todo 未通过 https://leetcode.cn/contest/season/2021-spring/
func minRecSize(lines [][]int) float64 {
	n := len(lines)
	k, b := make([]int, n), make([]int, n)
	sort.Slice(lines, func(i, j int) bool {
		return lines[i][0] < lines[j][0] ||
			lines[i][0] == lines[j][0] && lines[i][1] < lines[j][1]
	})

	for i := 0; i < n; i++ {
		k[i], b[i] = lines[i][0], lines[i][1]
	}

	p, q := 0, 0
	for q < n && k[q] == k[p] {
		q++
	}
	if q >= n {
		return 0
	}

	min := func(a, b float64) float64 {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b float64) float64 {
		if a > b {
			return a
		}
		return b
	}

	x_min, x_max, y_min, y_max := 1e100, -1e100, 1e100, -1e100
	for q < n {
		r := q
		for r+1 < n && k[r+1] == k[q] {
			r++
		}
		cx1 := 1.0 * float64(b[r]-b[p]) / float64(k[r]-k[p])
		cx2 := 1.0 * float64(b[q]-b[q-1]) / float64(k[q]-k[q-1])
		x_min = min(x_min, min(cx1, cx2))
		x_max = max(x_max, max(cx1, cx2))

		cy1 := 1.0 * float64(b[r]*k[p]-b[p]*k[r]) / float64(k[r]-k[p])
		cy2 := 1.0 * float64(b[q]*k[q-1]-b[q-1]*k[q]) / float64(k[q]-k[q-1])
		y_min = min(y_min, min(cy1, cy2))
		y_max = max(y_max, max(cy1, cy2))
		p = q
		for q < n && k[q] == k[p] {
			q++
		}
	}
	return (x_max - x_min) * (y_max - y_min)
}

// 6
func guardCastle(grid []string) int {
	l := len(grid[0]) + 2
	g0, g1, g0t, g1t := make([]byte, l), make([]byte, l), make([]byte, l), make([]byte, l)
	copy(g0[1:], grid[0])
	copy(g1[1:], grid[1])
	g0[0], g0[l-1], g1[0], g1[l-1] = '.', '.', '.', '.'
	copy(g0t, g0)
	copy(g1t, g1)

	cal1 := func(i, p int) int {
		for j := p + 1; j <= i; j++ {
			if (g0[j] == '#' && (g1[j] == '#' || g1[j-1] == '#')) ||
				(g1[j] == '#' && (g0[j] == '#' || g0[j-1] == '#')) {
				return 0
			}
		}
		for j := i; j >= p; j-- {
			if g0[j] == '#' {
				if j < i && g1[j+1] == '.' {
					g1[j+1] = '#'
				}
				return 1
			}
			if g1[j] == '#' {
				if j < i && g0[j+1] == '.' {
					g0[j+1] = '#'
				}
				return 1
			}
		}
		if g0[i] == '.' {
			g0[i] = '#'
		}
		if g1[i] == '.' {
			g1[i] = '#'
		}
		return 2
	}

	cal := func() int {
		for i := 1; i < l; i++ {
			if g0[i]+g1[i] == 3 || g0[i-1]+g0[i] == 3 || g1[i-1]+g1[i] == 3 {
				return 1e5
			}
		}
		o, p, t := 0, 0, -1
		for i := 1; i < l; i++ {
			if g0[i] == 1 || g1[i] == 1 {
				if t == 2 {
					o += cal1(i, p)
				}
				p, t = i, 1
			} else if g0[i] == 2 || g1[i] == 2 {
				if t == 1 {
					o += cal1(i, p)
				}
				p, t = i, 2
			}
		}
		return o
	}

	for i := 1; i < l-1; i++ {
		switch g0[i] {
		case 'S', 'P':
			g0[i] = 1
		case 'C':
			g0[i] = 2
		}
		switch g1[i] {
		case 'S', 'P':
			g1[i] = 1
		case 'C':
			g1[i] = 2
		}
	}
	o := cal()

	g0, g1 = g0t, g1t
	for i := 1; i < l-1; i++ {
		switch g0[i] {
		case 'S':
			g0[i] = 1
		case 'C', 'P':
			g0[i] = 2
		}
		switch g1[i] {
		case 'S':
			g1[i] = 1
		case 'C', 'P':
			g1[i] = 2
		}
	}
	o = min(o, cal())

	if o == 1e5 {
		return -1
	}
	return o
}
