//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"math/bits"
	"os"
	"strconv"
)

func CF962C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	ans := 0
	Fscan(in, &s)
	for i := uint(1); i < 1<<len(s); i++ {
		if s[bits.TrailingZeros(i)] == '0' {
			continue
		}
		x := 0
		for j := i; j > 0; j &= j - 1 {
			x = x*10 + int(s[bits.TrailingZeros(j)]-'0')
		}
		if q := int(math.Sqrt(float64(x))); q*q == x {
			ans = max(ans, len(strconv.Itoa(x)))
		}
	}

	if ans == 0 {
		Fprintln(out, -1)
	} else {
		Fprintln(out, len(s)-ans)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF962C(os.Stdin, os.Stdout) }
