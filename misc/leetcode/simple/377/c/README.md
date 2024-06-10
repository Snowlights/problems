#### 题目

<p>给你两个下标从 <strong>0</strong> 开始的字符串 <code>source</code> 和 <code>target</code> ，它们的长度均为 <code>n</code> 并且由 <strong>小写 </strong>英文字母组成。</p>

<p>另给你两个下标从 <strong>0</strong> 开始的字符数组 <code>original</code> 和 <code>changed</code> ，以及一个整数数组 <code>cost</code> ，其中 <code>cost[i]</code> 代表将字符 <code>original[i]</code> 更改为字符 <code>changed[i]</code> 的成本。</p>

<p>你从字符串 <code>source</code> 开始。在一次操作中，<strong>如果 </strong>存在 <strong>任意</strong> 下标 <code>j</code> 满足 <code>cost[j] == z</code>  、<code>original[j] == x</code> 以及 <code>changed[j] == y</code> 。你就可以选择字符串中的一个字符 <code>x</code> 并以 <code>z</code> 的成本将其更改为字符 <code>y</code> 。</p>

<p>返回将字符串 <code>source</code> 转换为字符串 <code>target</code> 所需的<strong> 最小 </strong>成本。如果不可能完成转换，则返回 <code>-1</code> 。</p>

<p><strong>注意</strong>，可能存在下标 <code>i</code> 、<code>j</code> 使得 <code>original[j] == original[i]</code> 且 <code>changed[j] == changed[i]</code> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>source = "abcd", target = "acbe", original = ["a","b","c","c","e","d"], changed = ["b","c","b","e","b","e"], cost = [2,5,5,1,2,20]
<strong>输出：</strong>28
<strong>解释：</strong>将字符串 "abcd" 转换为字符串 "acbe" ：
- 更改下标 1 处的值 'b' 为 'c' ，成本为 5 。
- 更改下标 2 处的值 'c' 为 'e' ，成本为 1 。
- 更改下标 2 处的值 'e' 为 'b' ，成本为 2 。
- 更改下标 3 处的值 'd' 为 'e' ，成本为 20 。
产生的总成本是 5 + 1 + 2 + 20 = 28 。
可以证明这是可能的最小成本。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>source = "aaaa", target = "bbbb", original = ["a","c"], changed = ["c","b"], cost = [1,2]
<strong>输出：</strong>12
<strong>解释：</strong>要将字符 'a' 更改为 'b'：
- 将字符 'a' 更改为 'c'，成本为 1 
- 将字符 'c' 更改为 'b'，成本为 2 
产生的总成本是 1 + 2 = 3。
将所有 'a' 更改为 'b'，产生的总成本是 3 * 4 = 12 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>source = "abcd", target = "abce", original = ["a"], changed = ["e"], cost = [10000]
<strong>输出：</strong>-1
<strong>解释：</strong>无法将 source 字符串转换为 target 字符串，因为下标 3 处的值无法从 'd' 更改为 'e' 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= source.length == target.length <= 10<sup>5</sup></code></li>
	<li><code>source</code>、<code>target</code> 均由小写英文字母组成</li>
	<li><code>1 <= cost.length== original.length == changed.length <= 2000</code></li>
	<li><code>original[i]</code>、<code>changed[i]</code> 是小写英文字母</li>
	<li><code>1 <= cost[i] <= 10<sup>6</sup></code></li>
	<li><code>original[i] != changed[i]</code></li>
</ul>

#### 思路

建图，从 $\textit{original}[i]$ 向 $\textit{changed}[i]$ 连边，边权为 $\textit{cost}[i]$。
然后用 Floyd 算法求图中任意两点最短路，得到 $\textit{dis}$ 矩阵，这里得到的 $\textit{dis}[i][j]$ 表示字母 $i$ 通过若干次替换操作变成字母 $j$ 的最小成本。
最后累加所有 $\textit{dis}[\textit{original}[i]][\textit{changed}[i]]$，即为答案。如果答案为无穷大，返回 $-1$。

```go  [sol]
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	dis := make([][]int, 26)
	for i := range dis {
		dis[i] = make([]int, 26)
		for j := range dis[i] {
			if j != i {
				dis[i][j] = 1e12
			}
		}
	}

	for i := range original {
		c, d := int(original[i]-'a'), int(changed[i]-'a')
		dis[c][d] = min(dis[c][d], cost[i])
	}

	for k := range dis {
		for i := range dis {
			for j := range dis {
				// 不选 k，选 k
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}
	ans := 0
	for i := range source {
		if source[i] != target[i] {
			ans += dis[int(source[i]-'a')][(target[i] - 'a')]
		}
	}
	if ans >= 1e12 {
		return -1
	}
	return int64(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m+|\Sigma|^3)$，其中 $n$ 为 $\textit{source}$ 的长度，$m$ 为 $\textit{cost}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|^2)$。

#### 相似题目

- [2642. 设计可以求最短路径的图类](https://leetcode.cn/problems/design-graph-with-shortest-path-calculator/) 1811
- [1334. 阈值距离内邻居最少的城市](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/) 1855
- [2101. 引爆最多的炸弹](https://leetcode.cn/problems/detonate-the-maximum-bombs/) 1880
