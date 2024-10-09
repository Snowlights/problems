### 题目  

<p>给你一个整数数组 <code>prices</code> ，它表示一个商店里若干巧克力的价格。同时给你一个整数 <code>money</code> ，表示你一开始拥有的钱数。</p>

<p>你必须购买 <strong>恰好 </strong>两块巧克力，而且剩余的钱数必须是 <strong>非负数</strong> 。同时你想最小化购买两块巧克力的总花费。</p>

<p>请你返回在购买两块巧克力后，最多能剩下多少钱。如果购买任意两块巧克力都超过了你拥有的钱，请你返回 <code>money</code> 。注意剩余钱数必须是非负数。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>prices = [1,2,2], money = 3
<b>输出：</b>0
<b>解释：</b>分别购买价格为 1 和 2 的巧克力。你剩下 3 - 3 = 0 块钱。所以我们返回 0 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>prices = [3,2,3], money = 3
<b>输出：</b>3
<b>解释：</b>购买任意 2 块巧克力都会超过你拥有的钱数，所以我们返回 3 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= prices.length &lt;= 50</code></li>
	<li><code>1 &lt;= prices[i] &lt;= 100</code></li>
	<li><code>1 &lt;= money &lt;= 100</code></li>
</ul>
 
### 思路

可以通过一次遍历维护最小的两个值，最后进行判断是否大于money

```go 
func buyChoco(a []int, money int) (ans int) {
	sort.Ints(a)
	if a[0]+a[1] <= money {
		return money - a[0] - a[1]
	}
	return money
}
```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{prices}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。