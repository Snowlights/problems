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

	var n, ans int
	Fscan(in, &n)
	l, r := make([]int, n), make([]int, n)
	for i := range l {
		Fscan(in, &l[i], &r[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(l)))
	sort.Ints(r)

	for i := 0; l[i] > r[i]; i++ {
		ans += (n - 1 - 2*i) * (l[i] - r[i])
	}
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
