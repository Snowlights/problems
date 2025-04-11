//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

func CF2047A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		ans := 0
		s := 0
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			s += v
			rt := int(math.Sqrt(float64(s)))
			if rt%2 > 0 && rt*rt == s {
				ans++
			}
		}
		Fprintln(out, ans)
	}

}

func main() { CF2047A(os.Stdin, os.Stdout) }
