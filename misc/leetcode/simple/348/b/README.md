#### 题目  

<p>给你一个下标从 <strong>0</strong> 开始、长度为 <code>n</code> 的整数排列 <code>nums</code> 。</p>

<p>如果排列的第一个数字等于 <code>1</code> 且最后一个数字等于 <code>n</code> ，则称其为 <strong>半有序排列</strong> 。你可以执行多次下述操作，直到将 <code>nums</code> 变成一个 <strong>半有序排列</strong> ：</p>

<ul>
	<li>选择 <code>nums</code> 中相邻的两个元素，然后交换它们。</li>
</ul>

<p>返回使 <code>nums</code> 变成 <strong>半有序排列</strong> 所需的最小操作次数。</p>

<p><strong>排列</strong> 是一个长度为 <code>n</code> 的整数序列，其中包含从 <code>1</code> 到 <code>n</code> 的每个数字恰好一次。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [2,1,4,3]
<strong>输出：</strong>2
<strong>解释：</strong>可以依次执行下述操作得到半有序排列：
1 - 交换下标 0 和下标 1 对应元素。排列变为 [1,2,4,3] 。
2 - 交换下标 2 和下标 3 对应元素。排列变为 [1,2,3,4] 。
可以证明，要让 nums 成为半有序排列，不存在执行操作少于 2 次的方案。</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [2,4,1,3]
<strong>输出：</strong>3
<strong>解释：
</strong>可以依次执行下述操作得到半有序排列：
1 - 交换下标 1 和下标 2 对应元素。排列变为 [2,1,4,3] 。
2 - 交换下标 0 和下标 1 对应元素。排列变为 [1,2,4,3] 。
3 - 交换下标 2 和下标 3 对应元素。排列变为 [1,2,3,4] 。
可以证明，要让 nums 成为半有序排列，不存在执行操作少于 3 次的方案。
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>nums = [1,3,4,2,5]
<strong>输出：</strong>0
<strong>解释：</strong>这个排列已经是一个半有序排列，无需执行任何操作。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length == n &lt;= 50</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 50</code></li>
	<li><code>nums</code> 是一个 <strong>排列</strong></li>
</ul>
 
#### 思路  

记录 $1$ 和 $n$ 所在的位置为 $first$ 和 $last$ :  
- 如果 $1$ 的位置比 $n$ 的位置左，$1$ 移动 $first$ 步， $n$ 移动 $n-1-last$步
- 如果 $1$ 的位置比 $n$ 的位置右，$1$ 移动 $first$ 步， 因为左移 $1$的时候已经移动了 $n$， 最终移动 $n-2-last$步

```go 
func semiOrderedPermutation(a []int) (ans int) {
	n, first, last := len(a), 0, 0
	for i, v := range a {
		if v == 1 {
			first = i
		}
		if v == n {
			last = i
		}
	}

	if first < last {
		return first + n - 1 - last
	}

	return first + n - 2 - last
}
```

```python  
class Solution:
    def semiOrderedPermutation(self, nums: List[int]) -> int:
        n = len(nums)
        p = nums.index(1)
        q = nums.index(n)
        return p + n - 1 - q - (p > q)
```

#### 复杂度分析  

- 时间复杂度：${O}\mathcal(n)$。 $n$ 为 $\textit{a}$ 长度。
- 空间复杂度：${O}\mathcal(1)$。 仅用到部分变量