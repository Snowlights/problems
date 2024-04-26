//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1800F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	var s string
	Fscan(in, &n)
	a := make([]struct{ m, all uint32 }, n)
	for i := range a {
		Fscan(in, &s)
		for _, c := range s {
			b := uint32(1) << (c - 'a')
			a[i].m ^= b
			a[i].all |= b
		}
	}
	for k := 0; k < 26; k++ {
		cnt := map[uint32]int{}
		for _, p := range a {
			if p.all>>k&1 == 0 {
				ans += cnt[1<<26-1^1<<k^p.m]
				cnt[p.m]++
			}
		}
	}
	Fprint(out, ans)

}

func main() { CF1800F(os.Stdin, os.Stdout) }
