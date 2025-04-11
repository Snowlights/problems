### 题目

<p>给你两个长度相同的字符串&nbsp;<code>s</code> 和&nbsp;<code>t</code>&nbsp;，以及两个整数数组&nbsp;<code>nextCost</code> 和&nbsp;<code>previousCost</code>&nbsp;。</p>

<p>一次操作中，你可以选择 <code>s</code>&nbsp;中的一个下标 <code>i</code>&nbsp;，执行以下操作 <strong>之一</strong>&nbsp;：</p>

<ul>
	<li>将&nbsp;<code>s[i]</code>&nbsp;切换为字母表中的下一个字母，如果&nbsp;<code>s[i] == 'z'</code>&nbsp;，切换后得到&nbsp;<code>'a'</code>&nbsp;。操作的代价为&nbsp;<code>nextCost[j]</code>&nbsp;，其中&nbsp;<code>j</code>&nbsp;表示&nbsp;<code>s[i]</code>&nbsp;在字母表中的下标。</li>
	<li>将&nbsp;<code>s[i]</code>&nbsp;切换为字母表中的上一个字母，如果&nbsp;<code>s[i] == 'a'</code>&nbsp;，切换后得到&nbsp;<code>'z'</code>&nbsp;。操作的代价为&nbsp;<code>previousCost[j]</code>&nbsp;，其中&nbsp;<code>j</code> 是&nbsp;<code>s[i]</code>&nbsp;在字母表中的下标。</li>
</ul>

<p><strong>切换距离</strong>&nbsp;指的是将字符串 <code>s</code>&nbsp;变为字符串 <code>t</code>&nbsp;的 <strong>最少</strong>&nbsp;操作代价总和。</p>

<p>请你返回从 <code>s</code>&nbsp;到 <code>t</code>&nbsp;的 <strong>切换距离</strong>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "abab", t = "baba", nextCost = [100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0], previousCost = [1,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><b>解释：</b></p>

<ul>
	<li>选择下标&nbsp;<code>i = 0</code>&nbsp;并将&nbsp;<code>s[0]</code>&nbsp;向前切换 25 次，总代价为 1 。</li>
	<li>选择下标&nbsp;<code>i = 1</code>&nbsp;并将&nbsp;<code>s[1]</code>&nbsp;向后切换 25 次，总代价为 0 。</li>
	<li>选择下标&nbsp;<code>i = 2</code>&nbsp;并将&nbsp;<code>s[2]</code>&nbsp;向前切换 25 次，总代价为 1 。</li>
	<li>选择下标&nbsp;<code>i = 3</code>&nbsp;并将&nbsp;<code>s[3]</code>&nbsp;向后切换 25 次，总代价为 0 。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "leet", t = "code", nextCost = [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], previousCost = [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]</span></p>

<p><span class="example-io"><b>输出：</b>31</span></p>

<p><b>解释：</b></p>

<ul>
	<li>选择下标&nbsp;<code>i = 0</code>&nbsp;并将&nbsp;<code>s[0]</code>&nbsp;向前切换 9 次，总代价为 9 。</li>
	<li>选择下标&nbsp;<code>i = 1</code>&nbsp;并将&nbsp;<code>s[1]</code>&nbsp;向后切换 10 次，总代价为 10 。</li>
	<li>选择下标&nbsp;<code>i = 2</code> 并将&nbsp;<code>s[2]</code>&nbsp;向前切换 1 次，总代价为 1 。</li>
	<li>选择下标 <code>i = 3</code> 并将&nbsp;<code>s[3]</code>&nbsp;向后切换 11 次，总代价为 11 。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length == t.length &lt;= 10<sup>5</sup></code></li>
	<li><code>s</code> 和&nbsp;<code>t</code>&nbsp;都只包含小写英文字母。</li>
	<li><code>nextCost.length == previousCost.length == 26</code></li>
	<li><code>0 &lt;= nextCost[i], previousCost[i] &lt;= 10<sup>9</sup></code></li>
</ul>

### 思路

我们可以预处理 $\textit{nextCost}$ 和 $\textit{previousCost}$ 的 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) $\textit{nxtSum}$ 和 $\textit{preSum}$，从而加速操作代价和的计算。

考虑到字母表是环形的，可以把前缀和数组延长一倍，从而变成非环形的。

```
func shiftDistance(s, t string, nextCost, previousCost []int) (ans int64) {
	const sigma = 26
	var nxtSum, preSum [sigma*2 + 1]int64
	for i := 0; i < sigma*2; i++ {
		nxtSum[i+1] = nxtSum[i] + int64(nextCost[i%sigma])
		preSum[i+1] = preSum[i] + int64(previousCost[i%sigma])
	}
	for i := range s {
		x := s[i] - 'a'
		y := t[i] - 'a'
		if y < x {
			y += sigma
		}
		res1 := nxtSum[y] - nxtSum[x]
		y = t[i] - 'a'
		if x < y {
			x += sigma
		}
		res2 := preSum[x+1] - preSum[y+1]
		ans += min(res1, res2)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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