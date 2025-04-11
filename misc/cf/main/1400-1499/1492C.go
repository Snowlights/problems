//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1492C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, ans int
	var s, t []byte
	Fscan(in, &n, &m, &s, &t)

	l := make([]int, m)
	k := 0
	for i, b := range t {
		for s[k] != b {
			k++
		}
		l[i] = k
		k++
	}

	r := make([]int, m)
	k = n - 1
	for i := m - 1; i >= 0; i-- {
		b := t[i]
		for s[k] != b {
			k--
		}
		r[i] = k
		k--
	}

	for i := 1; i < m; i++ {
		ans = max(ans, l[i]-l[i-1], r[i]-r[i-1], r[i]-l[i-1])
	}
	Fprint(out, ans)

}

func main() { CF1492C(os.Stdin, os.Stdout) }
