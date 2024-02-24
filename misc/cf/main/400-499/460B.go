//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF460B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var a, b, c int
	Fscan(in, &a, &b, &c)
	ans := make([]int, 0, 10)
	for i := 1; i < 82; i++ {
		x := b
		for j := 0; j < a; j++ {
			x = x * i
		}
		x += c
		if x >= 1e9 {
			break
		}
		if x > 0 {
			s := i
			for x := x; x > 0; x /= 10 {
				s -= x % 10
			}
			if s == 0 {
				ans = append(ans, x)
			}
		}
	}
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

func main() { CF460B(os.Stdin, os.Stdout) }
