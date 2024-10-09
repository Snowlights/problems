package c_140

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc140/tasks/abc140_e
//
//输入 n(≤1e5) 和一个 1~n 的排列 p。
//输出 p 中所有长度至少 为 2 的子数组的第二大元素的和。

// https://atcoder.jp/contests/abc140/submissions/36373341
//
//方法一：贡献法，对每个 p[i]，求上个、上上个、下个、下下个更大元素的位置，分别记作 L LL R RR。
//具体求法见双周赛第四题（链接见右）。
//那么 p[i] 对答案的贡献为 p[i] * ((L-LL)*(R-i) + (RR-R)*(i-L))。
//
//方法二：对于排列，其实不需要单调栈。
//用双向链表思考（实现时用的数组）：
//把 p 转换成双向链表，按元素值从小到大遍历 p[i]，
//那么方法一中的 4 个位置就是 p[i] 左边两个节点和右边两个节点。
//算完 p[i] 后把 p[i] 从链表中删掉。
//怎么用数组实现见代码。
//
//顺带一提，循环结束后，L[i] 和 R[i] 就是 p[i] 的上个/下个更大元素的位置。
//如果要求更小元素位置，倒序遍历即可。

func AtCoderABC140E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	p := make([]int, n)
	left, lleft, right, rright := make([]int, n), make([]int, n), make([]int, n), make([]int, n)
	for i := range p {
		fmt.Fscan(r, &p[i])
		left[i], lleft[i] = -1, -1
		right[i], rright[i] = n, n
	}

	s, t := []int{}, []int{}
	for i, v := range p {
		for len(t) > 0 && v > p[t[len(t)-1]] {
			rright[t[len(t)-1]] = i
			t = t[:len(t)-1]
		}
		poped := []int{}
		for len(s) > 0 && v > p[s[len(s)-1]] {
			poped = append(poped, s[len(s)-1])
			right[s[len(s)-1]] = i
			s = s[:len(s)-1]
		}
		s = append(s, i)
		for len(poped) > 0 {
			t = append(t, poped[len(poped)-1])
			poped = poped[:len(poped)-1]
		}
	}

	s, t = []int{}, []int{}
	for i := n - 1; i >= 0; i-- {
		v := p[i]
		for len(t) > 0 && v > p[t[len(t)-1]] {
			lleft[t[len(t)-1]] = i
			t = t[:len(t)-1]
		}
		poped := []int{}
		for len(s) > 0 && v > p[s[len(s)-1]] {
			poped = append(poped, s[len(s)-1])
			left[s[len(s)-1]] = i
			s = s[:len(s)-1]
		}
		s = append(s, i)
		for len(poped) > 0 {
			t = append(t, poped[len(poped)-1])
			poped = poped[:len(poped)-1]
		}
	}
	ans := 0
	for i, v := range p {
		ans += v * (left[i] - lleft[i]) * (right[i] - i)
		ans += v * (i - left[i]) * (rright[i] - right[i])
	}
	fmt.Fprintln(w, ans)
}
