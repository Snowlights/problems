package _400

import (
	"bufio"
	"fmt"
	"io"
)

// https://codeforces.com/problemset/problem/1407/D
//
//输入 n(≤3e5) 和一个长为 n 的数组 h (1≤h[i]≤1e9)。
//满足如下三个条件之一，就可以从 i 跳到 j (i<j)：
//1. i+1=j
//2. max(h[i+1],...,h[j-1]) < min(h[i],h[j])
//3. min(h[i+1],...,h[j-1]) > max(h[i],h[j])
//输出从 1 跳到 n 最少需要多少步。

func CF1407D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, v int
	fmt.Fscan(r, n)
	fmt.Fscan(in, &n, &v)
	type pair struct{ v, f int }
	s := []pair{{v, 0}}
	t := []pair{{v, 0}}

	for i := 1; i < n; i++ {
		fmt.Fscan(r, &v)
		f := int(2e9)
		for len(s) > 0 && v > s[len(s)-1].v {
			f = min(f, s[len(s)-1].f)
			s = s[:len(s)-1]
		}

		if len(s) > 0 {
			f = min(f, s[len(s)-1].f)
			if v == s[len(s)-1].v {
				s = s[:len(s)-1]
			}
		}

		for len(t) > 0 && v < t[len(t)-1].v {
			f = min(f, t[len(t)-1].f)
			t = t[:len(t)-1]
		}

		if len(t) > 0 {
			f = min(f, t[len(t)-1].f)
			if v == t[len(t)-1].v {
				t = t[:len(t)-1]
			}
		}

		s = append(s, pair{v, f + 1})
		t = append(t, pair{v, f + 1})

	}

	fmt.Fprintln(w, t[len(t)-1].f)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
