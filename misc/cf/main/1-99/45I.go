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

func CF45I(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, zero int
	Fscan(in, &n)
	neg, pos := make([]int, 0), make([]int, 0)
	ans := make([]int, 0, n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		if v < 0 {
			neg = append(neg, v)
		} else if v > 0 {
			pos = append(pos, v)
		} else {
			zero++
		}
	}
	if n == 1 {
		Fprintln(out, v)
		return
	}
	sort.Ints(neg)
	sort.Ints(pos)
	for i := 0; i+1 < len(neg); i += 2 {
		ans = append(ans, neg[i], neg[i+1])
	}
	ans = append(ans, pos...)
	if len(ans) == 0 {
		ans = append(ans, 0)
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

func main() { CF45I(os.Stdin, os.Stdout) }
