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

func CF486E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	l := make([]int, n)
	dp := []int{}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		p := sort.SearchInts(dp, a[i])
		if p < len(dp) {
			dp[p] = a[i]
		} else {
			dp = append(dp, a[i])
		}
		l[i] = p + 1
	}
	r := make([]int, n)
	dp = nil
	for i := n - 1; i >= 0; i-- {
		v := -a[i]
		p := sort.SearchInts(dp, v)
		if p < len(dp) {
			dp[p] = v
		} else {
			dp = append(dp, v)
		}
		r[i] = p + 1
	}
	ans := make([]byte, n)
	cnt := make([]int, n+1)
	for i, l := range l {
		if l+r[i]-1 != len(dp) {
			ans[i] = '1'
		} else {
			ans[i] = '3'
			cnt[l]++
		}
	}
	for i, v := range ans {
		if v == '3' && cnt[l[i]] > 1 {
			ans[i] = '2'
		}
	}
	Fprint(out, string(ans))

}

func main() { CF486E(os.Stdin, os.Stdout) }
