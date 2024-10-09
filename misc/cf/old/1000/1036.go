package _000

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// 定义一个数字是“好数”，当且仅当它的十进制表示下有不超过3个数字1∼9
// 举个例子：4,200000,102034,200000,10203是“好数”，
// 然而4231,102306,72774200004231,102306,7277420000不是
// 给定[l,r]，问有多少个x使得 l <= x <= r，且x是“好数”
// 一共有T(1 <= T <= 10^{4} )组数据，对于每次的询问，输出一行一个整数表示答案
// 1 <= l <= r <= 10^18
var dp1, dp2 [][]int64
var T, m, n int
var ll, rr int64
var l, r string

func CF1036C(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	fmt.Fscan(_r, &T)
	for ; T > 0; T-- {
		fmt.Fscan(_r, &ll, &rr)
		ll--
		l, r = strconv.FormatInt(ll, 10), strconv.FormatInt(rr, 10)
		m, n = len(l), len(r)
		dp1, dp2 = make([][]int64, m), make([][]int64, n)
		for i := range dp1 {
			dp1[i] = make([]int64, m)
			for j := range dp1[i] {
				dp1[i][j] = -1
			}
		}
		for i := range dp2 {
			dp2[i] = make([]int64, n)
			for j := range dp2[i] {
				dp2[i][j] = -1
			}
		}
		fmt.Fprintln(_w, fr(0, 0, true)-fl(0, 0, true))
	}

}

func fl(i, pre int, limit bool) int64 {
	if i == m {
		if pre <= 3 {
			return 1
		}
		return 0
	}

	res := int64(0)
	if !limit {
		dv := &dp1[i][pre]
		if *dv >= 0 {
			return *dv
		}
		defer func() {
			*dv = res
		}()
	}

	up := byte('9')
	if limit {
		up = l[i]
	}

	res += fl(i+1, pre, limit && up == '0')
	for j := byte('1'); j <= up; j++ {
		res += fl(i+1, pre+1, limit && j == up)
	}
	return res
}

func fr(i, pre int, limit bool) int64 {
	if i == n {
		if pre <= 3 {
			return 1
		}
		return 0
	}

	res := int64(0)
	if !limit {
		dv := &dp2[i][pre]
		if *dv >= 0 {
			return *dv
		}
		defer func() {
			*dv = res
		}()
	}

	up := byte('9')
	if limit {
		up = r[i]
	}

	res += fr(i+1, pre, limit && up == '0')
	for j := byte('1'); j <= up; j++ {
		res += fr(i+1, pre+1, limit && j == up)
	}
	return res
}
