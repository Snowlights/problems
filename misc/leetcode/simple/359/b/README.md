#### 题目

<p>给你两个整数 <code>n</code> 和 <code>k</code> 。</p>

<p>对于一个由 <strong>不同</strong> 正整数组成的数组，如果其中不存在任何求和等于 k 的不同元素对，则称其为 <strong>k-avoiding</strong> 数组。</p>

<p>返回长度为 <code>n</code> 的 <strong>k-avoiding</strong> 数组的可能的最小总和。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>n = 5, k = 4
<strong>输出：</strong>18
<strong>解释：</strong>设若 k-avoiding 数组为 [1,2,4,5,6] ，其元素总和为 18 。
可以证明不存在总和小于 18 的 k-avoiding 数组。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>n = 2, k = 6
<strong>输出：</strong>3
<strong>解释：</strong>可以构造数组 [1,2] ，其元素总和为 3 。
可以证明不存在总和小于 3 的 k-avoiding 数组。 
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n, k &lt;= 50</code></li>
</ul>

#### 思路

对于 $[1,k-1]$ 内的数字：
- $1$ 和 $k-1$ 只能选其中一个；
- $2$ 和 $k-2$ 只能选其中一个；
- $3$ 和 $k-3$ 只能选其中一个；
- ……
- 一直到 $\left\lfloor\dfrac{k}{2}\right\rfloor$，无论 $k$ 是奇数还是偶数，它都可以选。
  
设 $m=\min\left(\left\lfloor\dfrac{k}{2}\right\rfloor, n\right)$，那么答案的第一段是从 $1$ 到 $m$，元素和为

$$
\dfrac{m(m+1)}{2}
$$

此时还剩下 $n-m$ 个数，只能从 $k$ 开始往后选，那么答案的第二段是从 $k$ 到 $k+n-m-1$，元素和为

$$
\dfrac{(2k+n-m-1)(n-m)}{2}
$$

所以答案为

$$
\dfrac{m(m+1) + (2k+n-m-1)(n-m)}{2}
$$


```go 
func minimumSum(n, k int) int {
	m := min(k/2, n)
	return (m*(m+1) + (k*2+n-m-1)*(n-m)) / 2
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$ 。