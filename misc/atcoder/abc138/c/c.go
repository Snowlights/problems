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

	var n int
	var ans float64
	Fscan(in, &n)
	a := make([]float64, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	sort.Float64s(a)
	ans = a[0]
	for i := 1; i < n; i++ {
		ans = (ans + a[i]) / 2
	}
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
