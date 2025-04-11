//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2028C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, low int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &low)
		a := make([]int, n)
		pre := []int{0}
		tot, s := 0, 0
		for i := range a {
			Fscan(in, &a[i])
			tot += a[i]
			s += a[i]
			if s >= low {
				pre = append(pre, tot)
				s = 0
			}
		}
		if len(pre) <= m {
			Fprintln(out, -1)
			continue
		}
		ans := tot - pre[m]
		s = 0
		j := m - 1
		for i := n - 1; i >= 0 && j >= 0; i-- {
			tot -= a[i]
			s += a[i]
			if s >= low {
				ans = max(ans, tot-pre[j])
				j--
				s = 0
			}
		}
		Fprintln(out, ans)
	}

}

func main() { CF2028C(os.Stdin, os.Stdout) }
