package _00

import (
	"bufio"
	"fmt"
	"io"
)

func CF543A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, b, mod, v int
	fmt.Fscan(r, &n, &m, &b, &mod)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, b+1)
	}
	// j行代码时k个bug的方案数
	f[0][0] = 1
	for ; n > 0; n-- {
		fmt.Fscan(r, &v)
		for j := 1; j <= m; j++ {
			for k := v; k <= b; k++ {
				f[j][k] = (f[j][k] + f[j-1][k-v]) % mod
			}
		}
	}
	ans := 0
	for _, v := range f[m] {
		ans = (ans + v) % mod
	}
	fmt.Fprintln(w, ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
