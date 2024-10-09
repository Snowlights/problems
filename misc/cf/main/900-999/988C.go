//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF988C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var k, n int
	type pair struct {
		k, i int
	}
	h := make(map[int]pair)
	Fscan(in, &k)
	for j := 1; j <= k; j++ {
		Fscan(in, &n)
		sum, a := 0, make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &a[i])
			sum += a[i]
		}
		for i := 0; i < n; i++ {
			if p, ok := h[sum-a[i]]; ok && p.k < j {
				Fprintln(out, "YES")
				Fprintln(out, p.k, p.i+1)
				Fprintln(out, j, i+1)
				return
			}
			h[sum-a[i]] = pair{
				k: j,
				i: i,
			}
		}

	}
	Fprintln(out, "NO")

}

func main() { CF988C(os.Stdin, os.Stdout) }
