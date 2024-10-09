#### 题目

<p>给你一个字符串 <code>word</code> 和一个 <strong>非负 </strong>整数 <code>k</code>。</p>

<p>返回 <code>word</code> 的 <span data-keyword="substring-nonempty">子字符串</span> 中，每个元音字母（<code>'a'</code>、<code>'e'</code>、<code>'i'</code>、<code>'o'</code>、<code>'u'</code>）<strong>至少</strong> 出现一次，并且 <strong>恰好 </strong>包含 <code>k</code> 个辅音字母的子字符串的总数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">word = "aeioqq", k = 1</span></p>

<p><strong>输出：</strong><span class="example-io">0</span></p>

<p><strong>解释：</strong></p>

<p>不存在包含所有元音字母的子字符串。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">word = "aeiou", k = 0</span></p>

<p><strong>输出：</strong><span class="example-io">1</span></p>

<p><strong>解释：</strong></p>

<p>唯一一个包含所有元音字母且不含辅音字母的子字符串是 <code>word[0..4]</code>，即 <code>"aeiou"</code>。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">word = "ieaouqqieaouqq", k = 1</span></p>

<p><strong>输出：</strong>3</p>

<p><strong>解释：</strong></p>

<p>包含所有元音字母并且恰好含有一个辅音字母的子字符串有：</p>

<ul>
	<li><code>word[0..5]</code>，即 <code>"ieaouq"</code>。</li>
	<li><code>word[6..11]</code>，即 <code>"qieaou"</code>。</li>
	<li><code>word[7..12]</code>，即 <code>"ieaouq"</code>。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>5 &lt;= word.length &lt;= 250</code></li>
	<li><code>word</code> 仅由小写英文字母组成。</li>
	<li><code>0 &lt;= k &lt;= word.length - 5</code></li>
</ul>

#### 思路

问题等价于如下两个问题：

- 每个元音字母至少出现一次，并且**至少**包含 $k$ 个辅音字母的子串个数。记作 $f_k$。
- 每个元音字母至少出现一次，并且**至少**包含 $k+1$ 个辅音字母的子串个数。记作 $f_{k+1}$。

二者相减，所表达的含义就是**恰好**包含 $k$ 个辅音字母了，所以答案为 $f_k - f_{k+1}$。

对于每个问题，由于子串越长，越满足要求，有单调性，所以可以用**滑动窗口**解决。

如果你之前没有做过统计子串/子数组个数的滑动窗口，推荐先完成 [2962. 统计最大元素出现至少 K 次的子数组](https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/)（[我的题解](https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/solutions/2560940/hua-dong-chuang-kou-fu-ti-dan-pythonjava-xvwg/)），这也是一道至少+统计个数的问题，且比本题要简单许多。

## 答疑

**问**：能不能把 $f_k$ 定义成「至多」？

**答**：至多和前面的「每个元音字母**至少**出现一次」相克，「至少」要求子串越长越好，而「至多」要求子串越短越好，这样必须分开求解（总共要计算四个滑动窗口），相比下面代码的直接求解要麻烦许多。

**问**：代码中的 `ans += left` 是什么意思？

**答**：滑动窗口的内层循环结束时，右端点**固定**在 $\textit{right}$，左端点在 $0,1,2,\cdots,\textit{left}-1$ 的所有子串都是合法的，这一共有 $\textit{left}$ 个。

```
func f(word string, k int) (ans int64) {
	// 这里用哈希表实现，替换成数组会更快
	cnt1 := map[byte]int{} // 每种元音的个数
	cnt2 := 0 // 辅音个数
	left := 0
	for _, b := range word {
		if strings.ContainsRune("aeiou", b) {
			cnt1[byte(b)]++
		} else {
			cnt2++
		}
		for len(cnt1) == 5 && cnt2 >= k {
			out := word[left]
			if strings.ContainsRune("aeiou", rune(out)) {
				cnt1[out]--
				if cnt1[out] == 0 {
					delete(cnt1, out)
				}
			} else {
				cnt2--
			}
			left++
		}
		ans += int64(left)
	}
	return
}

func countOfSubstrings(word string, k int) int64 {
	return f(word, k) - f(word, k+1)
}
```

## 优化一

1. 把哈希表改成数组，额外用一个变量 $\textit{size}_1$ 维护元音种类数。
2. 判断元音的代码可以用位运算优化，把 $\texttt{aeiou}$ 视作集合 $\{0,4,8,14,20\}$，根据 [从集合论到位运算](https://leetcode.cn/circle/discuss/CaOJ45/)，这等于 $1065233$。用 `(1065233 >> b & 1) > 0` 可以判断字母 $b$ 是否为元音。

```
func f(word string, k int) (ans int64) {
	const vowelMask = 1065233
	cnt1 := ['u' - 'a' + 1]int{}
	size1 := 0 // 元音种类数
	cnt2 := 0
	left := 0
	for _, b := range word {
		b -= 'a'
		if vowelMask>>b&1 > 0 {
			if cnt1[b] == 0 {
				size1++
			}
			cnt1[b]++
		} else {
			cnt2++
		}
		for size1 == 5 && cnt2 >= k {
			out := word[left] - 'a'
			if vowelMask>>out&1 > 0 {
				cnt1[out]--
				if cnt1[out] == 0 {
					size1--
				}
			} else {
				cnt2--
			}
			left++
		}
		ans += int64(left)
	}
	return
}

func countOfSubstrings(word string, k int) int64 {
	return f(word, k) - f(word, k+1)
}
```

## 优化二

把两个滑动窗口合并成一个。我一般把这种滑窗叫做**三指针滑窗**。

```
func countOfSubstrings(word string, k int) (ans int64) {
	const vowelMask = 1065233
	var cntVowel1, cntVowel2 ['u' - 'a' + 1]int
	sizeVowel1, sizeVowel2 := 0, 0 // 元音种类数
	cntConsonant1, cntConsonant2 := 0, 0
	left1, left2 := 0, 0
	for _, b := range word {
		b -= 'a'
		if vowelMask>>b&1 > 0 {
			if cntVowel1[b] == 0 {
				sizeVowel1++
			}
			cntVowel1[b]++
			if cntVowel2[b] == 0 {
				sizeVowel2++
			}
			cntVowel2[b]++
		} else {
			cntConsonant1++
			cntConsonant2++
		}

		for sizeVowel1 == 5 && cntConsonant1 >= k {
			out := word[left1] - 'a'
			if vowelMask>>out&1 > 0 {
				cntVowel1[out]--
				if cntVowel1[out] == 0 {
					sizeVowel1--
				}
			} else {
				cntConsonant1--
			}
			left1++
		}

		for sizeVowel2 == 5 && cntConsonant2 > k {
			out := word[left2] - 'a'
			if vowelMask>>out&1 > 0 {
				cntVowel2[out]--
				if cntVowel2[out] == 0 {
					sizeVowel2--
				}
			} else {
				cntConsonant2--
			}
			left2++
		}

		ans += int64(left1 - left2)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{word}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$ 或者 $\mathcal{O}(|\Sigma|)$，其中 $|\Sigma|=21$。


## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
- [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
- [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)