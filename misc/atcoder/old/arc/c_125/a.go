package c_125

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/arc125/tasks/arc125_c
//
//输入 n k (k≤n≤2e5) 和一个长为 k 的严格递增数组 a，元素范围 [1,n]。
//输出一个 1~n 的排列 p，使得 a 是 p 的一个 LIS。
//如果有多个这样的 p，输出字典序最小的那个。
//注：LIS 指最长上升子序列。

// 提示 1：先思考（手玩）一个简单的情况，如果没有字典序最小的要求，要如何构造答案？
// 提示 2：继续手玩，让你构造出的排列的字典序尽量小。
// 提示 3：构造方案如下：
// 把不在 a 中的数字记到 miss 数组中。
// 双指针遍历 a 的前 k-1 个数以及 miss，
// 如果 miss[j] < a[i]，那么 miss[j] 就可以插在 a[i] 和 a[i+1] 之间。
// 最后把 a[k-1] 加到剩余的 miss 中，从大到小输出。
func AtCoderARC1325C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, k, x int
	fmt.Fscan(r, &n, &k)
	a := make([]int, k)
	miss := make([]int, 0, n-k)
	for i := range a {
		fmt.Fscan(r, &a[i])
		for x++; x < a[i]; x++ {
			miss = append(miss, x)
		}
	}
	for x++; x <= n; x++ {
		miss = append(miss, x)
	}

	j := 0
	for _, v := range a[:k-1] {
		fmt.Fprint(w, v, " ")
		if j < len(miss) && miss[j] < v {
			fmt.Fprint(w, miss[j], " ")
			j++
		}
	}
	i := len(miss) - 1
	for ; i >= j && miss[i] > a[k-1]; i-- {
		fmt.Fprint(w, miss[i], " ")
	}
	fmt.Fprint(w, a[k-1], " ")
	for ; i >= j; i-- {
		fmt.Fprint(w, miss[i], " ")
	}

}
