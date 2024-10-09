//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF870C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n == 2 || n == 5 || n == 7 || n == 11 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, n/4-n%2) // 1 和 3 输出 -1
		}
	}
}

func main() { CF870C(os.Stdin, os.Stdout) }
