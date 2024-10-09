//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

func CF1759D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		two, five, k := bits.TrailingZeros(uint(n)), 0, 1

		for x := n; x%5 == 0; x /= 5 {
			five++
		}
		for two < five && k*2 <= m {
			two++
			k *= 2
		}
		for five < two && k*5 <= m {
			five++
			k *= 5
		}
		for k*10 <= m {
			k *= 10
		}
		Fprintln(out, n*(m-m%k))
	}

}

func main() { CF1759D(os.Stdin, os.Stdout) }
