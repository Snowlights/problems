//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF435B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s []byte
	var k int
	Fscan(in, &s, &k)
	for i := 0; i < len(s) && k > 0; i++ {
		swap := i
		for j := i + 1; j < len(s) && j <= i+k; j++ {
			if s[j] > s[swap] {
				swap = j
			}
		}
		for swap > i {
			s[swap], s[swap-1] = s[swap-1], s[swap]
			k--
			swap--
		}
	}
	Fprintln(out, string(s))
}

func main() { CF435B(os.Stdin, os.Stdout) }
