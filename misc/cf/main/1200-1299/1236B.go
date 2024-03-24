//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1236B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod int = 1e9 + 7
	pow := func(x, n int) int {
		// x %= mod
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var n, m int
	Fscan(in, &n, &m)
	Fprintln(out, pow(pow(2, m)-1, n))
}

func main() { CF1236B(os.Stdin, os.Stdout) }
