### 题目

<p>给定一个整数数组 <code>nums</code> 和一个整数 <code>k</code>，如果元素 <code>nums[i]</code> <strong>严格</strong> 大于下标&nbsp;<code>i - k</code> 和 <code>i + k</code> 处的元素（如果这些元素存在），则该元素 <code>nums[i]</code> 被认为是 <strong>好</strong> 的。如果这两个下标都不存在，那么 <code>nums[i]</code> 仍然被认为是 <strong>好</strong> 的。</p>

<p>返回数组中所有 <strong>好</strong> 元素的 <strong>和</strong>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,3,2,1,5,4], k = 2</span></p>

<p><strong>输出：</strong> <span class="example-io">12</span></p>

<p><strong>解释：</strong></p>

<p>好的数字包括&nbsp;<code>nums[1] = 3</code>，<code>nums[4] = 5</code> 和 <code>nums[5] = 4</code>，因为它们严格大于下标&nbsp;<code>i - k</code> 和 <code>i + k</code> 处的数字。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [2,1], k = 1</span></p>

<p><strong>输出：</strong> <span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<p>唯一的好数字是 <code>nums[0] = 2</code>，因为它严格大于 <code>nums[1]</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 100</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 1000</code></li>
	<li><code>1 &lt;= k &lt;= floor(nums.length / 2)</code></li>
</ul>

### 思路

遍历 $\textit{nums}$，累加符合要求的元素值。

把 $i-k<0$ 或者 $i+k\ge n$ 的情况当作「符合要求」。

```
func findValidPair(s string) (ans string) {
	cnt := [10]int{}
	for _, b := range s {
		cnt[b-'0']++
	}
	for i := 1; i < len(s); i++ {
		x, y := s[i-1]-'0', s[i]-'0'
		if x != y && cnt[x] == int(x) && cnt[y] == int(y) {
			return s[i-1 : i+1]
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
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
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)