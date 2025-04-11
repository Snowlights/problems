//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF363C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s []byte
	Fscan(in, &s)
	type pair struct {
		c int
		b byte
	}
	a := []pair{}
	cnt := 0
	for i, b := range s {
		cnt++
		if i == len(s)-1 || b != s[i+1] {
			a = append(a, pair{min(cnt, 2), b})
			cnt = 0
		}
	}

	for i, n := 0, len(a); i < n; i++ {
		if a[i].c == 1 {
			continue
		}
		st := i
		for ; i < n && a[i].c == 2; i++ {
		}
		for st++; st < i; st += 2 {
			a[st].c = 1
		}
	}

	ans := []byte{}
	for _, p := range a {
		ans = append(ans, bytes.Repeat([]byte{p.b}, p.c)...)
	}
	Fprintf(out, "%s", ans)

}

func main() { CF363C(os.Stdin, os.Stdout) }
