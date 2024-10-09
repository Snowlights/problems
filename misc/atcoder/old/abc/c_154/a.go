package c_154

import (
	"bufio"
	"fmt"
	"io"
)

// e
func AtCoderABC154E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n string
	var k int
	fmt.Fscan(r, &n, &k)
	dp := make([][]int64, len(n))
	for i := range dp {
		dp[i] = make([]int64, len(n))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var f func(i, pre int, limit bool) int64
	f = func(i, pre int, limit bool) int64 {
		if i == len(n) {
			if pre == k {
				return 1
			}
			return 0
		}

		res := int64(0)
		if !limit {
			dv := &dp[i][pre]
			if *dv >= 0 {
				return *dv
			}
			defer func() {
				*dv = res
			}()
		}

		up := byte('9')
		if limit {
			up = n[i]
		}

		res += f(i+1, pre, limit && up == '0')
		for j := byte('1'); j <= up; j++ {
			res += f(i+1, pre+1, limit && j == up)
		}
		return res
	}

	fmt.Fprintln(w, f(0, 0, true))
}
