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

	var n, m int
	ans := math.MaxInt64
	Fscan(in, &n, &m)
	for a := 1; a <= n; a++ {
		b := (m-1)/a + 1
		if b < a {
			break
		}
		if b <= n && a*b < ans {
			ans = a * b
		}
	}
	if ans == math.MaxInt64 {
		Fprintln(out, -1)
		return
	}
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
