//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1994E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, k, n int
	Fscan(in, &T)
	for range T {
		ans := 0
		Fscan(in, &k)
		for range k {
			Fscan(in, &n)
			c := ans | n
			for i := 19; i >= 0; i-- {
				if ans&n>>i&1 > 0 {
					c |= 1<<i - 1
					break
				}
			}
			ans = c
			for range n - 1 {
				Fscan(in, &n)
			}
		}
		Fprintln(out, ans)
	}

}

func main() { CF1994E(os.Stdin, os.Stdout) }
