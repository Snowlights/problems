### 题目

老师在黑板上写出一个字符串 `s` 表示方程，该方程仅包含变量 `x` 、其对应系数和 `'+' ， '-'` 操作。
假设方程中 `x` 为 `answer` ，将 `x` 以字符串 `"x=#answer"` 的形式返回。
题目保证，如果方程中只有一个解，则 `answer` 的值是一个整数。
如果方程没有解或存在的解不为整数，请返回 `"No solution"` 。
如果方程有无限解，则返回 `"Infinite solutions"` 。

**示例 1：**

> 输入: `s = "x+5-3+x=6+x-2"`
> 输出: `"x=2"`

**示例 2:**

> 输入: `s = "x=x"`
> 输出: `"Infinite solutions"`

**示例 3:**

> 输入: `s = "2x=x"`
> 输出: `"x=0"`

**提示:**

- `3 <= s.length <= 1000`
- `s` 仅包含 `'x','+','-','='` 且只有一个 `'='`
- 方程由绝对值在 `[0, 100]`  范围内且无任何前导零的整数和变量 `'x'` 组成。

### 思路

以 $=$ 分割字符串两端，对两端分别进行模拟,计算 $x$ 和 $sum$ 得值，之后做判断。

```go  
func mathProblem(cal string) string {

	handle := func(s string) (int, int) {
		i, x, sum, flag := 0, 0, 0, 1

		for i < len(s) {
			switch s[i] {
			case 'x':
				x += flag
				i++
			case '+':
				flag = 1
				i++
			case '-':
				flag = -1
				i++
			default:
				v := 0
				for i < len(s) && s[i] >= '0' && s[i] <= '9' {
					v = v*10 + int(s[i]-'0')
					i++
				}
				if i < len(s) && s[i] == 'x' {
					x += flag * v
					i++
				} else {
					sum += flag * v
				}
			}
		}
		return x, sum
	}
	parts := strings.Split(cal, "=")
	ax, av := handle(parts[0])
	bx, bv := handle(parts[1])

	if ax == bx {
		if av-bv != 0 {
			return "No solution"
		}
		return "Infinite solutions"
	}
	if (bv-av)%(ax-bx) != 0 {
		return "No solution"
	}
	return fmt.Sprintf("x=%d", (bv-av)/(ax-bx))
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{s}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到部分额外变量。