package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/big"
	"os"
	"strings"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	var x, y int
	Fscan(in, &s, &x, &y)
	a := strings.Split(s, "T")
	f := func(a []string, i, x int) bool {
		x = abs(x)
		f := big.NewInt(1)
		p := new(big.Int)
		for ; i < len(a); i += 2 {
			f.Or(f, p.Lsh(f, uint(len(a[i])*2)))
			x += len(a[i])
		}
		return f.Bit(x) > 0
	}
	if f(a, 2, x-len(a[0])) && f(a, 1, y) {
		Fprint(out, "Yes")
	} else {
		Fprint(out, "No")
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() { run(os.Stdin, os.Stdout) }
