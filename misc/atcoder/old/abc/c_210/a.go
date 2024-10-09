package c_210

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// https://atcoder.jp/contests/abc210/tasks/abc210_d
//
// 输入 n m (2≤n,m≤1000) c(≤1e9) 和一个 n 行 m 列的矩阵 a，元素范围 [1,1e9]。
// 对于两个不同位置 (i,j) 和 (i',j')，输出 a[i][j] + a[i'][j'] + c*(|i-i'|+|j-j'|) 的最小值。
func AtCoderABC210D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, c int
	fmt.Fscan(r, &n, &m, &c)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			fmt.Fscan(r, &a[i][j])
		}
	}

	pre, suf, ans := make([]int, m+1), make([]int, m+1), int(1e18)
	for i := range pre {
		pre[i] = 1e18
		suf[i] = 1e18
	}
	for i := range a {
		for j, v := range a[i] {
			ans = min(ans, v+c*(i+j)+min(pre[j+1], pre[j]))
			pre[j+1] = min(min(pre[j+1], pre[j]), v-c*(i+j))
		}
		for j := m - 1; j >= 0; j-- {
			ans = min(ans, a[i][j]+c*(i-j)+min(suf[j+1], suf[j]))
			suf[j] = min(min(suf[j+1], suf[j]), a[i][j]-c*(i-j))
		}
	}
	fmt.Fprintln(w, ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// https://atcoder.jp/contests/abc210/tasks/abc210_e
//
//输入 n(≤1e9) 和 m(≤1e5)，表示一个 n 个点,0条边的图(节点编号从 0 开始),以及 m 个操作。
//每个操作输入两个数 a(1≤a≤n-1) 和 c(≤1e9)，
//表示你可以花费 c，任意选择一个 [0,n-1] 内的数字 x，把 x 和 (x+a)%n 连边。
//这 m 个操作，每个都可以执行任意多次，可以按任意顺序执行。
//输出让图连通需要的最小花费。
//如果无法做到，输出 -1。

func AtCoderABC210E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	type op struct {
		a, c int
	}
	var n, m, ans int
	fmt.Fscan(r, &n, &m)
	opList := make([]op, m)
	for i := range opList {
		fmt.Fscan(r, &opList[i].a, &opList[i].c)
	}

	sort.Slice(opList, func(i, j int) bool {
		return opList[i].c < opList[j].c
	})

	for _, op := range opList {
		g := gcd(n, op.a)
		ans += (n - g) * op.c
		n = g
	}

	if n > 1 {
		ans = -1
	}
	fmt.Fprintln(w, ans)

}

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func findMaxForm(strs []string, m, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		zeros := strings.Count(s, "0")
		ones := len(s) - zeros
		for j := m; j >= zeros; j-- {
			for k := n; k >= ones; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeros][k-ones]+1)
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
