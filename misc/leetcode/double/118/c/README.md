### 题目

<p>你在一个水果超市里，货架上摆满了玲琅满目的奇珍异果。</p>

<p>给你一个下标从 <strong>1</strong>&nbsp;开始的数组&nbsp;<code>prices</code>&nbsp;，其中&nbsp;<code>prices[i]</code>&nbsp;表示你购买第 <code>i</code>&nbsp;个水果需要花费的金币数目。</p>

<p>水果超市有如下促销活动：</p>

<ul>
	<li>如果你花费 <code>price[i]</code>&nbsp;购买了水果&nbsp;<code>i</code>&nbsp;，那么接下来的 <code>i</code>&nbsp;个水果你都可以免费获得。</li>
</ul>

<p><strong>注意</strong>&nbsp;，即使你&nbsp;<strong>可以</strong>&nbsp;免费获得水果&nbsp;<code>j</code>&nbsp;，你仍然可以花费&nbsp;<code>prices[j]</code>&nbsp;个金币去购买它以便能免费获得接下来的 <code>j</code>&nbsp;个水果。</p>

<p>请你返回获得所有水果所需要的 <strong>最少</strong>&nbsp;金币数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>prices = [3,1,2]
<b>输出：</b>4
<b>解释</b><strong>：</strong>你可以按如下方法获得所有水果：
- 花 3 个金币购买水果 1 ，然后免费获得水果 2 。
- 花 1 个金币购买水果 2 ，然后免费获得水果 3 。
- 免费获得水果 3 。
注意，虽然你可以免费获得水果 2 ，但你还是花 1 个金币去购买它，因为这样的总花费最少。
购买所有水果需要最少花费 4 个金币。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>prices = [1,10,1,1]
<b>输出：</b>2
<b>解释：</b>你可以按如下方法获得所有水果：
- 花 1 个金币购买水果 1 ，然后免费获得水果 2 。
- 免费获得水果 2 。
- 花 1 个金币购买水果 3 ，然后免费获得水果 4 。
- 免费获得水果 4 。
购买所有水果需要最少花费 2 个金币。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= prices.length &lt;= 1000</code></li>
	<li><code>1 &lt;= prices[i] &lt;= 10<sup>5</sup></code></li>
</ul>

### 思路

## 一、启发思考：寻找子问题

我们需要解决的问题是：「获得第 $1$ 个及其后面的水果所需要的最少金币数」。

第 $1$ 个水果一定要买，然后呢？

第 $2$ 个水果可以购买，也可以免费获得：

- 如果购买，那么需要解决的问题为：「获得第 $2$ 个及其后面的水果所需要的最少金币数」。
- 如果免费获得，那么需要解决的问题为：「获得第 $3$ 个及其后面的水果所需要的最少金币数」。

无论哪种情况都会把原问题变成一个**和原问题相似的、规模更小的子问题**，所以可以用**递归**解决。

## 二、状态定义与状态转移方程

从上面的讨论可以知道，只需要一个 $i$ 就能表达子问题，即定义 $\textit{dfs}(i)$ 表示获得第 $i$ 个及其后面的水果所需要的最少金币数。注意 $i$ 从 $1$ 开始。

第 $i$ 个水果一定要买，那么从 $i+1$ 到 $2i$ 的水果都是免费的。枚举下一个需要**购买**的水果 $j$，它的范围是 $[i+1,2i+1]$。

所有情况取最小值，再加上购买第 $i$ 个水果需要的金币 $\textit{prices}[i]$，就得到了 $\textit{dfs}(i)$。写成式子就是

$$
\textit{dfs}(i) = \textit{prices}[i] + \min_{j=i+1}^{2i+1} \textit{dfs}(j)
$$

注意到当 $2i\ge n$，即 $i\ge \left\lceil\dfrac{n}{2}\right\rceil = \left\lfloor\dfrac{n+1}{2}\right\rfloor$ 时，后面的水果都可以免费获得了，所以递归边界为

$$
\textit{dfs}(i)=\textit{prices}[i]
$$

其中 $i\ge \left\lfloor\dfrac{n+1}{2}\right\rfloor$。

递归入口：$\textit{dfs}(1)$，也就是答案。

由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

```go  
func minimumCoins(prices []int) int {
	n := len(prices)
	memo := make([]int, (n+1)/2)
	var dfs func(int) int
	dfs = func(i int) int {
		if i*2 >= n {
			return prices[i-1] // i 从 1 开始
		}
		p := &memo[i]
		if *p != 0 { // 之前算过
			return *p
		}
		res := math.MaxInt
		for j := i + 1; j <= i*2+1; j++ {
			res = min(res, dfs(j))
		}
		res += prices[i-1]
		*p = res // 记忆化
		return res
	}
	return dfs(1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{prices}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(n)$。有多少个状态，$\textit{memo}$ 数组的大小就是多少。


```go  
func minimumCoins(prices []int) int {
	n := len(prices)
	for i := (n+1)/2 - 1; i > 0; i-- {
		prices[i-1] += slices.Min(prices[i : i*2+1])
	}
	return prices[0]
}

```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{prices}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。Python 忽略切片开销。
