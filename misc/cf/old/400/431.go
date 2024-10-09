package _00

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// 1<=n<=10^18
// 求满足n+1~2n之间恰有m个数二进制表示中有k个1的n，输出任意一个解即可

func CF431D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	search := func(l, r int64, f func(int64) bool) int64 {
		for l < r {
			m := (l + r) >> 1
			if f(m) {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}

	var m int64
	var k int
	fmt.Fscan(r, &m, &k)
	cal := func(s string) int64 {
		n := len(s)
		dp := make([][]int64, n)
		for i := range dp {
			dp[i] = make([]int64, n)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		var f func(i, pre int, limit bool) int64
		f = func(i, pre int, limit bool) int64 {
			if pre > k {
				return 0
			}
			if i == n {
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

			up := byte('1')
			if limit {
				up = s[i]
			}
			for j := byte('0'); j <= up; j++ {
				c := pre
				if j == '1' {
					c++
				}
				res += f(i+1, c, limit && j == up)
			}

			return res
		}
		return f(0, 0, true)
	}

	fmt.Fprintln(w, search(1, int64(1e18+1), func(i int64) bool {
		return cal(strconv.FormatInt(int64(2*i), 2))-cal(strconv.FormatInt(int64(i), 2)) >= m
	}))

}
