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

func CF608B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var (
		s, t        string
		ans, cnt, d int
	)
	Fscan(in, &s, &t)
	d = len(t) - len(s)
	cnt = strings.Count(t[:d], "1")
	for i, v := range s {
		cnt += int(t[i+d] - '0')
		if v == '0' {
			ans += cnt
		} else {
			ans += d - cnt + 1
		}
		cnt -= int(t[i] - '0')
	}
	Fprintln(out, ans)
}

func main() { CF608B(os.Stdin, os.Stdout) }
