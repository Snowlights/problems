//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF920C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var s string
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &s)
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			continue
		}
		st := i
		for i < len(s) && s[i] == '1' {
			i++
		}
		sort.Ints(a[st : i+1])
	}
	if sort.IntsAreSorted(a) {
		Fprintln(out, "YES")
	} else {
		Fprintln(out, "NO")
	}
}

func main() { CF920C(os.Stdin, os.Stdout) }
