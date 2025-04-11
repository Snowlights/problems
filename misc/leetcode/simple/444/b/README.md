#### 题目

<p>请你设计一个数据结构来高效管理网络路由器中的数据包。每个数据包包含以下属性：</p>

<ul>
	<li><code>source</code>：生成该数据包的机器的唯一标识符。</li>
	<li><code>destination</code>：目标机器的唯一标识符。</li>
	<li><code>timestamp</code>：该数据包到达路由器的时间戳。</li>
</ul>

<p>实现 <code>Router</code> 类：</p>

<p><code>Router(int memoryLimit)</code>：初始化路由器对象，并设置固定的内存限制。</p>

<ul>
	<li><code>memoryLimit</code> 是路由器在任意时间点可以存储的 <strong>最大</strong> 数据包数量。</li>
	<li>如果添加一个新数据包会超过这个限制，则必须移除 <strong>最旧的</strong> 数据包以腾出空间。</li>
</ul>

<p><code>bool addPacket(int source, int destination, int timestamp)</code>：将具有给定属性的数据包添加到路由器。</p>

<ul>
	<li>如果路由器中已经存在一个具有相同 <code>source</code>、<code>destination</code> 和 <code>timestamp</code> 的数据包，则视为重复数据包。</li>
	<li>如果数据包成功添加（即不是重复数据包），返回 <code>true</code>；否则返回 <code>false</code>。</li>
</ul>

<p><code>int[] forwardPacket()</code>：以 FIFO（先进先出）顺序转发下一个数据包。</p>

<ul>
	<li>从存储中移除该数据包。</li>
	<li>以数组 <code>[source, destination, timestamp]</code> 的形式返回该数据包。</li>
	<li>如果没有数据包可以转发，则返回空数组。</li>
</ul>

<p><code>int getCount(int destination, int startTime, int endTime)</code>：</p>

<ul>
	<li>返回当前存储在路由器中（即尚未转发）的，且目标地址为指定 <code>destination</code> 且时间戳在范围 <code>[startTime, endTime]</code>（包括两端）内的数据包数量。</li>
</ul>

<p><strong>注意</strong>：对于 <code>addPacket</code> 的查询会按照 <code>timestamp</code> 的递增顺序进行。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><br />
<span class="example-io">["Router", "addPacket", "addPacket", "addPacket", "addPacket", "addPacket", "forwardPacket", "addPacket", "getCount"]<br />
[[3], [1, 4, 90], [2, 5, 90], [1, 4, 90], [3, 5, 95], [4, 5, 105], [], [5, 2, 110], [5, 100, 110]]</span></p>

<p><strong>输出：</strong><br />
<span class="example-io">[null, true, true, false, true, true, [2, 5, 90], true, 1] </span></p>

<p><strong>解释：</strong></p>
<code>Router router = new Router(3);</code> // 初始化路由器，内存限制为 3。<br />
<code>router.addPacket(1, 4, 90);</code> // 数据包被添加，返回 True。<br />
<code>router.addPacket(2, 5, 90);</code> // 数据包被添加，返回 True。<br />
<code>router.addPacket(1, 4, 90);</code> // 这是一个重复数据包，返回 False。<br />
<code>router.addPacket(3, 5, 95);</code> // 数据包被添加，返回 True。<br />
<code>router.addPacket(4, 5, 105);</code> // 数据包被添加，<code>[1, 4, 90]</code> 被移除，因为数据包数量超过限制，返回 True。<br />
<code>router.forwardPacket();</code> // 转发数据包 <code>[2, 5, 90]</code> 并将其从路由器中移除。<br />
<code>router.addPacket(5, 2, 110);</code> // 数据包被添加，返回 True。<br />
<code>router.getCount(5, 100, 110);</code> // 唯一目标地址为 5 且时间在 <code>[100, 110]</code>&nbsp;范围内的数据包是 <code>[4, 5, 105]</code>，返回 1。</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><br />
<span class="example-io">["Router", "addPacket", "forwardPacket", "forwardPacket"]<br />
[[2], [7, 4, 90], [], []]</span></p>

<p><strong>输出：</strong><br />
<span class="example-io">[null, true, [7, 4, 90], []] </span></p>

<p><strong>解释：</strong></p>
<code>Router router = new Router(2);</code> // 初始化路由器，内存限制为 2。<br />
<code>router.addPacket(7, 4, 90);</code> // 返回 True。<br />
<code>router.forwardPacket();</code> // 返回 <code>[7, 4, 90]</code>。<br />
<code>router.forwardPacket();</code> // 没有数据包可以转发，返回 <code>[]</code>。</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= memoryLimit &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= source, destination &lt;= 2 * 10<sup>5</sup></code></li>
	<li><code>1 &lt;= timestamp &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= startTime &lt;= endTime &lt;= 10<sup>9</sup></code></li>
	<li><code>addPacket</code>、<code>forwardPacket</code> 和 <code>getCount</code> 方法的总调用次数最多为 <code>10<sup>5</sup></code>。</li>
	<li>对于 <code>addPacket</code> 的查询，<code>timestamp</code> 按递增顺序给出。</li>
</ul>

#### 思路

题目要求 FIFO（先进先出），这可以用**队列**实现。
为了判重，可以用哈希表记录数据包。
对于 $\texttt{getCount}$，需要按照 $\textit{destination}$ 分组，所以要用哈希表套队列。

具体来说，创建三个数据结构：
1. $\textit{packetQ}$：存储数据包的队列。
2. $\textit{packetSet}$：存储所有未转发的数据包，方便判重。
3. $\textit{destToTimestamps}$：哈希表套队列，key 是 $\textit{destination}$，value 是对应的由 $\textit{timestamp}$ 组成的队列。

> 注：可以只把 $\textit{destination}$ 保存到 $\textit{packetQ}$ 中，$\textit{destToTimestamps}$ 中额外保存 $\textit{source}$。但考虑到二分是本题的瓶颈，所以 $\textit{destToTimestamps}$ 只保存 $\textit{timestamp}$ 更好。

其他就是模拟了，具体见代码。

$\texttt{getCount}$ 可以用 [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/) 的做法， 二分查找。为了方便二分，可以用列表模拟队列。

```
type packet struct {
	source, destination, timestamp int
}

type Router struct {
	memoryLimit      int
	packetQ          []packet            // packet 队列
	packetSet        map[packet]struct{} // packet 集合
	destToTimestamps map[int][]int       // destination -> [timestamp]
}

func Constructor(memoryLimit int) Router {
	return Router{
		memoryLimit:      memoryLimit,
		packetSet:        map[packet]struct{}{},
		destToTimestamps: map[int][]int{},
	}
}

func (r *Router) AddPacket(source, destination, timestamp int) bool {
	pkt := packet{source, destination, timestamp}
	if _, ok := r.packetSet[pkt]; ok {
		return false
	}
	r.packetSet[pkt] = struct{}{}
	if len(r.packetQ) == r.memoryLimit { // 太多了
		r.ForwardPacket()
	}
	r.packetQ = append(r.packetQ, pkt) // 入队
	r.destToTimestamps[destination] = append(r.destToTimestamps[destination], timestamp)
	return true
}

func (r *Router) ForwardPacket() []int {
	if len(r.packetQ) == 0 {
		return nil
	}
	pkt := r.packetQ[0]
	r.packetQ = r.packetQ[1:] // 出队
	r.destToTimestamps[pkt.destination] = r.destToTimestamps[pkt.destination][1:]
	delete(r.packetSet, pkt)
	return []int{pkt.source, pkt.destination, pkt.timestamp}
}

func (r *Router) GetCount(destination, startTime, endTime int) int {
	timestamps := r.destToTimestamps[destination]
	return sort.SearchInts(timestamps, endTime+1) - sort.SearchInts(timestamps, startTime)
}
```

#### 复杂度分析

- 时间复杂度：$\texttt{GetCount}$ 是 $\mathcal{O}(\log \min(q, \textit{memoryLimit}))$，其中 $q$ 是 $\texttt{addPacket}$ 的调用次数。其余操作为 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(\min(q, \textit{memoryLimit}))$。

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