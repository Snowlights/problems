//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF1904C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		if k >= 3 {
			Fprintln(out, 0)
			continue
		}
		sort.Ints(a)
		ans := a[0]
		for i, v := range a[1:] {
			ans = min(ans, v-a[i])
		}
		if k == 1 {
			Fprintln(out, ans)
			continue
		}
		for i, v := range a {
			for _, vv := range a[i+1:] {
				val := abs(v - vv)
				j := sort.SearchInts(a, val)
				if j < n {
					ans = min(ans, a[j]-val)
				}
				if j > 0 {
					ans = min(ans, val-a[j-1])
				}
			}
		}
		Fprintln(out, ans)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() { CF1904C(os.Stdin, os.Stdout) }
