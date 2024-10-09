package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, l, cur, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for r := range a {
		Fscan(in, &a[r])
		for l < r && cur&a[r] > 0 {
			cur ^= a[l]
			l++
		}
		ans += r - l + 1
		cur |= a[r]
	}

	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
