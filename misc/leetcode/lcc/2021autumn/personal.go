package _021autumn

import (
	"math/bits"
	"sort"
)

// 1
func minimumSwitchingTimes(source [][]int, tar [][]int) (ans int) {
	c1 := map[int]int{}
	for i, r := range source {
		for j, v := range r {
			_, _ = i, j
			c1[v]++
		}
	}

	for i, r := range tar {
		for j, v := range r {
			_, _ = i, j
			c1[v]--
		}
	}

	for _, c := range c1 {
		ans += abs(c)
	}
	ans /= 2

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 2
func maxmiumScore(a []int, cnt int) (ans int) {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	sum := 0
	for _, v := range a[:cnt] {
		sum += v
	}

	if sum&1 == 0 {
		return sum
	}

	for _, v := range a[cnt:] {
		if v&1 != a[cnt-1]&1 {
			res := sum - a[cnt-1] + v
			ans = max(ans, res)
			break
		}
	}

	for i := cnt - 1; i >= 0; i-- {
		if a[i]&1 != a[cnt-1]&1 {
			for _, v := range a[cnt:] {
				if v&1 != a[i]&1 {
					res := sum - a[i] + v
					ans = max(ans, res)
					break
				}
			}
			break
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

// 3
func flipChess(g []string) (ans int) {
	type pair struct{ x, y int }
	dir8 := []pair{
		{1, 0}, {1, 1}, {0, 1}, {-1, 1},
		{-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

	n, m := len(g), len(g[0])
	for i0, r := range g {
		for j0, v := range r {
			if v == '.' {
				g2 := make([][]byte, n)
				for i, r := range g {
					g2[i] = []byte(r)
				}
				g2[i0][j0] = 'X'
				cnt := 0
				for {
					found := false
					for i, r := range g2 {
						for j, v := range r {
							if v == 'O' {
							o:
								for k := 0; k < 4; k++ {
									x, y := i, j
									d := dir8[k]
									for {
										x += d.x
										y += d.y
										if 0 <= x && x < n && 0 <= y && y < m && g2[x][y] != '.' {
											if g2[x][y] == 'X' {
												break
											}
										} else {
											continue o
										}
									}
									x, y = i, j
									d = dir8[k+4]
									for {
										x += d.x
										y += d.y
										if 0 <= x && x < n && 0 <= y && y < m && g2[x][y] != '.' {
											if g2[x][y] == 'X' {
												break
											}
										} else {
											continue o
										}
									}
									g2[i][j] = 'X'
									found = true
									cnt++
									break
								}
							}
						}
					}
					if !found {
						break
					}
				}
				ans = max(ans, cnt)
			}
		}
	}
	return
}

// 4
func circleGame(toys [][]int, a [][]int, r int) (ans int) {
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] })

	k := 0
	for _, w := range a[1:] {
		if a[k][0] != w[0] || a[k][1] != w[1] {
			k++
			a[k] = w
		}
	}
	a = a[:k+1]
	type pair struct{ x, y int }
	cs := [][]pair{}
	for _, p := range a {
		if len(cs) == 0 || p[0] != cs[len(cs)-1][0].x {
			cs = append(cs, []pair{{p[0], p[1]}})
		} else {
			cs[len(cs)-1] = append(cs[len(cs)-1], pair{p[0], p[1]})
		}
	}

o:
	for _, p := range toys {
		x, y, rp := p[0], p[1], p[2]
		if rp > r {
			continue
		}
		i := sort.Search(len(cs), func(ii int) bool { return cs[ii][0].x+r >= x+rp })
		if i == len(cs) {
			continue
		}
		for ; i < len(cs); i++ {
			cs := cs[i]
			if cs[0].x-r > x-rp {
				break
			}
			j := sort.Search(len(cs), func(ii int) bool { return cs[ii].y+r >= y+rp })
			for ; j < len(cs); j++ {
				q := cs[j]
				if q.y-r > y-rp {
					break
				}
				if (x-q.x)*(x-q.x)+(y-q.y)*(y-q.y) <= (r-rp)*(r-rp) {
					ans++
					continue o
				}
			}

		}
	}

	return
}

// 5
func trafficCommand(ds []string) (ans int) {
	n0, n1, n2, n3 := len(ds[0]), len(ds[1]), len(ds[2]), len(ds[3])
	dp := make([][][][]int, n0+1)
	for i := range dp {
		dp[i] = make([][][]int, n1+1)
		for j := range dp[i] {
			dp[i][j] = make([][]int, n2+1)
			for k := range dp[i][j] {
				dp[i][j][k] = make([]int, n3+1)
				for i2 := range dp[i][j][k] {
					dp[i][j][k][i2] = -1
				}
			}
		}
	}
	mp := [...]int{'e': 0, 'S': 1, 'W': 2, 'N': 3}

	// 去，另来，另去
	valid := [4][4][4]int{
		1: {
			1: {1},
			2: {0, 0, 0, 1},
			3: {0, 0, 1},
		},
		2: {
			1: {1},
			2: {1, 1},
			3: {},
		},
		3: {
			1: {1, 0, 1},
			2: {1, 1},
			3: {1, 1, 1},
		},
	}
	ok := func(v, w, tov, tow int) bool {
		w -= v
		if w < 0 {
			w += 4
		}
		tov -= v
		if tov < 0 {
			tov += 4
		}
		tow -= v
		if tow < 0 {
			tow += 4
		}
		return valid[tov][w][tow] == 1
	}
	var f func(int, int, int, int) int
	f = func(p0, p1, p2, p3 int) (res int) {
		if p0 == n0 && p1 == n1 && p2 == n2 && p3 == n3 {
			return
		}
		dv := &dp[p0][p1][p2][p3]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		ps := [4]int{p0, p1, p2, p3}

		res = 1e9
		calc := func(sub uint) (res int) {
			//if ps == [4]int{} && sub == 15 {
			//	print()
			//}
			cs := [][2]int{}
			for sub := sub; sub > 0; sub &= sub - 1 {
				p := bits.TrailingZeros(sub)
				if ps[p] == len(ds[p]) {
					return 1e9
				}
				to := mp[ds[p][ps[p]]]
				for _, q := range cs {
					if !ok(p, q[0], to, q[1]) {
						return 1e9
					}
				}
				cs = append(cs, [2]int{p, to})
			}
			for _, p := range cs {
				ps[p[0]]++
			}
			res = f(ps[0], ps[1], ps[2], ps[3]) + 1
			for _, p := range cs {
				ps[p[0]]--
			}
			return
		}
		for sub := uint(1); sub < 1<<4; sub++ {
			r := calc(sub)
			res = min(res, r)
		}

		return
	}
	ans = f(0, 0, 0, 0)
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
