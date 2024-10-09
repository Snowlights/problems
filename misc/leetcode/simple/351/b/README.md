#### 题目  

<p>给你两个整数：<code>num1</code> 和 <code>num2</code> 。</p>

<p>在一步操作中，你需要从范围 <code>[0, 60]</code> 中选出一个整数 <code>i</code> ，并从 <code>num1</code> 减去 <code>2<sup>i</sup> + num2</code> 。</p>

<p>请你计算，要想使 <code>num1</code> 等于 <code>0</code> 需要执行的最少操作数，并以整数形式返回。</p>

<p>如果无法使 <code>num1</code> 等于 <code>0</code> ，返回 <code>-1</code> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>num1 = 3, num2 = -2
<strong>输出：</strong>3
<strong>解释：</strong>可以执行下述步骤使 3 等于 0 ：
- 选择 i = 2 ，并从 3 减去 2<sup>2</sup> + (-2) ，num1 = 3 - (4 + (-2)) = 1 。
- 选择 i = 2 ，并从 1 减去 2<sup>2</sup> + (-2) ，num1 = 1 - (4 + (-2)) = -1 。
- 选择 i = 0 ，并从 -1 减去 2<sup>0</sup> + (-2) ，num1 = (-1) - (1 + (-2)) = 0 。
可以证明 3 是需要执行的最少操作数。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>num1 = 5, num2 = 7
<strong>输出：</strong>-1
<strong>解释：</strong>可以证明，执行操作无法使 5 等于 0 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= num1 &lt;= 10<sup>9</sup></code></li>
	<li><code>-10<sup>9</sup> &lt;= num2 &lt;= 10<sup>9</sup></code></li>
</ul>
 
#### 思路  

#### 提示 1
从小到大枚举答案。
#### 提示 2
假设操作了 $k$ 次，那么操作后 $\textit{num}_1$ 变成 $\textit{num}_1 - \textit{num}_2\cdot k$ 再减去 $k$ 个 $2^i$。
此时问题变成：$\textit{num}_1 - \textit{num}_2\cdot k$ 能否拆分成 $k$ 个 $2^i$ 之和？
#### 提示 3
设 $x=\textit{num}_1 - \textit{num}_2\cdot k$。
- 如果 $x<0$，无解。
- 否则如果 $x<k$，那么即使每次操作取 $i=0$，也至少要把 $x$ 拆分成 $k$ 个 $1$ 之和，这是不可能的。
- 否则如果 $x$ 中二进制 $1$ 的个数大于 $k$，也无法拆分成 $k$ 个 $2^i$ 之和，无解。
- 否则分解方案一定存在，返回 $k$。（因为可以把一个 $2^j$ 分解成两个 $2^{j-1}$，所以分解出的 $2^i$ 的**个数**可以从「$x$ 中二进制 $1$ 的个数」一直到 $x$，在这个范围内的分解方案都是存在的。）
  代码实现时，如果出现 $x<k$ 的情况，说明 $\textit{num}_2\ge 0$，那么对于更大的 $k$，$x$ 无法变得更大，所以后面都无解，直接退出循环。

```go 
func makeTheIntegerZero(x, y int) int {

	for k := 1; k <= x-k*y; k++ {
		if k >= bits.OnesCount(uint(x-k*y)) {
			return k
		}
	}

	return -1
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(\log(x+|y|))$，其中 $x=\textit{num}_1,y=\textit{num}_2$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。