#### 题目  

<p>给你一个长度为 <code>n</code> 的字符串 <code>moves</code> ，该字符串仅由字符 <code>&#39;L&#39;</code>、<code>&#39;R&#39;</code> 和 <code>&#39;_&#39;</code> 组成。字符串表示你在一条原点为 <code>0</code> 的数轴上的若干次移动。</p>

<p>你的初始位置就在原点（<code>0</code>），第 <code>i</code> 次移动过程中，你可以根据对应字符选择移动方向：</p>

<ul>
	<li>如果 <code>moves[i] = &#39;L&#39;</code> 或 <code>moves[i] = &#39;_&#39;</code> ，可以选择向左移动一个单位距离</li>
	<li>如果 <code>moves[i] = &#39;R&#39;</code> 或 <code>moves[i] = &#39;_&#39;</code> ，可以选择向右移动一个单位距离</li>
</ul>

<p>移动 <code>n</code> 次之后，请你找出可以到达的距离原点 <strong>最远</strong> 的点，并返回 <strong>从原点到这一点的距离</strong> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>moves = &#34;L_RL__R&#34;
<strong>输出：</strong>3
<strong>解释：</strong>可以到达的距离原点 0 最远的点是 -3 ，移动的序列为 &#34;LLRLLLR&#34; 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>moves = &#34;_R__LL_&#34;
<strong>输出：</strong>5
<strong>解释：</strong>可以到达的距离原点 0 最远的点是 -5 ，移动的序列为 &#34;LRLLLLL&#34; 。
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>moves = &#34;_______&#34;
<strong>输出：</strong>7
<strong>解释：</strong>可以到达的距离原点 0 最远的点是 7 ，移动的序列为 &#34;RRRRRRR&#34; 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= moves.length == n &lt;= 50</code></li>
	<li><code>moves</code> 仅由字符 <code>&#39;L&#39;</code>、<code>&#39;R&#39;</code> 和 <code>&#39;_&#39;</code> 组成</li>
</ul>
 
#### 思路  

先只考虑 L 和 R，也就是用 R 的个数减去 L 的个数，得到当前位置 $x$。
- 如果 $x>0$，那么所有的下划线都应该向右走；
- 如果 $x<0$，那么所有的下划线都应该向左走；
- 如果 $x=0$，向左向右都可以。

```go 
func furthestDistanceFromOrigin(moves string) int {
	return abs(strings.Count(moves, "R")-strings.Count(moves, "L")) + strings.Count(moves, "_")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{moves}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。