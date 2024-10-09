package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	h, r := map[int]int{
		1: 1,
	}, 2
	for i := 2; i < 55000; i++ {
		h[i] = h[i-1]
		if r*r == i {
			h[i] += h[r]
			r++
		}
	}

	var n, m, v int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &m)
		v = int(math.Sqrt(float64(m)))
		if v*v > m {
			v--
		}
		ans := 0
		for i := 1; i*i <= v; i++ {
			ans += h[i] * (v - i*i + 1)
		}
		Fprintln(out, ans)
	}

}

func main() { run(os.Stdin, os.Stdout) }
