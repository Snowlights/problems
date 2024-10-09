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

func CF1765D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([]int, n)
	ans = n
	for i := range a {
		Fscan(in, &a[i])
		ans += a[i]
	}
	sort.Ints(a)
	l, r := 0, n-1
	for l < r {
		if a[l]+a[r] <= m {
			l++
			if l < r {
				ans -= 2
			} else {
				ans--
			}
		}
		r--
	}
	Fprintln(out, ans)

}

func main() { CF1765D(os.Stdin, os.Stdout) }
