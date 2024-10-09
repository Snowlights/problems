#### 题目

<p>给你一个下标从 <strong>0</strong>&nbsp;开始的字符串&nbsp;<code>s</code>&nbsp;、字符串&nbsp;<code>a</code>&nbsp;、字符串&nbsp;<code>b</code>&nbsp;和一个整数&nbsp;<code>k</code>&nbsp;。</p>

<p>如果下标 <code>i</code>&nbsp;满足以下条件，则认为它是一个 <strong>美丽下标</strong>&nbsp;：</p>

<ul>
	<li><code>0 &lt;= i &lt;= s.length - a.length</code></li>
	<li><code>s[i..(i + a.length - 1)] == a</code></li>
	<li>存在下标&nbsp;<code>j</code>&nbsp;使得：
	<ul>
		<li><code>0 &lt;= j &lt;= s.length - b.length</code></li>
		<li><code>s[j..(j + b.length - 1)] == b</code></li>
		<li><code>|j - i| &lt;= k</code></li>
	</ul>
	</li>
</ul>

<p>以数组形式按<strong>&nbsp;从小到大排序&nbsp;</strong>返回美丽下标。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "isawsquirrelnearmysquirrelhouseohmy", a = "my", b = "squirrel", k = 15
<strong>输出：</strong>[16,33]
<strong>解释：</strong>存在 2 个美丽下标：[16,33]。
- 下标 16 是美丽下标，因为 s[16..17] == "my" ，且存在下标 4 ，满足 s[4..11] == "squirrel" 且 |16 - 4| &lt;= 15 。
- 下标 33 是美丽下标，因为 s[33..34] == "my" ，且存在下标 18 ，满足 s[18..25] == "squirrel" 且 |33 - 18| &lt;= 15 。
因此返回 [16,33] 作为结果。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>s = "abcd", a = "a", b = "a", k = 4
<b>输出：</b>[0]
<strong>解释：</strong>存在 1 个美丽下标：[0]。
- 下标 0 是美丽下标，因为 s[0..0] == "a" ，且存在下标 0 ，满足 s[0..0] == "a" 且 |0 - 0| &lt;= 4 。
因此返回 [0] 作为结果。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= k &lt;= s.length &lt;= 5 * 10<sup>5</sup></code></li>
	<li><code>1 &lt;= a.length, b.length &lt;= 5 * 10<sup>5</sup></code></li>
	<li><code>s</code>、<code>a</code>、和&nbsp;<code>b</code>&nbsp;只包含小写英文字母。</li>
</ul>


#### 思路

#### KMP+二分查找

1.用 KMP 求出 $a$ 在 $s$ 中的所有出现位置，记作 $\textit{posA}$。
2. 用 KMP 求出 $b$ 在 $s$ 中的所有出现位置，记作 $\textit{posB}$。
3. 遍历 $\textit{posA}$ 中的下标 $i$，在 $\textit{posB}$ 中二分查找离 $i$ 最近的 $j$。如果 $|i-j|\le k$，则把 $i$ 加入答案。

```go [sol]
func beautifulIndices(s, a, b string, k int) []int {
	calcMaxMatchLengths := func(s string) []int {
		match := make([]int, len(s))
		for i, c := 1, 0; i < len(s); i++ {
			v := s[i]
			for c > 0 && s[c] != v {
				c = match[c-1]
			}
			if s[c] == v {
				c++
			}
			match[i] = c
		}
		return match
	}
	kmpSearch := func(text, pattern string) (pos []int) {
		match := calcMaxMatchLengths(pattern)
		lenP := len(pattern)
		c := 0
		for i, v := range text {
			for c > 0 && pattern[c] != byte(v) {
				c = match[c-1]
			}
			if pattern[c] == byte(v) {
				c++
			}
			if c == lenP {
				pos = append(pos, i-lenP+1)
				c = match[c-1] // 不允许重叠时 c = 0
			}
		}
		return
	}

	aPos, bPos := kmpSearch(s, a), kmpSearch(s, b)
	ans := make([]int, 0)

	for _, v := range aPos {
		idx := sort.SearchInts(bPos, v)
		if idx < len(bPos) && abs(bPos[idx] - v) <= k ||
			idx > 0 && abs(bPos[idx-1] - v) <= k{
			ans = append(ans, v)
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n \log n)$，$n$ 为 $\textit{s}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

#### KMP+双指针

由于 $\\textit{posA}$ 和 $\\textit{posB}$ 都是有序的，也可以用双指针做。

```go [sol]
func beautifulIndices(s, a, b string, k int) []int {
	calcMaxMatchLengths := func(s string) []int {
		match := make([]int, len(s))
		for i, c := 1, 0; i < len(s); i++ {
			v := s[i]
			for c > 0 && s[c] != v {
				c = match[c-1]
			}
			if s[c] == v {
				c++
			}
			match[i] = c
		}
		return match
	}
	kmpSearch := func(text, pattern string) (pos []int) {
		match := calcMaxMatchLengths(pattern)
		lenP := len(pattern)
		c := 0
		for i, v := range text {
			for c > 0 && pattern[c] != byte(v) {
				c = match[c-1]
			}
			if pattern[c] == byte(v) {
				c++
			}
			if c == lenP {
				pos = append(pos, i-lenP+1)
				c = match[c-1] // 不允许重叠时 c = 0
			}
		}
		return
	}

	aPos, bPos := kmpSearch(s, a), kmpSearch(s, b)
	ans := make([]int, 0)

	j, m := 0, len(bPos)
	for _, i := range aPos {
		for j < m && bPos[j] < i-k {
			j++
		}
		if j < m && abs(bPos[j] - i) <= k {
			ans = append(ans, i)
		}
	}
	
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，$n$ 为 $\textit{s}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
