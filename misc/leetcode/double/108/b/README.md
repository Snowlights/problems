### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> ，表示一些石块的初始位置。再给你两个长度<strong> 相等</strong> 下标从 <strong>0</strong> 开始的整数数组 <code>moveFrom</code> 和 <code>moveTo</code> 。</p>

<p>在 <code>moveFrom.length</code> 次操作内，你可以改变石块的位置。在第 <code>i</code> 次操作中，你将位置在 <code>moveFrom[i]</code> 的所有石块移到位置 <code>moveTo[i]</code> 。</p>

<p>完成这些操作后，请你按升序返回所有 <strong>有</strong> 石块的位置。</p>

<p><strong>注意：</strong></p>

<ul>
	<li>如果一个位置至少有一个石块，我们称这个位置 <strong>有</strong> 石块。</li>
	<li>一个位置可能会有多个石块。</li>
</ul>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [1,6,7,8], moveFrom = [1,7,2], moveTo = [2,9,5]
<b>输出：</b>[5,6,8,9]
<b>解释：</b>一开始，石块在位置 1,6,7,8 。
第 i = 0 步操作中，我们将位置 1 处的石块移到位置 2 处，位置 2,6,7,8 有石块。
第 i = 1 步操作中，我们将位置 7 处的石块移到位置 9 处，位置 2,6,8,9 有石块。
第 i = 2 步操作中，我们将位置 2 处的石块移到位置 5 处，位置 5,6,8,9 有石块。
最后，至少有一个石块的位置为 [5,6,8,9] 。</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [1,1,3,3], moveFrom = [1,3], moveTo = [2,2]
<b>输出：</b>[2]
<b>解释：</b>一开始，石块在位置 [1,1,3,3] 。
第 i = 0 步操作中，我们将位置 1 处的石块移到位置 2 处，有石块的位置为 [2,2,3,3] 。
第 i = 1 步操作中，我们将位置 3 处的石块移到位置 2 处，有石块的位置为 [2,2,2,2] 。
由于 2 是唯一有石块的位置，我们返回 [2] 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= moveFrom.length &lt;= 10<sup>5</sup></code></li>
	<li><code>moveFrom.length == moveTo.length</code></li>
	<li><code>1 &lt;= nums[i], moveFrom[i], moveTo[i] &lt;= 10<sup>9</sup></code></li>
	<li>测试数据保证在进行第 <code>i</code> 步操作时，<code>moveFrom[i]</code> 处至少有一个石块。</li>
</ul>
 
### 思路  

注意每次操作，将 $\textit{moveFrom}[i]$ 的**所有石块**都移走，所以无需维护 $\textit{nums}[i]$ 的个数。

```go 
func relocateMarbles(a []int, moveFrom []int, moveTo []int) (ans []int) {

	h := make(map[int]bool)
	for _, v := range a {
		h[v] = true
	}

	for i, v := range moveFrom {
		delete(h, v)
		h[moveTo[i]] = true
	}
	a = []int{}
	for i := range h {
		a = append(a, i)
	}
	sort.Ints(a)
	return a
}
```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(m + n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度，$m$ 为 $\textit{moveFrom}$ 的长度。排序是瓶颈。
- 空间复杂度：$\mathcal{O}(n)$。