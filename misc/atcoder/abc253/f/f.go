package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type node struct {
	lo, ro    *node
	l, r, sum int
}

func buildPST(l, r int) *node {
	o := &node{l: l, r: r}
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = buildPST(l, m)
	o.ro = buildPST(m+1, r)
	return o
}

func (o *node) update(l, r, add int) {
	if l <= o.l && o.r <= r {
		o.sum += add
		return
	}
	lo := *(o.lo)
	o.lo = &lo
	ro := *(o.ro)
	o.ro = &ro
	if add := o.sum; add > 0 {
		o.lo.sum += add
		o.ro.sum += add
		o.sum = 0
	}
	if l <= o.lo.r {
		o.lo.update(l, r, add)
	}
	if r >= o.ro.l {
		o.ro.update(l, r, add)
	}
}

func (o *node) query(i int) int {
	if o.l == o.r {
		return o.sum
	}
	if i <= o.lo.r {
		return o.sum + o.lo.query(i)
	}
	return o.sum + o.ro.query(i)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, q, op, l, r, v int
	Fscan(in, &n, &m, &q)
	t := []*node{buildPST(1, m)}
	type pair struct{ i, v int }
	last := make([]pair, n+1)
	for ; q > 0; q-- {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			Fscan(in, &v)
			_o := *t[len(t)-1]
			o := &_o
			o.update(l, r, v)
			t = append(t, o)
		} else if op == 2 {
			last[l] = pair{len(t) - 1, r}
		} else {
			res := t[len(t)-1].query(r)
			p := last[l]
			res2 := t[p.i].query(r)
			Fprintln(out, p.v+res-res2)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
