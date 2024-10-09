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

	gcd := func(a, b int) int {
		for a%b != 0 {
			a, b = b, a%b
		}
		return b
	}

	var l, r int
	Fscan(in, &l, &r)
	for i := r - l; i > 0; i-- {
		for j := l; j+i <= r; j++ {
			if gcd(j, i+j) == 1 {
				Fprintln(out, i)
				return
			}
		}
	}

}

func main() { run(os.Stdin, os.Stdout) }
