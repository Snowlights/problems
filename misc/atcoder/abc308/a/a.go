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
		No  = "No"
		Yes = "Yes"
	)
	ans := Yes
	var pre, v int
	for i := 0; i < 8; i++ {
		Fscan(in, &v)
		if v < 100 || v > 675 || v%25 > 0 || v < pre {
			ans = No
		}
		pre = v
	}

	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
