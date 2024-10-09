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

func CF985E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, d int32
	Fscan(in, &n, &k, &d)
	a := make([]int32, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)

	f := make([]bool, n+1)
	f[0] = true
	left, cnt := 0, 0
	for i := k; i <= n; i++ {
		if f[i-k] {
			cnt++
		}
		for a[i-1]-a[left] > d {
			if f[left] {
				cnt--
			}
			left++
		}
		f[i] = cnt > 0
	}
	if f[n] {
		Fprintln(out, "YES")
	} else {
		Fprintln(out, "NO")
	}

}

func main() { CF985E(os.Stdin, os.Stdout) }
