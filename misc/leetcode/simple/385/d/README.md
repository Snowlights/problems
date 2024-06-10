#### 题目

<p>给你一个大小为 <code>m x n</code> 、下标从 <strong>0</strong> 开始的二维矩阵 <code>mat</code> 。在每个单元格，你可以按以下方式生成数字：</p>

<ul>
	<li>最多有 <code>8</code> 条路径可以选择：东，东南，南，西南，西，西北，北，东北。</li>
	<li>选择其中一条路径，沿着这个方向移动，并且将路径上的数字添加到正在形成的数字后面。</li>
	<li>注意，每一步都会生成数字，例如，如果路径上的数字是 <code>1, 9, 1</code>，那么在这个方向上会生成三个数字：<code>1, 19, 191</code> 。</li>
</ul>

<p>返回在遍历矩阵所创建的所有数字中，出现频率最高的、<strong>大于</strong> <code>10</code>的<span data-keyword="prime-number">质数</span>；如果不存在这样的质数，则返回 <code>-1</code><em> </em>。如果存在多个出现频率最高的质数，那么返回其中最大的那个。</p>

<p><strong>注意：</strong>移动过程中不允许改变方向。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>
<strong><img alt="" src="https://assets.leetcode.com/uploads/2024/02/15/south" style="width: 641px; height: 291px;" /> </strong>

<pre>
<strong>
输入：</strong>mat = [[1,1],[9,9],[1,1]]
<strong>输出：</strong>19
<strong>解释：</strong> 
从单元格 (0,0) 出发，有 3 个可能的方向，这些方向上可以生成的大于 10 的数字有：
东方向: [11], 东南方向: [19], 南方向: [19,191] 。
从单元格 (0,1) 出发，所有可能方向上生成的大于 10 的数字有：[19,191,19,11] 。
从单元格 (1,0) 出发，所有可能方向上生成的大于 10 的数字有：[99,91,91,91,91] 。
从单元格 (1,1) 出发，所有可能方向上生成的大于 10 的数字有：[91,91,99,91,91] 。
从单元格 (2,0) 出发，所有可能方向上生成的大于 10 的数字有：[11,19,191,19] 。
从单元格 (2,1) 出发，所有可能方向上生成的大于 10 的数字有：[11,19,19,191] 。
在所有生成的数字中，出现频率最高的质数是 19 。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>mat = [[7]]
<strong>输出：</strong>-1
<strong>解释：</strong>唯一可以生成的数字是 7 。它是一个质数，但不大于 10 ，所以返回 -1 。</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>mat = [[9,7,8],[4,6,5],[2,8,6]]
<strong>输出：</strong>97
<strong>解释：</strong> 
从单元格 (0,0) 出发，所有可能方向上生成的大于 10 的数字有: [97,978,96,966,94,942] 。
从单元格 (0,1) 出发，所有可能方向上生成的大于 10 的数字有: [78,75,76,768,74,79] 。
从单元格 (0,2) 出发，所有可能方向上生成的大于 10 的数字有: [85,856,86,862,87,879] 。
从单元格 (1,0) 出发，所有可能方向上生成的大于 10 的数字有: [46,465,48,42,49,47] 。
从单元格 (1,1) 出发，所有可能方向上生成的大于 10 的数字有: [65,66,68,62,64,69,67,68] 。
从单元格 (1,2) 出发，所有可能方向上生成的大于 10 的数字有: [56,58,56,564,57,58] 。
从单元格 (2,0) 出发，所有可能方向上生成的大于 10 的数字有: [28,286,24,249,26,268] 。
从单元格 (2,1) 出发，所有可能方向上生成的大于 10 的数字有: [86,82,84,86,867,85] 。
从单元格 (2,2) 出发，所有可能方向上生成的大于 10 的数字有: [68,682,66,669,65,658] 。
在所有生成的数字中，出现频率最高的质数是 97 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>m == mat.length</code></li>
	<li><code>n == mat[i].length</code></li>
	<li><code>1 &lt;= m, n &lt;= 6</code></li>
	<li><code>1 &lt;= mat[i][j] &lt;= 9</code></li>
</ul>

#### 思路

对于每个单元格，枚举八个方向，生成数字，统计其中质数个数。  
最后返回出现次数最多的质数，如果有多个这样的质数，返回最大的那个。

```go [sol]
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func mostFrequentPrime(mat [][]int) int {
	dirs := []struct{ x, y int }{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	m, n := len(mat), len(mat[0])
	cnt := map[int]int{}
	for i, row := range mat {
		for j, v := range row {
			for _, d := range dirs {
				x, y, v := i+d.x, j+d.y, v
				for 0 <= x && x < m && 0 <= y && y < n {
					v = v*10 + mat[x][y]
					// 如果 v 在 cnt 中，那么 v 一定是质数
					if cnt[v] > 0 || isPrime(v) {
						cnt[v]++
					}
					x += d.x
					y += d.y
				}
			}
		}
	}

	ans, maxCnt := -1, 0
	for v, c := range cnt {
		if c > maxCnt {
			ans, maxCnt = v, c
		} else if c == maxCnt {
			ans = max(ans, v)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mnk\cdot 10^{k/2})$，其中 $m$ 和 $n$ 分别为 $\textit{mat}$ 的行数和列数，$k=\max(m,n)$。总共有 $\mathcal{O}(mnk)$ 个数，判断质数需要 $\mathcal{O}(10^{k/2})$ 的时间。
- 空间复杂度：$\mathcal{O}(mnk)$。

