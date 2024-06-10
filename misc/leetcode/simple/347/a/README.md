#### 题目  

<p>给你一个用字符串表示的正整数 <code>num</code> ，请你以字符串形式返回不含尾随零的整数<em> </em><code>num</code><em> </em>。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>num = &#34;51230100&#34;
<strong>输出：</strong>&#34;512301&#34;
<strong>解释：</strong>整数 &#34;51230100&#34; 有 2 个尾随零，移除并返回整数 &#34;512301&#34; 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>num = &#34;123&#34;
<strong>输出：</strong>&#34;123&#34;
<strong>解释：</strong>整数 &#34;123&#34; 不含尾随零，返回整数 &#34;123&#34; 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= num.length &lt;= 1000</code></li>
	<li><code>num</code> 仅由数字 <code>0</code> 到 <code>9</code> 组成</li>
	<li><code>num</code> 不含前导零</li>
</ul>
 
#### 思路  

直接去除字符串末尾的 $0$ 即可

```go 
func removeTrailingZeros(n string) (ans string) {
	return strings.TrimRight(n, "0")
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。