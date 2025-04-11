//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2019B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, q, pre, v, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q, &pre)
		cnt := map[int]int{n - 1: 1}
		for i := 1; i < n; i++ {
			Fscan(in, &v)
			cnt[i*(n-i)] += v - pre - 1
			cnt[(i+1)*(n-i)-1]++
			pre = v
		}
		for range q {
			Fscan(in, &k)
			Fprint(out, cnt[k], " ")
		}
		Fprintln(out)
	}

}

func main() { CF2019B(os.Stdin, os.Stdout) }
