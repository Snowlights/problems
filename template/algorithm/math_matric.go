package algorithm

import "math"

// 矩阵快速幂

type matrix [][]int64

func newMatrix(n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int64, m)
	}
	return a
}

func newIdentityMatrix(n int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int64, n)
		a[i][i] = 1
	}
	return a
}

func (a matrix) mul(b matrix) matrix {
	const mod int64 = 1e9 + 7 // 998244353
	c := newMatrix(len(a), len(b[0]))
	for i, row := range a {
		for j := range b[0] {
			for k, v := range row {
				c[i][j] = (c[i][j] + v*b[k][j]) % mod // 注：此处不能化简
			}
			if c[i][j] < 0 {
				c[i][j] += mod
			}
		}
	}
	return c
}

func (a matrix) pow(n int64) matrix {
	res := newIdentityMatrix(len(a))
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}

// 比如 n*n 的国际象棋的马，从 (sx,sy) 走 k 步到 (tx,ty)，需要多少步
// 这里可以先 O(n^2) 预处理走一步的转移，构建矩阵 a
// 然后用一个 [1 * (n^2)] 的矩阵初始矩阵乘 a^k
// 得到一个 [1 * (n^2)] 的结果矩阵 res
// res[0][tx*n+ty] 就是答案
func (a matrix) solve(n, sx, sy, tx, ty int, k int64) int64 {
	b := matrix{make([]int64, n*n)}
	b[0][sx*n+sy] = 1
	res := b.mul(a.pow(k))
	return res[0][tx*n+ty]
}

// a(n) = p*a(n-1) + q*a(n-2)
// 注意：数列从 0 开始，若题目从 1 开始则输入的 n 为 n-1
// https://ac.nowcoder.com/acm/contest/9247/A
// m 项递推式，以及包含常数项的情况见《挑战》P201
// a(n) = a(n-1) + a(n-m) https://codeforces.com/problemset/problem/1117/D
func calcFibonacci(p, q, a0, a1, n int64) int64 {
	const mod int64 = 1e9 + 7 // 998244353
	//n--
	if n == 0 {
		return (a0%mod + mod) % mod
	}
	if n == 1 {
		return (a1%mod + mod) % mod
	}
	m := matrix{
		{p, q},
		{1, 0},
	}.pow(n - 1)
	return ((m[0][0]*a1+m[0][1]*a0)%mod + mod) % mod
	//return m[0][0]
}

//

func (a matrix) add(b matrix) matrix {
	c := newMatrix(len(a), len(a[0]))
	for i, row := range a {
		for j, aij := range row {
			c[i][j] = aij + b[i][j] // % mod
		}
	}
	return c
}

func (a matrix) sub(b matrix) matrix {
	c := newMatrix(len(a), len(a[0]))
	for i, row := range a {
		for j, aij := range row {
			c[i][j] = aij - b[i][j] // % mod) + mod) % mod
		}
	}
	return c
}

func (a matrix) swapRows(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a matrix) swapCols(i, j int) {
	for k := range a {
		a[k][i], a[k][j] = a[k][j], a[k][i]
	}
}

func (a matrix) mulRow(i int, k int64) {
	for j := range a[i] {
		a[i][j] *= k // % mod
	}
}

func (a matrix) trace() (sum int64) {
	for i, row := range a {
		sum += row[i]
	}
	return
}

// NxN 矩阵求逆
// 模板题 https://www.luogu.com.cn/problem/P4783
func (matrix) inv(A matrix) matrix {
	const mod int64 = 1e9 + 7
	pow := func(x int64) (res int64) {
		//x %= mod
		res = 1
		for n := mod - 2; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}

	// 增广一个单位矩阵
	n := len(A)
	m := 2 * n
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int64, m)
		for j := range a {
			a[i][j] = A[i][j] // or read
		}
		a[i][n+i] = 1
	}

	for i := range a {
		for j := i; j < n; j++ {
			if a[j][i] != 0 {
				a[i], a[j] = a[j], a[i]
				break
			}
		}
		if a[i][i] == 0 {
			// 矩阵不是满秩的
			return nil
		}
		inv := pow(a[i][i])
		for j := i; j < m; j++ {
			a[i][j] = a[i][j] * inv % mod
		}
		for j := range a {
			if j != i {
				inv := a[j][i]
				for k := i; k < m; k++ {
					a[j][k] = (a[j][k] - inv*a[i][k]%mod + mod) % mod
				}
			}
		}
	}

	// 结果保存在 a 右侧
	res := make(matrix, n)
	for i, row := range a {
		res[i] = row[n:]
	}
	return res
}

// 高斯消元 Gaussian elimination O(n^3)   列主元消去法
// 求解 Ax=B，A 为方阵，返回解（无解或有无穷多组解）
// https://en.wikipedia.org/wiki/Gaussian_elimination
// https://en.wikipedia.org/wiki/Pivot_element#Partial_and_complete_pivoting
// https://oi-wiki.org/math/gauss/
// 总结 https://cloud.tencent.com/developer/article/1087352
// https://cp-algorithms.com/linear_algebra/determinant-gauss.html
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/GaussianElimination.java.html
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/GaussJordanElimination.java.html
// 模板题 https://www.luogu.com.cn/problem/P3389 https://www.luogu.com.cn/problem/P2455
//
//	https://codeforces.com/problemset/problem/21/B
//
// 与 SCC 结合 https://www.luogu.com.cn/problem/P6030
func gaussJordanElimination(A matrix, B []int64) (sol []float64, infSol bool) {
	const eps = 1e-8
	n := len(A)
	// 构造增广矩阵 (or read)
	a := make([][]float64, n)
	for i, row := range A {
		a[i] = make([]float64, n+1)
		for j, v := range row {
			a[i][j] = float64(v)
		}
		a[i][n] = float64(B[i])
	}
	row := 0
	for col := 0; col < n; col++ {
		// 列主元消去法：减小误差，把正在处理的未知数系数的绝对值最大的式子换到第 row 行
		pivot := row
		for i := row; i < n; i++ {
			if math.Abs(a[i][col]) > math.Abs(a[pivot][col]) {
				pivot = i
			}
		}
		// 这一列全为 0，表明无解或有无穷多解，具体是哪一种需要消元完成后才知道
		if math.Abs(a[pivot][col]) < eps {
			continue
		}
		a[row], a[pivot] = a[pivot], a[row]
		// 初等行变换：把正在处理的未知数的系数变为 1
		for j := col + 1; j <= n; j++ {
			a[row][j] /= a[row][col]
		}
		// 消元，构造简化行梯阵式
		for i := range a {
			if i != row {
				// 用当前行对其余行进行消元：从第 i 个式子中消去第 col 个未知数
				for j := col + 1; j <= n; j++ {
					a[i][j] -= a[i][col] * a[row][j]
				}
			}
		}
		row++
	}
	if row < n {
		for _, r := range a[row:] {
			if math.Abs(r[n]) > eps {
				return nil, false
			}
		}
		return nil, true
	}
	res := make([]float64, n)
	for i, r := range a {
		res[i] = r[n]
	}
	return res, false
}

// EXTRA: 求行列式（对结果模 mod）
// https://en.wikipedia.org/wiki/Determinant
// 参考 https://www.luogu.com.cn/blog/Stormy-Rey/calculate-det
func (a matrix) determinant(mod int64) int64 {
	n := len(a)
	res, sign := int64(1), 1
	for i := range a {
		for j := i + 1; j < n; j++ {
			for a[i][i] != 0 {
				div := a[j][i] / a[i][i]
				for k := i; k < n; k++ {
					a[j][k] = (a[j][k] - a[i][k]*div%mod + mod) % mod
				}
				a[i], a[j], sign = a[j], a[i], -sign
			}
			a[i], a[j], sign = a[j], a[i], -sign
		}
	}
	for i, r := range a {
		res = res * r[i] % mod
	}
	res = (res*int64(sign) + mod) % mod
	return res
}
