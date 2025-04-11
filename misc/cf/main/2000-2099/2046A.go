//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2046A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0
		mx := int(-1e18)
		for _, x := range a {
			Fscan(in, &y)
			ans += max(x, y)
			mx = max(mx, min(x, y))
		}
		Fprintln(out, ans+mx)
	}

}

func main() { CF2046A(os.Stdin, os.Stdout) }
