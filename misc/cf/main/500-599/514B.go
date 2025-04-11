//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF514B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x0, y0, x, y, ans int
	Fscan(in, &n, &x0, &y0)
	has := map[float64]bool{}
	for range n {
		Fscan(in, &x, &y)
		if x == x0 {
			ans = 1
		} else {
			has[float64(y-y0)/float64(x-x0)] = true
		}
	}
	Fprintln(out, ans+len(has))

}

func main() { CF514B(os.Stdin, os.Stdout) }
