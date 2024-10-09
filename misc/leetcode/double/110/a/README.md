### 题目  

<p>一开始，你的银行账户里有 <code>100</code> 块钱。</p>

<p>给你一个整数<code>purchaseAmount</code> ，它表示你在一次购买中愿意支出的金额。</p>

<p>在一个商店里，你进行一次购买，实际支出的金额会向 <strong>最近</strong> 的 <code>10</code> 的 <strong>倍数</strong> 取整。换句话说，你实际会支付一个 <strong>非负</strong> 金额 <code>roundedAmount</code> ，满足 <code>roundedAmount</code> 是 <code>10</code> 的倍数且 <code>abs(roundedAmount - purchaseAmount)</code> 的值 <strong>最小</strong> 。</p>

<p>如果存在多于一个最接近的 <code>10</code> 的倍数，<strong>较大的倍数</strong> 是你的实际支出金额。</p>

<p>请你返回一个整数，表示你在愿意支出金额为<em> </em><code>purchaseAmount</code><em> </em>块钱的前提下，购买之后剩下的余额。</p>

<p><strong>注意：</strong> <code>0</code> 也是 <code>10</code> 的倍数。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>purchaseAmount = 9
<b>输出：</b>90
<b>解释：</b>这个例子中，最接近 9 的 10 的倍数是 10 。所以你的账户余额为 100 - 10 = 90 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>purchaseAmount = 15
<b>输出：</b>80
<b>解释：</b>这个例子中，有 2 个最接近 15 的 10 的倍数：10 和 20，较大的数 20 是你的实际开销。
所以你的账户余额为 100 - 20 = 80 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 &lt;= purchaseAmount &lt;= 100</code></li>
</ul>
 
### 思路  

题目意思是把 $\textit{purchaseAmount}$ 的个位数「四舍五入」，这可以用如下公式

$$
\left\lceil\dfrac{\textit{purchaseAmount}+5}{10}\right\rceil\cdot 10
$$

```go 
func accountBalanceAfterPurchase(purchaseAmount int) (ans int) {
	return 100 - (purchaseAmount+5)/10*10
}
```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$ 。