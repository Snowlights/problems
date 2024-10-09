package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const m int = 1e10
	var n int
	var s string
	Fscan(in, &n, &s)
	p := func(num int) {
		Fprintln(out, num)
	}
	if n == 1 {
		p(int(s[0]&1+1) * m)
	} else if n == 2 {
		if s == "00" {
			p(0)
		} else if s == "01" {
			p(m - 1)
		} else {
			p(m)
		}
	} else if strings.Count(s[:3], "0") != 1 || !strings.HasPrefix(s, s[n-n%3:]) || strings.Count(s[:n-n%3], s[:3]) != n/3 {
		p(0)
	} else if s[:3] == "110" {
		p(m - (n-1)/3)
	} else if s[:3] == "101" {
		p(m - n/3)
	} else {
		p(m - (n+1)/3)
	}

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() { run(os.Stdin, os.Stdout) }
