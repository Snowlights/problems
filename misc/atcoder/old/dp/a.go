package dp

import (
	"bufio"
	"fmt"
	"io"
)

// Find the number of integers between 1 and K (inclusive)
// satisfying the following condition,  modulo 10^9 +7:
//
//The sum of the digits in base ten is a multiple of D.

func AtCoderDPTaskS(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	const (
		mod = 1e9 + 7
	)

	var s string
	var d int
	fmt.Fscan(r, &s, &d)

	n := len(s)
	dp := make([][]int64, n)
	for i := range dp {
		dp[i] = make([]int64, d)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var f func(i, pre int, limit, num bool) int64
	f = func(i, pre int, limit, num bool) int64 {
		if i == n {
			if num && pre == 0 {
				return 1
			}
			return 0
		}

		res := int64(0)
		if !limit && num {
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
			up = s[i]
		}
		res = res + f(i+1, pre, limit && up == '0', num)%mod
		for j := byte('1'); j <= up; j++ {
			res = (res + f(i+1, (pre+int(j-'0'))%d, limit && j == up, true)) % mod
		}
		return res
	}

	fmt.Fprintln(w, f(0, 0, true, false))

}
