#### 题目

<p>给你两个长度为 <code>n</code> 下标从 <strong>0</strong> 开始的整数数组 <code>cost</code> 和 <code>time</code> ，分别表示给 <code>n</code> 堵不同的墙刷油漆需要的开销和时间。你有两名油漆匠：</p>

<ul>
	<li>一位需要 <strong>付费</strong> 的油漆匠，刷第 <code>i</code> 堵墙需要花费 <code>time[i]</code> 单位的时间，开销为 <code>cost[i]</code> 单位的钱。</li>
	<li>一位 <strong>免费</strong> 的油漆匠，刷 <strong>任意</strong> 一堵墙的时间为 <code>1</code> 单位，开销为 <code>0</code> 。但是必须在付费油漆匠 <strong>工作</strong> 时，免费油漆匠才会工作。</li>
</ul>

<p>请你返回刷完 <code>n</code> 堵墙最少开销为多少。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>cost = [1,2,3,2], time = [1,2,3,2]
<b>输出：</b>3
<strong>解释：</strong>下标为 0 和 1 的墙由付费油漆匠来刷，需要 3 单位时间。同时，免费油漆匠刷下标为 2 和 3 的墙，需要 2 单位时间，开销为 0 。总开销为 1 + 2 = 3 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>cost = [2,3,4,2], time = [1,1,1,1]
<b>输出：</b>4
<b>解释：</b>下标为 0 和 3 的墙由付费油漆匠来刷，需要 2 单位时间。同时，免费油漆匠刷下标为 1 和 2 的墙，需要 2 单位时间，开销为 0 。总开销为 2 + 2 = 4 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= cost.length &lt;= 500</code></li>
	<li><code>cost.length == time.length</code></li>
	<li><code>1 &lt;= cost[i] &lt;= 10<sup>6</sup></code></li>
	<li><code>1 &lt;= time[i] &lt;= 500</code></li>
</ul>

#### 思路

用「选或不选」的思路来思考：

- 如果付费刷第 $n-1$ 堵墙，那么问题变成：刷前 $n-2$ 堵墙，且付费时间和为 $\textit{time}[n-1]$，免费时间和 $0$ 的最少开销（开销指 $\textit{cost}$ 的和）。
- 如果免费刷第 $n-1$ 堵墙，那么问题变成：刷前 $n-2$ 堵墙，且付费时间和为 $0$，免费时间和为 $1$ 的最少开销。

因此，定义 $\textit{dfs}(i,j,k)$ 表示刷前 $i$ 堵墙，且付费时间和为 $j$，免费时间和为 $k$ 的最少开销。递归到终点时，如果 $j\ge k$，说明这种方案是合法的，否则不合法。

但是，这样定义的话，状态个数就太多了，需要优化。由于最后是比较的 $j$ 和 $k$ 的「相对大小」，那么不妨把 $j$ 重新定义为【付费时间和】减去【免费时间和】，这样递归到终点时，如果 $j\ge 0$，说明这种方案是合法的，否则不合法。这样一来，状态个数就大大减少了。

分类讨论：

- 如果付费刷第 $i$ 堵墙：$\textit{dfs}(i,j) = \textit{dfs}(i-1,j+\textit{time}[i]) +\textit{cost}[i]$。
- 如果免费刷第 $i$ 堵墙：$\textit{dfs}(i,j) = \textit{dfs}(i-1,j-1)$。

两种情况取最小值：

$$
\textit{dfs}(i,j) = \min(\textit{dfs}(i-1,j+\textit{time}[i]) +\textit{cost}[i], \textit{dfs}(i-1,j-1))
$$

递归边界：如果 $j>i$，那么剩余的墙都可以免费刷，即 $\textit{dfs}(i,j) = 0$，否则 $\textit{dfs}(-1,j) = \infty$。

递归入口：$\textit{dfs}(n-1,0)$。

```go 

```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{cost}$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n^2)$，单个状态的计算时间为 $\mathcal{O}(1)$，因此时间复杂度为 $\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(n^2)$。
