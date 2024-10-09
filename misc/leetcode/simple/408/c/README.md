#### 题目

<p>给你一个二进制字符串 <code>s</code>。</p>

<p>请你统计并返回其中 <strong>1 显著 </strong> 的 <span data-keyword="substring-nonempty">子字符串</span> 的数量。</p>

<p>如果字符串中 1 的数量 <strong>大于或等于</strong> 0 的数量的 <strong>平方</strong>，则认为该字符串是一个 <strong>1 显著</strong> 的字符串 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">s = "00011"</span></p>

<p><strong>输出：</strong><span class="example-io">5</span></p>

<p><strong>解释：</strong></p>

<p>1 显著的子字符串如下表所示。</p>
</div>

<table>
	<thead>
		<tr>
			<th>i</th>
			<th>j</th>
			<th>s[i..j]</th>
			<th>0 的数量</th>
			<th>1 的数量</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>3</td>
			<td>3</td>
			<td>1</td>
			<td>0</td>
			<td>1</td>
		</tr>
		<tr>
			<td>4</td>
			<td>4</td>
			<td>1</td>
			<td>0</td>
			<td>1</td>
		</tr>
		<tr>
			<td>2</td>
			<td>3</td>
			<td>01</td>
			<td>1</td>
			<td>1</td>
		</tr>
		<tr>
			<td>3</td>
			<td>4</td>
			<td>11</td>
			<td>0</td>
			<td>2</td>
		</tr>
		<tr>
			<td>2</td>
			<td>4</td>
			<td>011</td>
			<td>1</td>
			<td>2</td>
		</tr>
	</tbody>
</table>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">s = "101101"</span></p>

<p><strong>输出：</strong><span class="example-io">16</span></p>

<p><strong>解释：</strong></p>

<p>1 不显著的子字符串如下表所示。</p>

<p>总共有 21 个子字符串，其中 5 个是 1 不显著字符串，因此有 16 个 1 显著子字符串。</p>
</div>

<table>
	<thead>
		<tr>
			<th>i</th>
			<th>j</th>
			<th>s[i..j]</th>
			<th>0 的数量</th>
			<th>1 的数量</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>1</td>
			<td>1</td>
			<td>0</td>
			<td>1</td>
			<td>0</td>
		</tr>
		<tr>
			<td>4</td>
			<td>4</td>
			<td>0</td>
			<td>1</td>
			<td>0</td>
		</tr>
		<tr>
			<td>1</td>
			<td>4</td>
			<td>0110</td>
			<td>2</td>
			<td>2</td>
		</tr>
		<tr>
			<td>0</td>
			<td>4</td>
			<td>10110</td>
			<td>2</td>
			<td>3</td>
		</tr>
		<tr>
			<td>1</td>
			<td>5</td>
			<td>01101</td>
			<td>2</td>
			<td>3</td>
		</tr>
	</tbody>
</table>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 4 * 10<sup>4</sup></code></li>
	<li><code>s</code> 仅包含字符 <code>'0'</code> 和 <code>'1'</code>。</li>
</ul>

#### 思路

注意到，如果子串中的 $0$ 非常多，多到 $0$ 的个数的平方比 $1$ 的个数都要大，那么这样的子串必然不是 $1$ 显著子串。

设 $\textit{cnt}_0$ 为子串中的 $0$ 的个数，$\textit{cnt}_1$ 为子串中的 $1$ 的个数，那么必须满足

$$
\textit{cnt}_0^2 \le \textit{cnt}_1 \le n
$$

解得

$$
\textit{cnt}_0 \le \sqrt{n}
$$

所以子串中的 $0$ 的个数不会超过 $\sqrt{n}$。

据此：
- 枚举子串左端点 $\textit{left}$。
- 枚举子串中有**恰好** $\textit{cnt}_0 = 0,1,2,3,\cdots$ 个 $0$，如果 $\textit{cnt}_0^2$ 超过了 $s$ 中的 $1$ 的个数，则退出循环。这只会枚举 $\mathcal{O}(\sqrt{n})$ 次。

定义：
- 子串左端点为 $\textit{left}$。
- 子串中**恰好**有 $\textit{cnt}_0$ 个 $0$。那么子串右端点的下标范围就是 $[p,q-1]$，其中 $p$ 为从 $\textit{left}$ 开始的第 $\textit{cnt}_0$ 个 $0$ 的下标，$q$ 为从 $\textit{left}$ 开始的第 $\textit{cnt}_0+1$ 个 $0$ 的下标。
- 从子串左端点 $\textit{left}$ 到 $p$ 之间有 $\textit{cnt}_1$ 个 $1$。

分类讨论：
- 如果 $\textit{cnt}_0^2\le \textit{cnt}_1$，那么子串右端点可以是 $p,p+1,p+2,\cdots, q-1$，一共有 $q-p$ 个。注意一定要保证子串中**恰好**有 $\textit{cnt}_0$ 个 $0$。
- 如果 $\textit{cnt}_0^2> \textit{cnt}_1$，那么为了补足 $1$ 的个数，子串右端点的最小值就不是 $p$ 了，而是 $p + (\textit{cnt}_0^2 - \textit{cnt}_1)$，一共有 $q - p - (\textit{cnt}_0^2 - \textit{cnt}_1)$ 个。如果这个值是负数，则说明没有符合要求的子串。

综上所述，当子串左端点为 $\textit{left}$ 且子串中有恰好 $\textit{cnt}_0$ 个 $0$ 时，一共有

$$
\max(q-p-d, 0)
$$

个 $1$ 显著子串。其中 

$$
d = \max(\textit{cnt}_0^2 - \textit{cnt}_1, 0)
$$

## 写法一：枚举左端点

```
func numberOfSubstrings(s string) (ans int) {
	a := []int{}
	for i, b := range s {
		if b == '0' {
			a = append(a, i)
		}
	}

	n := len(s)
	tot1 := n - len(a)
	a = append(a, n) // 哨兵

	for left, b := range s {
		if b == '1' {
			ans += a[0] - left // 不含 0 的子串个数
		}
		for k, j := range a[:len(a)-1] {
			cnt0 := k + 1
			if cnt0*cnt0 > tot1 {
				break
			}
			cnt1 := j - left - k
			ans += max(a[k+1]-j-max(cnt0*cnt0-cnt1, 0), 0)
		}
		if b == '0' {
			a = a[1:] // 这个 0 后面不会再枚举到了
		}
	}
	return
} 
```

## 写法二：枚举右端点

``` 
func numberOfSubstrings(s string) (ans int) {
	tot1 := 0
	a := []int{-1} // 哨兵
	for right, b := range s {
		if b == '0' {
			a = append(a, right)
		} else {
			ans += right - a[len(a)-1]
			tot1++
		}
		for k := len(a) - 1; k > 0 && (len(a)-k)*(len(a)-k) <= tot1; k-- {
			cnt0 := len(a) - k
			cnt1 := right - a[k] + 1 - cnt0
			ans += max(a[k]-a[k-1]-max(cnt0*cnt0-cnt1, 0), 0)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt{n})$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。\n\n注：使用队列，只保存 $\textit{right}$ 及其左侧的 $\mathcal{O}(\sqrt{n})$ 个 $0$ 的下标，可以把空间复杂度优化到 $\mathcal{O}(\sqrt{n})$。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
