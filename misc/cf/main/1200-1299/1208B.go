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

func CF1208B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans := sort.Search(n, func(m int) bool {
		cnt := map[int]int{}
		for _, v := range a[m:] {
			cnt[v]++
		}
		if len(cnt) == n-m {
			return true
		}
		for i := m; i < n; i++ {
			v := a[i]
			cnt[v]--
			if cnt[v] == 0 {
				delete(cnt, v)
			}
			cnt[a[i-m]]++
			if len(cnt) == n-m {
				return true
			}
		}
		return false
	})
	Fprint(out, ans)

}

func main() { CF1208B(os.Stdin, os.Stdout) }
