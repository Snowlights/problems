//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1362B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		has := [1024]bool{}
		for i := range a {
			Fscan(in, &a[i])
			has[a[i]] = true
		}
		if n%2 > 0 {
			Fprintln(out, -1)
			continue
		}
		ans := int(1e9)
	o:
		for i := 1; i < n; i++ {
			xor := a[0] ^ a[i]
			for _, v := range a {
				if !has[v^xor] {
					continue o
				}
			}
			ans = min(ans, xor)
		}
		if ans == 1e9 {
			ans = -1
		}
		Fprintln(out, ans)
	}

}

func main() { CF1362B(os.Stdin, os.Stdout) }
