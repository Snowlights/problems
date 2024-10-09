#### 题目

<p>给你一个整数数组 <code>nums</code> 和一个二维数组 <code>queries</code>，其中 <code>queries[i] = [pos<sub>i</sub>, x<sub>i</sub>]</code>。</p>

<p>对于每个查询 <code>i</code>，首先将 <code>nums[pos<sub>i</sub>]</code> 设置为 <code>x<sub>i</sub></code>，然后计算查询 <code>i</code> 的答案，该答案为 <code>nums</code> 中 <strong>不包含相邻元素 </strong>的子序列的 <strong>最大 </strong>和。</p>

<p>返回所有查询的答案之和。</p>

<p>由于最终答案可能非常大，返回其对 <code>10<sup>9</sup> + 7</code> <strong>取余</strong> 的结果。</p>

<p><strong>子序列</strong> 是指从另一个数组中删除一些或不删除元素而不改变剩余元素顺序得到的数组。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [3,5,9], queries = [[1,-2],[0,-3]]</span></p>

<p><strong>输出：</strong><span class="example-io">21</span></p>

<p><strong>解释：</strong><br />
执行第 1 个查询后，<code>nums = [3,-2,9]</code>，不包含相邻元素的子序列的最大和为 <code>3 + 9 = 12</code>。<br />
执行第 2 个查询后，<code>nums = [-3,-2,9]</code>，不包含相邻元素的子序列的最大和为 9 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [0,-1], queries = [[0,-5]]</span></p>

<p><strong>输出：</strong><span class="example-io">0</span></p>

<p><strong>解释：</strong><br />
执行第 1 个查询后，<code>nums = [-5,-1]</code>，不包含相邻元素的子序列的最大和为 0（选择空子序列）。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>-10<sup>5</sup> &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= queries.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>queries[i] == [pos<sub>i</sub>, x<sub>i</sub>]</code></li>
	<li><code>0 &lt;= pos<sub>i</sub> &lt;= nums.length - 1</code></li>
	<li><code>-10<sup>5</sup> &lt;= x<sub>i</sub> &lt;= 10<sup>5</sup></code></li>
</ul>


#### 思路

本题是可以修改数组元素值的 [198. 打家劫舍](https://leetcode.cn/problems/house-robber/)。

为了解决本题，首先来换一个角度，用**分治**的思想解决打家劫舍。

把 $\textit{nums}$ 从中间切开，分成两个子数组，分别记作 $a$ 和 $b$。

设 $f(A)$ 为数组 $A$ 的打家劫舍的答案。要计算 $f(\textit{nums})$，看上去，我们只需要分别计算 $f(a)$ 和 $f(b)$，但这是不对的。万一我们同时选了 $a$ 的最后一个数和 $b$ 的第一个数，就不满足题目要求了。

怎么办？加个约束，分类讨论：

- 约束 $a$ 的最后一个数一定不选，即 $f(a') + f(b)$，其中 $a'$ 是去掉最后一个数的 $a$。
- 约束 $b$ 的第一个数一定不选，即 $f(a) + f(b')$，其中 $b'$ 表示去掉第一个数的 $b$。
- 这两种情况取最大值，即为 $f(\textit{nums})$。

为方便我们继续讨论，定义：

- $f_{00}(A)$ 表示在 $A$ 第一个数一定不选，最后一个数也一定不选的情况下，打家劫舍的答案。
- $f_{01}(A)$ 表示在 $A$ 第一个数一定不选，最后一个数可选可不选的情况下，打家劫舍的答案。
- $f_{10}(A)$ 表示在 $A$ 第一个数可选可不选，最后一个数一定不选的情况下，打家劫舍的答案。
- $f_{11}(A)$ 表示在 $A$ 第一个数可选可不选，最后一个数也可选可不选的情况下，打家劫舍的答案，这等于上面定义的 $f(A)$。

按照该定义，上面的分类讨论可以表述为

$$
f_{11}(\textit{nums}) = \max(f_{10}(a) + f_{11}(b),\ f_{11}(a) + f_{01}(b))
$$

要计算 $f_{10}$ 和 $f_{01}$，得继续分类讨论。以 $f_{10}(a)$ 为例，把 $a$ 分成的左右两个数组 $p$ 和 $q$，那么：

- $p$ 的最后一个数一定不选，即 $f_{10}(p) + f_{10}(q)$，注意 $q$ 的最后一个数也不能选，因为我们计算的是 $f_{10}(a)$，$a$ 的最后一个数一定不能选。
- $q$ 的第一个数一定不选，即 $f_{11}(p) + f_{00}(q)$。
- 这两种情况取最大值，得

$$
f_{10}(a) = \max(f_{10}(p) + f_{10}(q),\ f_{11}(p) + f_{00}(q))
$$

同理可以得到 $f_{00}$ 和 $f_{01}$。

综上所述：

$$
\begin{align}
&f_{00}(a) = \max(f_{00}(p) + f_{10}(q),\ f_{01}(p) + f_{00}(q))\\
&f_{01}(a) = \max(f_{00}(p) + f_{11}(q),\ f_{01}(p) + f_{01}(q))\\
&f_{10}(a) = \max(f_{10}(p) + f_{10}(q),\ f_{11}(p) + f_{00}(q))\\
&f_{11}(a) = \max(f_{10}(p) + f_{11}(q),\ f_{11}(p) + f_{01}(q))
\end{align}
$$

这样就可以分治计算 $\textit{nums}$ 的打家劫舍了。

递归边界：如果 $a$ 的长度等于 $1$，那么按照定义，$f_{11}(a) = \max(a[0], 0)$，其余 $f_{00},f_{01},f_{10}$ 均为 $0$。

回到本题，对于下标 $i$ 的修改操作，我们可以用**线段树**的单点修改实现，按照上面列出的四个式子，合并左右区间。对于查询操作，由于询问的是整个数组，询问结果就是线段树根节点的 $f_{11}$，加入答案。

```
// f00 表示第一个数一定不选，最后一个数一定不选
// f01 表示第一个数一定不选，最后一个数可选可不选
// f10 表示第一个数可选可不选，最后一个数一定不选
// f11 表示第一个数可选可不选，最后一个数可选可不选，也就是没有任何限制
type data struct{ f00, f01, f10, f11 int }
type seg []data

func (t seg) maintain(o int) {
	a, b := t[o<<1], t[o<<1|1]
	t[o] = data{
		max(a.f00+b.f10, a.f01+b.f00),
		max(a.f00+b.f11, a.f01+b.f01),
		max(a.f10+b.f10, a.f11+b.f00),
		max(a.f10+b.f11, a.f11+b.f01),
	}
}

func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o].f11 = max(a[l], 0)
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r, i, val int) {
	if l == r {
		t[o].f11 = max(val, 0)
		return
	}
	m := (l + r) >> 1
	if i <= m {
		t.update(o<<1, l, m, i, val)
	} else {
		t.update(o<<1|1, m+1, r, i, val)
	}
	t.maintain(o)
}

func maximumSumSubsequence(nums []int, queries [][]int) (ans int) {
	n := len(nums)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(nums, 1, 0, n-1)
	for _, q := range queries {
		t.update(1, 0, n-1, q[0], q[1])
		ans += t[1].f11 // 注意 f11 没有任何限制，也就是整个数组的打家劫舍
	}
	return ans % 1_000_000_007
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
- [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)