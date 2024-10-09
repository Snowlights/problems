#### 题目  

<p>袋子中装有一些物品，每个物品上都标记着数字 <code>1</code> 、<code>0</code> 或 <code>-1</code> 。</p>

<p>给你四个非负整数 <code>numOnes</code> 、<code>numZeros</code> 、<code>numNegOnes</code> 和 <code>k</code> 。</p>

<p>袋子最初包含：</p>

<ul>
	<li><code>numOnes</code> 件标记为 <code>1</code> 的物品。</li>
	<li><code>numZeroes</code> 件标记为 <code>0</code> 的物品。</li>
	<li><code>numNegOnes</code> 件标记为 <code>-1</code> 的物品。</li>
</ul>

<p>现计划从这些物品中恰好选出 <code>k</code> 件物品。返回所有可行方案中，物品上所标记数字之和的最大值。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>numOnes = 3, numZeros = 2, numNegOnes = 0, k = 2
<strong>输出：</strong>2
<strong>解释：</strong>袋子中的物品分别标记为 {1, 1, 1, 0, 0} 。取 2 件标记为 1 的物品，得到的数字之和为 2 。
可以证明 2 是所有可行方案中的最大值。</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>numOnes = 3, numZeros = 2, numNegOnes = 0, k = 4
<strong>输出：</strong>3
<strong>解释：</strong>袋子中的物品分别标记为 {1, 1, 1, 0, 0} 。取 3 件标记为 1 的物品，1 件标记为 0 的物品，得到的数字之和为 3 。
可以证明 3 是所有可行方案中的最大值。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 &lt;= numOnes, numZeros, numNegOnes &lt;= 50</code></li>
	<li><code>0 &lt;= k &lt;= numOnes + numZeros + numNegOnes</code></li>
</ul>
 
#### 思路  

贪心，按照 $1,0,-1$ 的顺序选：
- 先选 $1$，如果 $k\le \textit{numOnes}$ 那么答案就是 $k$。
- 再选 $0$，如果 $k\le \textit{numOnes}+\textit{numZeros}$ 那么答案为 $\textit{numOnes}$。
- 最后选 $-1$（题目要求恰好选 $k$ 个），那么剩余必须选 $k-\textit{numOnes}-\textit{numZeros}$ 个 $-1$，答案为

$$
\textit{numOnes} + (k-\textit{numOnes}-\textit{numZeros})  \cdot  (-1)= \textit{numOnes} \cdot 2 + \textit{numZeros} - k
$$

```go 

```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。