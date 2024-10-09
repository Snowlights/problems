package _50

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

// https://codeforces.com/problemset/problem/863/E
//
// 输入 n(≤2e5) 和 n 个闭区间，区间左右端点在 [0,1e9] 内，区间的编号从 1 开始。
// 你需要从中删除一个区间，使得删除前后，所有区间的并集不变（只考虑整数）。
// 如果不存在这样的区间，输出 -1；否则输出该区间的编号。
func CF863E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	type pair struct {
		i, l, r int
	}
	pList := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &pList[i].l, &pList[i].r)
		pList[i].i = i
	}
	sort.Slice(pList, func(i, j int) bool {
		return pList[i].l < pList[j].l || (pList[i].l == pList[j].l && pList[i].r > pList[j].r)
	})

	for i := 1; i < len(pList); i++ {
		if pList[i-1].r >= pList[i].r {
			fmt.Fprintln(w, pList[i].i+1)
			return
		}
	}

	for i := 1; i < len(pList)-1; i++ {
		if pList[i-1].r+1 >= pList[i+1].l {
			fmt.Fprintln(w, pList[i].i+1)
			return
		}
	}
	fmt.Fprintln(w, -1)
}
