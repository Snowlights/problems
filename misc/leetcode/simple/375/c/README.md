#### 题目

<p>给你一个整数数组 <code>nums</code> 和一个 <strong>正整数</strong> <code>k</code> 。</p>

<p>请你统计有多少满足 「&nbsp;<code>nums</code> 中的 <strong>最大</strong> 元素」至少出现 <code>k</code> 次的子数组，并返回满足这一条件的子数组的数目。</p>

<p>子数组是数组中的一个连续元素序列。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,3,2,3,3], k = 2
<strong>输出：</strong>6
<strong>解释：</strong>包含元素 3 至少 2 次的子数组为：[1,3,2,3]、[1,3,2,3,3]、[3,2,3]、[3,2,3,3]、[2,3,3] 和 [3,3] 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,4,2,1], k = 3
<strong>输出：</strong>0
<strong>解释：</strong>没有子数组包含元素 4 至少 3 次。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>6</sup></code></li>
	<li><code>1 &lt;= k &lt;= 10<sup>5</sup></code></li>
</ul>

#### 思路

本题算法如下：
1. 设 $\textit{mx} = \max(\textit{nums})$。
2. 右端点 $\textit{right}$ 从左到右遍历 $\textit{nums}$。遍历到元素 $x=\textit{nums}[\textit{right}]$ 时，如果 $x=\textit{mx}$，就把计数器 $\textit{cntMx}$ 加一。
3. 如果此时 $\textit{cntMx}=k$，则不断右移左指针 $\textit{left}$，直到窗口内的 $\textit{mx}$ 的出现次数**小于** $k$ 为止。此时，对于右端点为 $\textit{right}$ 且左端点小于 $\textit{left}$ 的子数组，$\textit{mx}$ 的出现次数都至少为 $k$，把答案增加 $\textit{left}$。

例如示例 1，当右端点移到第二个 $3$ 时，左端点移到 $2$，此时 $[1,3,2,3]$ 和 $[3,2,3]$ 是满足要求的。
当右端点移到第三个 $3$ 时，左端点也移到第三个 $3$，此时 $[1,3,2,3,3], [3,2,3,3], [2,3,3], [3,3]$ 都是满足要求的。
所以答案为 $2+4=6$。

```go  
func countSubarrays(nums []int, k int) int64 {
	mx := nums[0]
	for _, v := range nums {
		mx = max(mx, v)
	}
	l, ans, h := 0, 0, make(map[int]int)
	for _, v := range nums {
		h[v]++
		for h[mx] >= k {
			h[nums[l]]--
			l++
		}
		ans += l
	}
	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。