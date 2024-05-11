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

func CF166C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x, ans int
	Fscan(in, &n, &x)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	i := sort.SearchInts(a, x)
	j := sort.SearchInts(a, x+1) - 1
	if i == n || a[i] != x {
		ans = 1
		j = i
		n++
	}
	m := (n - 1) / 2
	if i > m {
		ans += i*2 - n + 1
	} else if j < m {
		ans += n - j*2 - 2
	}
	Fprint(out, ans)

}

func main() { CF166C(os.Stdin, os.Stdout) }
