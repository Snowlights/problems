#### 题目

<p>给你一个下标从 <strong>0</strong>&nbsp;开始的整数数组&nbsp;<code>nums</code>&nbsp;。</p>

<p><code>nums</code>&nbsp;一个长度为 <code>k</code>&nbsp;的 <strong>子序列</strong>&nbsp;指的是选出 <code>k</code>&nbsp;个 <strong>下标</strong>&nbsp;<code>i<sub>0</sub>&nbsp;&lt;&nbsp;i<sub>1</sub> &lt;&nbsp;... &lt; i<sub>k-1</sub></code>&nbsp;，如果这个子序列满足以下条件，我们说它是 <strong>平衡的</strong>&nbsp;：</p>

<ul>
	<li>对于范围&nbsp;<code>[1, k - 1]</code>&nbsp;内的所有&nbsp;<code>j</code>&nbsp;，<code>nums[i<sub>j</sub>] - nums[i<sub>j-1</sub>] &gt;= i<sub>j</sub> - i<sub>j-1</sub></code>&nbsp;都成立。</li>
</ul>

<p><code>nums</code>&nbsp;长度为 <code>1</code>&nbsp;的 <strong>子序列</strong>&nbsp;是平衡的。</p>

<p>请你返回一个整数，表示 <code>nums</code>&nbsp;<strong>平衡</strong>&nbsp;子序列里面的 <strong>最大元素和</strong>&nbsp;。</p>

<p>一个数组的 <strong>子序列</strong>&nbsp;指的是从原数组中删除一些元素（<strong>也可能一个元素也不删除</strong>）后，剩余元素保持相对顺序得到的 <strong>非空</strong>&nbsp;新数组。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [3,3,5,6]
<b>输出：</b>14
<b>解释：</b>这个例子中，选择子序列 [3,5,6] ，下标为 0 ，2 和 3 的元素被选中。
nums[2] - nums[0] &gt;= 2 - 0 。
nums[3] - nums[2] &gt;= 3 - 2 。
所以，这是一个平衡子序列，且它的和是所有平衡子序列里最大的。
包含下标 1 ，2 和 3 的子序列也是一个平衡的子序列。
最大平衡子序列和为 14 。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [5,-1,-3,8]
<b>输出：</b>13
<b>解释：</b>这个例子中，选择子序列 [5,8] ，下标为 0 和 3 的元素被选中。
nums[3] - nums[0] &gt;= 3 - 0 。
所以，这是一个平衡子序列，且它的和是所有平衡子序列里最大的。
最大平衡子序列和为 13 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>nums = [-2,-1]
<b>输出：</b>-1
<b>解释：</b>这个例子中，选择子序列 [-1] 。
这是一个平衡子序列，而且它的和是 nums 所有平衡子序列里最大的。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-10<sup>9</sup> &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>

#### 思路

定义 $b[i] = \textit{nums}[i] - i$，问题变成从 $b$ 中选出一个非降子序列，求对应的 $\textit{nums}$ 的元素和的最大值。  
如果 $i$ 是子序列最后一个数的下标，考虑倒数第二个数的下标 $j$，如果 $b[j]\le b[i]$，那么就找到了一个子问题：
子序列最后一个数的下标是 $j$ 时，对应的 $\textit{nums}$ 的元素和的最大值。  
据此，定义 $f[i]$ 表示子序列最后一个数的下标是 $i$ 时，对应的 $\textit{nums}$ 的元素和的最大值。  
那么答案就是 $\max(f)$。

状态转移方程为

$$
f[i] = \max_{j} \max(f[j],0) + \textit{nums}[i]
$$

其中 $j$ 满足 $j < i$ 且 $b[j]\le b[i]$。如果 $f[j]<0$，则和 $0$ 取最大值，  
表示只选 $\textit{nums}[i]$ 一个数，前面的数都不选。  
这可以用**权值树状数组**（或者权值线段树）来优化。树状数组用来维护前缀最大值：
设下标为 $x=b[i]$，维护的值为 $\max(f[x], f[x-1], f[x-2], \cdots)$。
代码实现时需要先把 $\textit{nums}[i]-i$ **离散化**，再使用树状数组。

> 离散化：把最小元素映射为 $1$，次小元素映射为 $2$，依此类推。


```go  
// 树状数组模板（维护前缀最大值）
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}

func (f fenwick) preMax(i int) int {
	mx := math.MinInt
	for ; i > 0; i &= i - 1 {
		mx = max(mx, f[i])
	}
	return mx
}

func maxBalancedSubsequenceSum(nums []int) int64 {
	// 离散化 nums[i]-i
	b := slices.Clone(nums)
	for i := range b {
		b[i] -= i
	}
	slices.Sort(b)
	b = slices.Compact(b) // 去重

	// 初始化树状数组
	t := make(fenwick, len(b)+1)
	for i := range t {
		t[i] = math.MinInt
	}

	for i, x := range nums {
		j := sort.SearchInts(b, x-i) + 1 // nums[i]-i 离散化后的值（从 1 开始）
		f := max(t.preMax(j), 0) + x
		t.update(j, f)
	}
	return int64(t.preMax(len(b)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
