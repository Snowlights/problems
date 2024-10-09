//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF804B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	var b, ans int
	const mod int = 1e9+7
	Fscan(in, &s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'b' {
			b++
		} else {
			ans += b
			b = b * 2 % mod
		}
	}
	Fprintln(out, ans% mod)
}

func main() { CF804B(os.Stdin, os.Stdout) }
