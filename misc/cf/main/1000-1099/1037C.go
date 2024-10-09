//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1037C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s, t []byte
	Fscan(in, &s, &s, &t)
	ans := 0
	preI := -2
	for i, b := range s {
		if b == t[i] {
			continue
		}
		if preI == i-1 && b != s[preI] {
			preI = -2
		} else {
			ans++
			preI = i
		}
	}
	Fprint(out, ans)

}

func main() { CF1037C(os.Stdin, os.Stdout) }
