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

	var n, m int
	Fscan(in, &n, &m)
	c, cp, p := make([]string, n), make([]string, m), make([]int, m+1)
	for i := range c {
		Fscan(in, &c[i])
	}
	for i := range cp {
		Fscan(in, &cp[i])
	}
	for i := range p {
		Fscan(in, &p[i])
	}
	pm := make(map[string]int)
	for i, v := range cp {
		pm[v] = p[i+1]
	}
	ans := 0
	for _, v := range c {
		val, ok := pm[v]
		if ok {
			ans += val
		} else {
			ans += p[0]
		}
	}

	Fprintln(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
