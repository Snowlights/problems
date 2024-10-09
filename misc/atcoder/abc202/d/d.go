package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 61
	C := [mx][mx]int{}
	for i := 0; i < mx; i++ {
		C[i][0], C[i][i] = 1, 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}

	var a, b, k int
	Fscan(in, &a, &b, &k)
	ans := bytes.Repeat([]byte{'a'}, a+b)
	for i := range ans {
		if k > C[a+b-1][b] {
			k -= C[a+b-1][b]
			ans[i] = 'b'
			b--
		} else {
			a--
		}
	}
	Fprintln(out, string(ans))
}

func main() { run(os.Stdin, os.Stdout) }
