package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, i, j int
	Fscan(in, &n)
	m, b := make(map[int][]int), make(map[int][]int)
	card := make(map[int]map[int]bool)
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &op)
		switch op {
		case 1:
			Fscan(in, &i, &j)
			c, ok := card[i]
			if !ok {
				c = make(map[int]bool)
				card[i] = c
			}
			if !c[j] {
				b[i] = append(b[i], j)
				c[j] = true
			}

			m[j] = append(m[j], i)
		case 2:
			Fscan(in, &i)
			sort.Slice(m[i], func(a, b int) bool {
				return m[i][a] < m[i][b]
			})
			for _, v := range m[i] {
				Fprint(out, v, " ")
			}
			Fprintln(out)
		case 3:
			Fscan(in, &i)
			sort.Slice(b[i], func(c, d int) bool {
				return b[i][c] < b[i][d]
			})
			for _, v := range b[i] {
				Fprint(out, v, " ")
			}
			Fprintln(out)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
