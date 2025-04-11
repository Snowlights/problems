#### 题目

<p>给你一个字符串 <code>s</code> 和一个整数 <code>t</code>，表示要执行的<strong> 转换 </strong>次数。每次 <strong>转换 </strong>需要根据以下规则替换字符串 <code>s</code> 中的每个字符：</p>

<ul>
	<li>如果字符是 <code>'z'</code>，则将其替换为字符串 <code>"ab"</code>。</li>
	<li>否则，将其替换为字母表中的<strong>下一个</strong>字符。例如，<code>'a'</code> 替换为 <code>'b'</code>，<code>'b'</code> 替换为 <code>'c'</code>，依此类推。</li>
</ul>

<p>返回<strong> 恰好 </strong>执行 <code>t</code> 次转换后得到的字符串的 <strong>长度</strong>。</p>

<p>由于答案可能非常大，返回其对 <code>10<sup>9</sup> + 7</code> 取余的结果。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "abcyy", t = 2</span></p>

<p><strong>输出：</strong> <span class="example-io">7</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li><strong>第一次转换 (t = 1)</strong>
	<ul>
		<li><code>'a'</code> 变为 <code>'b'</code></li>
		<li><code>'b'</code> 变为 <code>'c'</code></li>
		<li><code>'c'</code> 变为 <code>'d'</code></li>
		<li><code>'y'</code> 变为 <code>'z'</code></li>
		<li><code>'y'</code> 变为 <code>'z'</code></li>
		<li>第一次转换后的字符串为：<code>"bcdzz"</code></li>
	</ul>
	</li>
	<li><strong>第二次转换 (t = 2)</strong>
	<ul>
		<li><code>'b'</code> 变为 <code>'c'</code></li>
		<li><code>'c'</code> 变为 <code>'d'</code></li>
		<li><code>'d'</code> 变为 <code>'e'</code></li>
		<li><code>'z'</code> 变为 <code>"ab"</code></li>
		<li><code>'z'</code> 变为 <code>"ab"</code></li>
		<li>第二次转换后的字符串为：<code>"cdeabab"</code></li>
	</ul>
	</li>
	<li><strong>最终字符串长度</strong>：字符串为 <code>"cdeabab"</code>，长度为 7 个字符。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "azbk", t = 1</span></p>

<p><strong>输出：</strong> <span class="example-io">5</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li><strong>第一次转换 (t = 1)</strong>
	<ul>
		<li><code>'a'</code> 变为 <code>'b'</code></li>
		<li><code>'z'</code> 变为 <code>"ab"</code></li>
		<li><code>'b'</code> 变为 <code>'c'</code></li>
		<li><code>'k'</code> 变为 <code>'l'</code></li>
		<li>第一次转换后的字符串为：<code>"babcl"</code></li>
	</ul>
	</li>
	<li><strong>最终字符串长度</strong>：字符串为 <code>"babcl"</code>，长度为 5 个字符。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>s</code> 仅由小写英文字母组成。</li>
	<li><code>1 &lt;= t &lt;= 10<sup>5</sup></code></li>
</ul>

#### 思路

先求出单个字母 $\texttt{a},\texttt{b},\cdots,\texttt{z}$ 替换 $t$ 次后的长度。

## 寻找子问题

例如字母 $\texttt{a}$ 替换一次变成 $\texttt{b}$ 和 $\texttt{c}$，问题变成计算 $\texttt{b}$ 替换 $t-1$ 次后的长度，$\texttt{c}$ 替换 $t-1$ 次后的长度。二者之和即为 $\texttt{a}$ 替换 $t$ 次后的长度。

## 状态定义和状态转移方程

据此，定义 $f[i][j]$ 表示字母 $j$ 替换 $i$ 次后的长度。

上面的例子，就是 $f[i][0] = f[i-1][1] + f[i-1][2]$。

一般地，设 $c=\textit{nums}[i]$，我们有

$$
f[i][j] = \sum_{j=i+1}^{i+c} f[i-1][j\bmod 26]
$$

初始值 $f[0][j] = 1$。

答案为 $\sum\limits_{j=0}^{25} f[t][j]\cdot \textit{cnt}[j]$。其中 $\textit{cnt}[j]$ 为 $s$ 中的字母 $j$ 的出现次数。

直接计算这个 DP 的话，时间复杂度为 $\mathcal{O}(t|\Sigma|)$，这可以解决周赛第二题。对于本题，还需继续优化。

## 矩阵快速幂优化

以示例 1 为例（也相当于周赛第二题），其中 $\textit{nums}=[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,2]$。

我们有

$$
\begin{aligned}
f[i][0] &= f[i-1][1]     \\
f[i][1] &= f[i-1][2]     \\
f[i][2] &= f[i-1][3]     \\
\vdots \\
f[i][23] &= f[i-1][24]     \\
f[i][24] &= f[i-1][25]     \\
f[i][25] &= f[i-1][0] + f[i-1][1]     \\
\end{aligned}
$$

用矩阵乘法表示，即

$$
\begin{bmatrix}
f[i][0] \\
f[i][1] \\
f[i][2] \\
\vdots \\
f[i][23] \\
f[i][24] \\
f[i][25] \\
\end{bmatrix}
= \begin{bmatrix}
0 & 1 & 0 & 0 & \cdots & 0 & 0 \\
0 & 0 & 1 & 0 & \cdots & 0 & 0 \\
0 & 0 & 0 & 1 & \cdots & 0 & 0 \\
\vdots & \vdots & \vdots & \vdots & \ddots & \vdots & \vdots \\
0 & 0 & 0 & 0 & \cdots & 1 & 0 \\
0 & 0 & 0 & 0 & \cdots & 0 & 1 \\
1 & 1 & 0 & 0 & \cdots & 0 & 0 \\
\end{bmatrix}
\begin{bmatrix}
f[i-1][0] \\
f[i-1][1] \\
f[i-1][2] \\
\vdots \\
f[i-1][23] \\
f[i-1][24] \\
f[i-1][25] \\
\end{bmatrix}
$$

把上式中的三个矩阵分别记作 $F[i],M,F[i-1]$，即

$$
F[i] = M\times F[i-1]
$$

那么有

$$
\begin{aligned}
F[t] ={} & M\times F[t-1]      \\
={} & M\times M\times F[t-2]        \\
={} & M\times M\times M\times  F[t-3]        \\
\vdots & \\
={} & M^t\times F[0]
\end{aligned}
$$

其中 $F[0]$ 是一个长为 $26$ 的列向量，值全为 $1$（对应着 $f$ 数组的初始值 $f[0][j] = 1$）。

$M^t$ 可以用**快速幂**计算，原理请看[【图解】一张图秒懂快速幂](https://leetcode.cn/problems/powx-n/solution/tu-jie-yi-zhang-tu-miao-dong-kuai-su-mi-ykp3i/)。

根据矩阵乘法的运算法则，$f[t][j]$ 等于矩阵 $M^t$ 的第 $j$ 行与列向量 $F[0]$ 计算点积。由于 $F[0]$ 全为 $1$，所以 $f[t][j]$ 也等于 $M^t$ 第 $j$ 行的元素和。

一般地，枚举 $i=0,1,\cdots, 25$ 以及 $j=i+1,i+2,\cdots,i+\textit{nums}[i]$，初始化 $M[i][j\bmod 26]=1$。

```
const mod = 1_000_000_007

type matrix [][]int

func newMatrix(n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func (a matrix) mul(b matrix) matrix {
	c := newMatrix(len(a), len(b[0]))
	for i, row := range a {
		for j := range b[0] {
			for k, v := range row {
				c[i][j] = (c[i][j] + v*b[k][j]) % mod
			}
		}
	}
	return c
}

func (a matrix) pow(n int) matrix {
	res := make(matrix, len(a))
	for i := range res {
		res[i] = make([]int, len(a))
		res[i][i] = 1
	}
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}

func lengthAfterTransformations(s string, t int, nums []int) (ans int) {
	const size = 26
	m := newMatrix(size, size)
	for i, c := range nums {
		for j := i + 1; j <= i+c; j++ {
			m[i][j%size] = 1
		}
	}
	m = m.pow(t)

	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	for i, row := range m {
		// m 第 i 行的和就是 f[t][i]
		fti := 0
		for _, x := range row {
			fti += x
		}
		ans += fti * cnt[i]
	}
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|^3\log t)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|^2)$。

更多相似题目，见下面动态规划题单中的「**§7.3 矩阵快速幂优化 DP**」。

## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
- [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
- [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)