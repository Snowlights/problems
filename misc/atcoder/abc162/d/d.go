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

	const (
		R = 'R'
		G = 'G'
		B = 'B'
	)
	// 先求R，G，B不相等时候的总个数
	// 减去R，G，B不相等且间距相等的个数即为答案
	var n, r, g, b, ans int
	var s string
	Fscan(in, &n, &s)
	for i := 0; i < n; i++ {
		switch s[i] {
		case R:
			r++
		case G:
			g++
		case B:
			b++
		}
		left, right := i-1, i+1
		for left >= 0 && right < n {
			if s[left] != s[i] && s[i] != s[right] && s[left] != s[right] {
				ans--
			}
			left--
			right++
		}
	}
	Fprintln(out, ans+r*g*b)
}

func main() { run(os.Stdin, os.Stdout) }
