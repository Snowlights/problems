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

func CF1895D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, a0 uint
	Fscan(in, &n)
	a := make([]uint, n)
	c1 := make([]int, 64)
	for i := uint(1); i < n; i++ {
		Fscan(in, &a[i])
		a[i] ^= a[i-1]
		for j := a[i]; j > 0; j &= j - 1 {
			c1[bits.TrailingZeros(j)]++
		}
		for j := i; j > 0; j &= j - 1 {
			c1[bits.TrailingZeros(j)]--
		}
	}
	for i, c := range c1 {
		if c != 0 {
			a0 |= 1 << i
		}
	}
	Fprint(out, a0)
	for _, v := range a[1:] {
		Fprint(out, " ", v^a0)
	}

}

func main() { CF1895D(os.Stdin, os.Stdout) }
