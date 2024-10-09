package c_133

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/arc133/tasks/arc133_c
//
//输入 n m k (≤2e5) 和长度分别为 n 和 m 的数组 a 和 b，元素范围 [0,k-1]。
//构造一个 n 行 m 列，元素范围在 [0,k-1] 的矩阵，
//使得第 i 行的元素和 % k = a[i]，第 j 列的元素和 % k = b[j]。
//你只需要输出这个矩阵的元素和的最大值。
//如果这个矩阵不存在，输出 -1。

func AtCoderARC133C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, k int
	fmt.Fscan(r, &n, &m, &k)

	a, b := make([]int, n), make([]int, m)
	sa, sb := 0, 0
	for i := range a {
		fmt.Fscan(r, &a[i])
		sa += ((k-1)*m - a[i]) % k
	}
	for i := range b {
		fmt.Fscan(r, &b[i])
		sb += ((k-1)*n - b[i]) % k
	}
	if sa%k != sb%k {
		fmt.Fprintln(w, -1)
	} else {
		fmt.Fprintln(w, n*m*(k-1)-max(sa, sb))
	}

}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
