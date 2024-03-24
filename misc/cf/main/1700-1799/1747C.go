//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

func CF1747C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, a, v, mm int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &a)
		mm = math.MaxInt
		for i := 1; i < n; i++ {
			Fscan(in, &v)
			mm = min(mm, v)
		}
		if a > mm {
			Fprintln(out, "Alice")
		} else {
			Fprintln(out, "Bob")
		}
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF1747C(os.Stdin, os.Stdout) }
