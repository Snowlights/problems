//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
)

func CF1811E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		s := []byte(strconv.FormatUint(uint64(v), 9))
		for i, v := range s {
			if v >= '4' {
				s[i]++
			}
		}
		Fprintln(out, string(s))
	}

}

func main() { CF1811E(os.Stdin, os.Stdout) }
