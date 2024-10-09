#### 题目

<p>给你两个 <strong>正整数</strong> <code>n</code> 和 <code>k</code>。</p>

<p>如果整数 <code>x</code> 满足以下全部条件，则该整数是一个 <strong>k 回文数</strong>：</p>

<ul>
	<li><code>x</code> 是一个 <span data-keyword="palindrome-integer">回文数</span>。</li>
	<li><code>x</code> 可以被 <code>k</code> 整除。</li>
</ul>

<p>以字符串形式返回 <strong>最大的&nbsp;</strong> <code>n</code> 位 <strong>k 回文数</strong>。</p>

<p><strong>注意</strong>，该整数 <strong>不 </strong>含前导零。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">n = 3, k = 5</span></p>

<p><strong>输出：</strong> <span class="example-io">"595"</span></p>

<p><strong>解释：</strong></p>

<p>595 是最大的 3 位 k 回文数。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">n = 1, k = 4</span></p>

<p><strong>输出：</strong> <span class="example-io">"8"</span></p>

<p><strong>解释：</strong></p>

<p>1 位 k 回文数只有 4 和 8。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">n = 5, k = 6</span></p>

<p><strong>输出：</strong> <span class="example-io">"89898"</span></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= k &lt;= 9</code></li>
</ul>

#### 思路

## 分析

举个例子，如果最左最右两个位置分别填 $3$ 和 $7$，我们会得到一个 $37000\cdots00073$ 的数字（中间的 $0$ 表示还没有填入数字），假设这个数模 $k$ 的结果为 $1$。如果最左最右两个位置分别填 $9$ 和 $5$，我们会得到一个 $95000\cdots00059$ 的数字（中间的 $0$ 表示还没有填入数字），假设这个数模 $k$ 的结果也为 $1$。
无论后面的数字怎么填，这两种情况其实是一样的，因为我们只关心回文数**最终**模 $k$ 的值是否为 $0$。
把「当前从右到左填到第 $i$ 位」和「已填入的数字模 $k$ 的值为 $j$」作为状态 $(i,j)$。如果我们从状态 $(i,j)$ 开始 DFS，无论后面怎么填，回文数**最终**模 $k$ 的值不为 $0$，那么当我们第二次 DFS 到状态 $(i,j)$ 时，根据上面的例子，我们可以直接得出，无论后面怎么填，回文数**最终**模 $k$ 的值仍然不会为 $0$。

这启发我们用图上的 DFS 思考。在 DFS 中的过程中，不访问之前访问过的点。

## 建图

枚举在第 $i$ 位填入数字 $d$（从大到小枚举 $d$），那么同时也在第 $n-1-i$ 位填入了 $d$。
填入数字后，回文数模 $k$ 的值变成了

$$
j_2 = (j + d\cdot (10^i + 10^{n-1-i}))\bmod k
$$

我们可以在 $(i,j)$ 和 $(i+1, j_2)$ 之间连边，得到一个有向图。
一开始什么数也没填，所以 $j=0$，最终模 $k$ 要等于 $0$，所以 $j$ 也等于 $0$。所以答案是一条从起点 $(0,0)$ 到终点 $(m,0)$ 的**字典序最大路径**，其中 $m=\left\lceil\dfrac{n}{2}\right\rceil$，因为我们只需填一半的数字，另一半可以镜像得到。
用 DFS 搜索，每次从 $d=9$ 开始倒着枚举，即可得到字典序最大路径。
注意特判 $n$ 为奇数且 $i=m-1$ 的情况，此时填入数字后，回文数模 $k$ 的值变成了

$$
(j + d\cdot 10^i)\bmod k
$$

代码实现时，可以预处理 $10^i \bmod k$。

``` 
func largestPalindrome(n, k int) string {
	pow10 := make([]int, n)
	pow10[0] = 1
	for i := 1; i < n; i++ {
		pow10[i] = pow10[i-1] * 10 % k
	}

	ans := make([]byte, n)
	m := (n + 1) / 2
	vis := make([][]bool, m+1)
	for i := range vis {
		vis[i] = make([]bool, k)
	}
	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i == m {
			return j == 0
		}
		vis[i][j] = true
		for d := 9; d >= 0; d-- { // 贪心：从大到小枚举
			var j2 int
			if n%2 > 0 && i == m-1 { // 正中间
				j2 = (j + d*pow10[i]) % k
			} else {
				j2 = (j + d*(pow10[i]+pow10[n-1-i])) % k
			}
			if !vis[i+1][j2] && dfs(i+1, j2) {
				ans[i] = '0' + byte(d)
				ans[n-1-i] = ans[i]
				return true
			}
		}
		return false
	}
	dfs(0, 0)
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nkD)$，其中 $D=10$。
- 空间复杂度：$\mathcal{O}(nk)$。

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
