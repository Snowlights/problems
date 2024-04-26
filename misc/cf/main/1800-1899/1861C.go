//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1861C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, s := 0, ""
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		l := 0
		finishedSort := 1
		unfinishedSort := int(1e9)
		for _, b := range s {
			if b == '+' {
				l++
			} else if b == '-' {
				l--
				if l < unfinishedSort {
					unfinishedSort = 1e9
				}
				if l < finishedSort {
					finishedSort = max(l, 1)
				}
			} else if b == '0' {
				if l <= finishedSort {
					Fprintln(out, "NO")
					continue o
				}
				unfinishedSort = min(unfinishedSort, l)
			} else {
				if l >= unfinishedSort {
					Fprintln(out, "NO")
					continue o
				}
				finishedSort = max(l, 1)
			}
		}
		Fprintln(out, "YES")
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF1861C(os.Stdin, os.Stdout) }
