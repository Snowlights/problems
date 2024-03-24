### 题目

<p>给你一个长度为 <code>n</code> 的 <strong>正</strong> 整数数组 <code>nums</code> 。</p>

<p><strong>多边形</strong> 指的是一个至少有 <code>3</code> 条边的封闭二维图形。多边形的 <strong>最长边</strong> 一定 <strong>小于</strong> 所有其他边长度之和。</p>

<p>如果你有 <code>k</code> （<code>k >= 3</code>）个 <strong>正</strong> 数 <code>a<sub>1</sub></code>，<code>a<sub>2</sub></code>，<code>a<sub>3</sub></code>, ...，<code>a<sub>k</sub></code> 满足 <code>a<sub>1</sub> <= a<sub>2</sub> <= a<sub>3</sub> <= ... <= a<sub>k</sub></code> <strong>且</strong> <code>a<sub>1</sub> + a<sub>2</sub> + a<sub>3</sub> + ... + a<sub>k-1</sub> > a<sub>k</sub></code><sub> </sub>，那么 <strong>一定</strong> 存在一个 <code>k</code> 条边的多边形，每条边的长度分别为 <code>a<sub>1</sub></code> ，<code>a<sub>2</sub></code> ，<code>a<sub>3</sub></code> ， ...，<code>a<sub>k</sub></code> 。</p>

<p>一个多边形的 <strong>周长</strong> 指的是它所有边之和。</p>

<p>请你返回从 <code>nums</code> 中可以构造的 <strong>多边形 </strong>的 <strong>最大周长</strong> 。如果不能构造出任何多边形，请你返回 <code>-1</code> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [5,5,5]
<b>输出：</b>15
<b>解释：</b>nums 中唯一可以构造的多边形为三角形，每条边的长度分别为 5 ，5 和 5 ，周长为 5 + 5 + 5 = 15 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [1,12,1,2,5,50,3]
<b>输出：</b>12
<b>解释：</b>最大周长多边形为五边形，每条边的长度分别为 1 ，1 ，2 ，3 和 5 ，周长为 1 + 1 + 2 + 3 + 5 = 12 。
我们无法构造一个包含变长为 12 或者 50 的多边形，因为其他边之和没法大于两者中的任何一个。
所以最大周长为 12 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>nums = [5,5,50]
<b>输出：</b>-1
<b>解释：</b>无法构造任何多边形，因为多边形至少要有 3 条边且 50 > 5 + 5 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 <= n <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
</ul>

### 思路

把数组排序，枚举 $\textit{nums}[i]$ 作为最长的那条边，那么 $\textit{nums}[i]$ 左边的数之和越大越好，
这样才能尽可能地组成多边形，并且周长尽量长。
所以周长就是 $\textit{nums}$ 的某个前缀和。
从大到小枚举 $\textit{nums}[i]$，如果满足

$$
\textit{nums}[0] + \textit{nums}[1] + \cdots + \textit{nums}[i-1] > \textit{nums}[i]

$$

那么答案就是

$$
\textit{nums}[0] + \textit{nums}[1] + \cdots + \textit{nums}[i-1] + \textit{nums}[i]

$$

```go [sol]
func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	s := 0
	for _, v := range nums {
		s += v
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if s > 2 * nums[i] {
			return int64(s)
		}
		s -= nums[i]
	}
	return -1
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。不计入排序的栈开销。
