package didi

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func maxFloat(a, b float64) float64 {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

// 判断线段是否相交
type Line struct {
	id             int
	x1, y1, x2, y2 float64
}

func helper(a, b, c []float64) float64 {
	w1, v1, w2, v2 := b[0]-a[0], b[1]-a[1], c[0]-a[0], c[1]-a[1]
	return w1*v2 - w2*v1
}

func check(l1, l2 *Line) bool {
	p1, p2 := []float64{l1.x1, l1.y1}, []float64{l1.x2, l1.y2}
	p3, p4 := []float64{l2.x1, l2.y1}, []float64{l2.x2, l2.y2}
	if maxFloat(p1[0], p2[0]) >= minFloat(p3[0], p4[0]) &&
		maxFloat(p3[0], p4[0]) >= minFloat(p1[0], p2[0]) &&
		maxFloat(p1[1], p2[1]) >= minFloat(p3[1], p4[1]) &&
		maxFloat(p3[1], p4[1]) >= minFloat(p2[1], p1[1]) {
		if helper(p1, p2, p3)*helper(p1, p2, p4) <= 0 &&
			helper(p3, p4, p1)*helper(p3, p4, p2) <= 0 {
			return true
		}
		return false
	} else {
		return false
	}
}

// 并查集板子
type uf struct {
	fa, sz  []int
	groups  int
	maxSize int
}

func newUnionFind(n int, a []int) uf {
	fa := make([]int, n)
	sz := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = a[i]
	}
	return uf{fa, sz, n, 1}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *uf) merge(from, to int) (isNewMerge int) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return 0
	}
	u.fa[x] = y
	u.sz[y] += u.sz[x]
	if u.sz[y] > u.maxSize {
		u.maxSize = u.sz[y]
	}
	u.groups--
	return u.sz[y]
}

func (u uf) same(x, y int) bool { return u.find(x) == u.find(y) }
func (u uf) size(x int) int     { return u.sz[u.find(x)] }
