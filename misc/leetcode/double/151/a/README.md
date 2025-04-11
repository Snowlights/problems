### 题目

<p>给你一个整数数组 <code>nums</code>。请你按照以下顺序 <strong>依次</strong>&nbsp;执行操作，转换 <code>nums</code>：</p>

<ol>
	<li>将每个偶数替换为 0。</li>
	<li>将每个奇数替换为 1。</li>
	<li>按&nbsp;<strong>非递减&nbsp;</strong>顺序排序修改后的数组。</li>
</ol>

<p>执行完这些操作后，返回结果数组。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1:</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [4,3,2,1]</span></p>

<p><strong>输出：</strong><span class="example-io">[0,0,1,1]</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>将偶数（4 和 2）替换为 0，将奇数（3 和 1）替换为 1。现在，<code>nums = [0, 1, 0, 1]</code>。</li>
	<li>按非递减顺序排序 <code>nums</code>，得到 <code>nums = [0, 0, 1, 1]</code>。</li>
</ul>
</div>

<p><strong class="example">示例 2:</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [1,5,1,4,2]</span></p>

<p><strong>输出：</strong><span class="example-io">[0,0,1,1,1]</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>将偶数（4 和 2）替换为 0，将奇数（1, 5 和 1）替换为 1。现在，<code>nums = [1, 1, 1, 0, 0]</code>。</li>
	<li>按非递减顺序排序&nbsp;<code>nums</code>，得到 <code>nums = [0, 0, 1, 1, 1]</code>。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 1000</code></li>
</ul>

### 思路

统计 $\textit{nums}$ 中的偶数个数 $\textit{cnt}_0$ 和奇数个数 $\textit{cnt}_1$，那么答案就是 $\textit{cnt}_0$ 个 $0$，后跟 $\textit{cnt}_1$ 个 $1$。


```
func transformArray(nums []int) []int {
	cnt := [2]int{}
	for _, x := range nums {
		cnt[x%2]++
	}
	clear(nums[:cnt[0]]) // 置 0
	for i := cnt[0]; i < len(nums); i++ {
		nums[i] = 1
	}
	return nums
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 写法二

也可以只统计 $\textit{cnt}_1$，那么 $\textit{cnt}_0 = n-\textit{cnt}_1$。


```
func transformArray(nums []int) []int {
	cnt1 := 0
	for _, x := range nums {
		cnt1 += x % 2
	}
	n := len(nums)
	cnt0 := n - cnt1
	clear(nums[:cnt0])
	for i := cnt0; i < n; i++ {
		nums[i] = 1
	}
	return nums
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

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