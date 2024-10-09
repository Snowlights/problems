package template

import (
	"bufio"
	"fmt"
	"io"
)

func LuoGuP3367(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, q, op, from, to int
	fmt.Fscan(r, &n, &q)
	uf := newUnionFind(n, make([]int, n+1))
	for i := 0; i < q; i++ {
		fmt.Fscan(r, &op, &from, &to)
		switch op {
		case 1:
			uf.merge(from, to)
		case 2:
			fmt.Fprintln(w, uf.find(from) == uf.find(to))
		}
	}
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
