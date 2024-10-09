package dp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// 311 如果一个十进制数能够被它的各位数字之和整除,则称这个数为“月之数”。
// 给定整数 L 和 R，你需要计算闭区间 [L,R] 中有多少个“月之数”

func ACWingContest313(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	clac := func(s string) int64 {
		n, m := len(s), 18*9+1
		dp, ans := make([][][]int64, n), int64(0)
		for target := 1; target < m; target++ {
			for i := range dp {
				dp[i] = make([][]int64, m)
				for j := range dp[i] {
					dp[i][j] = make([]int64, m)
					for k := range dp[i][j] {
						dp[i][j][k] = -1
					}
				}
			}

			var f func(i, sum, dsum int, limit bool) int64
			f = func(i, sum, dsum int, limit bool) int64 {
				if i == n {
					if sum == 0 && dsum == target {
						return 1
					}
					return 0
				}
				res := int64(0)
				if !limit {
					dv := &dp[i][sum][dsum]
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
				for j := 0; j <= up; j++ {
					res += f(i+1, (10*sum+j)%target, dsum+j, limit && j == up)
				}
				return res
			}
			ans += f(0, 0, 0, true)
		}
		return ans
	}

	var left, right int
	fmt.Fscan(r, &left, &right)
	fmt.Fprintln(w, clac(strconv.Itoa(right))-clac(strconv.Itoa(left-1)))
}
