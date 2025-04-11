#### 题目

<p>给你一个<strong>下标从 0 开始</strong>且大小为 <code>m x n</code> 的整数矩阵 <code>mat</code> 和一个整数 <code>k</code> 。请你将矩阵中的<strong> 奇数</strong> 行循环 <strong>右</strong> 移 <code>k</code> 次，<strong>偶数</strong> 行循环 <strong>左</strong> 移 <code>k</code> 次。</p>

<p>如果初始矩阵和最终矩阵完全相同，则返回 <code>true</code> ，否则返回 <code>false</code> 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>mat = [[1,2,1,2],[5,5,5,5],[6,3,6,3]], k = 2
<strong>输出：</strong>true
<strong>解释：</strong>
<img alt="" src="https://assets.leetcode.com/uploads/2023/10/29/similarmatrix.png" style="width: 500px; height: 117px;" />

初始矩阵如图一所示。
图二表示对奇数行右移一次且对偶数行左移一次后的矩阵状态。
图三是经过两次循环移位后的最终矩阵状态，与初始矩阵相同。
因此，返回 true 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>mat = [[2,2],[2,2]], k = 3
<strong>输出：</strong>true
<strong>解释：</strong>由于矩阵中的所有值都相等，即使进行循环移位，矩阵仍然保持不变。因此，返回 true 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>mat = [[1,2]], k = 1
<strong>输出：</strong>false
<strong>解释：</strong>循环移位一次后，mat = [[2,1]]，与初始矩阵不相等。因此，返回 false 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= mat.length &lt;= 25</code></li>
	<li><code>1 &lt;= mat[i].length &lt;= 25</code></li>
	<li><code>1 &lt;= mat[i][j] &lt;= 25</code></li>
	<li><code>1 &lt;= k &lt;= 50</code></li>
</ul>

#### 思路

由于循环左移 $n$ 次等价于循环左移 $0$ 次，左移 $n+1$ 次等价于循环左移 $1$ 次，……，所以可以先把 $k$ 模上 $m$。

如果此时 $k=0$，那么操作不会影响矩阵，直接返回 `true`。

否则模拟就行。

如果左移 $k$ 次可以让数组相等，那么右移 $k$ 次也可以让数组相等，所以怎么写都是对的，可以全部改成左移 $k$ 次或者右移 $k$ 次。


```go  
func areSimilar(mat [][]int, k int) bool {
	n := len(mat[0])
	k %= n
	if k == 0 {
		return true
	}
	for _, r := range mat {
		if !slices.Equal(r, append(r[k:], r[:k]...)) {
			return false
		}
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{mat}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。

