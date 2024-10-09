package _200

import (
	"bufio"
	"fmt"
	"io"
)

// https://codeforces.com/problemset/problem/1208/E
//
//输入 n(≤1e6) 和 w(≤1e6)，表示一个 n 行 w 列的表格。
//然后输入 n 个数组，第 i 个数组放在第 i 行中。
//输入的格式为：第一个数字表示数组的长度 m(≤w)，然后输入一个长为 m 的数组，元素范围 [-1e9,1e9]。
//保证所有数组的长度之和不超过 1e6。
//
//你可以滑动任意一行的整个数组。
//对表格的每一列，输出这一列的元素和的最大值。
//注意：每一列是单独计算的，不同列可以有不同的滑动方案。

//
//考虑每行能给答案带来多少贡献。
//
//如果 2m <= w，那么对于最左边的 m-1 列和最右边的 m-1 列，
//选择方案是有约束的（前缀最大值/后缀最大值），其余列可以取 max(max(a),0)，这可以用差分数组统计。
//
//如果 2m > w，那么问题变成长为 w-m+1 的滑动窗口最大值。
//这可以用单调队列实现，做法见力扣 239 题 https://leetcode.cn/problems/sliding-window-maximum/
//
//最后求差分数组的前缀和，即为答案。

func CF1208E(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var n, w, m int
	fmt.Fscan(_r, &n, &w)
	d := make([]int64, w+1)
	for ; n > 0; n-- {
		fmt.Fscan(_r, &m)
		a := make([]int64, m)
		for i := range a {
			fmt.Fscan(_r, &a[i])
		}
		if m*2 <= w {
			var pre, suf int64
			for i, v := range a {
				pre = max(pre, v)
				d[i] += pre
				d[i+1] -= pre
				suf = max(suf, a[m-1-i])
				d[w-1-i] += suf
				d[w-i] -= suf
			}
			d[m] += suf
			d[w-m] -= suf
		} else {
			sz := w - m + 1
			q := []int{}
			for i, v := range a {
				for len(q) > 0 && v >= a[q[len(q)-1]] {
					q = q[:len(q)-1]
				}
				q = append(q, i)
				if q[0] <= i-sz {
					q = q[1:]
				}
				mx := a[q[0]]
				if mx < 0 && i < w-m {
					mx = 0
				}
				d[i] += mx
				d[i+1] -= mx
			}
			suf := int64(0)
			for i := 0; i < w-m; i++ {
				suf = max(suf, a[m-1-i])
				d[w-1-i] += suf
				d[w-i] -= suf
			}
		}
	}
	s := int64(0)
	for _, v := range d[:w] {
		s += v
		fmt.Fprint(_w, s, " ")
	}

}
