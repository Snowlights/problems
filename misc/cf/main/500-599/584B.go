//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF584B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod int = 1e9 + 7
	pow := func(x, n int) int {
		//x %= mod
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var n int
	Fscan(in, &n)
	Fprintln(out, (pow(27, n) - pow(7, n) + mod) % mod )
}

func main() { CF584B(os.Stdin, os.Stdout) }
