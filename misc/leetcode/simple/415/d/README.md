#### 题目

<p>给你一个字符串数组 <code>words</code> 和一个字符串 <code>target</code>。</p>

<p>如果字符串 <code>x</code> 是 <code>words</code> 中<strong> 任意 </strong>字符串的 <span data-keyword="string-prefix">前缀</span>，则认为 <code>x</code> 是一个 <strong>有效</strong> 字符串。</p>

<p>现计划通过 <strong>连接 </strong>有效字符串形成 <code>target</code> ，请你计算并返回需要连接的 <strong>最少 </strong>字符串数量。如果无法通过这种方式形成 <code>target</code>，则返回 <code>-1</code>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">words = ["abc","aaaaa","bcdef"], target = "aabcdabc"</span></p>

<p><strong>输出：</strong> <span class="example-io">3</span></p>

<p><strong>解释：</strong></p>

<p>target 字符串可以通过连接以下有效字符串形成：</p>

<ul>
	<li><code>words[1]</code> 的长度为 2 的前缀，即 <code>"aa"</code>。</li>
	<li><code>words[2]</code> 的长度为 3 的前缀，即 <code>"bcd"</code>。</li>
	<li><code>words[0]</code> 的长度为 3 的前缀，即 <code>"abc"</code>。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">words = ["abababab","ab"], target = "ababaababa"</span></p>

<p><strong>输出：</strong> <span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<p>target 字符串可以通过连接以下有效字符串形成：</p>

<ul>
	<li><code>words[0]</code> 的长度为 5 的前缀，即 <code>"ababa"</code>。</li>
	<li><code>words[0]</code> 的长度为 5 的前缀，即 <code>"ababa"</code>。</li>
</ul>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">words = ["abcdef"], target = "xyz"</span></p>

<p><strong>输出：</strong> <span class="example-io">-1</span></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= words.length &lt;= 100</code></li>
	<li><code>1 &lt;= words[i].length &lt;= 5 * 10<sup>4</sup></code></li>
	<li>输入确保 <code>sum(words[i].length) &lt;= 10<sup>5</sup></code>.</li>
	<li><code>words[i]</code> &nbsp;只包含小写英文字母。</li>
	<li><code>1 &lt;= target.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>target</code> &nbsp;只包含小写英文字母。</li>
</ul>

#### 思路


把 $\textit{target}$ 划分成若干段，使得每一段都是某个 $\textit{words}[i]$ 的前缀。

返回**最小划分个数**。如果无法划分，返回 $-1$。

## 方法一：跳跃游戏 + 字符串哈希 + 二分

用划分型 DP 思考，考虑第一段划分多长。

比如示例 1 的 $\textit{target}=\texttt{aabcdabc}$。

如果第一段能划分出一个长为 $2$ 的子串，即 $\texttt{aa}$，那么也可以划分出一个更短的，长为 $1$ 的子串，即 $\texttt{a}$。

所以只需考虑**最大**的 $r$，满足 $\textit{target}$ 下标从 $0$ 到 $r$ 这一段是某个 $\textit{words}[i]$ 的前缀。

如果第一段划分出一个长为 $1$ 的子串，那么接下来要解决的问题为：$\textit{target}$ 的长为 $n-1$ 的后缀的最小划分个数，也就是从下标 $1$ 开始，剩余字符串的最小划分个数。

一般地，对于每个 $i$，都计算一个 $r_i$，满足 $\textit{target}$ 下标从 $i$ 到 $r_i$ 这一段是某个 $\textit{words}[i]$ 的前缀。

算出 $r_i$ 后，我们可以枚举当前这一段的长度：

- 长为 $1$，那么接下来思考从 $i+1$ 开始，剩余字符串的最小划分个数。
- 长为 $2$，那么接下来思考从 $i+2$ 开始，剩余字符串的最小划分个数。
- ……
- 长为 $r_i$，那么接下来思考从 $i+r_i$ 开始，剩余字符串的最小划分个数。

相当于我们可以从下标 $i$「跳到」下标 $i+1,i+2,\cdots,r_i+1$ 中的任意位置。如果跳到 $n$，表示 $\textit{target}$ 划分完毕。

问题转换成：

1. 计算 $r_i$。
2. 计算从 $0$ 跳到 $n$ 的最小跳跃次数。这类似 [45. 跳跃游戏 II](https://leetcode.cn/problems/jump-game-ii/) 或者 [1326. 灌溉花园的最少水龙头数目](https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/)，请看 [我的图解（1326 题）](https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/solutions/2123855/yi-zhang-tu-miao-dong-pythonjavacgo-by-e-wqry/)。

如何计算 $r_i$？

> 可以用字典树 + 枚举，这可以通过 [周赛第三题](https://leetcode.cn/problems/minimum-number-of-valid-strings-to-form-target-i/)，但无法通过本题。

预处理每个 $\textit{words}[i]$ 的每个前缀的字符串哈希值，**按照前缀长度分组**，保存到不同的集合中。每个集合保存的是相同前缀长度的哈希值。

由于 $\textit{words}$ 的长度至多为 $100$，所以每个集合至多保存 $100$ 个哈希值，根据生日攻击理论，单模哈希绰绰有余，碰撞概率很小。

对于 $i$，我们需要计算最大的长度 $\textit{sz}$，满足 $\textit{target}$ 从下标 $i$ 开始的长为 $\textit{sz}$ 的子串的哈希值是否在对应的集合中。

这可以用**二分**算出来，原理见 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

算出 $\textit{sz}$，就可以算出从 $i$ 向右，最远可以跳到 $i+\textit{sz}$。


```
func minValidStrings(words []string, target string) (ans int) {
	n := len(target)

	// 多项式字符串哈希（方便计算子串哈希值）
	// 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
	const mod = 1_070_777_777
	base := 9e8 - rand.Intn(1e8) // 随机 base，防止 hack（注意 Go1.20 之后的版本，每次随机的数都不一样）
	powBase := make([]int, n+1)  // powBase[i] = base^i
	preHash := make([]int, n+1)  // 前缀哈希值 preHash[i] = hash(s[:i])
	powBase[0] = 1
	for i, b := range target {
		powBase[i+1] = powBase[i] * base % mod
		preHash[i+1] = (preHash[i]*base + int(b)) % mod // 秦九韶算法计算多项式哈希
	}
	// 计算子串 target[l:r] 的哈希值，注意这是左闭右开区间 [l,r)
	// 计算方法类似前缀和
	subHash := func(l, r int) int {
		return ((preHash[r]-preHash[l]*powBase[r-l])%mod + mod) % mod
	}

	maxLen := 0
	for _, w := range words {
		maxLen = max(maxLen, len(w))
	}
	sets := make([]map[int]bool, maxLen)
	for i := range sets {
		sets[i] = map[int]bool{}
	}
	for _, w := range words {
		h := 0
		for j, b := range w {
			h = (h*base + int(b)) % mod
			sets[j][h] = true // 注意 j 从 0 开始
		}
	}

	curR := 0 // 已建造的桥的右端点
	nxtR := 0 // 下一座桥的右端点的最大值
	for i := range target {
		sz := sort.Search(min(n-i, maxLen), func(sz int) bool { return !sets[sz][subHash(i, i+sz+1)] })
		nxtR = max(nxtR, i+sz)
		if i == curR { // 到达已建造的桥的右端点
			if i == nxtR { // 无论怎么造桥，都无法从 i 到 i+1
				return -1
			}
			curR = nxtR // 建造下一座桥
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L + n\log n)$，其中 $n$ 是 $\textit{target}$ 的长度，$L$ 是 $\textit{words}$ 中所有字符串的长度之和。
- 空间复杂度：$\mathcal{O}(L + n)$。

## 方法二：AC 自动机优化 DP

方法一用到了字符串哈希，并不能保证 100% 正确。正解是 AC 自动机。

看示例 1，对比以下两个 $\textit{target}$ 的前缀：

- $\texttt{aabcd}$，需要连接 $2$ 个 $\textit{words}[i]$ 的前缀 $\texttt{aa}$ 和 $\texttt{bcd}$。
- $\texttt{aab}$，需要连接多少个前缀？可以肯定的是，答案一定不会比 $2$ 还大，因为我们把 $\texttt{aabcd}$ 末尾的 $\texttt{cd}$ 去掉就可以得到 $\texttt{aab}$。这仍然是 $2$ 个前缀的连接。

根据上述讨论，如果用 $f[i]$ 表示 $\textit{target}$ 的长为 $i$ 的前缀需要连接的最少字符串数量，那么 $f[i]\le f[i+1]$ 一定成立。

既然 $f$ 是有序数组，那么对于 $f[i]$，我们只需要知道最小的 $j$，满足从 $\textit{target}[j]$ 到 $\textit{target}[i-1]$ 是某个 $\textit{words}[i]$ 的前缀。

也就是说，匹配的 $\textit{words}[i]$ 的前缀要尽量长。这正是 **AC 自动机**的应用。原理见 [OI Wiki](https://oi-wiki.org/string/ac-automaton/)。学习之前推荐先看看我的 [KMP 原理讲解](https://www.zhihu.com/question/21923021/answer/37475572)。

算出了 $j$，那么有

$$
f[i] = f[j] + 1
$$

初始值 $f[0]=0$。

答案为 $f[n]$。

如果 AC 自动机没法匹配任何 $\textit{words}[i]$ 的非空前缀，返回 $-1$。

```
// 从根到 node 的字符串是某个（某些）words[i] 的前缀
type node struct {
	son  [26]*node
	fail *node // 当 cur.son[i] 不能匹配 target 中的某个字符时，cur.fail.son[i] 即为下一个待匹配节点（等于 root 则表示没有匹配）
	len  int   // 从根到 node 的字符串的长度，也是 node 在 trie 中的深度
}

type acam struct {
	root *node
}

func (ac *acam) put(s string) {
	cur := ac.root
	for _, b := range s {
		b -= 'a'
		if cur.son[b] == nil {
			cur.son[b] = &node{len: cur.len + 1}
		}
		cur = cur.son[b]
	}
}

func (ac *acam) buildFail() {
	ac.root.fail = ac.root
	q := []*node{}
	for i, son := range ac.root.son[:] {
		if son == nil {
			ac.root.son[i] = ac.root
		} else {
			son.fail = ac.root // 第一层的失配指针，都指向根节点 ∅
			q = append(q, son)
		}
	}
	// BFS
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for i, son := range cur.son[:] {
			if son == nil {
				// 虚拟子节点 cur.son[i]，和 cur.fail.son[i] 是同一个
				// 方便失配时直接跳到下一个可能匹配的位置（但不一定是某个 words[k] 的最后一个字母）
				cur.son[i] = cur.fail.son[i]
				continue
			}
			son.fail = cur.fail.son[i] // 计算失配位置
			q = append(q, son)
		}
	}
}

func minValidStrings(words []string, target string) int {
	ac := &acam{root: &node{}}
	for _, w := range words {
		ac.put(w)
	}
	ac.buildFail()

	n := len(target)
	f := make([]int, n+1)
	cur := ac.root
	for i, b := range target {
		// 如果没有匹配相当于移动到 fail 的 son[b-'a']
		cur = cur.son[b-'a']
		// 没有任何字符串的前缀与 target[:i+1] 的后缀匹配
		if cur == ac.root {
			return -1
		}
		f[i+1] = f[i+1-cur.len] + 1
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L|\Sigma| + n)$，其中 $n$ 是 $\textit{target}$ 的长度，$L$ 是 $\textit{words}$ 中所有字符串的长度之和，$|\Sigma|$ 是字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(L|\Sigma| + n)$。

## 分类题单

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