//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2033D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		ans, s := 0, 0
		has := map[int]bool{0: true}
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			s += v
			if has[s] {
				ans++
				has = map[int]bool{}
			}
			has[s] = true
		}
		Fprintln(out, ans)
	}

}

func main() { CF2033D(os.Stdin, os.Stdout) }
