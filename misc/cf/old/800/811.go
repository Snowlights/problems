package _00

import (
	"bufio"
	"fmt"
	"io"
)

// https://codeforces.com/problemset/problem/811/C
//
//输入 n(≤5000) 和一个长为 n 的数组 a (0≤a[i]≤5000)。
//你需要从 a 中选取若干个互不相交的子数组，且对于每个子数组 b，b 中的元素不会出现在 b 外面。
//对每个子数组去重，求异或和，然后把所有异或和相加，输出相加后的最大值。
//
//注：子数组是连续的。
//注：你不需要把每个 a[i] 都选上。

func CF811C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n+1)
	l := [5001]int{}
	rr := [5001]int{}
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &a[i])
		if l[a[i]] == 0 {
			l[a[i]] = i
		}
		rr[a[i]] = i
	}
	f := make([]int, n+1)
	vis := [5001]int{}
	for i := 1; i <= n; i++ {
		f[i] = f[i-1]
		for j, ll, xor := i, l[a[i]], 0; j > 0; j-- {
			v := a[j]
			if vis[v] < i {
				if rr[v] > j {
					break
				}
				vis[v] = i
				ll = min(ll, l[v])
				xor ^= v
			}
			if ll == j {
				f[i] = max(f[i], f[j-1]+xor)
			}
		}
	}
	fmt.Fprint(w, f[n])
}
