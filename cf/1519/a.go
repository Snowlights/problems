package _519

import (
	"bufio"
	"fmt"
	"io"
)

func CF1519D(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var n int64
	fmt.Fscan(_r, &n)
	a, b := make([]int64, n), make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(_r, &a[i])
	}
	sum := int64(0)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(_r, &b[i])
		sum += a[i] * b[i]
	}

	ans := sum
	for i := int64(0); i < n; i++ {
		cur := sum
		for l, r := i-1, i+1; l >= 0 && r < n; l-- {
			cur -= (a[l] - a[r]) * (b[l] - b[r])
			ans = max(ans, cur)
			r++
		}
		cur = sum
		for l, r := i-1, i; l >= 0 && r < n; l-- {
			cur -= (a[l] - a[r]) * (b[l] - b[r])
			ans = max(ans, cur)
			r++
		}
	}
	fmt.Fprintln(_w, ans)
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
