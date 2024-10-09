package keyence2020

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/keyence2020/tasks/keyence2020_d
//
//输入 n(≤18) 和两个长为 n 的数组 a b，元素范围在 [1,50]。
//a[i] 和 b[i] 表示第 i 张牌的正面数字和背面数字。
//初始所有牌正面朝上。
//每次操作你可以交换第 i 和 i+1 张牌的位置，同时翻转这两张牌。
//输出让看到的数字从左往右是递增（允许相等）所需要的最小操作次数。
//如果无法做到，输出 -1。

func AtCoderKeyence2020D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	a, b := make([]int, n), make([]int, n)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	for i := range b {
		fmt.Fscan(r, &b[i])
	}

}
