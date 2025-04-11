//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func isqrt9(x uint) int {
	rt := uint(math.Sqrt(float64(x)))
	if rt*rt > x {
		rt--
	}
	return int(rt)
}

func CF2009E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		s := (k*2 + n - 1) * n / 2
		b := k*2 - 1
		i := (isqrt9(uint(b*b)+uint(s)*4) - b) / 2
		f := func(i int) int { return (k*2+i-1)*i - s }
		Fprintln(out, min(-f(i), f(i+1)))
	}

}

func main() { CF2009E(os.Stdin, os.Stdout) }
