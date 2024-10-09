### 题目

<p>给你一个字符串&nbsp;<code>s</code>&nbsp;。<code>s[i]</code>&nbsp;要么是小写英文字母，要么是问号&nbsp;<code>'?'</code>&nbsp;。</p>

<p>对于长度为 <code>m</code>&nbsp;且 <strong>只</strong>&nbsp;含有小写英文字母的字符串 <code>t</code>&nbsp;，我们定义函数&nbsp;<code>cost(i)</code>&nbsp;为下标 <code>i</code>&nbsp;之前（也就是范围 <code>[0, i - 1]</code>&nbsp;中）出现过与&nbsp;<code>t[i]</code>&nbsp;<strong>相同</strong>&nbsp;字符出现的次数。</p>

<p>字符串 <code>t</code>&nbsp;的&nbsp;<strong>分数</strong>&nbsp;为所有下标&nbsp;<code>i</code>&nbsp;的&nbsp;<code>cost(i)</code>&nbsp;之 <strong>和</strong>&nbsp;。</p>

<p>比方说，字符串&nbsp;<code>t = "aab"</code>&nbsp;：</p>

<ul>
	<li><code>cost(0) = 0</code></li>
	<li><code>cost(1) = 1</code></li>
	<li><code>cost(2) = 0</code></li>
	<li>所以，字符串&nbsp;<code>"aab"</code>&nbsp;的分数为&nbsp;<code>0 + 1 + 0 = 1</code>&nbsp;。</li>
</ul>

<p>你的任务是用小写英文字母&nbsp;<strong>替换</strong> <code>s</code>&nbsp;中 <strong>所有</strong> 问号，使 <code>s</code>&nbsp;的 <strong>分数</strong><strong>最小&nbsp;</strong>。</p>

<p>请你返回替换所有问号<em>&nbsp;</em><code>'?'</code>&nbsp;之后且分数最小的字符串。如果有多个字符串的&nbsp;<strong>分数最小</strong>&nbsp;，那么返回字典序最小的一个。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">s = "???" </span></p>

<p><strong>输出：</strong>&nbsp;<span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">"abc" </span></p>

<p><strong>解释：</strong>这个例子中，我们将 <code>s</code>&nbsp;中的问号&nbsp;<code>'?'</code>&nbsp;替换得到&nbsp;<code>"abc"</code>&nbsp;。</p>

<p>对于字符串&nbsp;<code>"abc"</code>&nbsp;，<code>cost(0) = 0</code>&nbsp;，<code>cost(1) = 0</code>&nbsp;和&nbsp;<code>cost(2) = 0</code>&nbsp;。</p>

<p><code>"abc"</code>&nbsp;的分数为&nbsp;<code>0</code>&nbsp;。</p>

<p>其他修改 <code>s</code>&nbsp;得到分数 <code>0</code>&nbsp;的字符串为&nbsp;<code>"cba"</code>&nbsp;，<code>"abz"</code>&nbsp;和&nbsp;<code>"hey"</code>&nbsp;。</p>

<p>这些字符串中，我们返回字典序最小的。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong> <span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">s = "a?a?"</span></p>

<p><strong>输出：</strong> <span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">"abac"</span></p>

<p><strong>解释：</strong>这个例子中，我们将&nbsp;<code>s</code>&nbsp;中的问号&nbsp;<code>'?'</code>&nbsp;替换得到&nbsp;<code>"abac"</code>&nbsp;。</p>

<p>对于字符串&nbsp;<code>"abac"</code>&nbsp;，<code>cost(0) = 0</code>&nbsp;，<code>cost(1) = 0</code>&nbsp;，<code>cost(2) = 1</code>&nbsp;和&nbsp;<code>cost(3) = 0</code>&nbsp;。</p>

<p><code>"abac"</code>&nbsp;的分数为&nbsp;<code>1</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>s[i]</code>&nbsp;要么是小写英文字母，要么是&nbsp;<code>'?'</code> 。</li>
</ul>

### 思路

## 分析

对于每个字母 $i$，统计其出现次数 $\textit{freq}[i]$。

对于 $\texttt{cost}$ 函数的结果 $\textit{res}$ 来说，字母 $i$ 对 $\texttt{res}$ 的贡献就是 $1+2+\cdots + (\textit{freq}[i]-1) = \dfrac{\textit{freq}[i](\textit{freq}[i]-1)}{2}$。

由于问号的出现次数是一个定值，由基本不等式可知，要使 $\textit{res}$ 尽量小，这 $26$ 种字母的出现次数应当尽量接近。

## 方法一：最小堆

1. 统计字母出现次数 $\textit{freq}$，和字母组成 pair 加到一个最小堆中。
2. 设问号出现了 $q$ 次。循环 $q$ 次，每次取出堆顶字母（它是目前出现次数最小的）加入一个列表 $t$ 中，然后把该字母的出现次数加一，重新入堆。
3. 把 $t$ 从小到大排序，因为题目要求字典序最小。
4. 遍历 $s$ 中的问号，按顺序填入 $t$ 中的字母。

```go [sol-Go]
func minimizeStringValue(s string) string {
	h := make(hp, 26)
	for i := byte(0); i < 26; i++ {
		h[i].c = 'a' + i
	}
	for _, b := range s {
		if b != '?' {
			h[b-'a'].f++
		}
	}
	heap.Init(&h)

	t := make([]byte, strings.Count(s, "?"))
	for i := range t {
		t[i] = h[0].c
		h[0].f++ // 出现次数加一
		heap.Fix(&h, 0)
	}
	slices.Sort(t) // 排序，因为要求字典序最小

	ans := []byte(s)
	j := 0
	for i, b := range ans {
		if b == '?' {
			ans[i] = t[j] // 填入字母
			j++
		}
	}
	return string(ans)
}

type pair struct {
	f int
	c byte
}

type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { a, b := h[i], h[j]; return a.f < b.f || a.f == b.f && a.c < b.c }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(interface{})             {}
func (hp) Pop() (_ interface{})         { return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + n\log |\Sigma|)$，其中 $n$ 为 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(n + |\Sigma|)$。

## 方法二：贪心

下面介绍更快的做法，时间复杂度可以视作线性。

举例说明。为方便描述，假设字母表只有 $\texttt{`a'}$ 到 $\texttt{`d'}$ 四种字母。

1. 统计每种字母的出现次数 $\textit{freq}$，假设为 $[6,8,4,2]$。
2. 复制一份 $f = \textit{freq}$，将 $f$ 从小到大排序，得 $f=[2,4,6,8]$。
3. 设问号的个数 $q=5$。按照方法一的做法，我们实际上要找到一个最小的数 $\textit{limit}$，把小于 $\textit{limit}$ 的出现次数改成 $\textit{limit}$。遍历 $f$ 数组可以计算出，当 $\textit{limit}=5$ 时，我们可以把 $f$ 中的 $2$ 变成 $5$，把 $4$ 变成 $5$，这一共会消耗 $4$ 个问号。多出的 $\textit{extra}=1$ 个问号，给到目前出现次数不超过 $\textit{limit}$ 的最小字母，即 $\texttt{`c'}$。
4. 按照上述做法，创建一个数组 $\textit{target}$，作为替换完所有问号后，每种字母的出现次数。
5. 遍历字符串 $s$，比较 $\textit{freq}$ 和 $\textit{target}$ 来替换字母。如果 $\textit{freq}[j] < \textit{target}[j]$，则可以把问号替换成第 $j$ 个字母，然后把 $\textit{freq}[j]$ 增加一。如果 $\textit{freq}[j] = \textit{target}[j]$ 则增大 $j$。

```go [sol-Go]
func minimizeStringValue(s string) string {
	freq := [27]int{26: math.MaxInt / 26} // 哨兵
	for _, c := range s {
		if c != '?' {
			freq[c-'a']++
		}
	}

	f := freq
	slices.Sort(f[:])

	var limit, extra int
	q := strings.Count(s, "?")
	for i := 1; ; i++ {
		sum := i * (f[i] - f[i-1])
		if q <= sum {
			limit, extra = f[i-1]+q/i, q%i
			break
		}
		q -= sum
	}

	target := freq
	for i, c := range freq[:26] {
		if c > limit {
			continue
		}
		target[i] = limit
		if extra > 0 { // 还可以多分配一个
			extra--
			target[i]++
		}
	}

	ans := []byte(s)
	j := byte(0)
	for i, c := range ans {
		if c != '?' {
			continue
		}
		for freq[j] == target[j] {
			j++
		}
		freq[j]++
		ans[i] = 'a' + j
	}
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|\log |\Sigma|)$，其中 $n$ 为 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(n + |\Sigma|)$，或者 $\mathcal{O}(|\Sigma|)$，如果可以像 C++ 那样原地修改的话。

## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)

