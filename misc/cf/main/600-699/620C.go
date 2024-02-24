//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF620C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a, h := make([]int, n), make(map[int]bool)
	ans, cur := make([][]int, 0), make([]int, 2)
	for i := range a {
		Fscan(in, &a[i])
		if h[a[i]] {
			cur[1] = i + 1
			ans = append(ans, cur)
			cur = make([]int, 2)
			h = make(map[int]bool)
			continue
		}
		if cur[0] == 0 {
			cur[0] = i + 1
		}
		h[a[i]] = true
	}
	if len(ans) == 0 {
		Fprintln(out, -1)
	} else {
		ans[len(ans)-1][1] = n
		Fprintln(out, len(ans))
		for _, v := range ans {
			Fprintln(out, v[0], v[1])
		}
	}

}

func main() { CF620C(os.Stdin, os.Stdout) }
