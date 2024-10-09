### 题目

<p>给你一个长度为 <code>n</code>&nbsp;的整数数组&nbsp;<code>nums</code>&nbsp;和一个正整数&nbsp;<code>k</code>&nbsp;。</p>

<p>一个数组的 <strong>能量值</strong> 定义为：</p>

<ul>
	<li>如果 <strong>所有</strong>&nbsp;元素都是依次&nbsp;<strong>连续</strong> 且 <strong>上升</strong> 的，那么能量值为 <strong>最大</strong>&nbsp;的元素。</li>
	<li>否则为 -1 。</li>
</ul>

<p>你需要求出 <code>nums</code>&nbsp;中所有长度为 <code>k</code>&nbsp;的&nbsp;<span data-keyword="subarray-nonempty">子数组</span>&nbsp;的能量值。</p>

<p>请你返回一个长度为 <code>n - k + 1</code>&nbsp;的整数数组&nbsp;<code>results</code>&nbsp;，其中&nbsp;<code>results[i]</code>&nbsp;是子数组&nbsp;<code>nums[i..(i + k - 1)]</code>&nbsp;的能量值。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,2,3,4,3,2,5], k = 3</span></p>

<p><b>输出：</b>[3,4,-1,-1,-1]</p>

<p><strong>解释：</strong></p>

<p><code>nums</code>&nbsp;中总共有 5 个长度为 3 的子数组：</p>

<ul>
	<li><code>[1, 2, 3]</code>&nbsp;中最大元素为 3 。</li>
	<li><code>[2, 3, 4]</code>&nbsp;中最大元素为 4 。</li>
	<li><code>[3, 4, 3]</code>&nbsp;中元素 <strong>不是</strong>&nbsp;连续的。</li>
	<li><code>[4, 3, 2]</code>&nbsp;中元素 <b>不是</b>&nbsp;上升的。</li>
	<li><code>[3, 2, 5]</code>&nbsp;中元素 <strong>不是</strong>&nbsp;连续的。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [2,2,2,2,2], k = 4</span></p>

<p><span class="example-io"><b>输出：</b>[-1,-1]</span></p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [3,2,3,2,3,2], k = 2</span></p>

<p><span class="example-io"><b>输出：</b>[-1,3,-1,3,-1]</span></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == nums.length &lt;= 500</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= k &lt;= n</code></li>
</ul>


### 思路

遍历数组的同时，用一个计数器 $\textit{cnt}$ 统计连续递增的元素个数：

- 初始化 $\textit{cnt}=0$。
- 如果 $i=0$ 或者 $\textit{nums}[i]= \textit{nums}[i-1]+1$，则把 $\textit{cnt}$ 增加 $1$。
- 否则，把 $\textit{cnt}$ 置为 $1$。\n\n如果 $\textit{cnt}\ge k$，则更新 $\textit{ans}[i-k+1]=\textit{nums}[i]$。

```
func resultsArray(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	for i := range ans {
		ans[i] = -1
	}
	cnt := 0
	for i, x := range nums {
		if i == 0 || x == nums[i-1]+1 {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= k {
			ans[i-k+1] = x
		}
	}
	return ans
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