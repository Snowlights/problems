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

	var n, ans int
	var s string
	Fscan(in, &n)
	ab, a, b := 0, 0, 0
	for i := 0; i < n; i++ {
		Fscan(in, &s)
		if s[len(s)-1] == 'A' && s[0] == 'B' {
			ab++
		} else if s[0] == 'B' {
			b++
		} else if s[len(s)-1] == 'A' {
			a++
		}
		ans += strings.Count(s, "AB")
	}

	if ab > 0 && a == 0 && b == 0 {
		ans--
	}
	ans += ab + min(a, b)
	Fprintln(out, ans)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
