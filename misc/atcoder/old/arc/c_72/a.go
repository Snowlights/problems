package c_72

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc059/tasks/arc072_a
//
//输入 n (2≤n≤1e5) 和一个长为 n 的数组 a (-1e9≤a[i]≤1e9)。
//每次操作你可以把一个 a[i] 加一或减一。
//如果要让 a 的所有相邻前缀和的乘积都小于 0，至少需要操作多少次？

func AtCoderARC72A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}

	f := func(to int) int {
		s, ans := 0, 0
		for _, v := range a {
			s += v

			if s*to <= 0 {
				ans += abs(s - to)
				s = to
			}

			to *= -1
		}

		return ans
	}

	fmt.Fprintln(w, min(f(1), f(-1)))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
