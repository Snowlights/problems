//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1097C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	var s string
	cnt := make(map[int]int)
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		c, mm := 0, 0
		for _, v := range s {
			switch v {
			case '(':
				c++
			default:
				c--
				mm = min(mm, c)
			}
		}
		if mm == 0 || mm == c {
			if cnt[-c] > 0 {
				ans++
				cnt[-c]--
			} else {
				cnt[c]++
			}
		}
	}
	Fprintln(out, ans)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF1097C(os.Stdin, os.Stdout) }
