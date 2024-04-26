//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF91A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s, t string
	Fscan(in, &s, &t)

	n := len(s)
	// build nxt
	// nxt[i][j] 表示在 i 右侧的字符 j 的最近位置
	pos := [26]int{}
	for i := range pos {
		pos[i] = n
	}
	nxt := make([][26]int, n)
	for k := 0; k < 2; k++ {
		for i := n - 1; i >= 0; i-- {
			nxt[i] = pos
			pos[s[i]-'a'] = i
		}
	}

	ans := 1
	i, j := 0, 0
	if t[0] == s[0] {
		j = 1
	}
	for ; j < len(t); j++ {
		next := nxt[i][t[j]-'a']
		if next == n {
			Fprintln(out, -1)
			return
		}
		if next <= i {
			ans++
		}
		i = next
	}
	Fprintln(out, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF91A(os.Stdin, os.Stdout) }
