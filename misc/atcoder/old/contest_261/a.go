package contest_261

import (
	"bufio"
	"fmt"
	"io"
)

func AtCoder261A(in io.Reader, out io.Writer) {

	var l1, l2 int64
	var l3, l4 int64
	fmt.Fscan(in, &l1, &l2, &l3, &l4)

	l := max(l1, l3)
	r := min(l2, l4)

	if r-l < 0 {
		fmt.Fprintln(out, 0)
	} else {
		fmt.Fprintln(out, r-l)
	}
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func AtCoder261B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int64
	fmt.Fscan(r, &n)
	rec := make([]string, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &rec[i])
	}

	for i := int64(0); i < n; i++ {
		for j := int64(0); j < n; j++ {
			if i == j {
				continue
			}
			switch rec[i][j] {
			case 'W':
				if rec[j][i] != 'L' {
					fmt.Fprintln(w, "incorrect")
					return
				}
			case 'L':
				if rec[j][i] != 'W' {
					fmt.Fprintln(w, "incorrect")
					return
				}
			case 'D':
				if rec[j][i] != 'D' {
					fmt.Fprintln(w, "incorrect")
					return
				}
			}
		}
	}
	fmt.Fprintln(w, "correct")
}

func AtCoder261C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int64
	var s string
	fmt.Fscan(r, &n)
	h := make(map[string]int)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &s)
		v, ok := h[s]
		if ok && v > 0 {
			fmt.Fprintln(w, s+fmt.Sprintf("(%d)", v))
		} else {
			fmt.Fprintln(w, s)
		}
		h[s]++
	}
}

func AtCoder261D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int64
	fmt.Fscan(r, &n, &m)
	x, bonus := make([]int64, n), make(map[int64]int64)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &x[i])
	}
	for i := int64(0); i < m; i++ {
		var q, s int64
		fmt.Fscan(r, &q, &s)
		bonus[q] = s
	}

	dp := [5005][5005]int64{}
	dp[0][1] = x[0] + bonus[1]
	for i := int64(1); i < n; i++ {
		for j := int64(0); j <= i; j++ {
			dp[i][0] = max(dp[i][0], dp[i-1][j])
		}
		for j := int64(1); j <= i+1; j++ {
			dp[i][j] = bonus[j] + x[i] + dp[i-1][j-1]
		}
	}
	ans := int64(-1)
	for i := int64(0); i < 5005; i++ {
		for j := int64(0); j < 5005; j++ {
			ans = max(ans, dp[i][j])
		}
	}
	fmt.Fprintln(w, ans)
}

func AtCoder261E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, c, t, a int
	fmt.Fscan(r, &n, &c)
	m := 1<<30 - 1
	// s1初始化代表所有位全为1，  s0初始化表示所有位都为0
	s1, s0 := 1<<30-1, 0
	for i := int(0); i < n; i++ {
		fmt.Fscan(r, &t, &a)
		switch t {
		case 1:
			s1 &= a
			s0 &= a

		case 2:
			s1 |= a
			s0 |= a

		case 3:
			s1 ^= a
			s0 ^= a
		}
		// 相比于变化之前，不变的1  ||  变化的0，  ===》 操作之后的值

		c = (c & s1) | ((c ^ m) & s0)
		fmt.Fprintln(w, c)
	}

}
