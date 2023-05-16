---

### 前置知识：异或运算的性质

`a a = 0`。

`a b = c` 两边同时异或 `a` 可以得到 `b = c a`。

### 思路

如果知道 `{original}[0]`，利用 `{derived}[i]` 可以推出其余 `{original}[i]` 的值，即

``
{original}[i+1] = {original}[i] ⊕ {derived}[i]
``

那么有 

``
{original}[n-1] = {original}[0] ⊕ {derived}[0] ⊕ {derived}[1] ⊕ {derived}[n-2]
``

由于 

``
{original}[0] ⊕ {original}[n-1] ={derived}[n-1]
``

联立得

``
{derived}[0] ⊕ {derived}[1] ⊕ {derived}[n-1] = 0
``

所以如果上式成立，`{original}` 必然存在。

```py [sol1-Python3]
class Solution:
    def doesValidArrayExist(self, derived: List[int]) -> bool:
        return reduce(xor, derived) == 0
```

```java [sol1-Java]
class Solution {
    public boolean doesValidArrayExist(int[] derived) {
        int xor = 0;
        for (int x : derived)
            xor ^= x;
        return xor == 0;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    bool doesValidArrayExist(vector<int> &derived) {
        int xor_ = 0;
        for (int x: derived)
            xor_ ^= x;
        return xor_ == 0;
    }
};
```

```go [sol1-Go]
func doesValidArrayExist(derived []int) bool {
	xor := 0
	for _, x := range derived {
		xor ^= x
	}
	return xor == 0
}
```

#### 复杂度分析

- 时间复杂度：`{O}(n)`，其中 `n` 为 `{derived}` 的长度。
- 空间复杂度：`{O}(1)`。仅用到若干额外变量。
