package c_237

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc237/tasks/abc237_f
//
//输入 n (3≤n≤1000) 和 m (3≤m≤10)。
//输出有多少个满足如下条件的数组：
//1. 数组长度为 n；
//2. 数组元素范围在 [1,m]；
//3. 数组的 LIS 的长度恰好等于 3。
//对答案模 998244353。
//
//注：LIS 指最长严格上升子序列。

//回忆下 LIS 的 O(nlogn) 做法，在那个做法中，我们需要维护一个有序数组 arr。
//定义 f[i][x][y][z] 表示长为 i 且 arr=[x,y,z] 的符合题目要求的数组个数。
//计算 f[i][x][y][z] 时，枚举第 i 个数是多少，按照 LIS 的 O(nlogn)
// 做法转移到对应的 x/y/z 上。
//代码实现时，可以初始化 f[0][m+1][m+1][m+1] = 1，表示 arr=[inf,inf,inf] 的初始状态。
//答案为 sum(f[n][x][y][z])，1≤x<y<z≤m。
//
//为了节省内存，可以将元素值改为从 0 开始。

func AtCoderABC237F(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	const mod = 998244353
	var n, m int
	fmt.Fscan(r, &n, &m)
	f := make([][11][11][11]int, n+1)
	f[0][m][m][m] = 1
	for i := 1; i <= n; i++ {
		for x := 0; x <= m; x++ {
			for y := x; y <= m; y++ {
				for z := y; z <= m; z++ {
					for j := 0; j < m; j++ {
						val := f[i-1][x][y][z]
						if j <= x {
							f[i][j][y][z] = (f[i][j][y][z] + val) % mod
						} else if j <= y {
							f[i][x][j][z] = (f[i][x][j][z] + val) % mod
						} else if j <= z {
							f[i][x][y][j] = (f[i][x][y][j] + val) % mod
						}
					}
				}
			}
		}
	}

	ans := 0
	for x := 0; x < m; x++ {
		for y := x + 1; y < m; y++ {
			for z := y + 1; z < m; z++ {
				ans += f[n][x][y][z]
				ans %= mod
			}
		}
	}
	fmt.Fprintln(w, ans%mod)
}
