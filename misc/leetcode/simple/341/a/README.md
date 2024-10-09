#### 题目  

<p>给你一个大小为 <code>m x n</code> 的二进制矩阵 <code>mat</code> ，请你找出包含最多 <strong>1</strong> 的行的下标（从 <strong>0</strong> 开始）以及这一行中 <strong>1</strong> 的数目。</p>

<p>如果有多行包含最多的 1 ，只需要选择 <strong>行下标最小</strong> 的那一行。</p>

<p>返回一个由行下标和该行中 1 的数量组成的数组。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>mat = [[0,1],[1,0]]
<strong>输出：</strong>[0,1]
<strong>解释：</strong>两行中 1 的数量相同。所以返回下标最小的行，下标为 0 。该行 1 的数量为 1 。所以，答案为 [0,1] 。 
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>mat = [[0,0,0],[0,1,1]]
<strong>输出：</strong>[1,2]
<strong>解释：</strong>下标为 1 的行中 1 的数量最多<code>。</code>该行 1 的数量<code>为 2 。所以，答案为</code> [1,2] 。
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>mat = [[0,0],[1,1],[0,0]]
<strong>输出：</strong>[1,2]
<strong>解释：</strong>下标为 1 的行中 1 的数量最多。该行 1 的数量<code>为 2 。所以，答案为</code> [1,2] 。</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>m == mat.length</code> </li>
	<li><code>n == mat[i].length</code> </li>
	<li><code>1 &lt;= m, n &lt;= 100</code> </li>
	<li><code>mat[i][j]</code> 为 <code>0</code> 或 <code>1</code></li>
</ul>
 
#### 思路  

遍历每一行，统计一整行的元素和的最大值及其下标。由于是从上到下遍历行的，所以找到的最大值的下标一定是相同最大值中最小的。

```go 
func rowAndMaximumOnes(mat [][]int) []int {
	maxSum, idx := -1, 0
	for i, row := range mat {
		sum := 0
		for _, x := range row {
			sum += x
		}
		if sum > maxSum {
			maxSum, idx = sum, i
		}
	}
	return []int{idx, maxSum}
}
```

#### 复杂度分析  

- 时间复杂度：$O(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{mat}$ 的行数和列数。
- 空间复杂度：$O(1)$。仅用到若干额外变量。