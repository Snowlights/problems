package _600

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//
//输入 t (≤1e4) 表示 t 组数据，每组数据输入 n(≤2e5) 和一个长为
//n 的数组 a (0≤a[i]≤n)。所有数据的 n 之和不超过 2e5。
//每次操作，你可以把数组中的一个数加一。
//定义 mex(a) 表示不在 a 中的最小非负整数。
//定义 f(i) 表示使 mex(a) = i 的最小操作次数。如果无法做到，则 f(i) = -1。
//输出 n+1 个数：f(0), f(1), ..., f(n)。

// 5
//3
//0 1 3
//7
//0 1 2 3 4 3 2
//4
//3 0 0 0
//7
//4 6 2 3 5 0 5
//5
//4 0 1 0 4

// 1 1 0 -1
//1 1 2 2 1 0 2 6
//3 0 1 4 3
//1 0 -1 -1 -1 -1 -1 -1
//2 1 0 2 -1 -1

func CF1619E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T, n, v int
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &n)
		cnt := make([]int, n+1)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &v)
			cnt[v]++
		}
		fmt.Fprint(w, cnt[0], " ")
		s, ex, fill := 0, []int{}, int64(0)
		for i, c := range cnt[:n] {
			if s += c; s <= i {
				fmt.Fprint(w, strings.Repeat("-1 ", n-i))
				break
			}
			if c == 0 {
				fill += int64(i - ex[len(ex)-1])
				ex = ex[:len(ex)-1]
			}
			fmt.Fprint(w, fill+int64(cnt[i+1]), " ")
			for ; c > 1; c-- {
				ex = append(ex, i)
			}
		}
		fmt.Fprintln(w)
	}

}
