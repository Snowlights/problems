### 题目  

<p>给你两个整数 <code>m</code> 和 <code>n</code> ，表示一个下标从 <strong>0</strong> 开始的 <code>m x n</code> 的网格图。</p>

<p>给你一个下标从 <strong>0</strong> 开始的二维整数矩阵 <code>coordinates</code> ，其中 <code>coordinates[i] = [x, y]</code> 表示坐标为 <code>[x, y]</code> 的格子是 <strong>黑色的</strong> ，所有没出现在 <code>coordinates</code> 中的格子都是 <strong>白色的</strong>。</p>

<p>一个块定义为网格图中 <code>2 x 2</code> 的一个子矩阵。更正式的，对于左上角格子为 <code>[x, y]</code> 的块，其中 <code>0 &lt;= x &lt; m - 1</code> 且 <code>0 &lt;= y &lt; n - 1</code> ，包含坐标为 <code>[x, y]</code> ，<code>[x + 1, y]</code> ，<code>[x, y + 1]</code> 和 <code>[x + 1, y + 1]</code> 的格子。</p>

<p>请你返回一个下标从 <strong>0</strong> 开始长度为 <code>5</code> 的整数数组 <code>arr</code> ，<code>arr[i]</code> 表示恰好包含 <code>i</code> 个 <strong>黑色</strong> 格子的块的数目。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>m = 3, n = 3, coordinates = [[0,0]]
<b>输出：</b>[3,1,0,0,0]
<b>解释：</b>网格图如下：
<img alt="" src="https://assets.leetcode.com/uploads/2023/06/18/screen-shot-2023-06-18-at-44656-am.png" style="width: 150px; height: 128px;"/>
只有 1 个块有一个黑色格子，这个块是左上角为 [0,0] 的块。
其他 3 个左上角分别为 [0,1] ，[1,0] 和 [1,1] 的块都有 0 个黑格子。
所以我们返回 [3,1,0,0,0] 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>m = 3, n = 3, coordinates = [[0,0],[1,1],[0,2]]
<b>输出：</b>[0,2,2,0,0]
<b>解释：</b>网格图如下：
<img alt="" src="https://assets.leetcode.com/uploads/2023/06/18/screen-shot-2023-06-18-at-45018-am.png" style="width: 150px; height: 128px;"/>
有 2 个块有 2 个黑色格子（左上角格子分别为 [0,0] 和 [0,1]）。
左上角为 [1,0] 和 [1,1] 的两个块，都有 1 个黑格子。
所以我们返回 [0,2,2,0,0] 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= m &lt;= 10<sup>5</sup></code></li>
	<li><code>2 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= coordinates.length &lt;= 10<sup>4</sup></code></li>
	<li><code>coordinates[i].length == 2</code></li>
	<li><code>0 &lt;= coordinates[i][0] &lt; m</code></li>
	<li><code>0 &lt;= coordinates[i][1] &lt; n</code></li>
	<li><code>coordinates</code> 中的坐标对两两互不相同。</li>
</ul>
 
### 思路  

在 $m$ 和 $n$ 都很大的情况下，网格中有大量的 $2\times 2$ 的子矩阵是没有黑色格子的。只需要考虑有黑色格子的子矩阵。  
如果 $(x,y)$ 处有黑色格子，那么子矩阵左上角在 $(x-1,y-1),(x-1,y),(x,y-1),(x,y)$ 都是包含这个黑色格子的，统计这些子矩阵中有多少黑色格子，加到答案中。  
代码实现时，注意不要重复统计，可以用哈希表 $\textit{vis}$ 来记录统计过的子矩阵左上角。  
最后不含黑色格子的子矩阵个数就是  

$$
(m-1)\cdot (n-1) - \text{len}(\textit{vis})
$$

```go 
func countBlackBlocks(m, n int, coordinates [][]int) []int64 {
	type pair struct{ x, y int }
	set := make(map[pair]int, len(coordinates))
	for _, p := range coordinates {
		set[pair{p[0], p[1]}] = 1
	}

	arr := make([]int64, 5)
	vis := make(map[pair]bool, len(set)*4)
	for _, p := range coordinates {
		x, y := p[0], p[1]
		for i := max(x-1, 0); i <= x && i < m-1; i++ {
			for j := max(y-1, 0); j <= y && j < n-1; j++ {
				if !vis[pair{i, j}] {
					vis[pair{i, j}] = true
					cnt := set[pair{i, j}] + set[pair{i, j + 1}] +
						set[pair{i + 1, j}] + set[pair{i + 1, j + 1}]
					arr[cnt]++
				}
			}
		}
	}
	arr[0] = int64(m-1)*int64(n-1) - int64(len(vis))
	return arr
}

func max(a, b int) int { if b > a { return b }; return a }
```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(k)$，其中 $k$ 为 $\textit{coordinates}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。