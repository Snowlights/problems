//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

func CF1883E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v, pre, ans, p int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &pre)
		ans, p = 0, 0
		for i := 1; i < n; i++ {
			Fscan(in, &v)
			if v <= pre {
				p += bits.Len(uint((pre - 1) / v))
			} else {
				p = max(p-bits.Len(uint(v/pre))+1, 0)
			}
			ans += p
			pre = v
		}
		Fprintln(out, ans)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF1883E(os.Stdin, os.Stdout) }
