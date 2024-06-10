#### 题目  

<p>给你一个正整数 <code>n</code> ，请你计算在 <code>[1，n]</code> 范围内能被 <code>3</code>、<code>5</code>、<code>7</code> 整除的所有整数之和。</p>

<p>返回一个整数，用于表示给定范围内所有满足约束条件的数字之和。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>n = 7
<strong>输出：</strong>21
<strong>解释：</strong>在 <code>[1, 7]</code> 范围内能被 3、<code>5、</code><code>7 整除的所有整数分别是</code><code> 3、5、6、7</code> 。数字之和为 <code>21</code> 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>n = 10
<strong>输出：</strong>40
<strong>解释：</strong>在 <code>[1, 10]</code> 范围内能被 3、<code>5、</code><code>7 整除的所有整数分别是</code><code> 3、5、6、7、9、10</code> 。数字之和为 <code>40</code> 。
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>n = 9
<strong>输出：</strong>30
<strong>解释：</strong>在 <code>[1, 9]</code> 范围内能被 3、<code>5、</code><code>7 整除的所有整数分别是</code><code> 3、5、6、7、9</code> 。数字之和为 <code>30</code> 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 10<sup>3</sup></code></li>
</ul>
 
#### 思路  

在 $[1,n]$ 中 $m$ 的倍数有 $c = \left\lfloor\dfrac{n}{m}\right\rfloor$ 个，结合等差数列求和公式，这些数的和为

$$
\dfrac{(1+c)\cdot c}{2} \cdot m
$$

再结合容斥原理，可以算出 $3$ 或 $5$ 或 $7$ 的倍数之和。

```go 

```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(1)$，
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。