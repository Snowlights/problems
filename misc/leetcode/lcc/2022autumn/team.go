package _022autumn

// 1
func minNumBooths(demand []string) (ans int) {
	mx := [26]int{}
	for _, s := range demand {
		cnt := make([]int, 26)
		for _, b := range s {
			cnt[b-'a']++
		}
		for i, c := range cnt {
			mx[i] = max(mx[i], c)
		}
	}
	for _, v := range mx {
		ans += v
	}
	return
}

// 2
func expandBinaryTree(root *TreeNode) (ans *TreeNode) {
	var dfs func(*TreeNode)
	dfs = func(o *TreeNode) {
		if o == nil {
			return
		}

		if o.Left == nil && o.Right == nil {

			return
		}

		ll, rr := o.Left, o.Right
		if o.Left != nil {
			o.Left = &TreeNode{Val: -1}
			o.Left.Left = ll
			dfs(ll)
		}

		if o.Right != nil {
			o.Right = &TreeNode{Val: -1}
			o.Right.Right = rr
			dfs(rr)
		}

		return
	}
	dfs(root)
	ans = root
	return
}

// 3
func beautifulBouquet(f []int, cnt int) int {
	n, mod := len(f), int(1e9+7)
	m, r, ans := make(map[int]int), 0, 0
	for i := 0; i < n; i++ {
		for r < n && m[f[r]] < cnt {
			m[f[r]]++
			r++
		}
		ans += r - i
		ans %= mod
		m[f[i]]--
	}
	return ans
}

// 4
func Leetcode(a []string) (ans int) {
	type pair struct{ pos, lim, m int }
	rule := ['z' + 1]pair{
		'e': {0, 4, 7},
		'l': {3, 3, 3},
		'o': {5, 2, 3},
		'h': {7, 1, 1},
		't': {8, 1, 1},
		'c': {9, 1, 1},
		'd': {10, 1, 1},
	}
	const keys = "elohtcd"
	const full = 2012

	fc := func(s string) []int {
		cost := make([]int, 1<<11)
		for _i := range cost {
			cost[_i] = 1e9
		}
		var f func(s string, mask, cst int)
		f = func(s string, mask, cst int) {
			cost[mask] = min(cost[mask], cst)
			if s == "" {
				//fmt.Println(mask, cst)
				return
			}
			for i, c := range s {
				p := rule[c]
				if p.lim == 0 {
					continue
				}
				cnt := mask >> p.pos & p.m
				if cnt < p.lim {
					f(s[:i]+s[i+1:], mask+1<<p.pos, cst+i*(len(s)-1-i))
				}
			}
		}
		f(s, 0, 0)
		return cost
	}

	valid := func(m, add int) int {
		for _, c := range keys {
			p := rule[c]
			c1 := m >> p.pos & p.m
			c2 := add >> p.pos & p.m
			if c1+c2 > p.lim {
				return -1
			}
			m += c2 << p.pos
		}
		return m
	}

	n := len(a)
	cost := make([][]int, n)
	for i, s := range a {
		cost[i] = fc(s)
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 1<<11)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(i, mask int) (res int) {
		if i == n {
			if mask == full {
				return 0
			}
			return 1e9
		}
		dv := &dp[i][mask]
		if *dv != -1 {
			return *dv
		}
		res = 1e9
		for add, cst := range cost[i] {
			if cst < 1e9 {
				//fmt.Println(i, mask, add, cst)
				m2 := valid(mask, add)
				if m2 >= 0 {
					r := f(i+1, m2)
					res = min(res, r+cst)
				}
			}
		}

		*dv = res
		return
	}
	ans = f(0, 0)
	if ans == 1e9 {
		return -1
	}
	return
}

// 5
func sandyLandManagement(size int) [][]int {
	ans := [][]int{}
	var work func(size, x, y int)
	work = func(size, offX, offY int) {
		push := func(x, y int) {
			ans = append(ans, []int{x + offX, y + offY})
		}
		push(1, 1)
		for i := 2; i <= size; i++ {
			push(i, 1)
			if i != 2 {
				push(i, 2*i-1)
			}
		}
		if size <= 3 {
			if size >= 2 {
				push(2, 3)
			}
			return
		}
		push(4, 4)
		if size == 4 {
			return
		}
		work(size-4, offX+4, offY+4)
	}
	work(size, 0, 0)
	return ans
}

// 6
func reservoir(a []string) (ans int) {
	n, m := len(a), len(a[0])

	//for _, s := range a {
	//	s = strings.ReplaceAll(s, "l", "\\")
	//	s = strings.ReplaceAll(s, "r", "/")
	//	//s = strings.ReplaceAll(s, ".", " ")
	//	fmt.Println(s)
	//}
	//
	//for _, s := range a {
	//	s0 := ""
	//	s1 := ""
	//	for _, c := range s {
	//		if c == 'l' {
	//			s0 += "\\ "
	//			s1 += " \\"
	//		} else if c == 'r' {
	//			s0 += " /"
	//			s1 += "/ "
	//		} else {
	//			s0 += "  "
	//			s1 += "  "
	//		}
	//	}
	//	//s = strings.ReplaceAll(s, "l", "\\")
	//	//s = strings.ReplaceAll(s, "r", "/")
	//	//s = strings.ReplaceAll(s, ".", " ")
	//	fmt.Println(s0)
	//	fmt.Println(s1)
	//}

	// 0 上 1 下
	//vis := make([][][2]bool, n)
	//for i := range vis {
	//	vis[i] = make([][2]bool, m)
	//}

	wt := make([][][2]bool, n)
	for i := range wt {
		wt[i] = make([][2]bool, m)
	}

	// v = (r,c) + (r,c+1)
	// 只能走 r <= lim 的
	//
	dfs := func(r, c, lim int, checkDown bool) (touchLR, touchU, touchD bool, vis [][][2]bool) {
		vis = make([][][2]bool, n)
		for i := range vis {
			vis[i] = make([][2]bool, m)
		}
		//vis[r][c][0] = true
		//vis[r][c+1][0] = true

		//if r == 3 {
		//	println()
		//}

		var f func(int, int, int)
		f = func(x, y, p int) {
			if vis[x][y][p] {
				return
			}
			vis[x][y][p] = true
			cur := a[x][y]

			// 上
			if p == 0 {
				if x > lim {
					b := a[x-1][y]
					if b == '.' {
						f(x-1, y, 0)
						f(x-1, y, 1)
					} else {
						f(x-1, y, 1)
					}
				}
				if x == 0 {
					touchU = true
				}
			}

			// 下
			if checkDown {
				if p == 1 {
					if x+1 < n {
						b := a[x+1][y]
						if b == '.' {
							f(x+1, y, 0)
							f(x+1, y, 1)
						} else {
							f(x+1, y, 0)
						}
					} else {
						touchD = true
					}
				}
			}

			// 左
			if y > 0 {
				if cur == '.' || p == 0 == (cur == 'r') {
					b := a[x][y-1]
					if b == '.' {
						f(x, y-1, 0)
						f(x, y-1, 1)
					} else if b == 'l' {
						f(x, y-1, 0)
					} else {
						f(x, y-1, 1)
					}
				}
			} else {
				if cur == '.' || p == 0 == (cur == 'r') {
					touchLR = true
				}
			}

			// 右
			if y+1 < m {
				if cur == '.' || p == 0 == (cur == 'l') {
					b := a[x][y+1]
					if b == '.' {
						f(x, y+1, 0)
						f(x, y+1, 1)
					} else if b == 'r' {
						f(x, y+1, 0)
					} else {
						f(x, y+1, 1)
					}
				}
			} else {
				if cur == '.' || p == 0 == (cur == 'l') {
					touchLR = true
				}
			}

		}
		f(r, c, 0)
		f(r, c+1, 0)

		//for len(q) > 0 {
		//	tmp := q
		//	q = nil
		//	for _, p := range tmp {
		//		wt[p.x][p.y][p.p] = true
		//		dirs := dss[p.p]
		//		if p.p == 0 {
		//			x, y := p.x-1, p.y
		//			if 0 <= x && x < n && 0 <= y && y < m && !wt[x][y][1] {
		//				q = append(q, pair{x, y})
		//			}
		//		}
		//		for _, d := range dirs {
		//			if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m && !wt[x][y][p.p^1] {
		//				dis[x][y] = step
		//				q = append(q, pair{x, y})
		//			}
		//		}
		//	}
		//}
		return
	}

	for i, r := range a {
		for j, v := range r {
			if v == 'l' && j+1 < m && r[j+1] == 'r' {
				//if !wt[i][j][0] {
				var tmp [][][2]bool

				lim := i
				for ; lim >= 0; lim-- {
					touchLR, touchU, touchD, vv := dfs(i, j, lim, true)
					if touchLR {
						break
					}
					tmp = vv
					if touchD {
						_, _, _, tmp = dfs(i, j, lim+1, false)
						break
					}
					if touchU {
						break
					}
				}
				if lim < 0 {
					//fmt.Println(i, j, "skip")
					continue
				}
				cnt := 0
				for i, r := range tmp {
					for j, v := range r {
						if v[0] {
							wt[i][j][0] = true
							cnt++
						}
						if v[1] {
							wt[i][j][1] = true
							cnt++
						}
					}
				}
				//fmt.Println(i, j, cnt)
				//}
			}
		}
	}

	for _, r := range wt {
		for _, v := range r {
			if v[0] {
				ans++
			}
			if v[1] {
				ans++
			}
		}
	}

	return
}
