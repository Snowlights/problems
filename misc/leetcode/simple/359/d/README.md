#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 和一个整数 <code>k</code> 。</p>

<p>如果子数组中所有元素都相等，则认为子数组是一个 <strong>等值子数组</strong> 。注意，空数组是 <strong>等值子数组</strong> 。</p>

<p>从 <code>nums</code> 中删除最多 <code>k</code> 个元素后，返回可能的最长等值子数组的长度。</p>

<p><strong>子数组</strong> 是数组中一个连续且可能为空的元素序列。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,3,2,3,1,3], k = 3
<strong>输出：</strong>3
<strong>解释：</strong>最优的方案是删除下标 2 和下标 4 的元素。
删除后，nums 等于 [1, 3, 3, 3] 。
最长等值子数组从 i = 1 开始到 j = 3 结束，长度等于 3 。
可以证明无法创建更长的等值子数组。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,1,2,2,1,1], k = 2
<strong>输出：</strong>4
<strong>解释：</strong>最优的方案是删除下标 2 和下标 3 的元素。 
删除后，nums 等于 [1, 1, 1, 1] 。 
数组自身就是等值子数组，长度等于 4 。 
可以证明无法创建更长的等值子数组。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= nums.length</code></li>
	<li><code>0 &lt;= k &lt;= nums.length</code></li>
</ul>

#### 思路

把相同元素分组，相同元素的下标记录到哈希表（或者数组）$\textit{pos}$ 中。  
遍历 $\textit{pos}$ 中的每个下标列表 $\textit{ps}$，用双指针处理：  
如果等值子数组的元素下标是从 $\textit{ps}[\textit{left}]$ 到 $\textit{ps}[\textit{right}]$，那么子数组的长度为

$$
\textit{ps}[\textit{right}] - \textit{ps}[\textit{left}] + 1
$$

其中无需删除的元素个数为

$$
\textit{right} - \textit{left} + 1
$$

那么需要删除的元素个数为

$$
\textit{ps}[\textit{right}] - \textit{ps}[\textit{left}] - (\textit{right} - \textit{left})
$$

如果上式大于 $k$，则需要移动左指针。  
满足条件时，用 $\textit{right}-\textit{left}+1$ 更新答案的最大值。

代码实现时，可以在哈希表中记录 $\textit{ps}[i]-i$，这样删除的元素个数就是

$$
\textit{ps}[\textit{right}] - \textit{ps}[\textit{left}]
$$

```go  
func longestEqualSubarray(nums []int, k int) (ans int) {
	pos := make([][]int, len(nums)+1)
	for i, x := range nums {
		pos[x] = append(pos[x], i-len(pos[x]))
	}
	for _, ps := range pos {
		left := 0
		for right, p := range ps {
			for p-ps[left] > k { // 要删除的数太多了
				left++
			}
			ans = max(ans, right-left+1)
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(nm)$。
