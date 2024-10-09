### 题目  

<p>给你两个 <strong>正</strong> 整数 <code>n</code> 和 <code>x</code> 。</p>

<p>请你返回将<em> </em><code>n</code> 表示成一些 <strong>互不相同</strong> 正整数的<em> </em><code>x</code> 次幂之和的方案数。换句话说，你需要返回互不相同整数 <code>[n<sub>1</sub>, n<sub>2</sub>, ..., n<sub>k</sub>]</code> 的集合数目，满足 <code>n = n<sub>1</sub><sup>x</sup> + n<sub>2</sub><sup>x</sup> + ... + n<sub>k</sub><sup>x</sup></code> 。</p>

<p>由于答案可能非常大，请你将它对 <code>10<sup>9</sup> + 7</code> 取余后返回。</p>

<p>比方说，<code>n = 160</code> 且 <code>x = 3</code> ，一个表示 <code>n</code> 的方法是 <code>n = 2<sup>3</sup> + 3<sup>3</sup> + 5<sup>3</sup></code><sup> </sup>。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>n = 10, x = 2
<b>输出：</b>1
<b>解释：</b>我们可以将 n 表示为：n = 3<sup>2</sup> + 1<sup>2</sup> = 10 。
这是唯一将 10 表达成不同整数 2 次方之和的方案。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>n = 4, x = 1
<b>输出：</b>2
<b>解释：</b>我们可以将 n 按以下方案表示：
- n = 4<sup>1</sup> = 4 。
- n = 3<sup>1</sup> + 1<sup>1</sup> = 4 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 300</code></li>
	<li><code>1 &lt;= x &lt;= 5</code></li>
</ul>
 
### 思路  

把 $n$ 看成背包容量，$n_i^x$ 看成物品，本题就是一个 0-1 背包模板题。  
代码实现时，由于 $n=300,x=1$ 算出来的结果不超过 $64$ 位整数的最大值，所以可以在计算结束后再取模。

```go 
func numberOfWays(n, x int) int {
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; pow(i, x) <= n; i++ {
		v := pow(i, x)
		for s := n; s >= v; s-- {
			f[s] += f[s-v]
		}
	}
	return f[n] % (1e9 + 7)
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x
		}
		x = x * x
	}
	return res
}
```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n^2\log x)$。
- 空间复杂度：$\mathcal{O}(n)$