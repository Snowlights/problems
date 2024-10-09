#### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> ，数组长度为 <code>n</code> 。</p>

<p>你可以执行无限次下述运算：</p>

<ul>
	<li>选择一个之前未选过的下标 <code>i</code> ，并选择一个 <strong>严格小于</strong> <code>nums[i]</code> 的质数 <code>p</code> ，从 <code>nums[i]</code> 中减去 <code>p</code> 。</li>
</ul>

<p>如果你能通过上述运算使得 <code>nums</code> 成为严格递增数组，则返回 <code>true</code> ；否则返回 <code>false</code> 。</p>

<p><strong>严格递增数组</strong> 中的每个元素都严格大于其前面的元素。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [4,9,6,10]
<strong>输出：</strong>true
<strong>解释：</strong>
在第一次运算中：选择 i = 0 和 p = 3 ，然后从 nums[0] 减去 3 ，nums 变为 [1,9,6,10] 。
在第二次运算中：选择 i = 1 和 p = 7 ，然后从 nums[1] 减去 7 ，nums 变为 [1,2,6,10] 。
第二次运算后，nums 按严格递增顺序排序，因此答案为 true 。</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [6,8,11,12]
<strong>输出：</strong>true
<strong>解释：</strong>nums 从一开始就按严格递增顺序排序，因此不需要执行任何运算。</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>nums = [5,8,3]
<strong>输出：</strong>false
<strong>解释：</strong>可以证明，执行运算无法使 nums 按严格递增顺序排序，因此答案是 false 。</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 1000</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 1000</code></li>
	<li><code>nums.length == n</code></li>
</ul>
 
#### 思路  

设 $\textit{pre}$ 是上一个减完后的数字，$x=\textit{nums}[i]$ 为当前数字。  
设 $p$ 是满足 $x-p>\textit{pre}$ 的最大质数，换言之，$p$ 是小于 $x-\textit{pre}$ 的最大质数， 这可以预处理质数列表后，用二分查找得到。

```go 
var p = []int{0} // 哨兵，避免二分越界

func init() {
	const mx = 1000
	np := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !np[i] {
			p = append(p, i) // 预处理质数列表
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func primeSubOperation(nums []int) bool {
	pre := 0
	for _, x := range nums {
		if x <= pre {
			return false
		}
		pre = x - p[sort.SearchInts(p, x-pre)-1] // 减去 < x-pre 的最大质数
	}
	return true
}
```

#### 复杂度分析  

- 时间复杂度：$O(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U$ 为 $1000$ 以内的质数个数。
- 空间复杂度：$O(1)$。忽略预处理的空间，仅用到若干额外变量。