#### 题目  

<p>给你一个下标从 <strong>0</strong> 开始、长度为 <code>n</code> 的二进制字符串 <code>s</code> ，你可以对其执行两种操作：</p>

<ul>
	<li>选中一个下标 <code>i</code> 并且反转从下标 <code>0</code> 到下标 <code>i</code>（包括下标 <code>0</code> 和下标 <code>i</code> ）的所有字符，成本为 <code>i + 1</code> 。</li>
	<li>选中一个下标 <code>i</code> 并且反转从下标 <code>i</code> 到下标 <code>n - 1</code>（包括下标 <code>i</code> 和下标 <code>n - 1</code> ）的所有字符，成本为 <code>n - i</code> 。</li>
</ul>

<p>返回使字符串内所有字符 <strong>相等</strong> 需要的 <strong>最小成本</strong> 。</p>

<p><strong>反转</strong> 字符意味着：如果原来的值是 &#39;0&#39; ，则反转后值变为 &#39;1&#39; ，反之亦然。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>s = &#34;0011&#34;
<strong>输出：</strong>2
<strong>解释：</strong>执行第二种操作，选中下标 <code>i = 2</code> ，可以得到 <code>s = &#34;0000&#34; ，成本为 2</code> 。可以证明 2 是使所有字符相等的最小成本。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>s = &#34;010101&#34;
<strong>输出：</strong>9
<strong>解释：</strong>执行第一种操作，选中下标 i = 2 ，可以得到 s = &#34;101101&#34; ，成本为 3 。
执行第一种操作，选中下标 i = 1 ，可以得到 s = &#34;011101&#34; ，成本为 2 。
执行第一种操作，选中下标 i = 0 ，可以得到 s = &#34;111101&#34; ，成本为 1 。
执行第二种操作，选中下标 i = 4 ，可以得到 s = &#34;111110&#34; ，成本为 2 。
执行第一种操作，选中下标 i = 5 ，可以得到 s = &#34;111111&#34; ，成本为 1 。
使所有字符相等的总成本等于 9 。可以证明 9 是使所有字符相等的最小成本。 </pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length == n &lt;= 10<sup>5</sup></code></li>
	<li><code>s[i]</code> 为 <code>&#39;0&#39;</code> 或 <code>&#39;1&#39;</code></li>
</ul>
 
#### 思路  

如果 $s[i-1]\ne s[i]$，那么必须反转，不然没法都相等：

- 要么翻转 $s[0]$ 到 $s[i-1]$，成本为 $i$；
- 要么翻转 $s[i]$ 到 $s[n-1]$，成本为 $n-i$。

这两种情况取最小值。

从左到右遍历 $s$，如果 $s[i-1]\ne s[i]$，那么必须反转，规则同上。

反转后：

- $s[i]$ 及其左边的字符，都已经相等了。
- $s[i]$ 右边的每对相邻字符，**反转前不同的，反转后仍然不同**。所以要继续反转。


```go 
func minimumCost(s string) int64 {
	ans, n := 0, len(s)
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			ans += min(i, n-i)
		}
	}
	return int64(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{s}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
