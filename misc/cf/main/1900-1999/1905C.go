//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

func CF1905C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		st := []int{}
		for i, b := range s {
			for len(st) > 0 && b > s[st[len(st)-1]] {
				st = st[:len(st)-1]
			}
			st = append(st, i)
		}
		t := slices.Clone(s)
		m := len(st)
		for i, idx := range st {
			t[idx] = s[st[m-1-i]]
		}
		if !slices.IsSorted(t) {
			Fprintln(out, -1)
			continue
		}
		for _, i := range st {
			if s[i] == s[st[0]] {
				m--
			}
		}
		Fprintln(out, m)
	}

}

func main() { CF1905C(os.Stdin, os.Stdout) }
