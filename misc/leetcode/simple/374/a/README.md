#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的数组 <code>mountain</code> 。你的任务是找出数组&nbsp;<code>mountain</code> 中的所有 <strong>峰值</strong>。</p>

<p>以数组形式返回给定数组中 <strong>峰值</strong> 的下标，<strong>顺序不限</strong> 。</p>

<p><strong>注意：</strong></p>

<ul>
	<li><strong>峰值</strong> 是指一个严格大于其相邻元素的元素。</li>
	<li>数组的第一个和最后一个元素 <strong>不</strong> 是峰值。</li>
</ul>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>mountain = [2,4,4]
<strong>输出：</strong>[]
<strong>解释：</strong>mountain[0] 和 mountain[2] 不可能是峰值，因为它们是数组的第一个和最后一个元素。
mountain[1] 也不可能是峰值，因为它不严格大于 mountain[2] 。
因此，答案为 [] 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>mountain = [1,4,3,8,5]
<strong>输出：</strong>[1,3]
<strong>解释：</strong>mountain[0] 和 mountain[4] 不可能是峰值，因为它们是数组的第一个和最后一个元素。
mountain[2] 也不可能是峰值，因为它不严格大于 mountain[3] 和 mountain[1] 。
但是 mountain[1] 和 mountain[3] 严格大于它们的相邻元素。
因此，答案是 [1,3] 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 &lt;= mountain.length &lt;= 100</code></li>
	<li><code>1 &lt;= mountain[i] &lt;= 100</code></li>
</ul>

#### 思路

遍历下标在 $[1,n-2]$ 内的所有数，如果其大于其左右两侧相邻数字，则把 $i$ 加入答案。

```go  
func findPeaks(mountain []int) (ans []int) {
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			ans = append(ans, i)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{mountain}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

