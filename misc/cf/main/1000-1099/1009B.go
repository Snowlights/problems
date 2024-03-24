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

func CF1009B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(in, &s)
	a := strings.Repeat("1", strings.Count(s, "1"))
	s = strings.ReplaceAll(s, "1", "")
	i := strings.Index(s, "2")
	if i < 0 {
		Fprintln(out, s+a)
	} else {
		Fprintln(out, s[:i] + a + s[i:])
	}
}

func main() { CF1009B(os.Stdin, os.Stdout) }
