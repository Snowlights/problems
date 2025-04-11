//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF614A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var l, r, k int
	Fscan(in, &l, &r, &k)
	ok := false
	for powK := 1; ; powK *= k {
		if powK >= l {
			ok = true
			Fprint(out, " ", powK)
		}
		if powK > r/k {
			break
		}
	}
	if !ok {
		Fprint(out, -1)
	}

}

func main() { CF614A(os.Stdin, os.Stdout) }
