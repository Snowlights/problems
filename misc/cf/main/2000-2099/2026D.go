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

func CF2026D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, q, l, r int
	Fscan(in, &n)
	sum := make([]int, n+1)
	iSum := make([]int, n+1)
	for i := range n {
		Fscan(in, &v)
		sum[i+1] = sum[i] + v
		iSum[i+1] = iSum[i] + (n-i)*v
	}

	iSumSum := make([]int, n+1)
	for i := range n {
		iSumSum[i+1] = iSumSum[i] + iSum[n] - iSum[i]
	}

	m := n*2 + 1
	f := func(k int) int {
		i := (m - int(math.Ceil(math.Sqrt(float64(m*m-k*8))))) / 2
		k -= (m - i) * i / 2
		return iSumSum[i] + iSum[i+k] - iSum[i] - (n-i-k)*(sum[i+k]-sum[i])
	}

	Fscan(in, &q)
	for range q {
		Fscan(in, &l, &r)
		Fprintln(out, f(r)-f(l-1))
	}
}

func main() { CF2026D(os.Stdin, os.Stdout) }
