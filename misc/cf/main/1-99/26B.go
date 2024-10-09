//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF26B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	var ans, cnt int
	Fscan(in, &s)
	ans = len(s)
	for _, v := range s {
		switch v {
		case '(':
			cnt++
		case ')':
			if cnt > 0 {
				cnt--
			} else {
				ans--
			}
		}
	}
	Fprintln(out, ans-cnt)
}

func main() { CF26B(os.Stdin, os.Stdout) }
