package algorithm

/* 并查集
只有路径压缩的并查集复杂度是 O(nlogn) 的，这也是大多数情况下的实现方案
只有启发式合并（按深度合并）的并查集的复杂度也是 O(nlogn) 的，适用于可持久化的场景
*/

// 普通并查集
// 可视化 https://visualgo.net/zh/ufds
// https://oi-wiki.org/ds/dsu/
// https://cp-algorithms.com/data_structures/disjoint_set_union.html
// 并查集时间复杂度证明 https://oi-wiki.org/ds/dsu-complexity/
//
// 模板题 https://www.luogu.com.cn/problem/P3367

// 快速复制
func _(n int) {
	fa := make([]int, n) // n+1
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) { fa[find(from)] = find(to) }
	same := func(x, y int) bool { return find(x) == find(y) }

	// 总是合并到代表元更大的树上
	mergeBig := func(from, to int) int {
		ff, ft := find(from), find(to)
		if ff > ft {
			ff, ft = ft, ff
		}
		fa[ff] = ft
		return ft
	}

	_ = []interface{}{merge, same, mergeBig}
}

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

// 二维并查集
type ufPoint struct{ x, y int } // int32
type uf2d map[ufPoint]ufPoint

func (u uf2d) find(x ufPoint) ufPoint {
	if f, ok := u[x]; ok && f != x {
		u[x] = u.find(f)
		return u[x]
	}
	return x
}
func (u uf2d) merge(from, to ufPoint) { u[u.find(from)] = u.find(to) }

// Kick Start 2019C - Wiggle Walk https://codingcompetitions.withgoogle.com/kickstart/round/0000000000050ff2/0000000000150aac
func moveRobot(start ufPoint, command string) ufPoint {
	p := start
	w, n, e, s := uf2d{}, uf2d{}, uf2d{}, uf2d{}
	for _, c := range command {
		// 注意这里是矩阵
		w.merge(p, ufPoint{p.x, p.y - 1})
		n.merge(p, ufPoint{p.x - 1, p.y})
		e.merge(p, ufPoint{p.x, p.y + 1})
		s.merge(p, ufPoint{p.x + 1, p.y})
		switch c {
		case 'W':
			p = w.find(p)
		case 'N':
			p = n.find(p)
		case 'e':
			p = e.find(p)
		default:
			p = s.find(p)
		}
	}
	return p
}
