#### 题目

<p>给你一个字符串 <code>s</code>，模拟每秒钟的事件 <code>i</code>：</p>

<ul>
	<li>如果 <code>s[i] == 'E'</code>，表示有一位顾客进入候诊室并占用一把椅子。</li>
	<li>如果 <code>s[i] == 'L'</code>，表示有一位顾客离开候诊室，从而释放一把椅子。</li>
</ul>

<p>返回保证每位进入候诊室的顾客都能有椅子坐的<strong> 最少 </strong>椅子数，假设候诊室最初是 <strong>空的 </strong>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">s = "EEEEEEE"</span></p>

<p><strong>输出：</strong><span class="example-io">7</span></p>

<p><strong>解释：</strong></p>

<p>每秒后都有一个顾客进入候诊室，没有人离开。因此，至少需要 7 把椅子。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">s = "ELELEEL"</span></p>

<p><strong>输出：</strong><span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<p>假设候诊室里有 2 把椅子。下表显示了每秒钟等候室的状态。</p>
</div>
<table>
	<tbody>
		<tr>
			<th>秒</th>
			<th>事件</th>
			<th>候诊室的人数</th>
			<th>可用的椅子数</th>
		</tr>
		<tr>
			<td>0</td>
			<td>Enter</td>
			<td>1</td>
			<td>1</td>
		</tr>
		<tr>
			<td>1</td>
			<td>Leave</td>
			<td>0</td>
			<td>2</td>
		</tr>
		<tr>
			<td>2</td>
			<td>Enter</td>
			<td>1</td>
			<td>1</td>
		</tr>
		<tr>
			<td>3</td>
			<td>Leave</td>
			<td>0</td>
			<td>2</td>
		</tr>
		<tr>
			<td>4</td>
			<td>Enter</td>
			<td>1</td>
			<td>1</td>
		</tr>
		<tr>
			<td>5</td>
			<td>Enter</td>
			<td>2</td>
			<td>0</td>
		</tr>
		<tr>
			<td>6</td>
			<td>Leave</td>
			<td>1</td>
			<td>1</td>
		</tr>
	</tbody>
</table>


<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">s = "ELEELEELLL"</span></p>

<p><strong>输出：</strong><span class="example-io">3</span></p>

<p><strong>解释：</strong></p>

<p>假设候诊室里有 3 把椅子。下表显示了每秒钟等候室的状态。</p>
</div>
<table>
	<tbody>
		<tr>
			<th>秒</th>
			<th>事件</th>
			<th>候诊室的人数</th>
			<th>可用的椅子数</th>
		</tr>
		<tr>
			<td>0</td>
			<td>Enter</td>
			<td>1</td>
			<td>2</td>
		</tr>
		<tr>
			<td>1</td>
			<td>Leave</td>
			<td>0</td>
			<td>3</td>
		</tr>
		<tr>
			<td>2</td>
			<td>Enter</td>
			<td>1</td>
			<td>2</td>
		</tr>
		<tr>
			<td>3</td>
			<td>Enter</td>
			<td>2</td>
			<td>1</td>
		</tr>
		<tr>
			<td>4</td>
			<td>Leave</td>
			<td>1</td>
			<td>2</td>
		</tr>
		<tr>
			<td>5</td>
			<td>Enter</td>
			<td>2</td>
			<td>1</td>
		</tr>
		<tr>
			<td>6</td>
			<td>Enter</td>
			<td>3</td>
			<td>0</td>
		</tr>
		<tr>
			<td>7</td>
			<td>Leave</td>
			<td>2</td>
			<td>1</td>
		</tr>
		<tr>
			<td>8</td>
			<td>Leave</td>
			<td>1</td>
			<td>2</td>
		</tr>
		<tr>
			<td>9</td>
			<td>Leave</td>
			<td>0</td>
			<td>3</td>
		</tr>
	</tbody>
</table>


<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 50</code></li>
	<li><code>s</code> 仅由字母 <code>'E'</code> 和 <code>'L'</code> 组成。</li>
	<li><code>s</code> 表示一个有效的进出序列。</li>
</ul>

#### 思路

把 E 看成上车，L 看成下车，问：车上最多有多少人？模拟即可。

```
func minimumChairs(s string) (ans int) {
	cnt := 0
	for _, c := range s {
		if c == 'E' {
			cnt++
			ans = max(ans, cnt)
		} else {
			cnt--
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{s}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。


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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)