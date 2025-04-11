package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, v, l, r, L, R int
	Fscan(in, &n, &q)
	mp := map[int]int{}
	f := func() []int {
		s := make([]int, n+1)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			if mp[v] == 0 {
				mp[v] = rand.Intn(1e13)
			}
			s[i+1] = s[i] + mp[v]
		}
		return s
	}
	sa := f()
	sb := f()
	for ; q > 0; q-- {
		Fscan(in, &l, &r, &L, &R)
		if sa[r]-sa[l-1] == sb[R]-sb[L-1] {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}

}

func main() { run(os.Stdin, os.Stdout) }
