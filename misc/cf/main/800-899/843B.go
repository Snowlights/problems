//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand"
	"os"
)

func CF843B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, st, tar, v, nxt int
	Fscan(in, &n, &st, &tar)
	cur := -1
	for _, i := range rand.Perm(n)[:min(999, n)] {
		Fprintln(out, "?", i+1)
		Fscan(in, &v, &nxt)
		if v <= tar && v > cur {
			cur, st = v, nxt
		}
	}
	for ; st > 0 && cur < tar; cur, st = v, nxt {
		Println("?", st)
		Fscan(in, &v, &nxt)
	}
	if cur < tar {
		cur = -1
	}
	Fprint(out, "! ", cur)

}

func main() { CF843B(os.Stdin, os.Stdout) }
