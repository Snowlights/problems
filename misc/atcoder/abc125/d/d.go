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

	var n, v, neg, ans, mm int
	mm = math.MaxInt32
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		if v < 0 {
			neg ^= 1
			v = -v
		}
		ans += v
		if v < mm {
			mm = v
		}
	}
	Fprint(out, ans-mm*neg*2)
}

func main() { run(os.Stdin, os.Stdout) }
