#### 题目

<p>给你一个长度为 <code>n</code>&nbsp;的整数数组 <code>pizzas</code>，其中 <code>pizzas[i]</code> 表示第 <code>i</code>&nbsp;个披萨的重量。每天你会吃&nbsp;<strong>恰好</strong> 4 个披萨。由于你的新陈代谢能力惊人，当你吃重量为 <code>W</code>、<code>X</code>、<code>Y</code> 和 <code>Z</code> 的披萨（其中 <code>W &lt;= X &lt;= Y &lt;= Z</code>）时，你只会增加 1 个披萨的重量！体重增加规则如下：</p>

<ul>
	<li>在&nbsp;<strong><span style="box-sizing: border-box; margin: 0px; padding: 0px;">奇数天</span></strong>（按 <strong>1 开始计数</strong>）你会增加 <code>Z</code> 的重量。</li>
	<li>在&nbsp;<strong>偶数天</strong>，你会增加 <code>Y</code> 的重量。</li>
</ul>

<p>请你设计吃掉&nbsp;<strong>所有&nbsp;</strong>披萨的最优方案，并计算你可以增加的&nbsp;<strong>最大&nbsp;</strong>总重量。</p>

<p><strong>注意：</strong>保证 <code>n</code> 是 4 的倍数，并且每个披萨只吃一次。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">pizzas = [1,2,3,4,5,6,7,8]</span></p>

<p><strong>输出：</strong> <span class="example-io">14</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>第 1 天，你吃掉下标为 <code>[1, 2, 4, 7] = [2, 3, 5, 8]</code> 的披萨。你增加的重量为 8。</li>
	<li>第 2 天，你吃掉下标为 <code>[0, 3, 5, 6] = [1, 4, 6, 7]</code> 的披萨。你增加的重量为 6。</li>
</ul>

<p>吃掉所有披萨后，你增加的总重量为 <code>8 + 6 = 14</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">pizzas = [2,1,1,1,1,1,1,1]</span></p>

<p><strong>输出：</strong> <span class="example-io">3</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>第 1 天，你吃掉下标为 <code>[4, 5, 6, 0] = [1, 1, 1, 2]</code> 的披萨。你增加的重量为 2。</li>
	<li>第 2 天，你吃掉下标为 <code>[1, 2, 3, 7] = [1, 1, 1, 1]</code> 的披萨。你增加的重量为 1。</li>
</ul>

<p>吃掉所有披萨后，你增加的总重量为 <code>2 + 1 = 3</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>4 &lt;= n == pizzas.length &lt;= 2 * 10<sup><span style="font-size: 10.8333px;">5</span></sup></code></li>
	<li><code>1 &lt;= pizzas[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>n</code> 是 4 的倍数。</li>
</ul>

#### 思路

推荐先完成本题的简单版本：[1561. 你可以获得的最大硬币数目](https://leetcode.cn/problems/maximum-number-of-coins-you-can-get/)。

一共有 $\textit{days}=\dfrac{n}{4}$ 天，其中有 $\textit{odd} = \left\lceil\dfrac{\textit{days}}{2}\right\rceil = \left\lfloor\dfrac{\textit{days}+1}{2}\right\rfloor$ 个奇数天，
$\textit{even}=\left\lfloor\dfrac{\textit{days}}{2}\right\rfloor$ 个偶数天。


根据题意，**奇数天选最大的，偶数天只能选次大的**。例如 $\textit{odd}=3,\ \textit{even}=2$，我们可以先思考奇数天怎么选（先选前 $3$ 个最大的），
再思考偶数天怎么选（再从剩余披萨中跳着选 $2$ 个最大的）。假设披萨从大到小排序，那么选择方案为

$$
选选选\underline{\phantom{选}}选\underline{\phantom{选}}选
$$

其中 $\underline{\phantom{选}}$ 表示跳过不选。

用**交换论证法**可以证明这是最优的。比如交换第二个奇数天和最后一个偶数天，也就是按照「奇偶奇偶奇」的方法选，那么选择方案为

$$
选\underline{\phantom{选}}选选\underline{\phantom{选}}选选
$$

对比可以发现，本质上这个交换把后面的一个 $\underline{\phantom{选}}$ 插入到了前面的某个位置中，于是这两个位置之间的所有「选」的下标都向后移动了一位，
所以这个方案的元素和一定不会比交换前的优。

所以答案为

$$
\sum_{i=0}^{\textit{odd}-1} \textit{pizzas}[i] + \sum_{i=0}^{\textit{even}-1} \textit{pizzas}[\textit{odd}+2i+1]
$$

其中 $\textit{pizzas}$ 按照从大到小的顺序排序。

```
func maxWeight(pizzas []int) (ans int64) {
	slices.SortFunc(pizzas, func(a, b int) int { return b - a })
	days := len(pizzas) / 4
	odd := (days + 1) / 2
	for _, x := range pizzas[:odd] {
		ans += int64(x)
	}
	for i := range days / 2 {
		ans += int64(pizzas[odd+i*2+1])
	}
	return
}
```


#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{pizzas}$ 的长度，瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略切片和排序的栈开销。

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