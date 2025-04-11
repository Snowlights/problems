#### 题目

<p>给你一个&nbsp;<strong>循环&nbsp;</strong>数组&nbsp;<code>nums</code>&nbsp;和一个数组&nbsp;<code>queries</code>&nbsp;。</p>

<p>对于每个查询&nbsp;<code>i</code>&nbsp;，你需要找到以下内容：</p>

<ul>
	<li>数组&nbsp;<code>nums</code>&nbsp;中下标&nbsp;<code>queries[i]</code>&nbsp;处的元素与&nbsp;<strong>任意&nbsp;</strong>其他下标&nbsp;<code>j</code>（满足&nbsp;<code>nums[j] == nums[queries[i]]</code>）之间的&nbsp;<strong>最小&nbsp;</strong>距离。如果不存在这样的下标&nbsp;<code>j</code>，则该查询的结果为 <code>-1</code> 。</li>
</ul>

<p>返回一个数组&nbsp;<code>answer</code>，其大小与&nbsp;<code>queries</code>&nbsp;相同，其中&nbsp;<code>answer[i]</code>&nbsp;表示查询<code>i</code>的结果。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,3,1,4,1,3,2], queries = [0,3,5]</span></p>

<p><strong>输出：</strong> <span class="example-io">[2,-1,3]</span></p>

<p><strong>解释：</strong></p>
<ul>
	<li>查询 0：下标&nbsp;<code>queries[0] = 0</code>&nbsp;处的元素为&nbsp;<code>nums[0] = 1</code>&nbsp;。最近的相同值下标为 2，距离为 2。</li>
	<li>查询 1：下标&nbsp;<code>queries[1] = 3</code>&nbsp;处的元素为&nbsp;<code>nums[3] = 4</code>&nbsp;。不存在其他包含值 4 的下标，因此结果为 -1。</li>
	<li>查询 2：下标&nbsp;<code>queries[2] = 5</code>&nbsp;处的元素为&nbsp;<code>nums[5] = 3</code>&nbsp;。最近的相同值下标为 1，距离为 3（沿着循环路径：<code>5 -&gt; 6 -&gt; 0 -&gt; 1</code>）。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,2,3,4], queries = [0,1,2,3]</span></p>

<p><strong>输出：</strong> <span class="example-io">[-1,-1,-1,-1]</span></p>

<p><strong>解释：</strong></p>

<p>数组&nbsp;<code>nums</code>&nbsp;中的每个值都是唯一的，因此没有下标与查询的元素值相同。所有查询的结果均为 -1。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>
<ul>
	<li><code>1 &lt;= queries.length &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>6</sup></code></li>
	<li><code>0 &lt;= queries[i] &lt; nums.length</code></li>
</ul>

#### 思路

## 方法一：二分查找

看示例 1，所有 $1$ 的下标列表是 $p=[0,2,4]$。
由于 $\textit{nums}$ 是循环数组：
- 在下标列表前面添加 $4-n=-3$，相当于认为在 $-3$ 下标处也有一个 $1$。
- 在下标列表末尾添加 $0+n=7$，相当于认为在 $7$ 下标处也有一个 $1$。

修改后的下标列表为 $p=[-3,0,2,4,7]$。
于是，我们在 $p$ 中二分查找下标 $i$，设二分返回值为 $j$，那么：
- $p[j-1]$ 就是在 $i$ 左边的最近位置。
- $p[j+1]$ 就是在 $i$ 右边的最近位置。

两个距离取最小值，答案为

$$
\min(i-p[j-1], p[j+1]-i)
$$

如果 $\textit{nums}[i]$ 在 $\textit{nums}$ 中只出现了一次，那么答案为 $-1$。
代码实现时，可以直接把答案记录在 $\textit{queries}$ 数组中。

```
func solveQueries(nums []int, queries []int) []int {
	indices := map[int][]int{}
	for i, x := range nums {
		indices[x] = append(indices[x], i)
	}

	n := len(nums)
	for x, p := range indices {
		// 前后各加一个哨兵
		i0 := p[0]
		p = slices.Insert(p, 0, p[len(p)-1]-n)
		indices[x] = append(p, i0+n)
	}

	for qi, i := range queries {
		p := indices[nums[i]]
		if len(p) == 3 {
			queries[qi] = -1
		} else {
			j := sort.SearchInts(p, i)
			queries[qi] = min(i-p[j-1], p[j+1]-i)
		}
	}
	return queries
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。每次二分需要 $\mathcal{O}(\log n)$ 的时间。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 方法二：预处理左右最近相同元素的下标

定义 $\textit{left}[i]$ 表示在 $i$ 左边的等于 $\textit{nums}[i]$ 的最近元素下标。
注意数组是循环数组，我们可以像方法一那样，用 $-1$ 表示最后一个数的下标，$-2$ 表示倒数第二个数的下标，依此类推。
定义 $\textit{right}[i]$ 表示在 $i$ 右边的等于 $\textit{nums}[i]$ 的最近元素下标。注意数组是循环数组，我们可以像方法一那样，用 $n$ 表示第一个数的下标，$n+1$ 表示第二个数的下标，依此类推。

计算方式：
- 从 $-n$ 循环到 $n-1$，同时用一个哈希表记录每个数的最新位置。当 $i\ge 0$ 时，$\textit{left}[i]$ 就是哈希中记录的 $\textit{nums}[i]$ 的位置。注意先更新 $\textit{left}[i]$ 再更新哈希表。
- 从 $2n-1$ 循环到 $0$，同时用一个哈希表记录每个数的最新位置。当 $i < n$ 时，$\textit{right}[i]$ 就是哈希中记录的 $\textit{nums}[i]$ 的位置。注意先更新 $\textit{right}[i]$ 再更新哈希表。

答案为：

$$
\min(i-\textit{left}[i], \textit{right}[i]-i)
$$

如果上式等于 $n$，说明只有一个 $\textit{nums}[i]$，答案为 $-1$。

```
func solveQueries(nums []int, queries []int) []int {
	n := len(nums)
	left := make([]int, n)
	pos := map[int]int{}
	for i := -n; i < n; i++ {
		if i >= 0 {
			left[i] = pos[nums[i]]
		}
		pos[nums[(i+n)%n]] = i
	}

	right := make([]int, n)
	clear(pos)
	for i := n*2 - 1; i >= 0; i-- {
		if i < n {
			right[i] = pos[nums[i]]
		}
		pos[nums[i%n]] = i
	}

	for qi, i := range queries {
		l := left[i]
		if i-l == n {
			queries[qi] = -1
		} else {
			queries[qi] = min(i-l, right[i]-i)
		}
	}
	return queries
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)