### 题目

<p>给你一个字符串&nbsp;<code>s</code>&nbsp;，你需要将它分割成一个或者更多的&nbsp;<strong>平衡</strong>&nbsp;子字符串。比方说，<code>s == "ababcc"</code>&nbsp;那么&nbsp;<code>("abab", "c", "c")</code>&nbsp;，<code>("ab", "abc", "c")</code>&nbsp;和&nbsp;<code>("ababcc")</code>&nbsp;都是合法分割，但是&nbsp;<code>("a", <strong>"bab"</strong>, "cc")</code>&nbsp;，<code>(<strong>"aba"</strong>, "bc", "c")</code>&nbsp;和&nbsp;<code>("ab", <strong>"abcc"</strong>)</code>&nbsp;不是，不平衡的子字符串用粗体表示。</p>

<p>请你返回 <code>s</code>&nbsp;<strong>最少</strong> 能分割成多少个平衡子字符串。</p>

<p><b>注意：</b>一个 <strong>平衡</strong>&nbsp;字符串指的是字符串中所有字符出现的次数都相同。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "fabccddg"</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><strong>解释：</strong></p>

<p>我们可以将 <code>s</code>&nbsp;分割成 3 个子字符串：<code>("fab, "ccdd", "g")</code>&nbsp;或者&nbsp;<code>("fabc", "cd", "dg")</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "abababaccddb"</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p>我们可以将&nbsp;<code>s</code>&nbsp;分割成 2 个子字符串：<code>("abab", "abaccddb")</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 1000</code></li>
	<li><code>s</code>&nbsp;只包含小写英文字母。</li>
</ul>


### 思路

划分型 DP **固定套路**，见 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的「§6.2 计算划分个数」。

注意本题答案是存在的，因为单个字母是平衡的，我们一定可以划分成 $n$ 个子串。

## 方法一：记忆化搜索

定义 $\textit{dfs}(i)$ 表示划分前缀 $s[0]$ 到 $s[i]$ 的最小划分个数，则有

$$
\textit{dfs}(i) = \min_{j=0}^{i} \textit{dfs}(j-1) + 1
$$

其中 $s[j]$ 到 $s[i]$ 是平衡子串。

我们可以在倒序枚举 $j$ 的同时，用一个哈希表（或者数组）统计每个字符的出现次数。如果子串中每个字母的出现次数都相等，那么子串是平衡的。

**优化**：设子串中有 $k$ 种字母，如果子串长度不是 $k$ 的倍数，那么子串一定不是平衡的。

递归边界：$\textit{dfs}(-1) = 0$。

递归入口：$\textit{dfs}(n-1)$，即答案。

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。本题由于 $\textit{dfs}(i) > 0$，可以初始化成 $0$。

```
func minimumSubstringsInPartition(s string) int {
	n := len(s)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1 // -1 表示没有计算过
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := math.MaxInt
		cnt := [26]int{}
		k := 0
	next:
		for j := i; j >= 0; j-- {
			b := s[j] - 'a'
			if cnt[b] == 0 {
				k++
			}
			cnt[b]++
			if (i-j+1)%k > 0 {
				continue
			}
			for _, c := range cnt {
				if c > 0 && c != cnt[b] {
					continue next
				}
			}
			res = min(res, dfs(j-1)+1)
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n - 1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2|\Sigma|)$，其中 $n$ 为 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n|\Sigma|)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^2|\Sigma|)$。
- 空间复杂度：$\mathcal{O}(n|\Sigma|)$。注意递归中至多会创建 $n$ 个长为 $|\Sigma|$ 的 $\textit{cnt}$ 数组。

## 方法二：递推（1:1 翻译）

定义 $f[i+1]$ 表示划分前缀 $s[0]$ 到 $s[i]$ 的最小划分个数，则有

$$
f[i+1] = \min_{j=0}^{i}f[j] + 1
$$

其中 $s[j]$ 到 $s[i]$ 是平衡子串。

初始值 $f[0]= 0$，翻译自递归边界 $\textit{dfs}(-1) = 0$。

答案为 $f[n]$，翻译自递归入口 $\textit{dfs}(n-1)$。

```
func minimumSubstringsInPartition(s string) int {
	n := len(s)
	f := make([]int, n+1)
	for i := range s {
		f[i+1] = math.MaxInt
		cnt := [26]int{}
		k := 0
	next:
		for j := i; j >= 0; j-- {
			b := s[j] - 'a'
			if cnt[b] == 0 {
				k++
			}
			cnt[b]++
			if (i-j+1)%k > 0 {
				continue
			}
			for _, c := range cnt {
				if c != 0 && c != cnt[b] {
					continue next
				}
			}
			f[i+1] = min(f[i+1], f[j]+1)
		}
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2|\Sigma|)$，其中 $n$ 为 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(n + |\Sigma|)$。

## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)