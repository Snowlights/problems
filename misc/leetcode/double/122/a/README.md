### 题目

<p>给你一个长度为 <code>n</code> 的整数数组 <code>nums</code> 。</p>

<p>一个数组的 <strong>代价</strong> 是它的 <strong>第一个</strong> 元素。比方说，<code>[1,2,3]</code> 的代价是 <code>1</code> ，<code>[3,4,1]</code> 的代价是 <code>3</code> 。</p>

<p>你需要将 <code>nums</code> 分成 <code>3</code> 个 <strong>连续且没有交集</strong> 的子数组。</p>

<p>请你返回这些子数组的 <strong>最小</strong> 代价 <b>总和</b> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [1,2,3,12]
<b>输出：</b>6
<b>解释：</b>最佳分割成 3 个子数组的方案是：[1] ，[2] 和 [3,12] ，总代价为 1 + 2 + 3 = 6 。
其他得到 3 个子数组的方案是：
- [1] ，[2,3] 和 [12] ，总代价是 1 + 2 + 12 = 15 。
- [1,2] ，[3] 和 [12] ，总代价是 1 + 3 + 12 = 16 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [5,4,3]
<b>输出：</b>12
<b>解释：</b>最佳分割成 3 个子数组的方案是：[5] ，[4] 和 [3] ，总代价为 5 + 4 + 3 = 12 。
12 是所有分割方案里的最小总代价。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>nums = [10,3,1,1]
<b>输出：</b>12
<b>解释：</b>最佳分割成 3 个子数组的方案是：[10,3] ，[1] 和 [1] ，总代价为 10 + 1 + 1 = 12 。
12 是所有分割方案里的最小总代价。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 <= n <= 50</code></li>
	<li><code>1 <= nums[i] <= 50</code></li>
</ul>

### 思路

把数组分成三段，每一段取第一个数再求和，问和的最小值是多少。
第一段的第一个数是确定的，即 $\textit{nums}[0]$。
如果知道了第二段的第一个数的位置，和第三段的第一个数的位置，那么这个划分方案也就确定了。
这两个下标可以在 $[1,n-1]$ 中随意取。所以问题变成求下标在 $[1,n-1]$ 中的两个最小的数。

```go [sol]
func minimumCost(nums []int) int {
	fi, se := math.MaxInt, math.MaxInt
	for _, x := range nums[1:] {
		if x < fi {
			se = fi
			fi = x
		} else if x < se {
			se = x
		}
	}
	return nums[0] + fi + se
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
