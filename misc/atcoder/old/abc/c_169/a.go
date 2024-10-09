package c_169

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

// https://atcoder.jp/contests/abc169/tasks/abc169_e
//
//输入 n (2≤n≤2e5) 和 n 个区间的左右端点，区间范围在 [1,1e9]。
//每个区间内选一个整数，然后计算这 n 个整数的中位数。你能得到多少个不同的中位数？
//注：偶数长度的中位数是中间两个数的平均值（没有下取整）。

//上下界思想，计算中位数的最小值和最大值，然后范围内的都可以取到。
//由于偶数长度的中位数会存在小数点后为 0.5 的情况，所以要做一个 *2 的处理。具体技巧见代码

func AtCoderABC169E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i], &b[i])
	}
	sort.Ints(a)
	sort.Ints(b)

	if n&1 == 1 {
		aMid, bMid := a[n/2], b[n/2]
		fmt.Fprintln(w, bMid-aMid+1)
	} else {
		aMid, bMid := a[n/2]+a[n/2-1], b[n/2]+b[n/2-1]
		fmt.Fprintln(w, bMid-aMid+1)
	}
}
