package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
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

	var n, m, q, j, sumB, ans int
	n, m, q = r(), r(), r()
	a, b := make([]int, n), make([]int, m)
	for i := range a {
		a[i] = r()
	}
	for i := range b {
		b[i] = r()
	}
	sort.Ints(a)
	sort.Ints(b)

	for i := n - 1; i >= 0; i-- {
		for j < m && a[i]+b[j] <= q {
			sumB += b[j]
			j++
		}
		ans += j*a[i] + sumB + (m-j)*q
	}
	Fprintln(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
