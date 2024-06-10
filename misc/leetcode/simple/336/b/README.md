#### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 。你可以将 <code>nums</code> 中的元素按 <strong>任意顺序</strong> 重排（包括给定顺序）。</p>

<p>令 <code>prefix</code> 为一个数组，它包含了 <code>nums</code> 重新排列后的前缀和。换句话说，<code>prefix[i]</code> 是 <code>nums</code> 重新排列后下标从 <code>0</code> 到 <code>i</code> 的元素之和。<code>nums</code> 的 <strong>分数</strong> 是 <code>prefix</code> 数组中正整数的个数。</p>

<p>返回可以得到的最大分数。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [2,-1,0,1,-3,3,-3]
<strong>输出：</strong>6
<strong>解释：</strong>数组重排为 nums = [2,3,1,-1,-3,0,-3] 。
prefix = [2,5,6,5,2,2,-1] ，分数为 6 。
可以证明 6 是能够得到的最大分数。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [-2,-3,0]
<strong>输出：</strong>0
<strong>解释：</strong>不管怎么重排数组得到的分数都是 0 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-10<sup>6</sup> &lt;= nums[i] &lt;= 10<sup>6</sup></code></li>
</ul>
 
#### 思路  

对于一个负数来说，它后面的前缀和都会把这个负数加进去。  
由于要统计的是正数前缀和，那么把负数尽量放在后面，能统计到尽量多的正数前缀和。  
同时，绝对值小的负数应该排在负数的前面，尽量在前缀和减为负数前还能多统计一些正数。

```go 
func maxScore(nums []int) (ans int) {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	sum := 0
	for _, x := range nums {
		sum += x
		if sum <= 0 {
			break
		}
		ans++
	}
	return
}
```

#### 复杂度分析  

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。忽略排序时的栈开销，仅用到若干额外变量。