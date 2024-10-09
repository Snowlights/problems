### 题目

一个二进制字符串是由 **0** 和 **1** 组成的字符串。如果一个长度为 `m`（`m` 为奇数）的二进制字符串 `s` 满足 **最后一位数字** 与 `s` 的「排序中间数」相同，那么它就是美丽的。如 `11001` 是美丽的，而 `11100` 不是美丽的。

> - 注：「排序中间数」指将二进制各数位按升序排列后，中间的数字。例如，`11001` 升序排列后为 `[0,0,1,1,1]`，中间的数字 `1` 即为排序中间数。

对于一个长度为 `n` 的二进制字符串 `a` 和一个长度为 `2*n − 1` 的二进制字符串 `b`，如果对于所有 `i`（下标从 `0` 开始） ，都有 `b[2*i] == a[i]`，那么称 `b` 是 `a` 的扩展。例如，$\,\underline{1}\,0\,\underline{0}\,1\,\underline{0}\,1\,\underline{1}\,$ 和 $\,\underline{1}\,1\,\underline{0}\,1\,\underline{0}\,0\,\underline{1}\,$ 都是字符串 `1001` 的扩展。

给定一个长度为 `n` 的二进制字符串 `s`，请返回 `s` 所有前缀的美丽扩展的个数之和。

由于答案可能很大，你只需要求出它模 `998244353` 的结果。

**示例 1：**

> 输入：`s = "11"`
> 输出：`3`
> 解释：前缀 `1`，美丽的扩展有 `1` ，共 1 个。
> 对于前缀 `11`，其扩展有 $\,\underline{1}\,0\,\underline{1}\,$ 和 $\,\underline{1}\,1\,\underline{1}\,$，
> `101` 的最后一位数字为 `1`，将 `101` 各位数升序排序后为 `[0,1,1]`，排序中间数为 `1`，二者相同，为美丽字符串；
> `111` 的最后一位数字为 `1`，将 `111` 各位数升序排序后为 `[1,1,1]`，排序中间数为 `1`，二者相同，为美丽字符串；
> 共有 `3` 种满足条件的扩展，返回 `3`

**示例 2：**

> 输入：`s = "01001"`
> 输出：`17`
> 解释：对于前缀 `0`，美丽的扩展有 $\,\underline{0}\,$ ，共 1 个。
> 对于前缀 `01`，美丽的扩展有 $\,\underline{0}\,1\,\underline{1}\,$ ，共 1 个。
> 对于前缀 `010`，美丽的扩展有 $\,\underline{0}\,0\,\underline{1}\,0\,\underline{0}\,$, $\,\underline{0}\,0\,\underline{1}\,1\,\underline{0}\,$, $\,\underline{0}\,1\,\underline{1}\,0\,\underline{0}\,$ ，共 3 个。
> 对于前缀 `0100`，美丽的扩展有 $\,\underline{0}\,0\,\underline{1}\,0\,\underline{0}\,0\,\underline{0}\,$, $\,\underline{0}\,0\,\underline{1}\,0\,\underline{0}\,1\,\underline{0}\,$, $\,\underline{0}\,0\,\underline{1}\,1\,\underline{0}\,0\,\underline{0}\,$, $\,\underline{0}\,0\,\underline{1}\,1\,\underline{0}\,1\,\underline{0}\,$, $\,\underline{0}\,1\,\underline{1}\,0\,\underline{0}\,0\,\underline{0}\,$, $\,\underline{0}\,1\,\underline{1}\,0\,\underline{0}\,1\,\underline{0}\,$, $\,\underline{0}\,1\,\underline{1}\,1\,\underline{0}\,0\,\underline{0}\,$，共 7 个。
> 对于前缀 `01001`，美丽的扩展有 $\,\underline{0}\,1\,\underline{1}\,1\,\underline{0}\,1\,\underline{0}\,1\,\underline{1}\,$, $\,\underline{0}\,0\,\underline{1}\,1\,\underline{0}\,1\,\underline{0}\,1\,\underline{1}\,$, $\,\underline{0}\,1\,\underline{1}\,0\,\underline{0}\,1\,\underline{0}\,1\,\underline{1}\,$, $\,\underline{0}\,1\,\underline{1}\,1\,\underline{0}\,0\,\underline{0}\,1\,\underline{1}\,$, $\,\underline{0}\,1\,\underline{1}\,1\,\underline{0}\,1\,\underline{0}\,0\,\underline{1}\,$，的个数是 5。
> 所以答案是 17。

**提示：**

- `1 <= s.length <= 2*10^5`
- `s` 仅包含 **0** 和 **1**

### 思路

```go

```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{s}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到部分额外变量。