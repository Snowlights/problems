//go:build main
// +build main

package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"os"
)

func CF1265B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pos := make([]int, n)
		for i := range n {
			Fscan(in, &v)
			pos[v-1] = i
		}
		mn, mx := n, -1
		ans := bytes.Repeat([]byte{'0'}, n)
		for i, p := range pos {
			mn = min(mn, p)
			mx = max(mx, p)
			if mx-mn == i {
				ans[i] = '1'
			}
		}
		Fprintf(out, "%s\n", ans)
	}

}

func main() { CF1265B(os.Stdin, os.Stdout) }
