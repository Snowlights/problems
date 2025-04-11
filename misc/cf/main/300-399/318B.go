//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF318B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(in, &s)
	ans, cnt := 0, 0
	for i := 5; i <= len(s); i++ {
		if s[i-5:i] == "heavy" {
			cnt++
		} else if s[i-5:i] == "metal" {
			ans += cnt
		}
	}
	Fprint(out, ans)

}

func main() { CF318B(os.Stdin, os.Stdout) }
