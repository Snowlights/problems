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

func CF1777C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx = 100001
	ds := [mx][]int{}
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i {
			ds[j] = append(ds[j], i)
		}
	}

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		a = slices.Compact(a)

		ans := int(1e9)
		todo := m
		l := 0
		cnt := make([]int, m+1)
		for _, v := range a {
			for _, d := range ds[v] {
				if d > m {
					break
				}
				if cnt[d] == 0 {
					todo--
				}
				cnt[d]++
			}
			for todo == 0 {
				ans = min(ans, v-a[l])
				for _, d := range ds[a[l]] {
					if d > m {
						break
					}
					cnt[d]--
					if cnt[d] == 0 {
						todo++
					}
				}
				l++
			}
		}
		if ans == 1e9 {
			ans = -1
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

func main() { CF1777C(os.Stdin, os.Stdout) }
