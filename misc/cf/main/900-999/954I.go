//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF954I(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s, t []byte
	Fscan(in, &s, &t)

	m := len(t)
	ans := make([]int, len(s)-m+1)
	mp := ['g']int{}
	var dfs func(int, int)
	dfs = func(c, sz int) {
		if c == 'g' {
			pi := make([]int, m)
			j := 0
			for i := 1; i < m; i++ {
				v := mp[t[i]]
				for j > 0 && mp[t[j]] != v {
					j = pi[j-1]
				}
				if mp[t[j]] == v {
					j++
				}
				pi[i] = j
			}

			j = 0
			for i := range s {
				v := mp[s[i]]
				for j > 0 && mp[t[j]] != v {
					j = pi[j-1]
				}
				if mp[t[j]] == v {
					j++
				}
				if j == m {
					st := i - m + 1
					ans[st] = max(ans[st], sz)
					j = pi[j-1]
				}
			}
			return
		}
		mp[c] = sz
		dfs(c+1, sz+1)
		for j := range sz {
			mp[c] = j
			dfs(c+1, sz)
		}
	}
	dfs('a', 0)
	for _, v := range ans {
		Fprint(out, 6-v, " ")
	}

}

func main() { CF954I(os.Stdin, os.Stdout) }
