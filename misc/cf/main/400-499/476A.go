//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF476A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	ans := (n + 1) / 2
	ans += (m - ans%m) % m
	if ans > n {
		ans = -1
	}
	Fprint(out, ans)

}

func main() { CF476A(os.Stdin, os.Stdout) }
