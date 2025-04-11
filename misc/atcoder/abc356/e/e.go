package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, ans int
	const mx int = 1e6
	s := [mx + 1]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		s[v]++
	}
	for i := 2; i <= mx; i++ {
		s[i] += s[i-1]
	}
	for i := 1; i <= mx; i++ {
		c := s[i] - s[i-1]
		if c == 0 {
			continue
		}
		for j := i; j <= mx; j += i {
			ans += j / i * (s[min(j+i-1, mx)] - s[j-1]) * c
		}
		ans -= c * (c + 1) / 2
	}
	Fprint(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
