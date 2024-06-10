#### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> ，它包含 <code>n</code> 个 <strong>互不相同</strong> 的正整数。如果 <code>nums</code> 的一个排列满足以下条件，我们称它是一个特别的排列：</p>

<ul>
	<li>对于 <code>0 &lt;= i &lt; n - 1</code> 的下标 <code>i</code> ，要么 <code>nums[i] % nums[i+1] == 0</code> ，要么 <code>nums[i+1] % nums[i] == 0</code> 。</li>
</ul>

<p>请你返回特别排列的总数目，由于答案可能很大，请将它对<strong> </strong><code>10<sup>9 </sup>+ 7</code> <strong>取余</strong> 后返回。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [2,3,6]
<b>输出：</b>2
<b>解释：</b>[3,6,2] 和 [2,6,3] 是 nums 两个特别的排列。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [1,4,3]
<b>输出：</b>2
<b>解释：</b>[3,1,4] 和 [4,1,3] 是 nums 两个特别的排列。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 14</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>
 
#### 思路  

定义 $\textit{dfs}(i,j)$ 表示当前可以选的下标集合为 $i$，上一个选的数的下标是 $j$ 时，可以构造出多少个特别排列。

枚举当前要选的数的下标 $k$，如果 $\textit{nums}[k]$ 与 $\textit{nums}[j]$ 满足题目整除的要求，则

$$
\textit{dfs}(i,j) = \sum_{k\in i} \textit{dfs}(i\setminus \{k\},k)
$$

递归边界：$\textit{dfs}(0,j) = 1$，表示找到了一个特别排列。

递归入口：$\textit{dfs}(U\setminus \{j\},j)$，其中全集 $U=\{0,1,2,\cdots,n-1\}$。枚举特别排列的第一个数的下标 $j$，累加所有 $\textit{dfs}(U\setminus \{j\},j)$，即为答案。


```go 
func specialPerm(a []int) (ans int) {
	n := len(a)
	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i == 0 {
			return 1
		}
		dv := &dp[i][j]
		if *dv >= 0 {
			return *dv
		}
		defer func() {
			*dv = res
		}()
		for k, x := range a {
			if i>>k&1 > 0 && (a[j]%x == 0 || x%a[j] == 0) {
				res = (res + dfs(i^(1<<k), k)) % mod
			}
		}
		return
	}

	for i := range a {
		ans = (ans + dfs(((1<<n)-1)^(1<<i), i)%mod) % mod
	}

	ans = (ans%mod + mod) % mod
	return
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n^22^n)$，其中 $n$ 为 $\textit{nums}$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n2^n)$，单个状态的计算时间为 $\mathcal{O}(n)$，因此时间复杂度为 $\mathcal{O}(n^22^n)$。
- 空间复杂度：$\mathcal{O}(n2^n)$。