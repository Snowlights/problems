//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1969B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	var l, t int
	var cost int
	for Fscan(in, &t); t != 0; t-- {
		cost, l = 0, 0
		Fscan(in, &s)
		for i := 0; i < len(s); i++ {
			if s[i] == '1' {
				l++
			} else if l != 0 {
				cost += l + 1
			}
		}

		Fprintln(out, cost)
	}

}

func main() { CF1969B(os.Stdin, os.Stdout) }
