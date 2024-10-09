//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

func CF702C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	//var n, m, ans int
	//Fscan(in, &n, &m)
	//a, b := make([]int, n), make([]int, m)
	//for i := range a {
	//	Fscan(in, &a[i])
	//}
	//for i := range b {
	//	Fscan(in, &b[i])
	//}
	//
	//for _, v := range a {
	//	i := sort.SearchInts(b, v)
	//	c, d := math.MaxInt32, math.MaxInt32
	//	if i < m {
	//		c = b[i] - v
	//	}
	//	if i > 0 {
	//		d = v - b[i-1]
	//	}
	//	ans = max(ans, min(c, d))
	//}

	var n, m, v, i, ans int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	pre := math.MinInt64 / 2
	for ; m > 0; m-- {
		Fscan(in, &v)
		for i < n && a[i] <= v {
			ans = max(ans, min(a[i]-pre, v-a[i]))
			i++
		}
		pre = v
	}
	if i < n {
		ans = max(ans, a[n-1]-pre)
	}
	Fprintln(out, ans)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF702C(os.Stdin, os.Stdout) }
