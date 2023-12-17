//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF730J(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, sa, sb, m, maxSave int
	Fscan(in, &n)
	a := make([]struct{ a, b int }, n)
	for i := range a {
		Fscan(in, &a[i].a)
		sa += a[i].a
	}
	for i := range a {
		Fscan(in, &a[i].b)
	}

	sort.Slice(a, func(i, j int) bool { return a[i].b > a[j].b })
	for i, p := range a {
		sb += p.b
		if sb >= sa {
			m = i + 1
			break
		}
	}

	f := [101][10001]int{}
	for i := range f {
		for j := range f[i] {
			f[i][j] = -1e9
		}
	}
	f[0][0] = 0
	s := 0
	for i := n - 1; i >= 0; i-- {
		p := a[i]
		s = min(s+p.b, sb)
		for j := m; j > 0; j-- {
			for k := s; k >= p.b; k-- {
				f[j][k] = max(f[j][k], f[j-1][k-p.b]+p.a)
			}
		}
	}
	Println(f[m][sa:])
	for _, save := range f[m][sa:] {
		maxSave = max(maxSave, save)
	}
	Fprint(out, m, sa-maxSave)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF730J(os.Stdin, os.Stdout) }
