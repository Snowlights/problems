### 题目

<p>给你一个长度为 <code>n</code>&nbsp;的整数数组&nbsp;<code>nums</code>&nbsp;和一个 <strong>正</strong>&nbsp;整数&nbsp;<code>k</code>&nbsp;。</p>

<p>一个子序列的 <strong>能量</strong>&nbsp;定义为子序列中&nbsp;<strong>任意</strong>&nbsp;两个元素的差值绝对值的 <strong>最小值</strong>&nbsp;。</p>

<p>请你返回 <code>nums</code>&nbsp;中长度 <strong>等于</strong>&nbsp;<code>k</code>&nbsp;的 <strong>所有</strong>&nbsp;子序列的 <strong>能量和</strong>&nbsp;。</p>

<p>由于答案可能会很大，将答案对&nbsp;<code>10<sup>9 </sup>+ 7</code>&nbsp;<strong>取余</strong>&nbsp;后返回。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,2,3,4], k = 3</span></p>

<p><span class="example-io"><b>输出：</b>4</span></p>

<p><strong>解释：</strong></p>

<p><code>nums</code>&nbsp;中总共有 4 个长度为 3 的子序列：<code>[1,2,3]</code>&nbsp;，<code>[1,3,4]</code>&nbsp;，<code>[1,2,4]</code>&nbsp;和&nbsp;<code>[2,3,4]</code>&nbsp;。能量和为 <code>|2 - 3| + |3 - 4| + |2 - 1| + |3 - 4| = 4</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [2,2], k = 2</span></p>

<p><span class="example-io"><b>输出：</b>0</span></p>

<p><strong>解释：</strong></p>

<p><code>nums</code>&nbsp;中唯一一个长度为 2 的子序列是&nbsp;<code>[2,2]</code>&nbsp;。能量和为&nbsp;<code>|2 - 2| = 0</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [4,3,-1], k = 2</span></p>

<p><span class="example-io"><b>输出：</b>10</span></p>

<p><strong>解释：</strong></p>

<p><code>nums</code>&nbsp;总共有 3 个长度为 2 的子序列：<code>[4,3]</code>&nbsp;，<code>[4,-1]</code>&nbsp;和&nbsp;<code>[3,-1]</code>&nbsp;。能量和为&nbsp;<code>|4 - 3| + |4 - (-1)| + |3 - (-1)| = 10</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n == nums.length &lt;= 50</code></li>
	<li><code>-10<sup>8</sup> &lt;= nums[i] &lt;= 10<sup>8</sup> </code></li>
	<li><code>2 &lt;= k &lt;= n</code></li>
</ul>

### 思路

选或不选，找最小值加起来，记忆化搜索

```
func sumOfPowers(nums []int, k int) int {
	const (
		MOD int = 1e9 + 7
		INF int = 1e18
	)

	n := len(nums)
	sort.Ints(nums)
	type pair struct {
		i, rem, lst, mn int
	}

	dp := make(map[pair]int, n+10)
	var dfs func(i, rem, lst, mn int) int
	dfs = func(i, rem, lst, mn int) int {
		if n-i < rem {
			return 0
		}
		if rem == 0 {
			return mn
		}
		p := pair{
			i:   i,
			rem: rem,
			lst: lst,
			mn:  mn,
		}
		if res, ok := dp[p]; ok {
			return res
		}
		res := dfs(i+1, rem-1, nums[i], min(mn, nums[i]-lst)) % MOD
		res = (res + dfs(i+1, rem, lst, mn)) % MOD
		dp[p] = res
		return res
	}

	return dfs(0, k, -INF, INF) % MOD
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^4 k)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^3 k)$。

## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
