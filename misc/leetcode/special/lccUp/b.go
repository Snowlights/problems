package lccUp

// 1
func minNumBooths(demand []string) int {

	ans, cnt := 0, [26]int{}

	for _, d := range demand {
		h := make(map[int]int)
		for j := range d {
			h[int(d[j]-'a')]++
		}
		for k, v := range h {
			cnt[k] = max(cnt[k], v)
		}
	}

	for _, v := range cnt {
		ans += v
	}
	return ans
}

// 2
func expandBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v.Left != nil {
				q = append(q, v.Left)
				tmp := v.Left
				v.Left = &TreeNode{
					Val:   -1,
					Left:  tmp,
					Right: nil,
				}
			}
			if v.Right != nil {
				q = append(q, v.Right)
				tmp := v.Right
				v.Right = &TreeNode{
					Val:   -1,
					Left:  nil,
					Right: tmp,
				}
			}
		}
	}

	return root
}

// 3
func beautifulBouquet(flowers []int, cnt int) int {

	r, ans, mod := 0, 0, int(1e9+7)
	h := make(map[int]int)
	for i := range flowers {
		for r < len(flowers) && h[flowers[r]] < cnt {
			h[flowers[r]]++
			r++
		}
		ans += r - i
		ans %= mod
		h[flowers[i]]--
	}

	return ans
}

// 4
func Leetcode(a []string) int {
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
	ans := f(0, 0)
	if ans == 1e9 {
		return -1
	}
	return ans
}

// 5
func sandyLandManagement(size int) [][]int {
	ans := [][]int{{1, 1}}

	for i := 2; i <= size; i++ {
		if (size-i)%4 == 0 {
			for j := 1; j <= 2*i-1; j += 2 {
				ans = append(ans, []int{i, j})
			}
		} else if (size-i)%4 == 3 {
			ans = append(ans, []int{i, 1})
		} else if (size-i)%4 == 2 {
			for j := 3; j <= i*2-1; j += 2 {
				ans = append(ans, []int{i, j})
			}
		} else {
			ans = append(ans, []int{i, 2})
		}
	}

	return ans
}

// 6
func reservoir(a []string) (ans int) {
	n, m := len(a), len(a[0])

	wt := make([][][2]bool, n)
	for i := range wt {
		wt[i] = make([][2]bool, m)
	}

	dfs := func(r, c, lim int, checkDown bool) (touchLR, touchU, touchD bool, vis [][][2]bool) {
		vis = make([][][2]bool, n)
		for i := range vis {
			vis[i] = make([][2]bool, m)
		}

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
		return
	}

	for i, r := range a {
		for j, v := range r {
			if v == 'l' && j+1 < m && r[j+1] == 'r' {
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
