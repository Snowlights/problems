#### 题目

<p>给你一个仅由小写英文字母组成的字符串 <code>s</code> 。</p>

<p>如果一个字符串仅由单一字符组成，那么它被称为 <strong>特殊 </strong>字符串。例如，字符串 <code>"abc"</code> 不是特殊字符串，而字符串 <code>"ddd"</code>、<code>"zz"</code> 和 <code>"f"</code> 是特殊字符串。</p>

<p>返回在 <code>s</code> 中出现 <strong>至少三次 </strong>的<strong> 最长特殊子字符串 </strong>的长度，如果不存在出现至少三次的特殊子字符串，则返回 <code>-1</code> 。</p>

<p><strong>子字符串 </strong>是字符串中的一个连续<strong> 非空 </strong>字符序列。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "aaaa"
<strong>输出：</strong>2
<strong>解释：</strong>出现三次的最长特殊子字符串是 "aa" ：子字符串 "<em><strong>aa</strong></em>aa"、"a<em><strong>aa</strong></em>a" 和 "aa<em><strong>aa</strong></em>"。
可以证明最大长度是 2 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>s = "abcdef"
<strong>输出：</strong>-1
<strong>解释：</strong>不存在出现至少三次的特殊子字符串。因此返回 -1 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>s = "abcaba"
<strong>输出：</strong>1
<strong>解释：</strong>出现三次的最长特殊子字符串是 "a" ：子字符串 "<em><strong>a</strong></em>bcaba"、"abc<em><strong>a</strong></em>ba" 和 "abcab<em><strong>a</strong></em>"。
可以证明最大长度是 1 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 <= s.length <= 50</code></li>
	<li><code>s</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

按照相同字母分组，每组统计相同字母连续出现的长度。例如 aaaabbbabb 把 a 分成一组，
组内有长度 $4$ 和长度 $1$；把 b 分成一组，组内有长度 $3$ 和长度 $2$。
单独考虑每一组，按照长度从大到小排序，设长度列表为 $a$。

分类讨论：
- 从最长的特殊子串（$a[0]$）中取三个长度均为 $a[0]-2$ 的特殊子串。例如示例 1 的 aaaa 可以取三个 aa。
- 从最长和次长的特殊子串（$a[0],a[1]$）中取三个长度一样的特殊子串： 
  - 如果 $a[0]=a[1]$，那么可以取三个长度均为 $a[0]-1$ 的特殊子串。
  - 如果 $a[0]>a[1]$，那么可以取三个长度均为 $a[1]$ 的特殊子串：从最长中取两个，从次长中取一个。
  - 这两种情况合并成 $\min(a[0]-1, a[1])$。
-从最长、次长、第三长的的特殊子串（$a[0],a[1],a[2]$）中各取一个长为 $a[2]$ 的特殊子串。

这三种情况取最大值，即

$$
\max(a[0]-2, \min(a[0]-1, a[1]), a[2])
$$

取每一组的最大值，即为答案。
如果答案是 $0$，返回 $-1$。\n\n代码实现时，无需特判 $a$ 数组长度小于 $3$ 的情况，我们只需要在数组末尾加两个 $0$ 即可。

```go  [sol]
func maximumLength(s string) int {
	h, ans := make(map[int][]int), 0
	i, n, cnt := 0, len(s), 0
	for ; i < n; i++ {
		cnt++
		if i == n-1 || s[i] != s[i+1] {
			h[int(s[i]-'a')] = append(h[int(s[i]-'a')], cnt)
			cnt = 0
		}
	}

	for _, v := range h {
		v = append(v, 0, 0)
		sort.Ints(v)
		n = len(v)
		ans = max(ans, v[n-1]-2, min(v[n-1]-1, v[n-2]), v[n-3])
	}

	if ans == 0 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $s$ 的长度。如果改用堆维护前三大可以做到 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。
