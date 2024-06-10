#### 题目

<p>给你一个数组 <code>arr</code> ，数组中有 <code>n</code> 个 <b>非空</b> 字符串。</p>

<p>请你求出一个长度为 <code>n</code> 的字符串 <code>answer</code> ，满足：</p>

<ul>
	<li><code>answer[i]</code> 是 <code>arr[i]</code> <strong>最短</strong> 的子字符串，且它不是 <code>arr</code> 中其他任何字符串的子字符串。如果有多个这样的子字符串存在，<code>answer[i]</code> 应该是它们中字典序最小的一个。如果不存在这样的子字符串，<code>answer[i]</code> 为空字符串。</li>
</ul>

<p>请你返回数组<em> </em><code>answer</code> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>arr = ["cab","ad","bad","c"]
<b>输出：</b>["ab","","ba",""]
<b>解释：</b>求解过程如下：
- 对于字符串 "cab" ，最短没有在其他字符串中出现过的子字符串是 "ca" 或者 "ab" ，我们选择字典序更小的子字符串，也就是 "ab" 。
- 对于字符串 "ad" ，不存在没有在其他字符串中出现过的子字符串。
- 对于字符串 "bad" ，最短没有在其他字符串中出现过的子字符串是 "ba" 。
- 对于字符串 "c" ，不存在没有在其他字符串中出现过的子字符串。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>arr = ["abc","bcd","abcd"]
<b>输出：</b>["","","abcd"]
<b>解释：</b>求解过程如下：
- 对于字符串 "abc" ，不存在没有在其他字符串中出现过的子字符串。
- 对于字符串 "bcd" ，不存在没有在其他字符串中出现过的子字符串。
- 对于字符串 "abcd" ，最短没有在其他字符串中出现过的子字符串是 "abcd" 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == arr.length</code></li>
	<li><code>2 <= n <= 100</code></li>
	<li><code>1 <= arr[i].length <= 20</code></li>
	<li><code>arr[i]</code> 只包含小写英文字母。</li>
</ul>

#### 思路

对于 $\textit{arr}[i]$，从小到大枚举长度 $\textit{size}$，
然后枚举 $\textit{arr}[i]$ 的长为 $\textit{size}$ 的所有子串 $t$，
判断 $t$ 是否在其它字符串中，如果不在，就更新 $\textit{answer}[i]$ 的最小值。

```go [sol]
func shortestSubstrings(arr []string) []string {
	ans := make([]string, len(arr))
	for i, s := range arr {
		m := len(s)
		res := ""
		for size := 1; size <= m && res == ""; size++ {
		next:
			for k := size; k <= m; k++ {
				sub := s[k-size : k]
				if res != "" && sub >= res {
					continue
				}
				for j, t := range arr {
					if j != i && strings.Contains(t, sub) {
						continue next
					}
				}
				res = sub
			}
		}
		ans[i] = res
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2m^4)$，其中 $n$ 为 $\textit{arr}$ 的长度，$m$ 为 $\textit{arr}[i]$ 的长度，不超过 $20$。
- 空间复杂度：$\mathcal{O}(m)$ 或 $\mathcal{O}(1)$。忽略返回值的空间。
