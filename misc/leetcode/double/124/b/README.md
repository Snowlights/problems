### 题目

<p>给你一个字符串 <code>s</code> 。</p>

<p>请你进行以下操作直到 <code>s</code> 为 <strong>空</strong> ：</p>

<ul>
	<li>每次操作 <strong>依次</strong> 遍历 <code>'a'</code> 到 <code>'z'</code>，如果当前字符出现在 <code>s</code> 中，那么删除出现位置 <strong>最早</strong> 的该字符。</li>
</ul>

<p>请你返回进行 <strong>最后</strong> 一次操作 <strong>之前</strong> 的字符串<em> </em><code>s</code><em> </em>。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>s = "aabcbbca"
<b>输出：</b>"ba"
<b>解释：</b>我们进行以下操作：
- 删除 s = "<em><strong>a</strong></em>a<em><strong>bc</strong></em>bbca" 中加粗加斜字符，得到字符串 s = "abbca" 。
- 删除 s = "<em><strong>ab</strong></em>b<em><strong>c</strong></em>a" 中加粗加斜字符，得到字符串 s = "ba" 。
- 删除 s = "<em><strong>ba</strong></em>" 中加粗加斜字符，得到字符串 s = "" 。
进行最后一次操作之前的字符串为 "ba" 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>s = "abcd"
<b>输出：</b>"abcd"
<b>解释：</b>我们进行以下操作：
- 删除 s = "<em><strong>abcd</strong></em>" 中加粗加斜字符，得到字符串 s = "" 。
进行最后一次操作之前的字符串为 "abcd" 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= s.length <= 5 * 10<sup>5</sup></code></li>
	<li><code>s</code> 只包含小写英文字母。</li>
</ul>

### 思路

## 提示 1

最后一次操作时，剩下的字母互不相同。因为如果有相同字母，那么操作后还有剩余字母。

## 提示 2

设字母的最大出现次数为 $\textit{mx}$。由于删除是从左到右进行的，最后剩下的就是出现次数等于 $\textit{mx}$ 的靠右字母（相同字母取出现位置最右的）。

```go [sol]
func lastNonEmptyString(s string) string {
	ans, cnt, mx := []byte{},make(map[byte]int), 0
	for i := range s {
		cnt[s[i]]++
		if cnt[s[i]] > mx  {
			mx = cnt[s[i]]
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if cnt[s[i]] == mx {
			ans = append(ans, s[i])
			cnt[s[i]] = 0
		}
	}
	for i := 0; i < len(ans) / 2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return string(ans)
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|\log |\Sigma|)$，其中 $n$ 为 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。
