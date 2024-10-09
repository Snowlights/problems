package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var a, b string
	Fscan(in, &a, &b)
	h := make(map[byte]int)
	for i := range a {
		h[a[i]]++
	}
	var cnt int
	for i := range b {
		if b[i] == '@' {
			cnt++
			continue
		}
		h[b[i]]--
	}
	isChar := func(b byte) bool {
		for _, v := range []byte{'a', 't', 'c', 'o', 'd', 'e', 'r'} {
			if v == b {
				return true
			}
		}
		return false
	}
	for k, v := range h {
		if v == 0 || k == '@' {
			continue
		}
		if !isChar(k) {
			Fprintln(out, "No")
			return
		}
		if v > 0 {
			if v > cnt {
				Fprintln(out, "No")
				return
			}
			cnt -= v
		} else if v < 0 {
			if abs(v) > h['@'] {
				Fprintln(out, "No")
				return
			}
			h['@'] += v
		}
	}
	if h['@'] == cnt {
		Fprintln(out, "Yes")
	} else {
		Fprintln(out, "No")
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() { run(os.Stdin, os.Stdout) }
