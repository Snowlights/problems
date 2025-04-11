### 题目

<p>给你一个整数数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>你的任务是找到 <code>nums</code>&nbsp;中的 <strong>最长子序列</strong>&nbsp;<code>seq</code>&nbsp;，这个子序列中相邻元素的 <strong>绝对差</strong>&nbsp;构成一个 <strong>非递增</strong>&nbsp;整数序列。换句话说，<code>nums</code>&nbsp;中的序列&nbsp;<code>seq<sub>0</sub></code>, <code>seq<sub>1</sub></code>, <code>seq<sub>2</sub></code>, ..., <code>seq<sub>m</sub></code>&nbsp;满足&nbsp;<code>|seq<sub>1</sub> - seq<sub>0</sub>| &gt;= |seq<sub>2</sub> - seq<sub>1</sub>| &gt;= ... &gt;= |seq<sub>m</sub> - seq<sub>m - 1</sub>|</code>&nbsp;。</p>

<p>请你返回这个子序列的长度。</p>

<p>一个&nbsp;<strong>子序列</strong>&nbsp;指的是从一个数组中删除零个或多个元素后，剩下元素不改变顺序得到的&nbsp;<strong>非空</strong>&nbsp;数组。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [16,6,3]</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><b>解释：</b></p>

<p>最长子序列是&nbsp;<code>[16, 6, 3]</code>&nbsp;，相邻绝对差值为&nbsp;<code>[10, 3]</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [6,5,3,4,2,1]</span></p>

<p><span class="example-io"><b>输出：</b>4</span></p>

<p><strong>解释：</strong></p>

<p>最长子序列是&nbsp;<code>[6, 4, 2, 1]</code>&nbsp;，相邻绝对差值为&nbsp;<code>[2, 2, 1]</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [10,20,10,19,10,20]</span></p>

<p><span class="example-io"><b>输出：</b>5</span></p>

<p><b>解释：</b></p>

<p>最长子序列是&nbsp;<code>[10, 20, 10, 19, 10]</code>&nbsp;，相邻绝对差值为&nbsp;<code>[10, 10, 9, 9]</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 300</code></li>
</ul>


### 思路


分类讨论：

- $x=\textit{nums}[i]$ 单独形成一个子序列，那么 $f[i][j] = 1$。
- 子序列最后两数之差严格大于 $j$，也就是至少为 $j+1$，那么 $f[i][j] = f[i][j+1]$。
- 子序列最后两数之差恰好等于 $j$，那么子序列倒数第二个数是 $x-j$ 或者 $x+j$，并且子序列倒数第二第三的两数之差至少为 $j$，那么 $f[i][j] = \max(f[\textit{last}[x-j]][j]+1, f[\textit{last}[x+j]][j]+1)$。其中 $\textit{last}[x]$ 表示 $x$ 上一次出现的下标。

> 注：只考虑上一次出现的下标，是因为 $i$ 越大，子序列就越长。

所有情况取最大值，得

$$
f[i][j] = \max(1, f[i][j+1], f[\textit{last}[x-j]][j]+1, f[\textit{last}[x+j]][j]+1)
$$

> 注：由于 $f[i][j]$ 需要从 $f[i][j+1]$ 转移过来，所以 $j$ 要倒序枚举。

最终答案为所有 $f[i][j]$ 的最大值。


```
func longestSubsequence(nums []int) (ans int) {
	mx := slices.Max(nums)
	maxD := mx - slices.Min(nums)
	f := make([][]int, mx+1)
	for i := range f {
		f[i] = make([]int, maxD+1)
		for j := range f[i] {
			f[i][j] = math.MinInt
		}
	}

	for _, x := range nums {
		fx := 1
		for j := maxD; j >= 0; j-- {
			if x-j >= 0 {
				fx = max(fx, f[x-j][j]+1)
			}
			if x+j <= mx {
				fx = max(fx, f[x+j][j]+1)
			}
			f[x][j] = fx
			ans = max(ans, fx)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nD)$，其中 $n$ 是 $\textit{nums}$ 的长度，$D=\max(\textit{nums})-\min(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(nD)$。

## 优化

$\textit{last}$ 可以去掉，改成定义 $f[x][j]$ 表示以元素 $x$ 结尾的、最后两个数之差**至少**为 $j$ 的子序列的最长长度。

状态转移方程改成

$$
f[x][j] = \max(1, f[x][j+1], f[x-j][j]+1, f[x+j][j]+1)
$$

⚠**注意**：$j=0$ 的时候要单独处理，否则转移方程中的 $+1$ 会重复累加。更简单的做法是，用一个变量 $\textit{fx}$ 表示 $f[x][j]$，具体见代码。

```
func longestSubsequence(nums []int) (ans int) {
	mx := slices.Max(nums)
	maxD := mx - slices.Min(nums)
	f := make([][]int, mx+1)
	for i := range f {
		f[i] = make([]int, maxD+1)
	}

	for _, x := range nums {
		fx := 1
		for j := maxD; j >= 0; j-- {
			if x-j >= 0 {
				fx = max(fx, f[x-j][j]+1)
			}
			if x+j <= mx {
				fx = max(fx, f[x+j][j]+1)
			}
			f[x][j] = fx
			ans = max(ans, fx)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nD)$，其中 $n$ 是 $\textit{nums}$ 的长度，$D=\max(\textit{nums})-\min(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(D^2)$。

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
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)