//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF1730C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		st := s[:0]
		var t []byte
		for _, b := range s {
			for len(st) > 0 && b < st[len(st)-1] {
				t = append(t, min(st[len(st)-1]+1, '9'))
				st = st[:len(st)-1]
			}
			st = append(st, b)
		}
		st = append(st, t...)
		sort.Slice(st, func(i, j int) bool { return st[i] < st[j] })
		Fprintln(out, string(st))
	}
	
}

func min(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}

func main() { CF1730C(os.Stdin, os.Stdout) }
