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

func CF607A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	Fscan(in, &n)
	a := make([]struct{ p, d int }, n)
	for i := range a {
		Fscan(in, &a[i].p, &a[i].d)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].p < a[j].p })

	f := make([]int, n)
	for i, p := range a {
		j := sort.Search(n, func(j int) bool { return a[j].p >= p.p-p.d })
		if j > 0 {
			f[i] = f[j-1] + 1
		} else {
			f[i] = 1
		}
		ans = max(ans, f[i])
	}
	Fprint(out, n-ans)
	
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF607A(os.Stdin, os.Stdout) }
