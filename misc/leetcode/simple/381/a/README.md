#### 题目

<p>给你一个字符串 <code>word</code>，由 <strong>不同 </strong>小写英文字母组成。</p>

<p>电话键盘上的按键与 <strong>不同 </strong>小写英文字母集合相映射，可以通过按压按键来组成单词。例如，按键 <code>2</code> 对应 <code>["a","b","c"]</code>，我们需要按一次键来输入 <code>"a"</code>，按两次键来输入 <code>"b"</code>，按三次键来输入 <code>"c"</code><em>。</em></p>

<p>现在允许你将编号为 <code>2</code> 到 <code>9</code> 的按键重新映射到 <strong>不同 </strong>字母集合。每个按键可以映射到<strong> 任意数量 </strong>的字母，但每个字母 <strong>必须</strong> <strong>恰好</strong> 映射到 <strong>一个 </strong>按键上。你需要找到输入字符串 <code>word</code> 所需的<strong> 最少 </strong>按键次数。</p>

<p>返回重新映射按键后输入 <code>word</code> 所需的 <strong>最少 </strong>按键次数。</p>

<p>下面给出了一种电话键盘上字母到按键的映射作为示例。注意 <code>1</code>，<code>*</code>，<code>#</code> 和 <code>0</code> <strong>不</strong> 对应任何字母。</p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/12/26/keypaddesc.png" style="width: 329px; height: 313px;" />
<p> </p>

<p><strong class="example">示例 1：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/12/26/keypadv1e1.png" style="width: 329px; height: 313px;" />
<pre>
<strong>输入：</strong>word = "abcde"
<strong>输出：</strong>5
<strong>解释：</strong>图片中给出的重新映射方案的输入成本最小。
"a" -> 在按键 2 上按一次
"b" -> 在按键 3 上按一次
"c" -> 在按键 4 上按一次
"d" -> 在按键 5 上按一次
"e" -> 在按键 6 上按一次
总成本为 1 + 1 + 1 + 1 + 1 = 5 。
可以证明不存在其他成本更低的映射方案。
</pre>

<p><strong class="example">示例 2：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/12/26/keypadv1e2.png" style="width: 329px; height: 313px;" />
<pre>
<strong>输入：</strong>word = "xycdefghij"
<strong>输出：</strong>12
<strong>解释：</strong>图片中给出的重新映射方案的输入成本最小。
"x" -> 在按键 2 上按一次
"y" -> 在按键 2 上按两次
"c" -> 在按键 3 上按一次
"d" -> 在按键 3 上按两次
"e" -> 在按键 4 上按一次
"f" -> 在按键 5 上按一次
"g" -> 在按键 6 上按一次
"h" -> 在按键 7 上按一次
"i" -> 在按键 8 上按一次
"j" -> 在按键 9 上按一次
总成本为 1 + 2 + 1 + 2 + 1 + 1 + 1 + 1 + 1 + 1 = 12 。
可以证明不存在其他成本更低的映射方案。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= word.length <= 26</code></li>
	<li><code>word</code> 仅由小写英文字母组成。</li>
	<li><code>word</code> 中的所有字母互不相同。</li>
</ul>

#### 思路

由于各个字母互不相同，所以均匀分配到这 $8$ 个按键。
设字符串长度为 $n$，$k=\left\lfloor\dfrac{n}{8}\right\rfloor$，
那么先分配给每个按键 $k$ 个字母，总按键次数为

$$
8\cdot(1+2+\cdots + k) = 4k(k+1)

$$

剩余的 $n\bmod 8$ 个字母需要按 $k+1$ 次。 所以答案为

$$
4k(k+1) + (n\bmod 8)(k+1) = (4k + n\bmod 8)(k+1)

$$

```go [sol]
func minimumPushes(word string) int {
	n := len(word)
	if n <= 8 {
		return n
	}
	if n <= 16 {
		return 8 + (n-8) * 2
	}
	if n <= 24 {
		return 8 + 16 + (n-16) * 3
	}
	return 8 + 16 + 24 + (n-24) * 4
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。
