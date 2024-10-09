package c_147

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/arc147/tasks/arc147_b
//
//输入 n(≤400) 和一个 1~n 的排列 p，下标从 1 开始。
//你有两种操作：
//操作 "A i" 可以交换 p[i] 和 p[i+1]；
//操作 "B i" 可以交换 p[i] 和 p[i+2]。
//
//你需要让 p 变为递增的。
//最小化操作 A 的次数，同时总操作次数不能超过 1e5。
//输出总操作次数，以及具体的操作内容，详见样例。
//如果有多个方案，输出任意一种。
//保证存在这样的操作方案。

func AtCoderARC147B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	ps := make([]int, n)
	for i := range ps {
		fmt.Fscan(r, &ps[i])
	}

	ans := make([]string, 0, 1e5)
	f := func(ope string, i int) {
		s := fmt.Sprintf("%v %d", ope, i+1)
		ans = append(ans, s)
		j := i + 1
		if ope == "B" {
			j++
		}
		ps[i], ps[j] = ps[j], ps[i]
	}

	for i, p := range ps {
		if p%2 == (i+1)%2 {
			continue
		}
		for j := i + 1; j < n; j += 2 {
			if ps[j]%2 == p%2 {
				continue
			}
			for l := j; l >= i+3; l-- {
				k := l - 2
				f("B", k)
			}
			f("A", i)
			break
		}
	}

	f2 := func(_d int) {
		for {
			sorted := true
			for i := _d; i < n-2; i += 2 {
				j := i + 2
				a, b := ps[i], ps[j]
				if a > b {
					sorted = false
					f("B", i)
				}
			}

			if sorted {
				break
			}
		}
	}

	f2(0)
	f2(1)

	fmt.Fprintln(w, len(ans))
	for _, s := range ans {
		fmt.Fprintln(w, s)
	}

}
