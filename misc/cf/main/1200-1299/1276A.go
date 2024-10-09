//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1276A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var s string
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		ans := make([]int, 0, 20)
		for i := 0; i+2 < len(s); {
			if i+4 < len(s) && s[i:i+5] == "twone" {
				ans = append(ans, i+3)
				i += 4
			} else if i+2 < len(s) && (s[i:i+3] == "one" || s[i:i+3] == "two") {
				ans = append(ans, i+2)
				i += 2
			} else {
				i++
			}
		}
		Fprintln(out, len(ans))
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}

}

func main() { CF1276A(os.Stdin, os.Stdout) }
