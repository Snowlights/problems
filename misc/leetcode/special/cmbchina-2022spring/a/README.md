### 题目

请你设计一个文本编辑程序，需要编辑的稿件 `article` 为仅由大写字母、小写字母与空格组成的字符串（下标从 0 开始），光标所在字符串下标位置记作 `index`，程序运行后，若光标停留位置为空格，不作操作，返回原文本；否则光标所在位置对应的整个单词将被删除，并返回编辑后经过**整理空格**的文本。

**注意：**
- 输入保证字符串不以空格开头和结尾且不包含连续的空格。
- 返回的字符串不以空格开头和结尾且不包含连续的空格。若删除单词后有多余的空格，需要删除。

**示例 1：**
>输入：`article = "Singing dancing in the rain", index = 10`
>
>输出：`"Singing in the rain"`
>
>解释：
>"Singing da**n**cing in the rain" 光标位于 "dancing" 单词的字符 'n'
>删除光标所在的单词 "dancing" 及其前或后的一个空格。

**示例 2：**
>输入：`article = "Hello World", index = 2`
>
>输出：`"World"`
>
>解释：删除 article[2] 所在的单词 "Hello" 及其后方空格。

**示例 3：**
>输入：`article = "Hello World", index = 5`
>
>输出：`"Hello World"`
>
>解释：光标位于空格处，不作处理。

**提示：**
- `1 <= article.length <= 10^5`
- `0 <= index < article.length`
- `article[i]` 仅包含大写字母、小写字母与空格

### 思路

按空格拆分字符串，统计在 $index$ 之前的空格数量，之后做字符串拼接

```go   
func deleteText(article string, id int) string {
	if article[id] == ' ' {
		return article
	}
	parts := strings.Split(article, " ")
	cnt := strings.Count(article[:id], " ")
	ans := append(parts[:cnt], parts[cnt+1:]...)
	return strings.Join(ans, " ")
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{article}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
