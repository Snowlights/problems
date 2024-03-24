### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的 <code>m x n</code> 整数矩阵 <code>grid</code> 。矩阵中某一列的宽度是这一列数字的最大 <strong>字符串长度</strong> 。</p>

<ul>
	<li>比方说，如果 <code>grid = [[-10], [3], [12]]</code> ，那么唯一一列的宽度是 <code>3</code> ，因为 <code>-10</code> 的字符串长度为 <code>3</code> 。</li>
</ul>

<p>请你返回一个大小为 <code>n</code> 的整数数组 <code>ans</code> ，其中 <code>ans[i]</code> 是第 <code>i</code> 列的宽度。</p>

<p>一个有 <code>len</code> 个数位的整数 <code>x</code> ，如果是非负数，那么 <strong>字符串</strong><strong>长度</strong> 为 <code>len</code> ，否则为 <code>len + 1</code> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>grid = [[1],[22],[333]]
<b>输出：</b>[3]
<b>解释：</b>第 0 列中，333 字符串长度为 3 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>grid = [[-15,1,3],[15,7,12],[5,6,-2]]
<b>输出：</b>[3,1,2]
<b>解释：</b>
第 0 列中，只有 -15 字符串长度为 3 。
第 1 列中，所有整数的字符串长度都是 1 。
第 2 列中，12 和 -2 的字符串长度都为 2 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>m == grid.length</code></li>
	<li><code>n == grid[i].length</code></li>
	<li><code>1 &lt;= m, n &lt;= 100 </code></li>
	<li><code>-10<sup>9</sup> &lt;= grid[r][c] &lt;= 10<sup>9</sup></code></li>
</ul>
 
### 思路  

遍历每一列，求出数字的字符串形式的长度的最大值

```go 
func findColumnWidth(grid [][]int) []int {
	ans := make([]int, len(grid[0]))
	for j := range grid[0] {
		for _, row := range grid {
			ans[j] = max(ans[j], len(strconv.Itoa(row[j])))
		}
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```

### 复杂度分析  

- 时间复杂度：$O(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$O(1)$。返回值不计入。Python 忽略 `zip(*grid)` 导致的空间。