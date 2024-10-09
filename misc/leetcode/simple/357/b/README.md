#### 题目  

<p>给你一个长度为 <code>n</code> 的数组 <code>nums</code> 和一个整数 <code>m</code> 。请你判断能否执行一系列操作，将数组拆分成 <code>n</code> 个 <strong>非空 </strong>数组。</p>

<p>在每一步操作中，你可以选择一个 <strong>长度至少为 2</strong> 的现有数组（之前步骤的结果） 并将其拆分成 <strong>2</strong> 个子数组，而得到的 <strong>每个</strong> 子数组，<strong>至少</strong> 需要满足以下条件之一：</p>

<ul>
	<li>子数组的长度为 1 ，或者</li>
	<li>子数组元素之和 <strong>大于或等于</strong>  <code>m</code> 。</li>
</ul>

<p>如果你可以将给定数组拆分成 <code>n</code> 个满足要求的数组，返回 <code>true</code><em> </em>；否则，返回 <code>false</code> 。</p>

<p><strong>注意：</strong>子数组是数组中的一个连续非空元素序列。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [2, 2, 1], m = 4
<strong>输出：</strong>true
<strong>解释：</strong>
第 1 步，将数组 nums 拆分成 [2, 2] 和 [1] 。
第 2 步，将数组 [2, 2] 拆分成 [2] 和 [2] 。
因此，答案为 true 。</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [2, 1, 3], m = 5 
<strong>输出：</strong>false
<strong>解释：
</strong>存在两种不同的拆分方法：
第 1 种，将数组 nums 拆分成 [2, 1] 和 [3] 。
第 2 种，将数组 nums 拆分成 [2] 和 [1, 3] 。
然而，这两种方法都不满足题意。因此，答案为 false 。</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>nums = [2, 3, 3, 2, 3], m = 6
<strong>输出：</strong>true
<strong>解释：</strong>
第 1 步，将数组 nums 拆分成 [2, 3, 3, 2] 和 [3] 。
第 2 步，将数组 [2, 3, 3, 2] 拆分成 [2, 3, 3] 和 [2] 。
第 3 步，将数组 [2, 3, 3] 拆分成 [2] 和 [3, 3] 。
第 4 步，将数组 [3, 3] 拆分成 [3] 和 [3] 。
因此，答案为 true 。 </pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == nums.length &lt;= 100</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 100</code></li>
	<li><code>1 &lt;= m &lt;= 200</code></li>
</ul>
 
#### 思路  

根据题意，分割成一个长为 $2$ 和一个长为 $n-2$ 的数组，不如分割成一个长为 $1$ 和一个长为 $n-1$ 的数组，因为后者的数组元素和更大，更容易满足题目要求。  
依此类推，当剩下一个长为 $3$ 的数组时，需要分成一个长为 $1$ 和一个长为 $2$ 的数组。如果这个长为 $2$ 的数组的元素和 $\ge m$，那么返回 `true`，否则返回 `false`。  
所以问题变成判断数组中是否有两个相邻数字 $\ge m$ 即可。  
注意特判 $n\le 2$ 的情况，这是满足要求的。

```go 
func canSplitArray(a []int, m int) (ans bool) {
	if len(a) <= 2 {
		return true
	}
	for i := 1; i < len(a); i ++ {
		if a[i] + a[i-1] >= m {
			return true
		}
	}
	return
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。