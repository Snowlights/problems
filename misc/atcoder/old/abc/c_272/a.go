package c_272

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc272/tasks/abc272_e
//
//输入 n(≤2e5) m(≤2e5) 和长为 n 的整数数组 a (-1e9≤a[i]≤1e9)，下标从 1 开始。
//执行如下操作 m 次：
//对 1~n 的每个 i，把 a[i] += i。
//每次操作后，输出 mex(a)，即不在 a 中的最小非负整数。

// https://atcoder.jp/contests/abc272/submissions/37376965
//提示 1：由于 a 中只有 n 个数，因此在 [0,n-1] 范围之外的 a[i] 是没有意义的。
//提示 2：对于每个 i，有多少次操作，使得操作后 a[i] 仍然在 [0,n-1] 范围内？
//提示 3：根据调和级数，总共有 O(nlogn) 次这样的操作。
//提示 4：对每个 a[i]，把满足提示 2 的操作次数以及操作后的结果存下来，按照操作次数分组归类。
//提示 5：每个操作次数所对应的结果记录到一个哈希表中，暴力枚举计算 mex（提示 3）。
//代码实现时可以用时间戳来优化。

func AtCoderABC272E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, v int
	fmt.Fscan(r, &n, &m)
	b := make([][]int, m+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &v)
		l := 1
		if v < 0 {
			l = (-v + i - 1) / i
		}
		r := min(m, (n-1-v)/i)
		for j := l; j <= r; j++ {
			b[j] = append(b[j], v+j*i)
		}
	}
	time := make([]int, n+1)
	for i := 1; i <= m; i++ {
		for _, v := range b[i] {
			time[v] = i
		}
		mex := 0
		for time[mex] == i {
			mex++
		}
		fmt.Fprintln(w, mex)
	}

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
