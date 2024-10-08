//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1292A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, cnt, r, c int
	ban := [2][1e5 + 2]int{}
	for Fscan(in, &n, &q); q > 0; q-- {
		Fscan(in, &r, &c)
		r--
		ban[r][c] ^= 1
		s := ban[r^1][c] + ban[r^1][c-1] + ban[r^1][c+1]
		if ban[r][c] > 0 {
			cnt += s
		} else {
			cnt -= s
		}
		if cnt > 0 || ban[0][1] > 0 || ban[1][n] > 0 {
			Fprintln(out, "No")
		} else {
			Fprintln(out, "Yes")
		}
	}

}

func main() { CF1292A(os.Stdin, os.Stdout) }
