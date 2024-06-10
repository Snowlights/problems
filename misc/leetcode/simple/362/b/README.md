#### 题目

<p>给你四个整数 <code>sx</code>、<code>sy</code>、<code>fx</code>、<code>fy</code>&nbsp; 以及一个 <strong>非负整数</strong> <code>t</code> 。</p>
<p>在一个无限的二维网格中，你从单元格 <code>(sx, sy)</code> 开始出发。每一秒，你 <strong>必须</strong> 移动到任一与之前所处单元格相邻的单元格中。</p>
<p>如果你能在 <strong>恰好 </strong><code>t</code><strong> 秒</strong> 后到达单元格<em> </em><code>(fx, fy)</code> ，返回 <code>true</code> ；否则，返回&nbsp; <code>false</code> 。</p>
<p>单元格的 <strong>相邻单元格</strong> 是指该单元格周围与其至少共享一个角的 8 个单元格。你可以多次访问同一个单元格。</p>
<p>&nbsp;</p>
<p><strong class="example">示例 1：</strong></p>
<img alt src="https://assets.leetcode.com/uploads/2023/08/05/example2.svg" style="width: 443px; height: 243px;" />
<pre>
<strong>输入：</strong>sx = 2, sy = 4, fx = 7, fy = 7, t = 6
<strong>输出：</strong>true
<strong>解释：</strong>从单元格 (2, 4) 开始出发，穿过上图标注的单元格，可以在恰好 6 秒后到达单元格 (7, 7) 。 
</pre>
<p><strong class="example">示例 2：</strong></p>
<img alt src="https://assets.leetcode.com/uploads/2023/08/05/example1.svg" style="width: 383px; height: 202px;" />
<pre>
<strong>输入：</strong>sx = 3, sy = 1, fx = 7, fy = 3, t = 3
<strong>输出：</strong>false
<strong>解释：</strong>从单元格 (3, 1) 开始出发，穿过上图标注的单元格，至少需要 4 秒后到达单元格 (7, 3) 。 因此，无法在 3 秒后到达单元格 (7, 3) 。
</pre>
<p>&nbsp;</p>
<p><strong>提示：</strong></p>
<ul>
<li><code>1 &lt;= sx, sy, fx, fy &lt;= 10<sup>9</sup></code></li>
<li><code>0 &lt;= t &lt;= 10<sup>9</sup></code></li>
</ul>

#### 思路

由于可以往 $8$ 个方向走，那么最快可以在  

$$
\max\{|sx-fx|, |sy-fy|\}
$$

秒后到达终点（先斜着走再直走）。  
上式只要小于等于 $t$ 就能恰好到达终点。因为我们可以在到达终点附近时，在终点周围不断「绕圈」消耗时间，这样可以直到最后一秒才走到终点。  
特殊情况：如果起点和终点重合，那么除了 $t=1$ 的情况是无法到达终点的，其余情况我们总是可以到达终点。

```go 
func isReachableAtTime(sx, sy, fx, fy, t int) bool {
	if sx == fx && sy == fy {
		return t != 1
	}
	return max(abs(sx-fx), abs(sy-fy)) <= t
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。