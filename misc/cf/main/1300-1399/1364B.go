//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1364B(_r io.Reader, _w io.Writer) {
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
		ans := []any{a[0]}
		for i := 0; i < n-1; {
			st := i
			for i++; i < n && a[i] > a[i-1] == (a[st+1] > a[st]); i++ {
			}
			i--
			ans = append(ans, a[i])
		}
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}

}

func main() { CF1364B(os.Stdin, os.Stdout) }
