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

	var n, k int
	Fscan(in, &n, &k)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Strings(a[:k])
	for _, v := range a[:k] {
		Fprintln(out, v)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
