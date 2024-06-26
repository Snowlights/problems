#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 和一个整数 <code>x</code> 。</p>

<p>请你找到数组中下标距离至少为 <code>x</code> 的两个元素的 <strong>差值绝对值</strong> 的 <strong>最小值</strong> 。</p>

<p>换言之，请你找到两个下标 <code>i</code> 和 <code>j</code> ，满足 <code>abs(i - j) >= x</code> 且 <code>abs(nums[i] - nums[j])</code> 的值最小。</p>

<p>请你返回一个整数，表示下标距离至少为 <code>x</code> 的两个元素之间的差值绝对值的 <strong>最小值</strong> 。</p>

<p> </p>

<p><b>示例 1：</b></p>

<pre><b>输入：</b>nums = [4,3,2,4], x = 2
<b>输出：</b>0
<b>解释：</b>我们选择 nums[0] = 4 和 nums[3] = 4 。
它们下标距离满足至少为 2 ，差值绝对值为最小值 0 。
0 是最优解。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre><b>输入：</b>nums = [5,3,2,10,15], x = 1
<b>输出：</b>1
<b>解释：</b>我们选择 nums[1] = 3 和 nums[2] = 2 。
它们下标距离满足至少为 1 ，差值绝对值为最小值 1 。
1 是最优解。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre><b>输入：</b>nums = [1,2,3,4], x = 3
<b>输出：</b>3
<strong>解释：</strong>我们选择 nums[0] = 1 和 nums[3] = 4 。
它们下标距离满足至少为 3 ，差值绝对值为最小值 3 。
3 是最优解。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
	<li><code>0 <= x < nums.length</code></li>
</ul>

#### 思路

平衡树+双指针

```go  
func minAbsoluteDifference(nums []int, k int) int {
	ans := math.MaxInt
	t := redblacktree.NewWithIntComparator()
	t.Put(math.MaxInt, nil) // 哨兵
	t.Put(math.MinInt/2, nil)
	for i, y := range nums[k:] {
		t.Put(nums[i], nil)
		c, _ := t.Ceiling(y)
		f, _ := t.Floor(y)
		ans = min(ans, min(c.Key.(int)-y, y-f.Key.(int)))
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n-x)\log (n-x))$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n-x)$。
