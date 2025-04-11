### 题目

<p>给你一个下标从 <strong>0</strong>&nbsp;开始的非负整数数组&nbsp;<code>nums</code>&nbsp;和两个整数&nbsp;<code>l</code> 和&nbsp;<code>r</code>&nbsp;。</p>

<p>请你返回&nbsp;<code>nums</code>&nbsp;中子多重集合的和在闭区间&nbsp;<code>[l, r]</code>&nbsp;之间的 <strong>子多重集合的数目</strong> 。</p>

<p>由于答案可能很大，请你将答案对&nbsp;<code>10<sup>9 </sup>+ 7</code>&nbsp;取余后返回。</p>

<p><strong>子多重集合</strong> 指的是从数组中选出一些元素构成的 <strong>无序</strong>&nbsp;集合，每个元素 <code>x</code>&nbsp;出现的次数可以是&nbsp;<code>0, 1, ..., occ[x]</code>&nbsp;次，其中&nbsp;<code>occ[x]</code>&nbsp;是元素&nbsp;<code>x</code>&nbsp;在数组中的出现次数。</p>

<p><b>注意：</b></p>

<ul>
	<li>如果两个子多重集合中的元素排序后一模一样，那么它们两个是相同的&nbsp;<strong>子多重集合</strong>&nbsp;。</li>
	<li><strong>空</strong>&nbsp;集合的和是 <code>0</code>&nbsp;。</li>
</ul>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [1,2,2,3], l = 6, r = 6
<b>输出：</b>1
<b>解释：</b>唯一和为 6 的子集合是 {1, 2, 3} 。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [2,1,4,2,7], l = 1, r = 5
<b>输出：</b>7
<b>解释：</b>和在闭区间 [1, 5] 之间的子多重集合为 {1} ，{2} ，{4} ，{2, 2} ，{1, 2} ，{1, 4} 和 {1, 2, 2} 。
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<b>输入：</b>nums = [1,2,1,3,5,2], l = 3, r = 5
<b>输出：</b>9
<b>解释：</b>和在闭区间 [3, 5] 之间的子多重集合为 {3} ，{5} ，{1, 2} ，{1, 3} ，{2, 2} ，{2, 3} ，{1, 1, 2} ，{1, 1, 3} 和 {1, 2, 2} 。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 2 * 10<sup>4</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 2 * 10<sup>4</sup></code></li>
	<li><code>nums</code> 的和不超过&nbsp;<code>2 * 10<sup>4</sup></code> 。</li>
	<li><code>0 &lt;= l &lt;= r &lt;= 2 * 10<sup>4</sup></code></li>
</ul>

### 思路

## 朴素 DP

注意到相同数字是不作区分的，所以本题属于**多重背包**。

> 如果区分相同数字就是 0-1 背包。

用哈希表统计每个数的出现次数，记在 $\textit{cnt}$ 中。

定义 $f[i+1][j]$ 表示从 $\textit{cnt}$ 的前 $i$ 种数中选择一些数，元素和**恰好**为 $j$ 的方案数。

设第 $i$ 种数的值为 $x$。

枚举第 $i$ 种数选了 $k$ 个，则有

$$
f[i+1][j] = \sum_{k=0}^{\textit{cnt}[x]} f[i][j-kx]
$$

其中 $j-kx\ge 0$。

初始值 $f[0][0] = \textit{cnt}[0] + 1$，表示「在什么也不选的情况下，元素和为 $0$」有 $\textit{cnt}[0] + 1$ 个方案。

答案为 

$$
\sum_{i=l}^r f[m][i]
$$

其中 $m$ 是 $\textit{cnt}$ 的大小（不包括 $0$）。

## 优化方法一：式子变形+滚动数组

举例说明，假设 $x=2$ 并且 $\textit{cnt}[x]=3$，那么选 $0,1,2,3$ 个 $x$ 都是可以的，即

$$
f[i+1][j-2] = f[i][j-2] + f[i][j-4] + f[i][j-6] + f[i][j-8]
$$

所以

$$
\begin{aligned}
f[i+1][j] =\ & f[i][j] + f[i][j-2] + f[i][j-4] + f[i][j-6]\\
=\ &f[i][j] + (f[i+1][j-2] - f[i][j-8])\\
=\ &f[i+1][j-2] + f[i][j] - f[i][j-8]\\
\end{aligned}
$$

一般地，我们有

$$
f[i+1][j] = f[i+1][j-x] + f[i][j] - f[i][j-(\textit{cnt}[x]+1)\cdot x]
$$

如果 $j-(\textit{cnt}[x]+1)\cdot x < 0$ 则

$$
f[i+1][j] = f[i+1][j-x] + f[i][j]
$$

这样就可以 $\mathcal{O}(1)$ 地算出每个 $f[i][j]$ 了。

代码实现时，可以用滚动数组优化。

```
func countSubMultisets(nums []int, l, r int) (ans int) {
	const mod = 1_000_000_007
	total := 0
	cnt := map[int]int{}
	for _, x := range nums {
		total += x
		cnt[x]++
	}
	if l > total {
		return
	}

	r = min(r, total)
	f := make([]int, r+1)
	f[0] = cnt[0] + 1
	delete(cnt, 0)

	sum := 0
	for x, c := range cnt {
		newF := append([]int{}, f...)
		sum = min(sum+x*c, r) // 到目前为止，能选的元素和至多为 sum
		for j := x; j <= sum; j++ { // 把循环上界从 r 改成 sum 可以快不少
			newF[j] += newF[j-x]
			if j >= (c+1)*x {
				newF[j] -= f[j-(c+1)*x] // 注意这里有减法，可能产生负数
			}
			newF[j] %= mod
		}
		f = newF
	}

	for _, v := range f[l:] {
		ans += v
	}
	return (ans%mod + mod) % mod // 调整成非负数
}

func min(a, b int) int { if b < a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(S\sqrt{\min(S,n)})$，其中 $n$ 为 $\textit{nums}$ 的长度，$S$ 为 $\textit{nums}$ 的元素和。由 $1+2+\cdots+m=\dfrac{m(m+1)}{2}\le S$ 可知，$\textit{nums}$ 中至多有 $\mathcal{O}(\sqrt S)$ 个**不同**元素，所以哈希表的大小至多为 $\mathcal{O}(\sqrt S)$。当然，哈希表的大小不会超过 $n$。
- 空间复杂度：$\mathcal{O}(S)$。

## 优化方法二：同余前缀和优化

再看一眼这个式子

$$
f[i+1][j] = \sum_{k=0}^{\textit{cnt}[x]} f[i][j-kx]
$$

如果求出 $f[i]$ 的**同余前缀和**，那么 $f[i+1][j]$ 就可以转换成两个同余前缀和的差了。

同余前缀和指 $s[p] = f[i][p] + f[i][p-x] + f[i][p-2x] +\cdots$。

代码实现时，可以只用一个一维数组。

```go [sol-Go]
func countSubMultisets(nums []int, l, r int) (ans int) {
	const mod = 1_000_000_007
	total := 0
	cnt := map[int]int{}
	for _, x := range nums {
		total += x
		cnt[x]++
	}
	if l > total {
		return
	}

	r = min(r, total)
	f := make([]int, r+1)
	f[0] = cnt[0] + 1
	delete(cnt, 0)

	sum := 0
	for x, c := range cnt {
		sum = min(sum+x*c, r)
		for j := x; j <= sum; j++ {
			f[j] = (f[j] + f[j-x]) % mod // 同余前缀和
		}
		for j := sum; j >= x*(c+1); j-- {
			f[j] = (f[j] - f[j-x*(c+1)]) % mod // 两个同余前缀和的差
		}
	}

	for _, v := range f[l:] {
		ans += v
	}
	return (ans%mod + mod) % mod // 调整成非负数
}

func min(a, b int) int { if b < a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(S\sqrt{\min(S,n)})$，其中 $n$ 为 $\textit{nums}$ 的长度，$S$ 为 $\textit{nums}$ 的元素和。由 $1+2+\cdots+m=\dfrac{m(m+1)}{2}\le S$ 可知，$\textit{nums}$ 中至多有 $\mathcal{O}(\sqrt S)$ 个**不同**元素，所以哈希表的大小至多为 $\mathcal{O}(\sqrt S)$。当然，哈希表的大小不会超过 $n$。
- 空间复杂度：$\mathcal{O}(S)$。

