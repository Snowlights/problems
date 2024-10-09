#### 题目  

<p>给你一个长度为 <code>n</code> 下标从 <strong>0</strong> 开始的整数数组 <code>receiver</code> 和一个整数 <code>k</code> 。</p>

<p>总共有 <code>n</code> 名玩家，玩家 <strong>编号</strong> 互不相同，且为 <code>[0, n - 1]</code> 中的整数。这些玩家玩一个传球游戏，<code>receiver[i]</code> 表示编号为 <code>i</code> 的玩家会传球给编号为 <code>receiver[i]</code> 的玩家。玩家可以传球给自己，也就是说 <code>receiver[i]</code> 可能等于 <code>i</code> 。</p>

<p>你需要从 <code>n</code> 名玩家中选择一名玩家作为游戏开始时唯一手中有球的玩家，球会被传 <strong>恰好</strong> <code>k</code> 次。</p>

<p>如果选择编号为 <code>x</code> 的玩家作为开始玩家，定义函数 <code>f(x)</code> 表示从编号为 <code>x</code> 的玩家开始，<code>k</code> 次传球内所有接触过球玩家的编号之 <strong>和</strong> ，如果有玩家多次触球，则 <strong>累加多次</strong> 。换句话说， <code>f(x) = x + receiver[x] + receiver[receiver[x]] + ... + receiver<sup>(k)</sup>[x]</code> 。</p>

<p>你的任务时选择开始玩家 <code>x</code> ，目的是<strong> 最大化</strong> <code>f(x)</code> 。</p>

<p>请你返回函数的 <strong>最大值</strong> 。</p>

<p><strong>注意：</strong><code>receiver</code> 可能含有重复元素。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<table border="1" cellspacing="3" style="border-collapse: separate; text-align: center;">
	<tbody>
		<tr>
			<th style="padding: 5px; border: 1px solid black;">传递次数</th>
			<th style="padding: 5px; border: 1px solid black;">传球者编号</th>
			<th style="padding: 5px; border: 1px solid black;">接球者编号</th>
			<th style="padding: 5px; border: 1px solid black;">x + 所有接球者编号</th>
		</tr>
		<tr>
			<td style="padding: 5px; border: 1px solid black;"> </td>
			<td style="padding: 5px; border: 1px solid black;"> </td>
			<td style="padding: 5px; border: 1px solid black;"> </td>
			<td style="padding: 5px; border: 1px solid black;">2</td>
		</tr>
		<tr>
			<td style="padding: 5px; border: 1px solid black;">1</td>
			<td style="padding: 5px; border: 1px solid black;">2</td>
			<td style="padding: 5px; border: 1px solid black;">1</td>
			<td style="padding: 5px; border: 1px solid black;">3</td>
		</tr>
		<tr>
			<td style="padding: 5px; border: 1px solid black;">2</td>
			<td style="padding: 5px; border: 1px solid black;">1</td>
			<td style="padding: 5px; border: 1px solid black;">0</td>
			<td style="padding: 5px; border: 1px solid black;">3</td>
		</tr>
		<tr>
			<td style="padding: 5px; border: 1px solid black;">3</td>
			<td style="padding: 5px; border: 1px solid black;">0</td>
			<td style="padding: 5px; border: 1px solid black;">2</td>
			<td style="padding: 5px; border: 1px solid black;">5</td>
		</tr>
		<tr>
			<td style="padding: 5px; border: 1px solid black;">4</td>
			<td style="padding: 5px; border: 1px solid black;">2</td>
			<td style="padding: 5px; border: 1px solid black;">1</td>
			<td style="padding: 5px; border: 1px solid black;">6</td>
		</tr>
	</tbody>
</table>

<p> </p>

<pre><b>输入：</b>receiver = [2,0,1], k = 4
<b>输出：</b>6
<b>解释：</b>上表展示了从编号为 x = 2 开始的游戏过程。
从表中可知，f(2) 等于 6 。
6 是能得到最大的函数值。
所以输出为 6 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<table border="1" cellspacing="3" style="border-collapse: separate; text-align: center;">
	<tbody>
		<tr>
			<th style="padding: 5px; border: 1px solid black;">传递次数</th>
			<th style="padding: 5px; border: 1px solid black;">传球者编号</th>
			<th style="padding: 5px; border: 1px solid black;">接球者编号</th>
			<th style="padding: 5px; border: 1px solid black;">x + 所有接球者编号</th>
		</tr>
		<tr>
			<td style="padding: 5px; border: 1px solid black;"> </td>
			<td style="padding: 5px; border: 1px solid black;"> </td>
			<td style="padding: 5px; border: 1px solid black;"> </td>
			<td style="padding: 5px; border: 1px solid black;">4</td>
		</tr>
		<tr>
			<td style="padding: 5px; border: 1px solid black;">1</td>
			<td style="padding: 5px; border: 1px solid black;">4</td>
			<td style="padding: 5px; border: 1px solid black;">3</td>
			<td style="padding: 5px; border: 1px solid black;">7</td>
		</tr>
		<tr>
			<td style="padding: 5px; border: 1px solid black;">2</td>
			<td style="padding: 5px; border: 1px solid black;">3</td>
			<td style="padding: 5px; border: 1px solid black;">2</td>
			<td style="padding: 5px; border: 1px solid black;">9</td>
		</tr>
		<tr>
			<td style="padding: 5px; border: 1px solid black;">3</td>
			<td style="padding: 5px; border: 1px solid black;">2</td>
			<td style="padding: 5px; border: 1px solid black;">1</td>
			<td style="padding: 5px; border: 1px solid black;">10</td>
		</tr>
	</tbody>
</table>

<p> </p>

<pre><b>输入：</b>receiver = [1,1,1,2,3], k = 3
<b>输出：</b>10
<b>解释：</b>上表展示了从编号为 x = 4 开始的游戏过程。
从表中可知，f(4) 等于 10 。
10 是能得到最大的函数值。
所以输出为 10 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= receiver.length == n &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= receiver[i] &lt;= n - 1</code></li>
	<li><code>1 &lt;= k &lt;= 10<sup>10</sup></code></li>
</ul>
 
#### 思路

利用倍增算法，预处理每个节点 $x$ 的第 $2^i$ 个祖先节点，以及从 $x$ 的父节点到 $x$ 的第 $2^i$ 个祖先节点的节点编号之和。
最后枚举起点 $x$，一边向上跳一边累加节点编号。

```go 
func getMaxFunctionValue(receiver []int, K int64) int64 {
	type pair struct{ pa, sum int }
	n, m := len(receiver), bits.Len(uint(K))
	pa := make([][]pair, n)
	for i, p := range receiver {
		pa[i] = make([]pair, m)
		pa[i][0] = pair{p, p}
	}
	for i := 0; i+1 < m; i++ {
		for x := range pa {
			p := pa[x][i]
			pp := pa[p.pa][i]
			pa[x][i+1] = pair{pp.pa, p.sum + pp.sum} // 合并节点值之和
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		x := i
		sum := i // 节点值之和，初始为节点 i
		for k := uint(K); k > 0; k &= k - 1 {
			p := pa[x][bits.TrailingZeros(k)]
			sum += p.sum
			x = p.pa
		}
		ans = max(ans, sum)
	}
	return int64(ans)
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n\log k)$，其中 $n$ 为 $\textit{receiver}$ 的长度。
- 空间复杂度：$\mathcal{O}(n\log k)$。