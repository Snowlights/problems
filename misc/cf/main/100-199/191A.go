//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF191A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	var s string
	f := [26][26]int{}
	for i := range f {
		for j := range f[i] {
			f[i][j] = -1e9
		}
	}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		m := len(s)
		v, w := s[0]-'a', s[m-1]-'a'
		for j := 0; j < 26; j++ {
			f[j][w] = max(f[j][w], f[j][v]+m)
		}
		f[v][w] = max(f[v][w], m)
	}
	for i := 0; i < 26; i++ {
		ans = max(ans, f[i][i])
	}
	Fprint(out, ans)

}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() { CF191A(os.Stdin, os.Stdout) }
