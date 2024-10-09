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

	var n, q, op, x, y int
	Fscan(in, &n, &q)
	fa := make([]int, n+q+1)
	newBox := make([]int, n+q+1)
	originBox := make([]int, n+q+1)
	for i := range fa {
		fa[i] = i
		newBox[i] = i
		originBox[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	newBallId := n + 1
	newBoxId := n + q
	for ; q > 0; q-- {
		Fscan(in, &op, &x)
		if op == 1 {
			Fscan(in, &y)
			fa[newBox[y]] = newBox[x]
			newBox[y] = newBoxId
			originBox[newBoxId] = y
			newBoxId--
		} else if op == 2 {
			fa[newBallId] = newBox[x]
			newBallId++
		} else {
			Fprintln(out, originBox[find(x)])
		}
	}

}

func main() { run(os.Stdin, os.Stdout) }
