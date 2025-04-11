//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2039B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, s := 0, ""
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		for i := 1; i < len(s); i++ {
			if s[i-1] == s[i] {
				Fprintln(out, s[i-1:i+1])
				continue o
			}
			if i > 1 && s[i-2] != s[i] {
				Fprintln(out, s[i-2:i+1])
				continue o
			}
		}
		Fprintln(out, -1)
	}
	
}

func main() { CF2039B(os.Stdin, os.Stdout) }
