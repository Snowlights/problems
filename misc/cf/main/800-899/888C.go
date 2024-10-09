//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF888C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(in, &s)
	pos := make(map[int][]int)
	for i, v := range s {
		pos[int(v-'a')] = append(pos[int(v-'a')], i)
	}
	n := len(s)
	ans := n
	for _, p := range pos {
		p = append(append([]int{-1}, p...), n)
		cnt := 0
		for i := 1; i < len(p); i++ {
			cnt = max(cnt, p[i]-p[i-1])
		}
		ans = min(ans, cnt)
	}

	Fprint(out, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF888C(os.Stdin, os.Stdout) }
