#### 题目  

<p>有两只老鼠和 <code>n</code> 块不同类型的奶酪，每块奶酪都只能被其中一只老鼠吃掉。</p>

<p>下标为 <code>i</code> 处的奶酪被吃掉的得分为：</p>

<ul>
	<li>如果第一只老鼠吃掉，则得分为 <code>reward1[i]</code> 。</li>
	<li>如果第二只老鼠吃掉，则得分为 <code>reward2[i]</code> 。</li>
</ul>

<p>给你一个正整数数组 <code>reward1</code> ，一个正整数数组 <code>reward2</code> ，和一个非负整数 <code>k</code> 。</p>

<p>请你返回第一只老鼠恰好吃掉 <code>k</code> 块奶酪的情况下，<strong>最大</strong> 得分为多少。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>reward1 = [1,1,3,4], reward2 = [4,4,1,1], k = 2
<b>输出：</b>15
<b>解释：</b>这个例子中，第一只老鼠吃掉第 2 和 3 块奶酪（下标从 0 开始），第二只老鼠吃掉第 0 和 1 块奶酪。
总得分为 4 + 4 + 3 + 4 = 15 。
15 是最高得分。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>reward1 = [1,1], reward2 = [1,1], k = 2
<b>输出：</b>2
<b>解释：</b>这个例子中，第一只老鼠吃掉第 0 和 1 块奶酪（下标从 0 开始），第二只老鼠不吃任何奶酪。
总得分为 1 + 1 = 2 。
2 是最高得分。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == reward1.length == reward2.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= reward1[i], reward2[i] &lt;= 1000</code></li>
	<li><code>0 &lt;= k &lt;= n</code></li>
</ul>
 
#### 思路 

## 思考

如果 $k=1$，应该如何选择呢？（思考问题可以先从一些简单的情况开始）  
不妨先把奶酪全部给第二只老鼠，然后「撤销」其中的一块奶酪，给第一只老鼠。如何选择可以使得分最大？  
你可以把这个结论推广到 $k>1$ 的情况吗？

## 解惑

为方便描述，将 $\textit{reward}$ 简记为 $r$。  
先把奶酪全部给第二只老鼠，然后撤销其中的第 $i$ 块奶酪，给第一只老鼠，那么得分增加了

$$
r_1[i] - r_2[i]
$$

在 $k=1$ 时，选上式最大的奶酪，给第一只老鼠，这样可以使得分最大。（注意第一只老鼠一定要吃**恰好** $k$ 块奶酪）  
对于 $k>1$ 的情况，可以按照 $r_1[i] - r_2[i]$ 从大到小排序，把得分加上 $r_1[i] - r_2[i]$ 的前 $k$ 大之和。这可以用快速选择优化到 $\mathcal{O}(n)$，

```go 
func miceAndCheese(reward1, reward2 []int, k int) (ans int) {
	for i, x := range reward2 {
		ans += x // 全部给第二只老鼠
		reward1[i] -= x
	}
	sort.Sort(sort.Reverse(sort.IntSlice(reward1)))
	for _, x := range reward1[:k] {
		ans += x
	}
	return
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 为 $\textit{reward}_1$ 的长度。快速选择可以做到 $\mathcal{O}(n)$
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。