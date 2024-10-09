#### 题目

<p>给你两个正整数 <code>n</code> 和 <code>m</code> 。</p>

<p>现定义两个整数 <code>num1</code> 和 <code>num2</code> ，如下所示：</p>

<ul>
	<li><code>num1</code>：范围 <code>[1, n]</code> 内所有 <strong>无法被 </strong><code>m</code><strong> 整除</strong> 的整数之和。</li>
	<li><code>num2</code>：范围 <code>[1, n]</code> 内所有 <strong>能够被 </strong><code>m</code><strong> 整除</strong> 的整数之和。</li>
</ul>

<p>返回整数 <code>num1 - num2</code> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre><strong>输入：</strong>n = 10, m = 3
<strong>输出：</strong>19
<strong>解释：</strong>在这个示例中：
- 范围 [1, 10] 内无法被 3 整除的整数为 [1,2,4,5,7,8,10] ，num1 = 这些整数之和 = 37 。
- 范围 [1, 10] 内能够被 3 整除的整数为 [3,6,9] ，num2 = 这些整数之和 = 18 。
返回 37 - 18 = 19 作为答案。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre><strong>输入：</strong>n = 5, m = 6
<strong>输出：</strong>15
<strong>解释：</strong>在这个示例中：
- 范围 [1, 5] 内无法被 6 整除的整数为 [1,2,3,4,5] ，num1 = 这些整数之和 =  15 。
- 范围 [1, 5] 内能够被 6 整除的整数为 [] ，num2 = 这些整数之和 = 0 。
返回 15 - 0 = 15 作为答案。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre><strong>输入：</strong>n = 5, m = 1
<strong>输出：</strong>-15
<strong>解释：</strong>在这个示例中：
- 范围 [1, 5] 内无法被 1 整除的整数为 [] ，num1 = 这些整数之和 = 0 。 
- 范围 [1, 5] 内能够被 1 整除的整数为 [1,2,3,4,5] ，num2 = 这些整数之和 = 15 。
返回 0 - 15 = -15 作为答案。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n, m <= 1000</code></li>
</ul>

#### 思路

设 $k = \left\lfloor\dfrac{n}{m}\right\rfloor$。
$\textit{num}_2$ 是 $[1,n]$ 内的 $m$ 的倍数之和，即

$$
\begin{aligned}
&m + 2m + \cdots + km\\
=\ & (1+2+\cdots+k)\cdot m\\
=\ & \dfrac{k(k+1)}{2}\cdot m
\end{aligned}

$$

$\textit{num}_1$ 相当于 $(1+2+\cdots+n) - \textit{num}_2$，所以

$$
\begin{aligned}
&\textit{num}_1 - \textit{num}_2\\
=\ & (1+2+\cdots+n) - \textit{num}_2 \cdot 2\\
=\ & \dfrac{n(n+1)}{2} - k(k+1)m
\end{aligned}

$$

```go  
func differenceOfSums(n int, m int) (ans int) {
	k := n / m
	return n*(n+1)/2 - k*(k+1)*m
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$ 。
