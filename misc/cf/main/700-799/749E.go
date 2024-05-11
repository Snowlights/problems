//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF749E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, s0 int
	Fscan(in, &n)
	s1 := 0.
	t := make([][2]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		v = n + 1 - v
		s := 0
		for j := v; j > 0; j &= j - 1 {
			s0 += t[j][0]
			s += t[j][1]
		}
		s1 += float64(s * (n + 1 - i))
		for j := v; j <= n; j += j & -j {
			t[j][0]++
			t[j][1] += i
		}
	}
	m := float64(n)
	Printf("%.9f", float64(s0)-s1/m/(m+1)*2+(m+2)*(m-1)/24)

}

func main() { CF749E(os.Stdin, os.Stdout) }
