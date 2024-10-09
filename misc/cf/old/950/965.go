package _50

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

// https://codeforces.com/problemset/problem/965/D
//
//输入 w(≤1e5) L(<w) 和一个长为 w-1 的数组 a (0≤a[i]≤1e4，下标从 1 开始)。
//有一条宽为 w 的河，青蛙们在位置 0 处，位置 i 处有 a[i] 个石头，位置 w 为河对岸。
//青蛙单次跳跃距离不超过 L。每个石头只能被一只青蛙用一次。
//输出最多有多少只青蛙能过河。
//
//相似题目：蓝桥杯2022年第十三届省赛真题-青蛙过河
//https://www.dotcpp.com/oj/problem2667.html

func CF965D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var W, l int
	fmt.Fscan(r, &W, &l)
	a := make([]int, W)
	for i := 1; i < W; i++ {
		fmt.Fscan(r, &a[i])
		a[i] += a[i-1]
	}
	ans := math.MaxInt32
	for i := 0; i < W-l; i++ {
		ans = min(ans, a[i+l]-a[i])
	}

	fmt.Fprintln(w, ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
