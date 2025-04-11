//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF937B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var p, y int
	Fscan(in, &p, &y)
o:
	for x := y; x > p; x-- {
		for i := 2; i <= p && i*i <= x; i++ {
			if x%i == 0 {
				continue o
			}
		}
		Fprint(out, x)
		return
	}
	Fprint(out, -1)

}

func main() { CF937B(os.Stdin, os.Stdout) }
