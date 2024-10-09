//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF13E(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	r := func() (x int) {
		in.Scan()
		data := in.Bytes()
		if data[0] == '-' {
			for _, b := range data[1:] {
				x = x*10 + int(b&15)
			}
			return -x
		}
		for _, b := range data {
			x = x*10 + int(b&15)
		}
		return
	}

	var n, q, op, p, v, last int
	n, q = r(), r()
	a := make([]int, n)
	for i := range a {
		a[i] = r()
	}
	const B = 316
	type pair struct{ last, jump int }
	data := make([]pair, n)
	f := func(i int) {
		if i+a[i] >= n || i/B != (i+a[i])/B {
			data[i] = pair{i, 0}
		} else {
			data[i] = data[i+a[i]]
			data[i].jump++
		}
	}
	for i := n - 1; i >= 0; i-- {
		f(i)
	}

	for ; q > 0; q-- {
		op, p = r(), r()
		p--
		if op == 0 {
			v = r()
			a[p] = v
			for i := p; i >= p-p%B; i-- {
				f(i)
			}
		} else {
			jump := 0
			for ; p < n; p = last + a[last] {
				jump += data[p].jump + 1
				last = data[p].last
			}
			Fprintln(out, last+1, jump)
		}
	}

}

func main() { CF13E(os.Stdin, os.Stdout) }
