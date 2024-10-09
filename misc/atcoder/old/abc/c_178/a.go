package c_178

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

// https://atcoder.jp/contests/abc178/tasks/abc178_f
//
//输入 n(≤2e5) 和两个非降数组 a 和 b，元素范围在 [1,n]。
//如果可以重排 b，使得 a[i] != b[i] 对每个 i 都成立，则输出 Yes 和重排后的 b，否则输出 No。

func AtCoderABC185E(in io.Reader, out io.Writer) {
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

	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})

	c, left, right := -1, 0, 0
	for i := range a {
		if c == -1 && a[i] == b[i] {
			c = a[i]
			left = i
			right = i
		} else if c != -1 && a[i] == b[i] {
			right = i
		}
	}

	for i := 0; i < n; i++ {
		if c != a[i] && c != b[i] && left <= right {
			b[left], b[i] = b[i], b[left]
			left++
		}
	}

	if left <= right {
		fmt.Fprintln(w, "No")
	} else {
		fmt.Fprintln(w, "Yes")
		for i := 0; i < n; i++ {
			fmt.Fprint(w, b[i], " ")
		}
	}

}
