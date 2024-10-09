package p2600

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// p2657
func P4127(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var a, b int64
	fmt.Fscan(r, &a, &b)

	clac := func(num int64) int64 {
		s := strconv.FormatInt(num, 10)
		n := len(s)
		dp := make([][]int64, n)
		for i := range dp {
			dp[i] = make([]int64, 10)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}

		var f func(i, pre int, limit, num bool) int64
		f = func(i, pre int, limit, num bool) int64 {
			if i == n {
				if num {
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

			up := 9
			if limit {
				up = int(s[i] - '0')
			}
			for d := 0; d <= up; d++ {
				if !num || abs(d-pre) > 1 {
					res += f(i+1, d, limit && d == up, num || d > 0)
				}
			}
			return res
		}
		return f(0, 0, true, false)
	}

	fmt.Fprintln(w, clac(b)-clac(a-1))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
