package _100

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func CF1168A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var m, n int
	fmt.Fscan(r, &n, &m)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	ans := sort.Search(m, func(lim int) bool {
		mx := 0
		for _, v := range a {
			// (mx-v+m)%m 表示把 v 改成 mx，需要的操作次数
			if (mx-v+m)%m > lim { // 无法修改成 mx
				if v < mx { // 无法保证单调非降
					return false
				}
				mx = v // 只能 v 不变了
			}
		}
		return true
	})

	fmt.Fprintln(w, ans)
}
