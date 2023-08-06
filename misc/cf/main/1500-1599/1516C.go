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

func CF1516C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a, total, minLb, idx := make([]int, n), 0, math.MaxInt32, 0
	for i := range a {
		Fscan(in, &a[i])
		total += a[i]
		lb := a[i] & -a[i]
		if lb < minLb {
			minLb = lb
			idx = i
		}
	}
	total /= minLb
	if total%2 > 0 {
		Fprintln(out, 0)
		return
	}

	f := make([]bool, total+1)
	f[0] = true
	for _, v := range a {
		v /= minLb
		for j := total; j >= v; j-- {
			f[j] = f[j] || f[j-v]
		}
	}
	if f[total/2] {
		Fprintln(out, 1)
		Fprintln(out, idx+1)
	} else {
		Fprintln(out, 0)
	}

}

func main() { CF1516C(os.Stdin, os.Stdout) }
