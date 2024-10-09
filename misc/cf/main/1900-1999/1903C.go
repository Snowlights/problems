//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1903C(_r io.Reader, _w io.Writer) {
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
		ans, s := 0, 0
		for i := n - 1; i >= 0; i-- {
			s += a[i]
			if i == 0 || s > 0 {
				ans += s
			}
		}
		Fprintln(out, ans)
	}

}

func main() { CF1903C(os.Stdin, os.Stdout) }
