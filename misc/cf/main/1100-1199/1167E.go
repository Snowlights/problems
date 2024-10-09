//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1167E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x, v int
	Fscan(in, &n, &x)
	ps := [1e6 + 1]struct{ l, r, v int }{}
	for i := range ps {
		ps[i].l = 1e9
		ps[i].v = i
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		if ps[v].l == 1e9 {
			ps[v].l = i
		}
		ps[v].r = i
	}

	b := ps[:0]
	for _, p := range ps {
		if p.r > 0 {
			b = append(b, p)
		}
	}
	n = len(b)

	i := 0
	for i < n-1 && b[i].r < b[i+1].l {
		i++
	}
	if i == n-1 {
		Fprint(out, x*(x+1)/2)
		return
	}

	ans := b[i+1].v * (x + 1 - b[n-1].v) // 去掉后缀 b[<=i+1] ~ b[n-1]
	for j := n - 1; j == n-1 || b[j].r < b[j+1].l; j-- {
		for i >= 0 && b[i].r >= b[j].l {
			i--
		}
		ans += b[i+1].v * (b[j].v - b[j-1].v) // 去掉 b[<=i+1] ~ b[j-1]
	}
	Fprint(out, ans)

}

func main() { CF1167E(os.Stdin, os.Stdout) }
