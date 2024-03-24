### 题目  

<p>给你两个字符串 <code>s1</code> 和 <code>s2</code> ，两个字符串长度都为 <code>n</code> ，且只包含 <strong>小写 </strong>英文字母。</p>

<p>你可以对两个字符串中的 <strong>任意一个</strong> 执行以下操作 <strong>任意</strong> 次：</p>

<ul>
	<li>选择两个下标 <code>i</code> 和 <code>j</code> ，满足 <code>i &lt; j</code> 且 <code>j - i</code> 是 <strong>偶数</strong>，然后 <strong>交换</strong> 这个字符串中两个下标对应的字符。</li>
</ul>

<p> </p>

<p>如果你可以让字符串<em> </em><code>s1</code><em> </em>和<em> </em><code>s2</code> 相等，那么返回 <code>true</code> ，否则返回 <code>false</code> 。</p>

<p> </p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre><b>输入：</b>s1 = &#34;abcdba&#34;, s2 = &#34;cabdab&#34;
<b>输出：</b>true
<b>解释：</b>我们可以对 s1 执行以下操作：
- 选择下标 i = 0 ，j = 2 ，得到字符串 s1 = &#34;cbadba&#34; 。
- 选择下标 i = 2 ，j = 4 ，得到字符串 s1 = &#34;cbbdaa&#34; 。
- 选择下标 i = 1 ，j = 5 ，得到字符串 s1 = &#34;cabdab&#34; = s2 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre><b>输入：</b>s1 = &#34;abe&#34;, s2 = &#34;bea&#34;
<b>输出：</b>false
<b>解释：</b>无法让两个字符串相等。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == s1.length == s2.length</code></li>
	<li><code>1 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>s1</code> 和 <code>s2</code> 只包含小写英文字母。</li>
</ul>
 
### 思路 

判断奇数位和偶数位的字符总数是否一致

```go 
func canBeEqual(x string, y string) (ans bool) {
	cnt := [2][26]int{}
	for i, b := range x {
		cnt[i&1][b-'a']++
	}
	for i, b := range y {
		cnt[i&1][b-'a']--
	}
	for _, c := range cnt {
		for _, c := range c {
			if c != 0 {
				return
			}
		}
	}
	return true
}
```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$ 。