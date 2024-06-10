#### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 和一个整数 <code>p</code> 。请你从 <code>nums</code> 中找到 <code>p</code> 个下标对，每个下标对对应数值取差值，你需要使得这 <code>p</code> 个差值的 <strong>最大值</strong> <strong>最小</strong>。同时，你需要确保每个下标在这 <code>p</code> 个下标对中最多出现一次。</p>

<p>对于一个下标对 <code>i</code> 和 <code>j</code> ，这一对的差值为 <code>|nums[i] - nums[j]|</code> ，其中 <code>|x|</code> 表示 <code>x</code> 的 <strong>绝对值</strong> 。</p>

<p>请你返回 <code>p</code> 个下标对对应数值 <strong>最大差值</strong> 的 <strong>最小值</strong> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [10,1,2,7,1,3], p = 2
<b>输出：</b>1
<b>解释：</b>第一个下标对选择 1 和 4 ，第二个下标对选择 2 和 5 。
最大差值为 max(|nums[1] - nums[4]|, |nums[2] - nums[5]|) = max(0, 1) = 1 。所以我们返回 1 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [4,2,1,2], p = 1
<b>输出：</b>0
<b>解释：</b>选择下标 1 和 3 构成下标对。差值为 |2 - 2| = 0 ，这是最大差值的最小值。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>0 &lt;= p &lt;= (nums.length)/2</code></li>
</ul>
 
#### 思路  

#### 提示 1

看到「最大化最小值」或者「最小化最大值」就要想到**二分答案**，这是一个固定的套路。
为什么？一般来说，二分的值越大，越能/不能满足要求；二分的值越小，越不能/能满足要求，有单调性，可以二分。

#### 提示 2

二分数对中的最大差值 $mx$。 由于下标和答案无关，可以先排序。为了让匹配的数对尽量多，应尽量选相邻的元素，这样更能满足要求。例如 $[1,2,3,4]$，如果 $1,3$ 匹配，$2,4$ 匹配，最大差值是 $2$；而如果 $1,2$ 相邻匹配，$3,4$ 相邻匹配，最大差值只有 $1$。  
我们来算一算最多能匹配多少个数对：
- 如果可以选 $\textit{nums}[0]$ 和 $\textit{nums}[1]$，那么答案等于「$n-2$ 个数的最多数对个数」$+1$。
- 如果不选 $\textit{nums}[0]$，那么答案等于「$n-1$ 个数的最多数对个数」。
- 这两种情况取最大值。

这看上去很像 [198. 打家劫舍](https://leetcode.cn/problems/house-robber/)，可以用动态规划实现。  
也可以用贪心做：
- 注意到，「$n-1$ 个数的最多数对个数」不会超过「$n-3$ 个数的最多数对个数」$+1$。这里 $+1$ 表示选 $\textit{nums}[1]$ 和 $\textit{nums}[2]$。
- 由于「$n-2$ 个数的最多数对个数」$\ge$「$n-3$ 个数的最多数对个数」，所以如果可以选 $\textit{nums}[0]$ 和 $\textit{nums}[1]$，那么直接选就行。
- 依此类推，不断缩小问题规模。所以遍历一遍数组就能求出最多数对个数，具体见代码。

```go 
func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	return sort.Search(1e9, func(mx int) bool {
		cnt := 0
		for i := 0; i < len(nums)-1; i++ {
			if nums[i+1]-nums[i] <= mx { // 都选
				cnt++
				i++
			}
		}
		return cnt >= p
	})
}
```

#### 复杂度分析  

- 时间复杂度：$O(n\log n + n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})-\min(\textit{nums})$。
- 空间复杂度：$O(1)$。忽略排序时的栈空间，仅用到若干额外变量。