#### 题目

<p>给你一个字符串 <code>word</code>。如果 <code>word</code> 中同时出现某个字母 <code>c</code> 的小写形式和大写形式，并且<strong> 每个 </strong>小写形式的 <code>c</code> 都出现在第一个大写形式的 <code>c</code> 之前，则称字母 <code>c</code> 是一个 <strong>特殊字母</strong> 。</p>

<p>返回 <code>word</code> 中 <strong>特殊字母</strong> 的数量。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1:</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">word = "aaAbcBC"</span></p>

<p><strong>输出：</strong><span class="example-io">3</span></p>

<p><strong>解释：</strong></p>

<p>特殊字母是 <code>'a'</code>、<code>'b'</code> 和 <code>'c'</code>。</p>
</div>

<p><strong class="example">示例 2:</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">word = "abc"</span></p>

<p><strong>输出：</strong><span class="example-io">0</span></p>

<p><strong>解释：</strong></p>

<p><code>word</code> 中不存在特殊字母。</p>
</div>

<p><strong class="example">示例 3:</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">word = "AbBCab"</span></p>

<p><strong>输出：</strong><span class="example-io">0</span></p>

<p><strong>解释：</strong></p>

<p><code>word</code> 中不存在特殊字母。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= word.length &lt;= 2 * 10<sup>5</sup></code></li>
	<li><code>word</code> 仅由小写和大写英文字母组成。</li>
</ul>

#### 思路

状态机

对于每种字母，定义如下四种状态：
- $0$：初始状态。如果遇到小写字母，转到状态 $1$，否则不合法，转到状态 $-1$。
- $1$：如果遇到小写字母则不变，遇到大写字母则转到状态 $2$。
- $2$：如果遇到大写字母则不变，遇到小写字母则转到状态 $-1$。
- $-1$：遇到任何字母都不再变化。

![394C.png](https://pic.leetcode.cn/1713671840-HgbYWt-394C.png)

答案为状态为 $2$ 的字母种数。

附 ASCII 的性质：
- 对于大写英文字母：其二进制从右往左第 $6$ 个比特值一定是 $0$。
- 对于小写英文字母：其二进制从右往左第 $6$ 个比特值一定是 $1$。
- 对于任何英文字母：其小写字母二进制低 $5$ 位，一定和其大写字母二进制低 $5$ 位相等。

``` go 
func numberOfSpecialChars(word string) int {
	ans, cnt := 0, [26]int{}
	for _, v := range word {
		if 'a' <= v && v <= 'z' {
			vi := int(v - 'a')
			switch cnt[vi] {
			case 0:
				cnt[vi] = 1
			case 2:
				cnt[vi] = 3
			}
		} else {
			vi := int(v - 'A')
			switch cnt[vi] {
			case 0:
				cnt[vi] = 3
			case 1:
				cnt[vi] = 2
			}
		}
	}
	for _, v := range cnt {
		if v == 2 {
			ans++
		}
	}
	return ans
}

```


#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 为 $\textit{word}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为英文字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)