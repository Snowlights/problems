package c_140

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/arc140/tasks/arc140_c
//
//输入 n 和 x (x≤n≤2e5)。
//你需要构造一个 1~n 的排列 p，其中第一项为 x。
//p 可以生成一个长为 n-1 的数组 a，使得 a[i]=abs(p[i]-p[i+1])。
//通过构造合适的 p，最大化 a 的 LIS（最长严格递增子序列）的长度。
//输出 p。

func AtCoderARC136B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, x int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &x)
	x--
	f := func(n, i int) []int {
		l, r := i, i
		ans := []int{i}
		for j := 1; j < n; j++ {
			if (j == 1 && l >= n-1-r) || (j > 1 && ans[len(ans)-1] > ans[len(ans)-2]) {
				l--
				ans = append(ans, l)
			} else {
				r++
				ans = append(ans, r)
			}
		}
		return ans
	}
	var ans []int
	if abs(x-(n-1-x)) <= 1 {
		ans = f(n, x)
	} else {
		ans = f(n-1, (n-1)/2)
		for i, v := range ans {
			if v >= x {
				ans[i]++
			}
		}
		ans = append([]int{x}, ans...)
	}
	for _, v := range ans {
		fmt.Fprint(w, v+1, " ")
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
