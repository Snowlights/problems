package c_162

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc162/tasks/abc162_f
//本题是力扣 198. 打家劫舍 的变形题。
//输入 n (2≤n≤2e5) 和长为 n 的数组 a (-1e9≤a[i]≤1e9)。
//数组 a 就是 198 题的房屋存放金额。
//在本题中，你必须恰好偷 floor(n/2) 个房子。
//输出你能偷窃到的最高金额。
//
//你能做到 O(1) 空间吗？

// https://atcoder.jp/contests/abc162/submissions/36315908
//定义 f[i] 表示从前 i 个数中选 floor(i/2) 个数的最大收益。
//下面讨论的下标从 1 开始。
//f[0]=f[1]=0。
//对于 i≥2，
//如果选，那么 f[i] = f[i-2] + a[i]。
//如果不选：
//- 如果 i 为偶数，那么前面的选法是固定的，也就是所有奇数下标的和，记作 s。
//- 如果 i 为奇数，那么需要选 (i-1)/2 个数，在不选的情况下，相当于从前 i-1 个数中选 (i-1)/2 个数，这恰好是 f[i-1]。
//选或不选取 max。
//答案为 f[n]。

func AtCoderABC162F(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, s, v int64
	fmt.Fscan(r, &n, &s)
	f := make([]int64, n+1)
	for i := int64(2); i <= n; i++ {
		fmt.Fscan(r, &v)
		if i&1 == 0 {
			f[i] = max(s, f[i-2]+v)
		} else {
			s += v
			f[i] = max(f[i-1], f[i-2]+v)
		}
	}

	fmt.Fprintln(w, f[n])
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}
