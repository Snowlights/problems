//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

func CF1837D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		cnt, mask, ans := 0, 0, make([]string, n)
		for i, v := range s {
			if v == '(' {
				cnt++
			} else {
				cnt--
			}

			if cnt > 0 {
				ans[i] = "1"
				mask |= 1
			} else if cnt < 0 {
				ans[i] = "2"
				mask |= 2
			} else {
				ans[i] = ans[i-1]
			}
		}

		if cnt != 0 {
			Fprintln(out, -1)
		} else if mask != 3 {
			Fprintln(out, 1)
			Fprintln(out, strings.Repeat("1 ", n))
		} else {
			Fprintln(out, 2)
			Fprintln(out, strings.Join(ans, " "))
		}
	}

}

func main() { CF1837D(os.Stdin, os.Stdout) }
