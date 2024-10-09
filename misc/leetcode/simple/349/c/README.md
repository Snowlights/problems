#### 题目  

<p>给你一个长度为 <code>n</code> 、下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> ，表示收集不同巧克力的成本。每个巧克力都对应一个不同的类型，最初，位于下标 <code>i</code> 的巧克力就对应第 <code>i</code> 个类型。</p>

<p>在一步操作中，你可以用成本 <code>x</code> 执行下述行为：</p>

<ul>
	<li>同时修改所有巧克力的类型，将巧克力的类型 <code>i<sup>th</sup></code> 修改为类型 <code>((i + 1) mod n)<sup>th</sup></code>。</li>
</ul>

<p>假设你可以执行任意次操作，请返回收集所有类型巧克力所需的最小成本。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [20,1,15], x = 5
<strong>输出：</strong>13
<strong>解释：</strong>最开始，巧克力的类型分别是 [0,1,2] 。我们可以用成本 1 购买第 1 个类型的巧克力。
接着，我们用成本 5 执行一次操作，巧克力的类型变更为 [1,2,0] 。我们可以用成本 1 购买第 2 个类型的巧克力。
然后，我们用成本 5 执行一次操作，巧克力的类型变更为 [2,0,1] 。我们可以用成本 1 购买第 0 个类型的巧克力。
因此，收集所有类型的巧克力需要的总成本是 (1 + 5 + 1 + 5 + 1) = 13 。可以证明这是一种最优方案。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [1,2,3], x = 4
<strong>输出：</strong>6
<strong>解释：</strong>我们将会按最初的成本收集全部三个类型的巧克力，而不需执行任何操作。因此，收集所有类型的巧克力需要的总成本是 1 + 2 + 3 = 6 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 1000</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= x &lt;= 10<sup>9</sup></code></li>
</ul>
 
#### 思路  

#### 提示 1

枚举操作次数。

#### 提示 2

如果不操作，第 $i$ 个巧克力必须花费 $\textit{nums}[i]$ 收集，总成本为所有 $\textit{nums}[i]$ 之和。

如果操作一次，第 $i$ 个巧克力可以花费 $\min(\textit{nums}[i], \textit{nums}[(i+1)\bmod n])$ 收集。**注意在求和的情况下，把题意理解成循环左移还是循环右移，算出的结果都是一样的。**（样例 1 解释中的类型变更是反过来的，但计算结果是正确的。）

如果操作两次，第 $i$ 个巧克力可以花费 $\min(\textit{nums}[i], \textit{nums}[(i+1)\bmod n],  \textit{nums}[(i+2) \bmod n])$ 收集。

依此类推。

#### 提示 3

如果暴力枚举，总的时间复杂度是 $\mathcal{O}(n^3)$。

优化办法有三种：

1. 用 $\mathcal{O}(n^2)$ 的时间预处理所有子数组的最小值，存到一个二维数组中。这样做需要 $\mathcal{O}(n^2)$ 的空间。
2. 用 ST 表优化上述过程。但还有更简单的做法。
3. 用一个长为 $n$ 的数组 $\textit{sum}$ 统计操作 $i$ 次的总花费，这样就可以一边枚举子数组，一边求最小值，一边累加成本了。该方法只需要 $\mathcal{O}(n)$ 的空间。


```go 
func minCost(a []int, x int) int64 {
	ans, n := math.MaxInt64, len(a)
	sum := make([]int, n)
	for i := range sum {
		sum[i] = i * x
	}

	for i, v := range a {
		for j := i; j < n+i; j++ {
			v = min(v, a[j%n])
			sum[j-i] += v
		}
	}
	for _, v := range sum {
		ans = min(ans, v)
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

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。