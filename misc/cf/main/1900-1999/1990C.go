//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1990C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		vis := make([]int8, n+1)
		ans, mx := 0, 0
		for i := range a {
			Fscan(in, &v)
			ans += v
			if vis[v] == 1 {
				mx = max(mx, v)
			} else {
				vis[v] = 1
			}
			a[i] = mx
		}

		mx = 0
		for i, v := range a {
			if vis[v] == 2 {
				mx = max(mx, v)
			} else {
				vis[v] = 2
			}
			ans += v + mx*(n-i)
		}
		Fprintln(out, ans)
	}

}

func main() { CF1990C(os.Stdin, os.Stdout) }
