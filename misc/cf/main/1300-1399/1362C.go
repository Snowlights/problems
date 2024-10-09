//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1362C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		ans := 0
		for Fscan(in, &n); n > 0; n /= 2 {
			ans += n
		}
		Fprintln(out, ans)
	}

}

func main() { CF1362C(os.Stdin, os.Stdout) }
