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

	var n, ans int
	Fscan(in, &n)
	a, cnt := make([]int, n), make(map[int]int)
	for i := range a {
		Fscan(in, &a[i])
		cnt[a[i]]++
	}
	for _, v := range cnt {
		ans += (v - 1) * v / 2
	}

	for _, v := range a {
		cnt[v]--
		Fprintln(out, ans-cnt[v])
		cnt[v]++
	}

}

func main() { run(os.Stdin, os.Stdout) }
