### 题目

<p>给你一个整数数组&nbsp;<code>nums</code>&nbsp;和两个整数&nbsp;<code>k</code> 和&nbsp;<code>numOperations</code>&nbsp;。</p>

<p>你必须对 <code>nums</code>&nbsp;执行 <strong>操作</strong>&nbsp; <code>numOperations</code>&nbsp;次。每次操作中，你可以：</p>

<ul>
	<li>选择一个下标&nbsp;<code>i</code>&nbsp;，它在之前的操作中 <strong>没有</strong>&nbsp;被选择过。</li>
	<li>将 <code>nums[i]</code>&nbsp;增加范围&nbsp;<code>[-k, k]</code>&nbsp;中的一个整数。</li>
</ul>

<p>在执行完所有操作以后，请你返回 <code>nums</code>&nbsp;中出现 <strong>频率最高</strong>&nbsp;元素的出现次数。</p>

<p>一个元素 <code>x</code>&nbsp;的 <strong>频率</strong>&nbsp;指的是它在数组中出现的次数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,4,5], k = 1, numOperations = 2</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p>通过以下操作得到最高频率 2 ：</p>

<ul>
	<li>将&nbsp;<code>nums[1]</code>&nbsp;增加 0 ，<code>nums</code> 变为&nbsp;<code>[1, 4, 5]</code>&nbsp;。</li>
	<li>将&nbsp;<code>nums[2]</code>&nbsp;增加 -1 ，<code>nums</code> 变为&nbsp;<code>[1, 4, 4]</code>&nbsp;。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [5,11,20,20], k = 5, numOperations = 1</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p>通过以下操作得到最高频率 2 ：</p>

<ul>
	<li>将 <code>nums[1]</code> 增加 0 。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= k &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= numOperations &lt;= nums.length</code></li>
</ul>

### 思路

## 方法一：差分

**前置知识**：[差分原理讲解](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)（推荐和[【图解】从一维差分到二维差分](https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/) 一起看）。

设 $x = \textit{nums}[i]$。

$x$ 可以变成 $[x-k,x+k]$ 中的整数。注意对于同一个 $\textit{nums}[i]$ 至多操作一次。

反过来想，对于一个整数 $y$，有多少个数可以变成 $y$？

这可以用**差分**计算，也就是把 $[x-k,x+k]$ 中的每个整数的出现次数都加一。

最后计算差分的前缀和，就得到了有 $\textit{sumD}$ 个数可以变成 $y$。

如果 $y$ 不在 $\textit{nums}$ 中，那么 $y$ 的最大频率为 $\min(\textit{sumD}, \textit{numOperations})$。

如果 $y$ 在 $\textit{nums}$ 中，且出现了 $\textit{cnt}$ 次，那么有 $\textit{sumD}-\textit{cnt}$ 个其他元素可以变成 $y$，但这不能超过 $\textit{numOperations}$，所以有

$$
\min(\textit{sumD}-\textit{cnt}, \textit{numOperations})
$$

个其他元素可以变成 $y$，再加上 $y$ 自身出现的次数 $\textit{cnt}$，得到 $y$ 的最大频率为

$$
\textit{cnt} + \min(\textit{sumD}-\textit{cnt}, \textit{numOperations}) = \min(\textit{sumD}, \textit{cnt}+\textit{numOperations})
$$

注意上式兼容 $y$ 不在 $\textit{nums}$ 中的情况，此时 $\textit{cnt}=0$。


### 答疑

**问**：为什么代码只考虑在 $\textit{diff}$ 中的数字？

**答**：比如 $x$ 在 $\textit{diff}$ 中，$x+1$ 不在 $\textit{diff}$ 中，那么 $x+1$ 的 $\textit{sumD}$ 和 $\textit{x}$ 的是一样的，无需重复计算。

```
func maxFrequency(nums []int, k, numOperations int) (ans int) {
	cnt := map[int]int{}
	diff := map[int]int{}
	for _, x := range nums {
		cnt[x]++
		diff[x] += 0 // 把 x 插入 diff，以保证下面能遍历到 x
		diff[x-k]++  // 把 [x-k, x+k] 中的每个整数的出现次数都加一
		diff[x+k+1]--
	}

	sumD := 0
	for _, x := range slices.Sorted(maps.Keys(diff)) {
		sumD += diff[x]
		ans = max(ans, min(sumD, cnt[x]+numOperations))
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：同向三指针 + 滑动窗口

**前置知识**：[滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

把 $\textit{nums}$ 从小到大排序。

方法一中的 $\textit{cnt}[x]$ 也可以用同向三指针/滑动窗口计算。

- 如果 $x$ 在 $\textit{nums}$ 中，用**同向三指针**计算。
- 如果 $x$ 不在 $\textit{nums}$ 中，用**滑动窗口**计算。

### 同向三指针

遍历排序后的 $\textit{nums}$，设 $x=\textit{nums}[i]$。计算元素值在 $[x-k,x+k]$ 中的元素个数。

遍历的同时，维护左指针 $\textit{left}$，它是最小的满足

$$
\textit{nums}[\textit{left}] \ge x - k
$$

的下标。

遍历的同时，维护右指针 $\textit{right}$，它是最小的满足

$$
\textit{nums}[\textit{right}] > x + k
$$

的下标。如果不存在，则 $\textit{right}=n$。

那么方法一中的 $\textit{sumD}$ 就是

$$
\textit{right} - \textit{left}
$$

### 滑动窗口

枚举 $x=\textit{nums}[\textit{right}]$ 作为被修改的最大元素。计算元素值在 $[x-2k,x]$ 中的元素个数。

设 $\textit{nums}[\textit{left}]$ 是被修改的最小元素，那么需要满足

$$
\textit{nums}[\textit{right}] - \textit{nums}[\textit{left}] \le 2k
$$

那么可以把

$$
\textit{right} - \textit{left} + 1
$$

个数都变成一样的。注意上式不能超过 $\textit{numOperations}$。

### 细节

如果同向三指针计算完毕后，发现答案已经 $\ge \textit{numOperations}$，那么无需计算滑动窗口。

```
func maxFrequency(nums []int, k, numOperations int) (ans int) {
	slices.Sort(nums)

	n := len(nums)
	var cnt, left, right int
	for i, x := range nums {
		cnt++
		// 循环直到连续相同段的末尾，这样可以统计出 x 的出现次数
		if i < n-1 && x == nums[i+1] {
			continue
		}
		for nums[left] < x-k {
			left++
		}
		for right < n && nums[right] <= x+k {
			right++
		}
		ans = max(ans, min(right-left, cnt+numOperations))
		cnt = 0
	}

	if ans >= numOperations {
		return ans
	}

	left = 0
	for right, x := range nums {
		for nums[left] < x-k*2 {
			left++
		}
		ans = max(ans, right-left+1)
	}
	return min(ans, numOperations) // 最后和 numOperations 取最小值
}
```

也可以把两个 for 循环合起来。

```
func maxFrequency(nums []int, k, numOperations int) (ans int) {
	slices.Sort(nums)

	n := len(nums)
	var cnt, left, right, left2 int
	for i, x := range nums {
		for nums[left2] < x-k*2 {
			left2++
		}
		ans = max(ans, min(i-left2+1, numOperations))

		cnt++
		// 循环直到连续相同段的末尾，这样可以统计出 x 的出现次数
		if i < n-1 && x == nums[i+1] {
			continue
		}
		for nums[left] < x-k {
			left++
		}
		for right < n && nums[right] <= x+k {
			right++
		}
		ans = max(ans, min(right-left, cnt+numOperations))
		cnt = 0
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

更多相似题目，见下面滑动窗口题单中的「**§2.1 求最长/最大**」，以及数据结构题单中的「**§2.1 一维差分**」。

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