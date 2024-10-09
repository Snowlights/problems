//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1077C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, sum, mx, mx2, mxI int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		v := a[i]
		sum += v
		if v > mx {
			mx2 = mx
			mx = v
			mxI = i
		} else if v > mx2 {
			mx2 = v
		}
	}

	ans := []any{}
	for i, v := range a {
		if i == mxI && sum-v == mx2*2 || i != mxI && sum-v == mx*2 {
			ans = append(ans, i+1)
		}
	}
	Fprintln(out, len(ans))
	Fprintln(out, ans...)

}

func main() { CF1077C(os.Stdin, os.Stdout) }
