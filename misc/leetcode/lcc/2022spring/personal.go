package _022spring

import (
	"math/bits"
	. "problems/testutil/leetcode"
	"sort"
)

// 1
func giveGem(a []int, operations [][]int) (ans int) {
	for _, p := range operations {
		v, w := p[0], p[1]
		c := a[v] / 2
		a[w] += c
		a[v] -= c
	}
	mx := int(-1e9)
	for _, v := range a {
		if v > mx {
			mx = v
		}
	}
	mi := int(1e9)
	for _, v := range a {
		if v < mi {
			mi = v
		}
	}
	ans = mx - mi
	return
}

// 2
func perfectMenu(materials []int, cookbooks [][]int, attribute [][]int, lim int) (ans int) {
	calc := func(sub int) (res int) {
		var cx, cy int
		cm := make([]int, len(materials))
		for _s := uint(sub); _s > 0; _s &= _s - 1 {
			p1 := bits.TrailingZeros(_s)
			a := cookbooks[p1]
			for i, v := range a {
				cm[i] += v
				if cm[i] > materials[i] {
					return -1
				}
			}
			cx += attribute[p1][0]
			cy += attribute[p1][1]
		}
		if cy < lim {
			return -1
		}
		res = cx
		return
	}
	ans = -1
	for sub := 0; sub < 1<<len(cookbooks); sub++ {
		res := calc(sub)
		ans = max(ans, res)
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
type seg []struct {
	l, r    int
	v, todo int
}

func (t seg) maintain(o int) {

}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = -1
	if l == r {
		t[o].v = 0
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) do(O int, v int) {
	o := &t[O]
	o.v = v
	o.todo = v
}

func (t seg) spread(o int) {
	if v := t[o].todo; v != -1 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = -1
	}
}

func (t seg) update(o, l, r, v int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

func (t seg) spreadAll(o int) {
	if t[o].l == t[o].r {
		return
	}
	t.spread(o)
	t.spreadAll(o << 1)
	t.spreadAll(o<<1 | 1)
}

func getNumber(root *TreeNode, ops [][]int) (ans int) {
	a := []int{}
	var dfs func(*TreeNode)
	dfs = func(o *TreeNode) {
		if o == nil {
			return
		}
		v := o.Val
		dfs(o.Left)
		a = append(a, v)
		dfs(o.Right)
	}
	dfs(root)

	has := map[int]bool{}
	for _, v := range a {
		has[v] = true
	}
	for _, op := range ops {
		v, w := op[1], op[2]
		has[v] = true
		has[w] = true
	}

	b := make([]int, 0, len(has))
	for k := range has {
		b = append(b, k)
	}
	sort.Ints(b)
	n := len(b)

	t := make(seg, n*4)
	t.build(1, 1, n)

	for _, op := range ops {
		v, l, r := op[0], op[1], op[2]
		l = sort.SearchInts(a, l) + 1
		r = sort.SearchInts(a, r) + 1
		//fmt.Println(l,r,v)
		t.update(1, l, r, v)
	}
	t.spreadAll(1)
	for _, p := range t {
		if p.l > 0 && p.l == p.r {
			if p.v == 1 {
				ans++
			}

		}
	}
	return
}

// 4
func defendSpaceCity(time []int, position []int) int {
	n := len(time)
	f, g := Init()
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		pos := position[i] - 1
		mp[pos] = mp[pos] | (1 << (time[i] - 1))
	}
	var h [101][32]int
	for i := 0; i < 32; i++ {
		h[0][i] = f[i]
	}
	for i := 1; i <= 100; i++ {
		for j := 0; j < 32; j++ {
			h[i][j] = 1000
			val := mp[i-1]
			for mask := 0; mask < 32; mask++ {
				if mask&val != val {
					continue
				}
				for sub := mask; sub >= 0; sub = mask & (sub - 1) {
					h[i][j] = min(h[i][j], h[i-1][sub]+g[mask-sub][j])
					if sub == 0 {
						break
					}
				}
			}
		}
	}
	return h[100][0]
}

func Init() ([32]int, [32][32]int) {
	var f [32]int
	for i := 0; i < 32; i++ {
		f[i] = bits.OnesCount(uint(i)) * 2
	}
	for i := 1; i <= 5; i++ {
		mask := (1 << i) - 1
		for j := 0; (mask << j) < 32; j++ {
			f[mask<<j] = i + 1
		}
	}
	for mask := 1; mask < 32; mask++ {
		for i := 0; i < 5; i++ {
			sub := (1 << i) & mask
			f[mask] = min(f[mask], f[sub]+f[mask-sub])
		}
	}

	var g [32][32]int
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			g[i][j] = f[i] + f[j]
		}
	}

	for i := 1; i <= 5; i++ {
		mask := (1 << i) - 1
		for j := 0; (mask << j) < 32; j++ {
			g[mask<<j][mask<<j] = i + 2
		}
	}

	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			for k := 0; k < 5; k++ {
				subI := (1 << k) & i
				subJ := (1 << k) & j
				g[i][j] = min(g[i][j], g[subI][subJ]+g[i-subI][j-subJ])
			}
		}
	}

	return f, g
}

// 5
func minimumCost(cost []int, roads [][]int) int64 {
	ans := 0

	n := len(cost)
	g := make([][]int, n)
	for _, e := range roads {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	bccIDs := make([]int, n)
	comps := [][]int{}
	idCnt := 0
	isCut := make([]bool, n)
	dfn := make([]int, n)
	dfsClock := 0
	type edge struct{ v, w int }
	stack := []edge{}
	var f func(v, fa int) int
	f = func(v, fa int) int {
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		childCnt := 0
		for _, w := range g[v] {
			e := edge{v, w}
			if dfn[w] == 0 {
				stack = append(stack, e)
				childCnt++
				lowW := f(w, v)
				if lowW >= dfn[v] {
					isCut[v] = true
					idCnt++
					comp := []int{}
					for {
						e, stack = stack[len(stack)-1], stack[:len(stack)-1]
						if bccIDs[e.v] != idCnt {
							bccIDs[e.v] = idCnt
							comp = append(comp, e.v)
						}
						if bccIDs[e.w] != idCnt {
							bccIDs[e.w] = idCnt
							comp = append(comp, e.w)
						}
						if e.v == v && e.w == w {
							break
						}
					}
					comps = append(comps, comp)
				}
				lowV = min(lowV, lowW)
			} else if w != fa && dfn[w] < dfn[v] {
				stack = append(stack, e)
				lowV = min(lowV, dfn[w])
			}
		}
		if fa == -1 && childCnt == 1 {
			isCut[v] = false
		}
		return lowV
	}
	for v, timestamp := range dfn {
		if timestamp == 0 {
			if len(g[v]) == 0 {
				idCnt++
				bccIDs[v] = idCnt
				comps = append(comps, []int{v})
				continue
			}
			f(v, -1)
		}
	}

	vid := idCnt
	cutIDs := make([]int, n)
	for i, is := range isCut {
		if is {
			vid++
			cutIDs[i] = vid
		}
	}

	g2 := make([][]int, vid+1)
	sz := make([]int, vid+1)
	mi := make([]int, vid+1)
	for i := range mi {
		mi[i] = 1e18
	}
	for i, is := range isCut {
		if is {
			sz[cutIDs[i]] = cost[i]
			mi[cutIDs[i]] = cost[i]
		}
	}
	for v, cp := range comps {
		v++
		for _, w := range cp {
			if !isCut[w] {
				sz[v] += cost[w]
				mi[v] = min(mi[v], cost[w])
			}
			if w := cutIDs[w]; w > 0 {
				g2[v] = append(g2[v], w)
				g2[w] = append(g2[w], v)
			}
		}
	}

	for i, v := range mi {
		if v == 1e18 {
			mi[i] = 0
		}
	}

	//g2[1] = append(g2[1], -1)

	if vid == 1 {
		ans = mi[1]
		return int64(ans)
	}

	s := 0
	mx := 0
	for v, vs := range g2 {
		if len(vs) == 1 {
			s += mi[v]
			mx = max(mx, mi[v])
		}
	}
	s -= mx
	ans = s

	//var f2 func(v, fa int) (int, int)
	//f2 = func(v, fa int) (int, int) {
	//	s := sz[v]
	//	if len(g2[v]) == 1 {
	//		return s - mi[v], s
	//	}
	//	d0mx := 0
	//	d1mx := 0
	//	for _, w := range g2[v] {
	//		if w != fa {
	//			d0, d1 := f2(w, v)
	//			if d0mx > 0 {
	//				ans = max(ans, d0mx+d1+s)
	//			}
	//			ans = max(ans, d1mx+d0+s)
	//			d1mx = max(d1mx, d1)
	//			d0mx = max(d0mx, d0)
	//		}
	//	}
	//	ans = max(ans, d1mx)
	//	return d0mx + s, d1mx + s
	//}
	//f2(1, -1)
	//sum := 0
	//for _, v := range cost {
	//	sum += v
	//}
	//ans = sum - ans
	return int64(ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
