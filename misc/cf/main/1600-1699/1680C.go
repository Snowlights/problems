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

func CF1680C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer func(out *bufio.Writer) {
		err := out.Flush()
		if err != nil {

		}
	}(out)

	var n int
	var s string
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		zero, one, l, ans := 0, strings.Count(s, "1"), 0, len(s)
		for _, c := range s {
			v := int(c & 1)
			zero += v ^ 1
			one -= v
			if zero > one {
				v = int(s[l] & 1)
				zero -= v ^ 1
				one += v
				l++
			}
			ans = min(ans, one)
		}
		Fprintln(out, ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF1680C(os.Stdin, os.Stdout) }
