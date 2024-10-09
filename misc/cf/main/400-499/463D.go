//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF463D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, ans int
	Fscan(in, &m, &n)
	p := make([][]int, n)
	a := make([]int, m)
	for i := range p {
		p[i] = make([]int, m+1)
		for j := range a {
			Fscan(in, &a[j])
			p[i][a[j]] = j
		}
	}

	f := make([]int, m)
	for i, v := range a {
	o:
		for j, w := range a[:i] {
			for _, p := range p {
				if p[w] > p[v] {
					continue o
				}
			}
			f[i] = max(f[i], f[j])
		}
		f[i]++
		ans = max(ans, f[i])
	}
	Fprint(out, ans)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF463D(os.Stdin, os.Stdout) }
