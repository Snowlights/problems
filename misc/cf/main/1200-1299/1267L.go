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

func CF1267L(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k, p, i0 int
	var s []byte
	Fscan(in, &n, &m, &k, &s)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	ans := make([][]byte, n)
	for i := range ans {
		ans[i] = make([]byte, m)
	}
	for j := 0; j < m; j++ {
		for i := i0; i < k; i++ {
			ans[i][j] = s[p]
			if i > i0 && s[p] != s[p-1] {
				i0 = i
			}
			p++
		}
	}

	for _, t := range ans {
		for j, b := range t {
			if b == 0 {
				t[j] = s[p]
				p++
			}
		}
		Fprintln(out, string(t))
	}
	
}

func main() { CF1267L(os.Stdin, os.Stdout) }
