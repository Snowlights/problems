package c_255

import (
	"bufio"
	"fmt"
	"os"
)

// https://atcoder.jp/contests/abc255/tasks/abc255_e
//
//输入 n (2≤n≤1e5) 和 m(≤10)，长为 n-1 的数组 s 和长为 m 的严格递增数组 x，元素值范围在 [-1e9,1e9]。
//数组 x 中的元素叫做幸运数。
//对于一个长为 n 的序列 a，如果所有相邻元素满足 a[i]+a[i+1]=s[i]，则称为一个好序列。
//输出好序列中最多能有多少个数是幸运数（重复数字也算，见样例）。

func AtCoderABC255E() {
	in := bufio.NewReader(os.Stdin)
	var n, m, sum, ans int
	fmt.Fscan(in, &n, &m)
	s := make([]int, n-1)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}
	cnt := make(map[int]int, n*m)
	x := make([]int, m)
	for i := range x {
		fmt.Fscan(in, &x[i])
		cnt[x[i]]++
	}
	for i, v := range s {
		if i&1 > 0 {
			sum -= v
			for _, w := range x {
				cnt[sum+w]++
			}
		} else {
			sum += v
			for _, w := range x {
				cnt[sum-w]++
			}
		}
	}
	for _, c := range cnt {
		if c > ans {
			ans = c
		}
	}
	fmt.Print(ans)
}
