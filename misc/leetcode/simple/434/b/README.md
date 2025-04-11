#### 题目

<p>给你一个整数&nbsp;<code>numberOfUsers</code>&nbsp;表示用户总数，另有一个大小为 <code>n x 3</code> 的数组&nbsp;<code>events</code>&nbsp;。</p>

<p>每个&nbsp;<code inline="">events[i]</code>&nbsp;都属于下述两种类型之一：</p>

<ol>
	<li><strong>消息事件（Message Event）：</strong><code>["MESSAGE", "timestamp<sub>i</sub>", "mentions_string<sub>i</sub>"]</code>
	<ul>
		<li>事件表示在&nbsp;<code>timestamp<sub>i</sub></code>&nbsp;时，一组用户被消息提及。</li>
		<li><code>mentions_string<sub>i</sub></code>&nbsp;字符串包含下述标识符之一：
		<ul>
			<li><code>id&lt;number&gt;</code>：其中&nbsp;<code>&lt;number&gt;</code>&nbsp;是一个区间&nbsp;<code>[0,numberOfUsers - 1]</code>&nbsp;内的整数。可以用单个空格分隔&nbsp;<strong>多个</strong> id ，并且 id 可能重复。此外，这种形式可以提及离线用户。</li>
			<li><code>ALL</code>：提及 <strong>所有</strong> 用户。</li>
			<li><code>HERE</code>：提及所有 <strong>在线</strong> 用户。</li>
		</ul>
		</li>
	</ul>
	</li>
	<li><strong>离线事件（Offline Event）：</strong><code>["OFFLINE", "timestamp<sub>i</sub>", "id<sub>i</sub>"]</code>
	<ul>
		<li>事件表示用户&nbsp;<code>id<sub>i</sub></code>&nbsp;在&nbsp;<code>timestamp<sub>i</sub></code>&nbsp;时变为离线状态 <strong>60 个单位时间</strong>。用户会在&nbsp;<code>timestamp<sub>i</sub> + 60</code>&nbsp;时自动再次上线。</li>
	</ul>
	</li>
</ol>

<p>返回数组&nbsp;<code>mentions</code>&nbsp;，其中&nbsp;<code>mentions[i]</code>&nbsp;表示 &nbsp;id 为 &nbsp;<code>i</code>&nbsp;的用户在所有&nbsp;<code>MESSAGE</code>&nbsp;事件中被提及的次数。</p>

<p>最初所有用户都处于在线状态，并且如果某个用户离线或者重新上线，其对应的状态变更将会在所有相同时间发生的消息事件之前进行处理和同步。</p>

<p><strong>注意 </strong>在单条消息中，同一个用户可能会被提及多次。每次提及都需要被 <strong>分别</strong> 统计。</p>

<p>&nbsp;</p>

<p><b>示例 1：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>numberOfUsers = 2, events = [["MESSAGE","10","id1 id0"],["OFFLINE","11","0"],["MESSAGE","71","HERE"]]</span></p>

<p><span class="example-io"><b>输出：</b>[2,2]</span></p>

<p><b>解释：</b></p>

<p>最初，所有用户都在线。</p>

<p>时间戳 10 ，<code>id1</code>&nbsp;和&nbsp;<code>id0</code>&nbsp;被提及，<code>mentions = [1,1]</code></p>

<p>时间戳 11 ，<code>id0</code>&nbsp;<strong>离线</strong> 。</p>

<p>时间戳 71 ，<code>id0</code>&nbsp;再次 <strong>上线</strong>&nbsp;并且&nbsp;<code>"HERE"</code>&nbsp;被提及，<code>mentions = [2,2]</code></p>
</div>

<p><b>示例 2：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>numberOfUsers = 2, events = [["MESSAGE","10","id1 id0"],["OFFLINE","11","0"],["MESSAGE","12","ALL"]]</span></p>

<p><span class="example-io"><b>输出：</b>[2,2]</span></p>

<p><b>解释：</b></p>

<p>最初，所有用户都在线。</p>

<p>时间戳 10 ，<code>id1</code>&nbsp;和&nbsp;<code>id0</code>&nbsp;被提及，<code>mentions = [1,1]</code></p>

<p>时间戳 11 ，<code>id0</code>&nbsp;<strong>离线</strong> 。</p>

<p>时间戳 12 ，<code>"ALL"</code>&nbsp;被提及。这种方式将会包括所有离线用户，所以&nbsp;<code>id0</code>&nbsp;和&nbsp;<code>id1</code>&nbsp;都被提及，<code>mentions = [2,2]</code></p>
</div>

<p><b>示例 3：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>numberOfUsers = 2, events = [["OFFLINE","10","0"],["MESSAGE","12","HERE"]]</span></p>

<p><span class="example-io"><b>输出：</b>[0,1]</span></p>

<p><b>解释：</b></p>

<p>最初，所有用户都在线。</p>

<p>时间戳 10 ，<code>id0</code>&nbsp;<strong>离线</strong>&nbsp;<b>。</b></p>

<p>时间戳 12 ，<code>"HERE"</code>&nbsp;被提及。由于&nbsp;<code>id0</code>&nbsp;仍处于离线状态，其将不会被提及，<code>mentions = [0,1]</code></p>
</div>

<p>&nbsp;</p>

<p><b>提示：</b></p>

<ul>
	<li><code>1 &lt;= numberOfUsers &lt;= 100</code></li>
	<li><code>1 &lt;= events.length &lt;= 100</code></li>
	<li><code>events[i].length == 3</code></li>
	<li><code>events[i][0]</code>&nbsp;的值为&nbsp;<code>MESSAGE</code>&nbsp;或&nbsp;<code>OFFLINE</code>&nbsp;。</li>
	<li><code>1 &lt;= int(events[i][1]) &lt;= 10<sup>5</sup></code></li>
	<li>在任意 <code>"MESSAGE"</code>&nbsp;事件中，以&nbsp;<code>id&lt;number&gt;</code>&nbsp;形式提及的用户数目介于&nbsp;<code>1</code>&nbsp;和&nbsp;<code>100</code>&nbsp;之间。</li>
	<li><code>0 &lt;= &lt;number&gt; &lt;= numberOfUsers - 1</code></li>
	<li>题目保证 <code>OFFLINE</code>&nbsp;引用的用户 id 在事件发生时处于 <strong>在线</strong> 状态。</li>
</ul>

#### 思路

按照时间戳从小到大排序，时间戳相同的，离线事件排在前面，因为题目要求「状态变更在所有相同时间发生的消息事件之前进行处理」。

**离线事件**：用一个数组 $\textit{onlineT}$ 标记用户下次在线的时间戳。如果 $\textit{onlineT}[i]\le$ 当前时间戳，则表示用户 $i$ 已在线。

**消息事件**：按照规则把相应用户的答案加一。

```
func countMentions(numberOfUsers int, events [][]string) []int {
	// 按照时间戳从小到大排序，时间戳相同的，离线事件排在前面
	slices.SortFunc(events, func(a, b []string) int {
		ta, _ := strconv.Atoi(a[1])
		tb, _ := strconv.Atoi(b[1])
		return cmp.Or(ta-tb, int(b[0][0])-int(a[0][0]))
	})

	ans := make([]int, numberOfUsers)
	onlineT := make([]int, numberOfUsers)
	for _, e := range events {
		curT, _ := strconv.Atoi(e[1])
		if e[0][0] == 'O' {
			i, _ := strconv.Atoi(e[2])
			onlineT[i] = curT + 60
		} else if e[2][0] == 'A' {
			for i := range ans {
				ans[i]++
			}
		} else if e[2][0] == 'H' {
			for i, t := range onlineT {
				if t <= curT { // 在线
					ans[i]++
				}
			}
		} else {
			for _, s := range strings.Split(e[2], " ") {
				i, _ := strconv.Atoi(s[2:])
				ans[i]++
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + m\log m\log U + L)$，其中 $n$ 是 $\textit{numberOfUsers}$，$m$ 是 $\textit{events}$ 的长度，$U$ 是时间戳的最大值，$L$ 是所有 `mentions_string` 的长度之和。排序需要 $\mathcal{O}(m\log m)$ 次比较，每次比较需要 $\mathcal{O}(\log U)$ 的时间把长为 $\mathcal{O}(\log U)$ 的字符串时间戳转成整数。注：如果预处理这个转换，可以把排序的过程优化至 $\mathcal{O}(m\log m)$。
- 空间复杂度：$\mathcal{O}(n)$。


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