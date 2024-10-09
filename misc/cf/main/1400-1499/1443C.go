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

func CF1443C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct {
			a, b int
		}
		p, s := make([]pair, n), 0
		for i := range p {
			Fscan(in, &p[i].a)
		}
		for i := range p {
			Fscan(in, &p[i].b)
			s += p[i].b
		}
		sort.Slice(p, func(i, j int) bool {
			return p[i].a < p[j].a
		})
		ans := s
		for i := range p {
			s -= p[i].b
			ans = min(ans, max(p[i].a, s))
		}
		Fprintln(out, ans)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF1443C(os.Stdin, os.Stdout) }
