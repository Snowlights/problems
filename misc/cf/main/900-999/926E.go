//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF926E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	Fscan(in, &n)
	st := []int{}
	for range n {
		Fscan(in, &v)
		st = append(st, v)
		for len(st) > 1 && st[len(st)-2] == st[len(st)-1] {
			st[len(st)-2]++
			st = st[:len(st)-1]
		}
	}
	Fprintln(out, len(st))
	for _, v := range st {
		Fprint(out, v, " ")
	}
}

func main() { CF926E(os.Stdin, os.Stdout) }
