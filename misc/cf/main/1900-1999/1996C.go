//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1996C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var T, n, q, l, r int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q, &s, &t)
		sum := make([][26]int, n+1)
		for i, b := range s {
			sum[i+1] = sum[i]
			sum[i+1][b-'a']++
			sum[i+1][t[i]-'a']--
		}
		for range q {
			Fscan(in, &l, &r)
			ans := 0
			for i, sl := range sum[l-1] {
				ans += abs(sum[r][i] - sl)
			}
			Fprintln(out, ans/2)
		}
	}

}

func main() { CF1996C(os.Stdin, os.Stdout) }
