//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2009D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		has := make([][2]bool, n+1)
		cnt := 0
		for range n {
			Fscan(in, &x, &y)
			has[x][y] = true
			if has[x][0] && has[x][1] {
				cnt++
			}
		}

		ans := 0
		for i := 1; i < n; i++ {
			if has[i][0] && has[i-1][1] && has[i+1][1] {
				ans++
			}
			if has[i][1] && has[i-1][0] && has[i+1][0] {
				ans++
			}
		}
		Fprintln(out, ans+cnt*(n-2))
	}

}

func main() { CF2009D(os.Stdin, os.Stdout) }
