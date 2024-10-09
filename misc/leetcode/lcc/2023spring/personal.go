package _023spring

import (
	"sort"
	"strings"
)

// 1
func supplyWagon(a []int) (ans []int) {
	m := len(a) / 2
	for len(a) > m {
		mii := 1
		for i := 1; i < len(a); i++ {
			if a[i]+a[i-1] < a[mii]+a[mii-1] {
				mii = i
			}
		}
		a[mii-1] += a[mii]
		a = append(a[:mii], a[mii+1:]...)
	}
	return a
}

// 2
func adventureCamp(a []string) (ans int) {
	vis := map[string]bool{}
	for _, s := range strings.Split(a[0], "->") {
		vis[s] = true
	}
	mxC := 0
	ans = -1
	for i := 1; i < len(a); i++ {
		if a[i] == "" {
			continue
		}
		c := 0
		sp := strings.Split(a[i], "->")
		for _, s := range sp {
			if !vis[s] {
				c++
				vis[s] = true
			}
		}
		if c > mxC {
			mxC = c
			ans = i
		}
	}
	return
}

// 3
func fieldOfGreatestBlessing(a [][]int) (ans int) {
	hasx := map[int]bool{}
	hasy := map[int]bool{}
	for _, p := range a {
		x, y, sz := p[0], p[1], p[2]
		xx := 2*x - sz
		yy := 2*y - sz
		hasx[xx] = true
		hasy[yy] = true
		xx = 2*x + sz
		yy = 2*y + sz
		hasx[xx] = true
		hasy[yy] = true
	}

	xs := make([]int, 0, len(hasx))
	for k := range hasx {
		xs = append(xs, k)
	}
	sort.Ints(xs)

	ys := make([]int, 0, len(hasy))
	for k := range hasy {
		ys = append(ys, k)
	}
	sort.Ints(ys)

	n := len(xs)
	m := len(ys)
	diff := make([][]int, n+2)
	for i := range diff {
		diff[i] = make([]int, m+2)
	}
	update := func(r1, c1, r2, c2, x int) {
		diff[r1+1][c1+1] += x
		diff[r1+1][c2+2] -= x
		diff[r2+2][c1+1] -= x
		diff[r2+2][c2+2] += x
	}

	for _, p := range a {
		x, y, sz := p[0], p[1], p[2]
		xx := 2*x - sz
		yy := 2*y - sz
		r1 := sort.SearchInts(xs, xx)
		c1 := sort.SearchInts(ys, yy)
		xx = 2*x + sz
		yy = 2*y + sz
		r2 := sort.SearchInts(xs, xx)
		c2 := sort.SearchInts(ys, yy)
		update(r1, c1, r2, c2, 1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			diff[i][j] += diff[i][j-1] + diff[i-1][j] - diff[i-1][j-1]
		}
	}
	diff = diff[1 : n+1]
	for i, row := range diff {
		diff[i] = row[1 : m+1]
	}

	for i, r := range diff {
		for j, v := range r {
			_, _ = i, j
			ans = max(ans, v)
		}
	}

	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 4
type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func challengeOfTheKeeper(g []string) (ans int) {
	getST := func(labelS, labelT rune) (pair, pair) {
		var s, t pair
		for i, row := range g {
			for j, b := range row {
				p := pair{i, j}
				if b == labelS {
					s = p
				} else if b == labelT {
					t = p
				}
			}
		}
		return s, t
	}

	disAll := func(sx, sy int) [][]int {
		n, m := len(g), len(g[0])
		dis := make([][]int, n)
		for i := range dis {
			dis[i] = make([]int, m)
			for j := range dis[i] {
				dis[i][j] = 1e9
			}
		}
		dis[sx][sy] = 0
		q := []pair{{sx, sy}}
		for step := 1; len(q) > 0; step++ {
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, d := range dir4 {
					if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m && g[x][y] != '#' && dis[x][y] == 1e9 {
						dis[x][y] = step
						q = append(q, pair{x, y})
					}
				}
			}
		}
		return dis
	}

	s, t := getST('S', 'T')

	disFromT := disAll(t.x, t.y)

	findTTT := func(sx, sy, mx int) bool {
		// 先判断能不能到达 T
		if disFromT[sx][sy] == 1e9 {
			return false
		}

		n, m := len(g), len(g[0])
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}

		findT := false
		type pair struct{ x, y int }
		var bad func(int, int)
		bad = func(x, y int) {
			if x < 0 || x >= n || y < 0 || y >= m || vis[x][y] || g[x][y] == '#' {
				return
			}
			if g[x][y] == 'T' {
				findT = true
				return
			}
			vis[x][y] = true
			if g[x][y] == '.' {
				if xx, yy := x, m-1-y; g[xx][yy] != '#' && disFromT[xx][yy] > mx {
					return
				}
				if xx, yy := n-1-x, y; g[xx][yy] != '#' && disFromT[xx][yy] > mx {
					return
				}
			}
			for _, d := range dir4 {
				xx, yy := x+d.x, y+d.y
				bad(xx, yy)
			}
			return
		}
		bad(sx, sy)
		return findT
		//vis[sx][sy] = false
		//vis[t.x][t.y] = false
		//
		//for i, r := range vis {
		//	for j, b := range r {
		//		if !b {
		//			continue
		//		}
		//		if g[i][m-1-j] != '#' && disFromT[i][m-1-j] < 0 {
		//			return false
		//		}
		//		if g[n-1-i][j] != '#' && disFromT[n-1-i][j] < 0 {
		//			return false
		//		}
		//	}
		//}
		//ans = disFromT[sx][sy]
		//
		//for i, r := range vis {
		//	for j, b := range r {
		//		if !b {
		//			continue
		//		}
		//		if g[i][m-1-j] != '#'{
		//			ans = max(ans, disFromT[i][m-1-j])
		//		}
		//		if g[n-1-i][j] != '#' {
		//			ans = max(ans, disFromT[n-1-i][j])
		//		}
		//	}
		//}
		//return true
	}
	if !findTTT(s.x, s.y, 1e9-1) {
		return -1
	}
	n, m := len(g), len(g[0])
	ans = sort.Search(n*m+100, func(mx int) bool {
		ok := findTTT(s.x, s.y, mx)
		return ok
	})

	// 注意：初始位置不能执行传送

	return
}

// 5
func getSchemeCount(_, _ int, g []string) int64 {
	ans := 0
	n, m := len(g), len(g[0])
	a := make([][]byte, n)
	for i := range a {
		a[i] = []byte(g[i])
	}
	rotateCopy := func(a [][]byte) [][]byte {
		n, m := len(a), len(a[0])
		b := make([][]byte, m)
		for i := range b {
			b[i] = make([]byte, n)
		}
		for i, r := range a {
			for j, v := range r {
				b[j][n-1-i] = v
			}
		}
		return b
	}
	if n < m {
		a = rotateCopy(a)
	}
	// n >= m
	n, m = len(a), len(a[0])

	// check 除去 ? 必须满足 5 种状态
	// B=1
	// R=2

	// 0 空
	// 1 *121
	// 2 *212
	// 3 1
	// 4 2
	// 5 *11
	// 6 *22

	// B=0   R=1
	trans := [7][2]int{
		{3, 4},  // 0
		{-1, 2}, // 1
		{1, -1}, // 2
		{5, 2},  // 3
		{1, 6},  // 4
		{5, -1}, // 5
		{-1, 6}, // 6
	}

	calc := func(a []byte) int {
		n := len(a)
		dp := make([][7]int, n)
		for i := range dp {
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		var f func(int, int) int
		f = func(i, j int) (res int) {
			if j < 0 {
				return 0
			}
			if i == n {
				return 1
			}
			dv := &dp[i][j]
			if *dv != -1 {
				return *dv
			}
			defer func() { *dv = res }()
			res = 0
			if a[i] == 'B' {
				k := trans[j][0]
				res += f(i+1, k)
			} else if a[i] == 'R' {
				k := trans[j][1]
				res += f(i+1, k)
			} else if a[i] == '?' {
				res += f(i+1, j) // 不做
				k := trans[j][0]
				res += f(i+1, k)
				k = trans[j][1]
				res += f(i+1, k)
			} else {

			}
			return
		}
		res := f(0, 0)
		return res
	}
	col := make([]int, m)
	for j := range a[0] {
		b := []byte{}
		for i, r := range a {
			_ = i
			v := r[j]
			if v != '.' {
				b = append(b, v)
			}
		}
		col[j] = calc(b)
	}

	if m == 1 {
		return int64(col[0])
	}

	if m == 2 {
		return int64(col[0]) * int64(col[1])
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 1<<(m*3))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var DFS func(int, int) int
	DFS = func(r, mask int) (RES int) {
		if r == n {
			return 1
		}
		dv := &dp[r][mask]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = RES }()

		row := a[r]
		// 同行爆搜
		var f func(int, int, int)
		f = func(c, pre, colMask int) {
			if pre < 0 {
				return
			}
			if c == m {
				re := DFS(r+1, colMask)
				RES += re
				return
			}

			f1 := func(tp int) {
				k := trans[pre][tp]
				if k >= 0 {
					ms := colMask &^ (7 << (3 * c))
					kk := trans[colMask>>(3*c)&7][tp]
					if kk >= 0 {
						ms |= kk << (3 * c)
						f(c+1, k, ms)
					}
				}
			}
			if row[c] == 'B' {
				f1(0)
			} else if row[c] == 'R' {
				f1(1)
			} else if row[c] == '?' {
				f(c+1, pre, colMask) // 不做
				f1(0)
				f1(1)
			} else { // '.'
				f(c+1, pre, colMask)
			}
		}
		f(0, 0, mask)
		return
	}
	ans = DFS(0, 0)

	return int64(ans)
}
