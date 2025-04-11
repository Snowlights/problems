//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1847B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans, and := 0, -1
		for range n {
			Fscan(in, &v)
			and &= v
			if and == 0 {
				ans++
				and = -1
			}
		}
		Fprintln(out, max(ans, 1))
	}

}

func main() { CF1847B(os.Stdin, os.Stdout) }
