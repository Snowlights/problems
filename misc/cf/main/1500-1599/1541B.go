//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1541B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, aj int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans, last := 0, make([]int, 2*n+1)
		for j := 1; j <= n; j++ {
			Fscan(in, &aj)
			for ai := 1; ai*aj < j*2; ai++ {
				if i := last[ai]; i > 0 && ai*aj == i+j {
					ans++
				}
			}
			last[aj] = j
		}
		Fprintln(out, ans)
	}
}

func main() { CF1541B(os.Stdin, os.Stdout) }
