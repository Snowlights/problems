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

	var n, x1, x2, x3 int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &x1, &x2, &x3)
		total := x1 + x2 + x3
		if (total%3 != 0) || (x1%2 != x2%2 || x1%2 != x3%2) {
			Fprintln(out, -1)
			continue
		}
		total /= 3
		f := abs(x1-total) + abs(x2-total) + abs(x3-total)
		if f == 0 {
			Fprintln(out, 0)
		} else {
			Fprintln(out, f/4)
		}
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() { run(os.Stdin, os.Stdout) }
