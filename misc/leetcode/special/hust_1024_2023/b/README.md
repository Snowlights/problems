### 题目

给你一个非负整数数组 `chopsticks` 表示筷子长度。如果存在一个数 `x` ，使得 `chopsticks` 中恰好有 `x` 根筷子的长度 **大于或者等于** `x` ，那么就称 `chopsticks` 是一组 **特殊筷子** ，而 `x` 是该筷子组的 **特征值** 。

> - 注： `x` **不必** 是 `chopsticks` 的中的元素。

如果数组 `chopsticks` 是一组 **特殊筷子** ，请返回它的特征值 `x` 。否则，返回 `-1` 。可以证明的是，如果 `chopsticks` 是一组特殊筷子，那么其特征值 `x` 是 唯一的 。

**示例 1：**

> 输入：`chopsticks = [3,5]`
> 输出：`2`
> 解释：有 `2` 根筷子的长度`（3 和 5）`大于或等于 `2` 。

**示例 2：**

> 输入：`chopsticks = [0,0]`
> 输出：`-1`
> 解释：没有满足题目要求的特殊筷子组，故而也不存在特征值 `x` 。
> 如果 `x = 0`，应该有 `0` 根筷子长度 >= x，但实际有 `2` 个。
> 如果 `x = 1`，应该有 `1` 根筷子长度 >= x，但实际有 `0` 个。
> 如果 `x = 2`，应该有 `2` 根筷子长度 >= x，但实际有 `0` 个。
> x 不能取更大的值，因为 `chopsticks` 中只有两根筷子。

**示例 3：**

> 输入：`chopsticks = [0,4,3,0,4]`
> 输出：`3`
> 解释：有 `3` 根筷子长度大于或等于 `3` 。

**示例 4：**

> 输入：`chopsticks = [3,6,7,7,0]`
> 输出：`-1`

**提示：**

- `1 <= chopsticks.length <= 100`
- `0 <= chopsticks[i] <= 1000`

### 思路

结果之多为 $n$, $n$ 为数组长度
排序后二分查找 $>=i$ 的数, 如果 $>=$ 的数量恰好等于 $i$, $i$ 即为结果, 否则不存在结果, 返回 -1.

```go  
func specialChopsticks(chopsticks []int) int {
	sort.Ints(chopsticks)
	n := len(chopsticks)
	for i := 0; i <= n; i++ {
		idx := sort.SearchInts(chopsticks, i)
		if n-idx == i {
			return i
		}
	}
	return -1
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{chopsticks}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
