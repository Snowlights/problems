#### 题目

<p>考试中有 <code>n</code> 种类型的题目。给你一个整数 <code>target</code> 和一个下标从 <strong>0</strong> 开始的二维整数数组 <code>types</code> ，其中 <code>types[i] = [count<sub>i</sub>, marks<sub>i</sub>] </code>表示第 <code>i</code> 种类型的题目有 <code>count<sub>i</sub></code> 道，每道题目对应 <code>marks<sub>i</sub></code> 分。</p>

<p>返回你在考试中恰好得到 <code>target</code> 分的方法数。由于答案可能很大，结果需要对 <code>10<sup>9</sup> +7</code> 取余。</p>

<p><strong>注意</strong>，同类型题目无法区分。</p>

<ul>
	<li>比如说，如果有 <code>3</code> 道同类型题目，那么解答第 <code>1</code> 和第 <code>2</code> 道题目与解答第 <code>1</code> 和第 <code>3</code> 道题目或者第 <code>2</code> 和第 <code>3</code> 道题目是相同的。</li>
</ul>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>target = 6, types = [[6,1],[3,2],[2,3]]
<strong>输出：</strong>7
<strong>解释：</strong>要获得 6 分，你可以选择以下七种方法之一：
- 解决 6 道第 0 种类型的题目：1 + 1 + 1 + 1 + 1 + 1 = 6
- 解决 4 道第 0 种类型的题目和 1 道第 1 种类型的题目：1 + 1 + 1 + 1 + 2 = 6
- 解决 2 道第 0 种类型的题目和 2 道第 1 种类型的题目：1 + 1 + 2 + 2 = 6
- 解决 3 道第 0 种类型的题目和 1 道第 2 种类型的题目：1 + 1 + 1 + 3 = 6
- 解决 1 道第 0 种类型的题目、1 道第 1 种类型的题目和 1 道第 2 种类型的题目：1 + 2 + 3 = 6
- 解决 3 道第 1 种类型的题目：2 + 2 + 2 = 6
- 解决 2 道第 2 种类型的题目：3 + 3 = 6
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>target = 5, types = [[50,1],[50,2],[50,5]]
<strong>输出：</strong>4
<strong>解释：</strong>要获得 5 分，你可以选择以下四种方法之一：
- 解决 5 道第 0 种类型的题目：1 + 1 + 1 + 1 + 1 = 5
- 解决 3 道第 0 种类型的题目和 1 道第 1 种类型的题目：1 + 1 + 1 + 2 = 5
- 解决 1 道第 0 种类型的题目和 2 道第 1 种类型的题目：1 + 2 + 2 = 5
- 解决 1 道第 2 种类型的题目：5
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>target = 18, types = [[6,1],[3,2],[2,3]]
<strong>输出：</strong>1
<strong>解释：</strong>只有回答所有题目才能获得 18 分。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= target <= 1000</code></li>
	<li><code>n == types.length</code></li>
	<li><code>1 <= n <= 50</code></li>
	<li><code>types[i].length == 2</code></li>
	<li><code>1 <= count<sub>i</sub>, marks<sub>i</sub> <= 50</code></li>
</ul>

#### 思路

分组背包模板题。
定义 $f[i][j]$ 表示用前 $i$ 种题目恰好组成 $j$ 分的方案数。
对于第 $i$ 种题目，枚举做 $k$ 道题目，则子问题为「前 $i-1$ 种题目恰好组成 $j-k\cdot \textit{marks}_i$ 分的方案数」，因此有

$$
f[i][j] = \sum\limits_{k=0} f[i-1][j-k\cdot \textit{marks}_i]

$$

注意 $k$ 不能超过 $\textit{count}_i$，且 $j-k\cdot \textit{marks}_i\ge 0$。代码实现时可以像 0-1 背包那样，压缩成一维，

> 注：滚动优化后，$k=0$ 就是 $f[j]$，无需计算。

```go

```

#### 复杂度分析

- 时间复杂度：$O(TS)$，其中 $T$ 为 $\textit{target}$，$S$ 为所有 $\textit{count}_i$ 之和。
- 空间复杂度：$O(T)$。
