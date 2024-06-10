//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type info struct{ max, i int }

func mergeInfo(a, b info) info {
	if a.max >= b.max {
		return a
	}
	return b
}

type seg []struct {
	l, r int
	val  info
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = info{a[l], l}
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) set0(o, i int) {
	if t[o].l == t[o].r {
		t[o].val.max = 0
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.set0(o<<1, i)
	} else {
		t.set0(o<<1|1, i)
	}
	t.maintain(o)
}

func (t seg) maintain(o int) {
	t[o].val = mergeInfo(t[o<<1].val, t[o<<1|1].val)
}

func (t seg) query(o, l, r int) info {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func CF1416D(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12)
	rc := func() byte {
		if _i == _n {
			_n, _ = _r.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n, m, q := r(), r(), r()
	a := make([]int, n)
	for i := range a {
		a[i] = r()
	}
	es := make([]struct{ v, w int }, m)
	for i := range es {
		es[i].v = r() - 1
		es[i].w = r() - 1
	}
	qs := make([]struct{ tp, v int }, q)
	del := make([]bool, m)
	for i := range qs {
		qs[i].tp = r()
		qs[i].v = r() - 1
		if qs[i].tp == 2 {
			del[qs[i].v] = true
		}
	}

	g := make([][]int, n*2)
	fa := make([]int, n*2)
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
	merge := func(v, w int) {
		v = find(v)
		w = find(w)
		if v == w {
			return
		}
		fa[v] = n
		fa[w] = n
		g[n] = append(g[n], v, w)
		n++
	}
	for i, d := range del {
		if !d {
			merge(es[i].v, es[i].w)
		}
	}

	for i := q - 1; i >= 0; i-- {
		p := &qs[i]
		if p.tp == 1 {
			p.v = find(p.v)
		} else {
			e := es[p.v]
			merge(e.v, e.w)
		}
	}

	timeIn := make([]int, n)
	timeOut := make([]int, n)
	at := make([]int, n)
	clock := -1
	var dfs func(int)
	dfs = func(v int) {
		clock++
		if v < len(a) {
			at[clock] = a[v]
		}
		timeIn[v] = clock
		for _, w := range g[v] {
			dfs(w)
		}
		timeOut[v] = clock
	}
	for i := range timeIn {
		if find(i) == i { // 根
			dfs(i)
		}
	}

	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(at, 1, 0, n-1)
	for _, p := range qs {
		if p.tp == 2 {
			continue
		}
		res := t.query(1, timeIn[p.v], timeOut[p.v])
		Fprintln(out, res.max)
		if res.max > 0 {
			t.set0(1, res.i)
		}
	}
}

func main() { CF1416D(os.Stdin, os.Stdout) }
