### 题目

<p>给你一个&nbsp;<code>m x n</code>&nbsp;的二进制矩阵&nbsp;<code>grid</code>&nbsp;。</p>

<p>如果矩阵中一行或者一列从前往后与从后往前读是一样的，那么我们称这一行或者这一列是 <strong>回文</strong>&nbsp;的。</p>

<p>你可以将 <code>grid</code>&nbsp;中任意格子的值 <strong>翻转</strong>&nbsp;，也就是将格子里的值从 <code>0</code>&nbsp;变成 <code>1</code>&nbsp;，或者从 <code>1</code>&nbsp;变成 <code>0</code>&nbsp;。</p>

<p>请你返回 <strong>最少</strong>&nbsp;翻转次数，使得矩阵中 <strong>所有</strong>&nbsp;行和列都是 <strong>回文的</strong>&nbsp;，且矩阵中 <code>1</code>&nbsp;的数目可以被 <code>4</code>&nbsp;<strong>整除</strong>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[1,0,0],[0,1,0],[0,0,1]]</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><strong>解释：</strong></p>

<p><img src="https://assets.leetcode.com/uploads/2024/08/01/image.png" style="width: 400px; height: 105px;" /></p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[0,1],[0,1],[0,0]]</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/07/08/screenshot-from-2024-07-09-01-37-48.png" style="width: 300px; height: 104px;" /></p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[1],[1]]</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/08/01/screenshot-from-2024-08-01-23-05-26.png" style="width: 200px; height: 70px;" /></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>m == grid.length</code></li>
	<li><code>n == grid[i].length</code></li>
	<li><code>1 &lt;= m * n &lt;= 2 * 10<sup>5</sup></code></li>
	<li><code>0 &lt;= grid[i][j] &lt;= 1</code></li>
</ul>

### 思路

为方便描述，把 $\textit{grid}$ 简称为 $a$。

由于所有行和列都必须是回文的，所以要满足

$$
a[i][j] = a[i][n-1-j] = a[m-1-i][j] = a[m-1-i][n-1-j]
$$

也就是这四个数要么都是 $1$，要么都是 $0$。其中 $0\le i \le \lfloor m/2 \rfloor,\ 0\le j \le \lfloor n/2 \rfloor$。

设

$$
\textit{cnt}_1 = a[i][j] + a[i][n-1-j] + a[m-1-i][j] + a[m-1-i][n-1-j]
$$

把这四个数都变成 $0$ 需要翻转 $\textit{cnt}_1$ 次，都变成 $1$ 需要翻转 $4-\textit{cnt}_1$ 次。

两种情况取最小值，把

$$
\min(\textit{cnt}_1, 4-\textit{cnt}_1)
$$

加入答案。

## 特殊情况

讨论**正中间一排**（如果 $m$ 是奇数）和**正中间一列**（如果 $n$ 是奇数）中的格子要如何翻转。

首先，如果 $m$ 和 $n$ 都是奇数，那么**正中央的格子** $(\lfloor m/2 \rfloor, \lfloor n/2 \rfloor)$ 必须是 $0$。

然后统计正中间一排（如果 $m$ 是奇数）和正中间一列（如果 $n$ 是奇数）中的格子：

- 设 $\textit{diff}$ 为镜像位置不同的数对个数。
- 设 $\textit{cnt}_1$ 为镜像位置相同的 $1$ 的个数。

分类讨论：

- 如果 $\textit{cnt}_1$ 是 $4$ 的倍数，那么可以把这 $\textit{diff}$ 对数全部变成 $0$，把 $\textit{diff}$ 加入答案。
- 如果 $\textit{cnt}_1$ 不是 $4$ 的倍数，由于 $\textit{cnt}_1$ 是偶数，其除以 $4$ 必定余 $2$。继续讨论：
    - 如果 $\textit{diff} > 0$，我们可以把其中一对数都变成 $1$，其余 $\textit{diff}-1$ 对数全部变成 $0$，这样 $1$ 的个数就是 $4$ 的倍数了，把 $\textit{diff}$ 加入答案。
    - 如果 $\textit{diff} = 0$，可以把 $\textit{cnt}_1$ 中的两个 $1$ 变成 $0$，这样 $1$ 的个数就是 $4$ 的倍数了，把答案增加 $2$。

综上所述：

- 如果 $\textit{diff} > 0$，额外把 $\textit{diff}$ 加入答案。
- 如果 $\textit{diff} = 0$，额外把 $\textit{cnt}_1\bmod 4$ 加入答案。

```
func minFlips(a [][]int) (ans int) {
	m, n := len(a), len(a[0])
	for i, row := range a[:m/2] {
		row2 := a[m-1-i]
		for j, x := range row[:n/2] {
			cnt1 := x + row[n-1-j] + row2[j] + row2[n-1-j]
			ans += min(cnt1, 4-cnt1) // 全为 1 或全为 0
		}
	}

	if m%2 > 0 && n%2 > 0 {
		// 正中间的数必须是 0
		ans += a[m/2][n/2]
	}

	diff, cnt1 := 0, 0
	if m%2 > 0 {
		// 统计正中间这一排
		row := a[m/2]
		for j, x := range row[:n/2] {
			if x != row[n-1-j] {
				diff++
			} else {
				cnt1 += x * 2
			}
		}
	}
	if n%2 > 0 {
		// 统计正中间这一列
		for i, row := range a[:m/2] {
			if row[n/2] != a[m-1-i][n/2] {
				diff++
			} else {
				cnt1 += row[n/2] * 2
			}
		}
	}

	if diff > 0 {
		ans += diff
	} else {
		ans += cnt1 % 4
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。

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