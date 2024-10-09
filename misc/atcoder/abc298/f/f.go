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

	var n, r, c, v, ans int
	type pair struct {
		r, c int
	}
	type cs struct {
		c, v int
	}
	grid, row, col := make(map[pair]int), make(map[int]int), make(map[int]int)

	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &r, &c, &v)
		grid[pair{r, c}] = v
		row[r] += v
		col[c] += v
	}
	csList := make([]cs, 0, len(col))
	for c, v := range col {
		csList = append(csList, cs{c, v})
	}
	sort.Slice(csList, func(i, j int) bool {
		return csList[i].v > csList[j].v
	})

	for r, rv := range row {
		for _, cv := range csList {
			val := rv + cv.v
			g, ok := grid[pair{r, cv.c}]
			if !ok {
				ans = max(ans, val)
				break
			}
			ans = max(ans, val-g)
		}
	}
	Fprint(out, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
