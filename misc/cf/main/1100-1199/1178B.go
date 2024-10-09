//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1178B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(in, &s)
	var pre, suf, ans int
	for i := 1; i < len(s)-1; i++ {
		if s[i] == 'v' && s[i+1] == 'v' {
			suf++
		}
	}
	for i := 1; i < len(s)-2; i++ {
		if s[i] == 'o' {
			ans += pre * suf
		} else {
			if s[i-1] == 'v' {
				pre++
			}
			if s[i+1] == 'v' {
				suf--
			}
		}
	}
	Fprintln(out, ans)
}

func main() { CF1178B(os.Stdin, os.Stdout) }
