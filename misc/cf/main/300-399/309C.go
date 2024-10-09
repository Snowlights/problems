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

func CF309C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, ans uint
	Fscan(in, &n, &m)
	var ca, cb [30]int
	for ; n > 0; n-- {
		for Fscan(in, &v); v > 0; v &= v - 1 {
			ca[bits.TrailingZeros(v)]++
		}
	}
	for ; m > 0; m-- {
		Fscan(in, &v)
		cb[v]++
	}

O:
	for i, c := range cb {
	o:
		for ; c > 0; c-- {
			for j := i; j < 30; j++ {
				if ca[j] > 0 {
					ans++
					ca[j]--
					for j--; j >= i; j-- {
						ca[j]++
					}
					continue o
				}
			}
			break O
		}
	}
	Fprint(out, ans)

}

func main() { CF309C(os.Stdin, os.Stdout) }
