#### 题目

<p>给你一个字符串 <code>s</code>。</p>

<p>英文字母中每个字母的&nbsp;<strong>镜像&nbsp;</strong>定义为反转字母表之后对应位置上的字母。例如，<code>'a'</code> 的镜像是 <code>'z'</code>，<code>'y'</code> 的镜像是 <code>'b'</code>。</p>

<p>最初，字符串 <code>s</code> 中的所有字符都&nbsp;<strong>未标记&nbsp;</strong>。</p>

<p>字符串 <code>s</code>&nbsp;的初始分数为 0 ，你需要对其执行以下过程：</p>

<ul>
	<li>从左到右遍历字符串。</li>
	<li>对于每个下标&nbsp;<code>i&nbsp;</code>，找到距离最近的&nbsp;<strong>未标记</strong> 下标&nbsp;<code>j</code>，下标 <code>j</code> 需要满足&nbsp;<code>j &lt; i</code> 且 <code>s[j]</code> 是 <code>s[i]</code> 的镜像。然后&nbsp;<strong>标记</strong> 下标&nbsp;<code>i</code> 和 <code>j</code>，总分加上&nbsp;<code>i - j</code>&nbsp;的值。</li>
	<li>如果对于下标&nbsp;<code>i</code>，不存在满足条件的下标&nbsp;<code>j</code>，则跳过该下标，继续处理下一个下标，不需要进行标记。</li>
</ul>

<p>返回最终的总分。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "aczzx"</span></p>

<p><strong>输出：</strong> <span class="example-io">5</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li><code>i = 0</code>。没有符合条件的下标&nbsp;<code>j</code>，跳过。</li>
	<li><code>i = 1</code>。没有符合条件的下标&nbsp;<code>j</code>，跳过。</li>
	<li><code>i = 2</code>。距离最近的符合条件的下标是 <code>j = 0</code>，因此标记下标&nbsp;0 和 2，然后将总分加上&nbsp;<code>2 - 0 = 2</code>&nbsp;。</li>
	<li><code>i = 3</code>。没有符合条件的下标&nbsp;<code>j</code>，跳过。</li>
	<li><code>i = 4</code>。距离最近的符合条件的下标是 <code>j = 1</code>，因此标记下标&nbsp;1 和 4，然后将总分加上&nbsp;<code>4 - 1 = 3</code>&nbsp;。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "abcdef"</span></p>

<p><strong>输出：</strong> <span class="example-io">0</span></p>

<p><strong>解释：</strong></p>

<p>对于每个下标&nbsp;<code>i</code>，都不存在满足条件的下标&nbsp;<code>j</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>s</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

对于每种字母，分别用一个栈维护未标记字母的下标。

遍历到 $s[i]$ 时，分类讨论：

- 如果在 $s[i]$ 左边没有 $s[i]$ 的镜像字母，那么把下标 $i$ 加到第 $i$ 个栈中。
- 否则，弹出 $s[i]$ 的镜像字母对应的栈顶，即为我们要找的 $j$。把 $i-j$ 加入答案。

```
func calculateScore(s string) (ans int64) {
	stk := [26][]int{}
	for i, c := range s {
		c -= 'a'
		if st := stk[25-c]; len(st) > 0 {
			ans += int64(i - st[len(st)-1])
			stk[25-c] = st[:len(st)-1]
		} else {
			stk[c] = append(stk[c], i)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 思路

如果固定子串的左端点，那么**子串越长，字典序越大**。
单个子串的长度不能超过多少？
由于其余 $k-1$ 个子串必须是非空的，取长度为 $1$，其余子串的长度之和**至少**为 $k-1$。
所以单个子串的长度**至多**为 $n-(k-1)$。
暴力枚举子串的左端点，计算满足长度上限的最大子串。
注意特判 $k=1$ 的情况，此时无法分割，答案就是 $s$。


```
func answerString(s string, k int) (ans string) {
	if k == 1 {
		return s
	}
	n := len(s)
	for i := range n {
		ans = max(ans, s[i:min(i+n-k+1, n)])
	}
	return
}
```


#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(n-k))$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n-k)$ 或 $\mathcal{O}(1)$。Go 的切片不会发生拷贝，只需要 $\mathcal{O}(1)$ 额外空间。

注：利用**字符串哈希+二分长度**或者**后缀数组**，可以快速比较 $s$ 任意两个子串的大小，做到 $\mathcal{O}(n\log (n-k))$ 或者 $\mathcal{O}(n\log n)$。


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
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)