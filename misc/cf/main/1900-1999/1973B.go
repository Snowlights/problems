//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1973B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 1
		for i := 0; i < 20; i++ {
			pre := -1
			for j, v := range a {
				if v>>i&1 > 0 {
					ans = max(ans, j-pre)
					pre = j
				}
			}
			if pre >= 0 {
				ans = max(ans, n-pre)
			}
		}
		Fprintln(out, ans)
	}
}

func main() { CF1973B(os.Stdin, os.Stdout) }
