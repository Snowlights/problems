//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

type seg []struct{ l, r, g int }

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].g = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t[o].g = gcd(t[o<<1].g, t[o<<1|1].g)
}

func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].g
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return gcd(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CF2050F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, q, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			if i > 0 {
				a[i-1] = abs(a[i-1] - a[i])
			}
		}
		t := make(seg, 2<<bits.Len(uint(n-1)))
		t.build(a, 1, 0, n-1)
		for range q {
			Fscan(in, &l, &r)
			if l == r {
				Fprint(out, "0 ")
			} else {
				Fprint(out, t.query(1, l-1, r-2), " ")
			}
		}
		Fprintln(out)
	}

}

func main() { CF2050F(os.Stdin, os.Stdout) }
