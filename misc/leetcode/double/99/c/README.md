### 题目

<p>给你一个二维整数数组 <code>ranges</code> ，其中 <code>ranges[i] = [start<sub>i</sub>, end<sub>i</sub>]</code> 表示 <code>start<sub>i</sub></code> 到 <code>end<sub>i</sub></code> 之间（包括二者）的所有整数都包含在第 <code>i</code> 个区间中。</p>

<p>你需要将 <code>ranges</code> 分成 <strong>两个</strong> 组（可以为空），满足：</p>

<ul>
	<li>每个区间只属于一个组。</li>
	<li>两个有 <strong>交集</strong> 的区间必须在 <strong>同一个 </strong>组内。</li>
</ul>

<p>如果两个区间有至少 <strong>一个</strong> 公共整数，那么这两个区间是 <b>有交集</b> 的。</p>

<ul>
	<li>比方说，区间 <code>[1, 3]</code> 和 <code>[2, 5]</code> 有交集，因为 <code>2</code> 和 <code>3</code> 在两个区间中都被包含。</li>
</ul>

<p>请你返回将 <code>ranges</code> 划分成两个组的 <strong>总方案数</strong> 。由于答案可能很大，将它对 <code>10<sup>9</sup> + 7</code> <strong>取余</strong> 后返回。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>ranges = [[6,10],[5,15]]
<b>输出：</b>2
<b>解释：</b>
两个区间有交集，所以它们必须在同一个组内。
所以有两种方案：
- 将两个区间都放在第 1 个组中。
- 将两个区间都放在第 2 个组中。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>ranges = [[1,3],[10,20],[2,5],[4,8]]
<b>输出：</b>4
<b>解释：</b>
区间 [1,3] 和 [2,5] 有交集，所以它们必须在同一个组中。
同理，区间 [2,5] 和 [4,8] 也有交集，所以它们也必须在同一个组中。
所以总共有 4 种分组方案：
- 所有区间都在第 1 组。
- 所有区间都在第 2 组。
- 区间 [1,3] ，[2,5] 和 [4,8] 在第 1 个组中，[10,20] 在第 2 个组中。
- 区间 [1,3] ，[2,5] 和 [4,8] 在第 2 个组中，[10,20] 在第 1 个组中。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= ranges.length <= 10<sup>5</sup></code></li>
	<li><code>ranges[i].length == 2</code></li>
	<li><code>0 <= start<sub>i</sub> <= end<sub>i</sub> <= 10<sup>9</sup></code></li>
</ul>

### 思路

本题和 [56. 合并区间](https://leetcode.cn/problems/merge-intervals/) 是类似的。我们需要把有交集的区间都归在同一个集合中，假设最后分成了 $m$ 个集合，那么每个集合都可以决定要在第一个组还是第二个组，所以方案数为 $2^m$。怎么求出 $m$ 呢？初始化 $m=1$。按照左端点排序，遍历数组，同时维护区间右端点的最大值 $\textit{maxR}$：

- 如果当前区间的左端点大于 $\textit{maxR}$，由于已经按照左端点排序了，那么后面任何区间都不会和之前的区间有交集，换句话说，产生了一个新的集合，$m$ 加一。
- 否则，当前区间与上一个区间在同一个集合。

```go  
const mod int = 1e9 + 7

func countWays(ranges [][]int) (ans int) {
	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })
	ans, maxR := 2, ranges[0][1]
	for _, p := range ranges[1:] {
		if p[0] > maxR { // 产生了一个新的集合
			ans = ans * 2 % mod
		}
		maxR = max(maxR, p[1])
	}
	ans = (ans%mod + mod) % mod
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。忽略排序时的栈开销，仅用到若干额外变量。
