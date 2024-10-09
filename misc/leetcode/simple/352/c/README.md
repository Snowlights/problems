#### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 。<code>nums</code> 的一个子数组如果满足以下条件，那么它是 <strong>不间断</strong> 的：</p>

<ul>
	<li><code>i</code>，<code>i + 1</code> ，...，<code>j</code><sub> </sub> 表示子数组中的下标。对于所有满足 <code>i &lt;= i<sub>1</sub>, i<sub>2</sub> &lt;= j</code> 的下标对，都有 <code>0 &lt;= |nums[i<sub>1</sub>] - nums[i<sub>2</sub>]| &lt;= 2</code> 。</li>
</ul>

<p>请你返回 <strong>不间断</strong> 子数组的总数目。</p>

<p>子数组是一个数组中一段连续 <strong>非空</strong> 的元素序列。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [5,4,2,4]
<strong>输出：</strong>8
<b>解释：</b>
大小为 1 的不间断子数组：[5], [4], [2], [4] 。
大小为 2 的不间断子数组：[5,4], [4,2], [2,4] 。
大小为 3 的不间断子数组：[4,2,4] 。
没有大小为 4 的不间断子数组。
不间断子数组的总数目为 4 + 3 + 1 = 8 。
除了这些以外，没有别的不间断子数组。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [1,2,3]
<b>输出：</b>6
<b>解释：</b>
大小为 1 的不间断子数组：[1], [2], [3] 。
大小为 2 的不间断子数组：[1,2], [2,3] 。
大小为 3 的不间断子数组：[1,2,3] 。
不间断子数组的总数目为 3 + 2 + 1 = 6 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>
 
#### 思路  

本题思路和 [1438. 绝对差不超过限制的最长连续子数组](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/) 类似。

由于本题绝对差至多为 $2$，所以做法可以简单点。用平衡树或者哈希表都行，反正至多维护 $3$ 个数，添加删除可以视作是 $\mathcal{O}(1)$ 的。

代码框架是 滑动窗口（双指针）。在遍历数组的同时，用多重集合维护窗口内的数字。（如果用哈希表，还需记录数字的出现次数。）

如果窗口内的最大值与最小值的差大于 $2$，就不断移动左端点 $\textit{left}$，减少窗口内的数字。

最后

$$
[\textit{left},\textit{right}],[\textit{left}+1,\textit{right}],\cdots,[\textit{right},\textit{right}]
$$

这一共 $\textit{right}-\textit{left}+1$ 个子数组都是合法的，加入答案。

```go 
func continuousSubarrays(a []int) int64 {
	left, ans := 0, 0
	cnt := make(map[int]int)
	max := func() int {
		m := 0
		for k := range cnt {
			if k > m {
				m = k
			}
		}
		return m
	}
	min := func() int {
		m := math.MaxInt32
		for k := range cnt {
			if k < m {
				m = k
			}
		}
		return m
	}

	for r, v := range a {
		cnt[v]++
		for max() - min() > 2 {
			y := a[left]
			cnt[y]--
			if cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		ans += r - left + 1
	}
	return int64(ans)
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。注意至多维护 $3$ 个数，仅用到常量额外空间。