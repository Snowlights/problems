
//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2004F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0
		cnt := map[int]int{}
		for i := range a {
			s := 0
			for j := i; j < n; j++ {
				s += a[j]
				ans += j - i - cnt[s]
				cnt[s]++
			}
		}
		Fprintln(out, ans)
	}

}

func main() { CF2004F(os.Stdin, os.Stdout) }
