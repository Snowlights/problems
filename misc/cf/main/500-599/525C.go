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

func CF525C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for i := -1; i < n; {
		for i += 2; i < n && a[i-1]-a[i] > 1; i++ {
		}
		if i >= n {
			break
		}
		v := a[i]
		for i += 2; i < n && a[i-1]-a[i] > 1; i++ {
		}
		if i >= n {
			break
		}
		w := a[i]
		ans += v * w
	}
	Fprint(out, ans)
}

func main() { CF525C(os.Stdin, os.Stdout) }
