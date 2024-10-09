//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func CF1870D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, c, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ i, c int }
		st := []pair{{0, 0}}
		for i := 1; i <= n; i++ {
			Fscan(in, &c)
			for c <= st[len(st)-1].c {
				st = st[:len(st)-1]
			}
			st = append(st, pair{i, c})
		}

		Fscan(in, &k)
		h := int(1e9)
		for i := 1; i < len(st); i++ {
			d := st[i].c - st[i-1].c
			h = min(h, k/d)
			k -= d * h
			Fprint(out, strings.Repeat(strconv.Itoa(h)+" ", st[i].i-st[i-1].i))
		}
		Fprintln(out)
	}

}

func main() { CF1870D(os.Stdin, os.Stdout) }
