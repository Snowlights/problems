//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1678B2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, ans, seg int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		ans, seg = 0, 0
		for i, pre := 0, byte(0); i < n; i += 2 {
			if s[i] != s[i+1] {
				ans++
			} else if s[i] != pre {
				seg++
				pre = s[i]
			}
		}
		Fprintln(out, ans, max(seg, 1))
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF1678B2(os.Stdin, os.Stdout) }
