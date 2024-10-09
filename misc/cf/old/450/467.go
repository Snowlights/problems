package _50

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
)

func CF467A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int64
	var p, t int64
	ans := 0
	for fmt.Fscan(r, &n); n > 0; n-- {
		fmt.Fscan(r, &p, &t)
		if t-p >= 2 {
			ans++
		}
	}

	fmt.Fprintln(w, ans)
}

func CF467B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, k int
	fmt.Fscan(r, &n, &m, &k)
	h := make([]int, m)
	var val int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &h[i])
	}
	fmt.Fscan(r, &val)
	ans := 0
	for _, v := range h {
		if bits.OnesCount(uint(v^val)) <= k {
			ans++
		}
	}

	fmt.Fprintln(w, ans)
}

func CF467C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	s := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
		s[i] += s[i-1]
	}
	pre := make([]int64, n+1)
	f := make([]int64, n+1)
	for ; k > 0; k-- {
		for j := m; j <= n; j++ {
			f[j] = max(f[j-1], pre[j-m]+s[j]-s[j-m])
		}
		pre, f = f, pre
	}
	fmt.Fprint(out, pre[n])
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}
