//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
	"sort"
)

func CF985C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, l, ans int
	Fscan(in, &n, &k, &l)
	a := make([]int, n*k)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	i := sort.SearchInts(a, a[0]+l+1)
	if i < n {
		Fprint(out, 0)
		return
	}

	x := 0
	if k > 1 {
		x = (i - n + k - 2) / (k - 1)
	}
	for j := 0; j <= (x-1)*k; j += k {
		ans += a[j]
	}
	for _, v := range a[i-n+x : i] {
		ans += v
	}
	Fprint(out, ans)

}

func main() { CF985C(os.Stdin, os.Stdout) }
