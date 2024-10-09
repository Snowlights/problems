//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type seg [][]int

func (seg) merge(a, b []int) []int {
	const k = 31
	i, n := 0, len(a)
	j, m := 0, len(b)
	res := make([]int, 0, min(n+m, k))
	for len(res) < cap(res) {
		if i == n {
			res = append(res, b[j:min(j+k-len(res), m)]...)
			break
		}
		if j == m {
			res = append(res, a[i:min(i+k-len(res), n)]...)
			break
		}
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	return res
}

func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o] = a[l-1 : l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t[o] = t.merge(t[o<<1], t[o<<1|1])
}

func (t seg) query(o, l, r, L, R int) []int {
	if L <= l && r <= R {
		return t[o]
	}
	m := (l + r) >> 1
	if R <= m {
		return t.query(o<<1, l, m, L, R)
	}
	if m < L {
		return t.query(o<<1|1, m+1, r, L, R)
	}
	return t.merge(t.query(o<<1, l, m, L, R), t.query(o<<1|1, m+1, r, L, R))
}

func CF1665E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, q, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		t := make(seg, n*4)
		t.build(a, 1, 1, n)
		for Fscan(in, &q); q > 0; q-- {
			Fscan(in, &l, &r)
			b := t.query(1, 1, n, l, r)
			ans := 1 << 30
			for i, v := range b {
				for _, w := range b[:i] {
					ans = min(ans, v|w)
				}
			}
			Fprintln(out, ans)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF1665E(os.Stdin, os.Stdout) }
