//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

func CF2032C(_r io.Reader, _w io.Writer) {
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
		slices.Sort(a)

		ans := n - 2
		j := 0
		for i := 2; i < n; i++ {
			for a[j]+a[j+1] <= a[i] {
				j++
			}
			ans = min(ans, n-1-i+j)
		}
		Fprintln(out, ans)
	}

}

func main() { CF2032C(os.Stdin, os.Stdout) }
