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

func CF845C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]struct{ l, r int }, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].l < a[j].l })

	end1, end2 := -1, -1
	for _, p := range a {
		if p.l > end1 {
			end1 = p.r
		} else if p.l > end2 {
			end2 = p.r
		} else {
			Fprint(out, "NO")
			return
		}
	}
	Fprint(out, "YES")
	
}

func main() { CF845C(os.Stdin, os.Stdout) }
