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

func CF115B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, ans, cur, down int
	var s string
	Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		Fscan(in, &s)
		l := strings.IndexByte(s, 'W')
		if l < 0 {
			continue
		}
		down = i
		r := strings.LastIndexByte(s, 'W')
		if i%2 == 0 {
			if cur <= l {
				ans += r - cur
			} else {
				ans += cur - l + r - l
			}
			cur = r
		} else {
			if cur >= r {
				ans += cur - l
			} else {
				ans += r - cur + r - l
			}
			cur = l
		}
	}
	Fprint(out, ans+down)

}

func main() { CF115B(os.Stdin, os.Stdout) }
