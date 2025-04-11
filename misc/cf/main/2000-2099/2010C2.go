//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2010C2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(in, &s)
	n := len(s)
	z := make([]int, n)
	boxL, boxR := 0, 0
	for i := 1; i < (n+1)/2; i++ {
		if i <= boxR {
			z[i] = min(z[i-boxL], boxR-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			boxL, boxR = i, i+z[i]
			z[i]++
		}
		if i+z[i] >= n {
			Fprintln(out, "YES", s[i:])
			return
		}
	}
	Fprint(out, "NO")

}

func main() { CF2010C2(os.Stdin, os.Stdout) }
