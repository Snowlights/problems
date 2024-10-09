#### 题目  

给你三个字符串 <code>a</code> ，<code>b</code> 和 <code>c</code> ， 你的任务是找到长度 <strong>最短</strong> 的字符串，且这三个字符串都是它的 <strong>子字符串</strong> 。
<p>如果有多个这样的字符串，请你返回 <strong>字典序最小</strong> 的一个。</p>

<p>请你返回满足题目要求的字符串。</p>

<p><strong>注意：</strong></p>

<ul>
	<li>两个长度相同的字符串 <code>a</code> 和 <code>b</code> ，如果在第一个不相同的字符处，<code>a</code> 的字母在字母表中比 <code>b</code> 的字母 <strong>靠前</strong> ，那么字符串 <code>a</code> 比字符串 <code>b</code> <strong>字典序小</strong> 。</li>
	<li><strong>子字符串</strong> 是一个字符串中一段连续的字符序列。</li>
</ul>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><code><span style=""><b>输入：</b></span>a</code> = &#34;abc&#34;, <code>b</code> = &#34;bca&#34;, <code>c</code> = &#34;aaa&#34;
<b>输出：</b>&#34;aaabca&#34;
<b>解释：</b>字符串 &#34;aaabca&#34; 包含所有三个字符串：a = ans[2...4] ，b = ans[3..5] ，c = ans[0..2] 。结果字符串的长度至少为 6 ，且&#34;aaabca&#34; 是字典序最小的一个。</pre>

<p><strong>示例 2：</strong></p>

<pre><code><span style=""><b>输入：</b></span>a</code> = &#34;ab&#34;, <code>b</code> = &#34;ba&#34;, <code>c</code> = &#34;aba&#34;
<b>输出：</b>&#34;aba&#34;
<strong>解释：</strong>字符串 &#34;aba&#34; 包含所有三个字符串：a = ans[0..1] ，b = ans[1..2] ，c = ans[0..2] 。由于 c 的长度为 3 ，结果字符串的长度至少为 3 。&#34;aba&#34; 是字典序最小的一个。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= a.length, b.length, c.length &lt;= 100</code></li>
	<li><code>a</code> ，<code>b</code> ，<code>c</code> 只包含小写英文字母。</li>
</ul>
 
#### 思路  

暴力枚举 $a,b,c$ 的全排列 $a',b',c'$，然后合并 $a',b'$ 得到最短字符串 $x$，再合并 $x,c'$ 得到最短字符串 $s$。
可能有的同学觉得，$a',b'$ 合并一个较长的会不会更优？  
你可以这样理解：在没有完全包含的前提下，我们相当于在 $b'$ 的左边添加一些字母，右边添加一些字母，得到一个最短字符串 $s$。  
可以先特判一下完全包含的情况，对于没有完全包含的情况，必然是合并得越短越好。  
取最短且字典序最小的 $s$ 作为答案。

```go 
func merge(s, t string) string {
	// 先特判完全包含的情况
	if strings.Contains(s, t) {
		return s
	}
	if strings.Contains(t, s) {
		return t
	}
	for i := min(len(s), len(t)); ; i-- {
		// 枚举：s 的后 i 个字母和 t 的前 i 个字母是一样的
		if s[len(s)-i:] == t[:i] {
			return s + t[i:]
		}
	}
}

func minimumString(a, b, c string) (ans string) {
	arr := []string{a, b, c}
	// 枚举 arr 的全排列
	for _, p := range [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}} {
		s := merge(merge(arr[p[0]], arr[p[1]]), arr[p[2]])
		if ans == "" || len(s) < len(ans) || len(s) == len(ans) && s < ans {
			ans = s
		}
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $a,b,c$ 的长度的最大值。
- 空间复杂度：$\mathcal{O}(n)$