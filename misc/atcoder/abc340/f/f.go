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

	var a, b int
	Fscan(in, &a, &b)
	a = -a
	g, x0, y0 := exgcd(b, a)
	c := 2
	if c%g != 0 {
		Fprint(out, -1)
		return
	}
	c /= g
	Fprintln(out, x0*c, y0*c)

}

func exgcd(a, b int) (gcd, x, y int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, y, x = exgcd(b, a%b)
	y -= a / b * x
	return
}

func main() { run(os.Stdin, os.Stdout) }
