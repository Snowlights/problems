package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, sum int
	Fscan(in, &n, &sum)
	a := make([]int, n)
	for i := range a[:n-1] {
		Fscan(in, &a[i])
	}
	check := func() bool {
		b := make([]int, n)
		copy(b, a)
		sort.Ints(b)
		s := 0
		for _, v := range b[1 : n-1] {
			s += v
		}
		return s >= sum
	}
	for i := 0; i <= 100; i++ {
		a[n-1] = i
		if check() {
			Fprintln(out, i)
			return
		}
	}

	Fprintln(out, -1)
}

func main() { run(os.Stdin, os.Stdout) }
