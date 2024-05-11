//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

var swap []int

type seg []int

func (t seg) maintain(o int) {
	t[o] = t[o<<1] + t[o<<1|1]
}

func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o] = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, ol, or, d, i, val int) {
	if ol == or {
		t[o] = val
		return
	}
	m := (ol + or) >> 1
	sw := swap[d]
	if i <= m {
		t.update(o<<1^sw, ol, m, d-1, i, val)
	} else {
		t.update(o<<1^1^sw, m+1, or, d-1, i, val)
	}
	t.maintain(o)
}

func (t seg) query(o, ol, or, d, l, r int) int {
	if l <= ol && or <= r {
		return t[o]
	}
	m := (ol + or) >> 1
	sw := swap[d]
	if r <= m {
		return t.query(o<<1^sw, ol, m, d-1, l, r)
	}
	if m < l {
		return t.query(o<<1^1^sw, m+1, or, d-1, l, r)
	}
	return t.query(o<<1^sw, ol, m, d-1, l, r) + t.query(o<<1^1^sw, m+1, or, d-1, l, r)
}

func CF1401F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var w, q, op, l, r int
	Fscan(in, &w, &q)
	swap = make([]int, w+1)
	n := 1 << w
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg, n*2)
	t.build(a, 1, 1, n)
	for ; q > 0; q-- {
		Fscan(in, &op, &l)
		if op == 1 {
			Fscan(in, &r)
			t.update(1, 1, n, w, l, r)
		} else if op == 2 {
			for i := 0; i <= l; i++ {
				swap[i] ^= 1
			}
		} else if op == 3 {
			swap[l+1] ^= 1
		} else {
			Fscan(in, &r)
			Fprintln(out, t.query(1, 1, n, w, l, r))
		}
	}
}

func main() { CF1401F(os.Stdin, os.Stdout) }
