### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 。</p>

<p>定义 <code>nums</code> 一个子数组的 <strong>不同计数</strong> 值如下：</p>

<ul>
	<li>令 <code>nums[i..j]</code> 表示 <code>nums</code> 中所有下标在 <code>i</code> 到 <code>j</code> 范围内的元素构成的子数组（满足 <code>0 <= i <= j < nums.length</code> ），那么我们称子数组 <code>nums[i..j]</code> 中不同值的数目为 <code>nums[i..j]</code> 的不同计数。</li>
</ul>

<p>请你返回 <code>nums</code> 中所有子数组的 <strong>不同计数</strong> 的 <strong>平方</strong> 和。</p>

<p>由于答案可能会很大，请你将它对 <code>10<sup>9</sup> + 7</code> <strong>取余</strong> 后返回。</p>

<p>子数组指的是一个数组里面一段连续 <strong>非空</strong> 的元素序列。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [1,2,1]
<b>输出：</b>15
<b>解释：</b>六个子数组分别为：
[1]: 1 个互不相同的元素。
[2]: 1 个互不相同的元素。
[1]: 1 个互不相同的元素。
[1,2]: 2 个互不相同的元素。
[2,1]: 2 个互不相同的元素。
[1,2,1]: 2 个互不相同的元素。
所有不同计数的平方和为 1<sup>2</sup> + 1<sup>2</sup> + 1<sup>2</sup> + 2<sup>2</sup> + 2<sup>2</sup> + 2<sup>2</sup> = 15 。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [2,2]
<b>输出：3</b>
<strong>解释：</strong>三个子数组分别为：
[2]: 1 个互不相同的元素。
[2]: 1 个互不相同的元素。
[2,2]: 1 个互不相同的元素。
所有不同计数的平方和为 1<sup>2</sup> + 1<sup>2</sup> + 1<sup>2</sup> = 3 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 100</code></li>
	<li><code>1 <= nums[i] <= 100</code></li>
</ul>

### 思路

推荐先完成本题的简单版本：[2262. 字符串的总引力](https://leetcode.cn/problems/total-appeal-of-a-string/)

为方便描述，下文将 $\textit{nums}$ 简记为 $a$。

## 提示 1

把右端点**相同的子数组，分为同一组。
右端点为 $i$ 的子数组，可以看成是右端点为 $i-1$ 的子数组，在末尾添加上 $a[i]$。
添加后，右端点为 $i-1$ 的这些子数组的「不同计数的平方」之和**增加**了多少？

## 提示 2

假设一个子数组的「不同计数」为 $x$，那么它的「不同计数的平方」为 $x^2$。
如果这个子数组的「不同计数」增加了 $1$，那么它的「不同计数的平方」的**增加量**为

$$
(x+1)^2 - x^2 = 2x+1

$$

## 提示 3

假设 $i=3$，那么右端点为 $i-1$ 的子数组有

- $a[0..2]$，设其「不同计数」为 $x_0$。
- $a[1..2]$，设其「不同计数」为 $x_1$。
- $a[2..2]$，设其「不同计数」为 $x_2$。

其中 $a[k..i]$ 表示从 $a[k]$ 到 $a[i]$ 的子数组。
考虑从子数组 $a[k..i-1]$ 到子数组 $a[k..i]$，分类讨论：

- 如果 $a[i]$ 之前没有遇到过（例如 $a=[1,2,3,4]$），那么这些子数组的「不同计数」都会增加 $1$。根据提示 2，「不同计数的平方」之和的**增加量**为 $(x_0+x_1+x_2)\cdot 2 + 3$。
- 如果 $a[i]$ 之前遇到过，设其上次出现的下标为 $j$，那么：
  - 对于子数组 $a[0..i-1],\ a[1..i-1],\ a[2..i-1],\cdots,a[j..i-1]$，在其末尾添加 $a[i]$ 后，这些子数组的「不同计数」是不会变化的，因为 $a[i]$ 已经在 $a[j]$ 处出现过了。
  - 对于子数组 $a[j+1..i-1],\ a[j+2..i-1],\cdots,a[i-1..i-1]$，由于不包含 $a[i]$，这些子数组的「不同计数」都会增加 $1$，「不同计数的平方」之和的**增加量**计算方式同上。
- 别忘了 $a[i..i]$ 也是一个子数组，把它的「不同计数」加一。

所以，我们需要一个这样的数据结构，用来维护子数组的「不同计数」：

1. 定义 $f[k]$ 表示左端点为 $k$ 的子数组的「不同计数」，如果当前遍历到 $\textit{nums}[i]$，那么 $f[k]$ 就对应着子数组 $a[k..i]$。
2. 区间加一：例如 $a[1..3]$，$a[2..3]$ 和 $a[3..3]$ 的「不同计数」都增加了 $1$，那么就把区间 $[1,3]$ 的「不同计数」加一。
3. 询问区间元素和：为了计算出「不同计数的平方」之和的**增加量**，需要知道从 $f[k+1]$ 到 $f[i-1]$ 的「不同计数」。

这可以用 lazy 线段树实现

## 提示 4

用一个变量 $s$ 维护右端点为 $i$ 的子数组的「不同计数的平方」之和。 遍历 $\textit{nums}$，每次循环按照上述规则更新 $s$：

1. 为了方便调用线段树，假设下标从 $1$ 开始。
2. 设 $a[i]$ 上次出现的下标为 $j$（不存在则为 $0$）。询问 $[j+1,i]$ 的元素和，设为 $s_1$。把 $s_1\cdot 2 + i-j$ 加到 $s$ 中。
3. 把区间 $[j+1,i]$ 都加一。\n4. 把 $s$ 加到答案中。\n5. 更新 $a[i]$ 的上一次出现位置为 $i$。

代码实现时，由于查询的区间和更新的区间是同一个，可以同时完成。

```go

type lazySeg []struct{ sum, todo int }

func (t lazySeg) do(o, l, r, add int) {
	t[o].sum += add * (r - l + 1)
	t[o].todo += add
}

// o=1  [l,r] 1<=l<=r<=n
// 把 [L,R] 加一，同时返回加一之前的区间和
func (t lazySeg) queryAndAdd1(o, l, r, L, R int) (res int) {
	if L <= l && r <= R {
		res = t[o].sum
		t.do(o, l, r, 1)
		return
	}
	m := (l + r) >> 1
	if add := t[o].todo; add != 0 {
		t.do(o<<1, l, m, add)
		t.do(o<<1|1, m+1, r, add)
		t[o].todo = 0
	}
	if L <= m {
		res = t.queryAndAdd1(o<<1, l, m, L, R)
	}
	if m < R {
		res += t.queryAndAdd1(o<<1|1, m+1, r, L, R)
	}
	t[o].sum = t[o<<1].sum + t[o<<1|1].sum
	return
}

func sumCounts(nums []int) (ans int) {
	last := map[int]int{}
	n := len(nums)
	t := make(lazySeg, n*4)
	s := 0
	for i, x := range nums {
		i++
		j := last[x]
		s += t.queryAndAdd1(1, 1, n, j+1, i)*2 + i - j
		ans = (ans + s) % 1_000_000_007
		last[x] = i
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\ log n)$，其中 $n$ 为 $\textit{nums}$ 的长度
- 空间复杂度：$\mathcal{O}(n)$。
