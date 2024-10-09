package c_250

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc250/tasks/abc250_e
//
//输入 n(≤2e5) 和两个长为 n 的数组 a 和 b，元素范围在 [1,1e9]。
//然后输入 q(≤2e5) 表示 q 个询问，每个询问输入两个数 x 和 y，范围在 [1,n]。
//对每个询问，设 a 的前 x 个元素去重得到集合 A，b 的前 y 个元素去重得到
// 集合 B，如果 A = B，输出 "Yes"，否则输出 "No"。

// https://atcoder.jp/contests/abc250/submissions/35814659
//
//为方便处理，首先把数组 a 转换成升序：
//例如，先把 31412 置换为 12324，然后求前缀最大值得到 12334（不影响答案的正确性）。
//
// 数组 b 也做同样的置换，
// 然后用 https://leetcode.cn/problems/max-chunks-to-make-sorted/
// 中提到的技巧，标记 b[i] 应该匹配到 a 中的哪个数字。

func AtCoderABC250E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, v, x, y, q int
	fmt.Fscan(r, &n)
	a, b := make([]int, n+1), make([]int, n+1)
	ah, bh := make(map[int]int), make(map[int]int)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &v)
		if ah[v] == 0 {
			ah[v] = len(ah) + 1
		}
		a[i] = max(a[i-1], ah[v])
	}
	mx := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &v)
		v = ah[v]
		if v == 0 {
			mx = 1e9
		}
		mx = max(mx, v)
		bh[v]++
		if mx == len(bh) {
			b[i] = mx
		}
	}
	for fmt.Fscan(r, &q); q > 0; q-- {
		fmt.Fscan(r, &x, &y)
		if a[x] == b[y] {
			fmt.Fprintln(w, "Yes")
		} else {
			fmt.Fprintln(w, "No")
		}
	}
}
