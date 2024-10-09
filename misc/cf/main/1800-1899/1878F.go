//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func factorize(x int) map[int]int {
	cnt := map[int]int{}
	for i := 2; i*i <= x; i++ {
		if x%i > 0 {
			continue
		}
		e := 1
		for x /= i; x%i == 0; x /= i {
			e++
		}
		cnt[i] = e
	}
	if x > 1 {
		cnt[x] = 1
	}
	return cnt
}

func CF1878F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, q, k, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		cnt := factorize(n)
	next:
		for ; q > 0; q-- {
			Fscan(in, &k)
			if k == 2 {
				cnt = factorize(n)
				continue
			}
			Fscan(in, &x)
			for p, e := range factorize(x) {
				cnt[p] += e
			}
			d := 1
			for _, e := range cnt {
				d *= e + 1
			}
			for p, e := range factorize(d) {
				if e > cnt[p] {
					Fprintln(out, "NO")
					continue next
				}
			}
			Fprintln(out, "YES")
		}
	}

}

func main() { CF1878F(os.Stdin, os.Stdout) }
