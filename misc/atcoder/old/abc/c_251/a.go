package c_251

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc251/tasks/abc251_e
//
//输入 n (2≤n≤3e5) 和长为 n 的数组 a (1≤a[i]≤1e9)，下标从 1 开始。
//有 n 只动物围成一圈，你可以花费 a[i] 喂食动物 i 和 i+1。
//特别地，你可以花费 a[n] 喂食动物 n 和 1。
//输出喂食所有动物需要的最小花费。

func AtCoderABC251E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}
	const inf int = 1e18
	pay, notPay := a[0], inf
	for _, v := range a[1:] {
		pay, notPay = min(pay, notPay)+v, pay
	}

	ans := min(pay, notPay)

	pay, notPay = a[n-1]+a[0], a[n-1]
	for _, v := range a[1 : n-1] {
		pay, notPay = min(pay, notPay)+v, pay
	}
	ans = min(ans, min(pay, notPay))
	fmt.Fprintln(w, ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
