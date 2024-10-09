package c_100

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc102/tasks/arc100_b
//
//输入 n (4≤n≤1e5) 和一个长为 n 的数组 a (1≤a[i]≤1e9)。
//将 a 分割成 4 个非空连续子数组，计算这 4 个子数组的元素和。
//你需要最小化这 4 个和中的最大值与最小值的差。
//输出这个最小值。

func AtCoderARC100B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &s[i])
		s[i] += s[i-1]
	}

	ans := int(1e18)
	for left, mid, right := 1, 2, 3; mid < n-1; mid++ {
		midS := s[mid]
		for abs(s[left]*2-midS) > abs(s[left+1]*2-midS) {
			left++
		}
		for abs(s[right]*2-midS-s[n]) > abs(s[right+1]*2-midS-s[n]) {
			right++
		}
		a, b, c, d := s[left], midS-s[left], s[right]-midS, s[n]-s[right]
		ans = min(ans, max(max(max(a, b), c), d)-min(min(min(a, b), c), d))
	}
	fmt.Fprintln(w, ans)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
