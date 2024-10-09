package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s []byte
	var k int
	Fscan(in, &s, &k)
	for i := range s {
		switch s[i] {
		case '1':
			k -= 1 << (len(s) - 1 - i)
		}
	}
	if k < 0 {
		Fprintln(out, -1)
		return
	}
	for i := range s {
		switch s[i] {
		case '?':
			if k >= 1<<(len(s)-1-i) {
				k -= 1 << (len(s) - 1 - i)
				s[i] = '1'
			} else {
				s[i] = '0'
			}
		}
	}
	ans, _ := strconv.ParseInt(string(s), 2, 64)
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
