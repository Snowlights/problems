#### 题目  

<p><code>n</code> 个人站成一排，按从 <code>1</code> 到 <code>n</code> 编号。</p>

<p>最初，排在队首的第一个人拿着一个枕头。每秒钟，拿着枕头的人会将枕头传递给队伍中的下一个人。一旦枕头到达队首或队尾，传递方向就会改变，队伍会继续沿相反方向传递枕头。</p>

<ul>
	<li>例如，当枕头到达第 <code>n</code> 个人时，TA 会将枕头传递给第 <code>n - 1</code> 个人，然后传递给第 <code>n - 2</code> 个人，依此类推。</li>
</ul>

<p>给你两个正整数 <code>n</code> 和 <code>time</code> ，返回 <code>time</code> 秒后拿着枕头的人的编号。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>n = 4, time = 5
<strong>输出：</strong>2
<strong>解释：</strong>队伍中枕头的传递情况为：1 -&gt; 2 -&gt; 3 -&gt; 4 -&gt; 3 -&gt; 2 。
5 秒后，枕头传递到第 2 个人手中。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>n = 3, time = 2
<strong>输出：</strong>3
<strong>解释：</strong>队伍中枕头的传递情况为：1 -&gt; 2 -&gt; 3 。
2 秒后，枕头传递到第 3 个人手中。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 1000</code></li>
	<li><code>1 &lt;= time &lt;= 1000</code></li>
</ul>
 
#### 思路  

到达端点需要 $n-1$ 时间。  
设 $t=\textit{time}\bmod (n-1)$，分类讨论：
- 如果 $\dfrac{\textit{time}}{n-1}$ 是偶数，说明在从 $1$ 到 $n$，答案为 $1+t$；
- 如果 $\dfrac{\textit{time}}{n-1}$ 是奇数，说明在从 $n$ 到 $1$，答案为 $n-t$。

```go 
func passThePillow(n, time int) int {
	t := time % (n - 1)
	if time/(n-1)%2 > 0 {
		return n - t
	}
	return 1 + t
}
```

#### 复杂度分析  

- 时间复杂度：$O(1)$。
- 空间复杂度：$O(1)$。仅用到若干变量。