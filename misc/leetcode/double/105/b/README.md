### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的字符串 <code>s</code> 和一个单词字典 <code>dictionary</code> 。你需要将 <code>s</code> 分割成若干个 <strong>互不重叠</strong> 的子字符串，每个子字符串都在 <code>dictionary</code> 中出现过。<code>s</code> 中可能会有一些 <strong>额外的字符</strong> 不在任何子字符串中。</p>

<p>请你采取最优策略分割 <code>s</code> ，使剩下的字符 <strong>最少</strong> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>s = &#34;leetscode&#34;, dictionary = [&#34;leet&#34;,&#34;code&#34;,&#34;leetcode&#34;]
<b>输出：</b>1
<b>解释：</b>将 s 分成两个子字符串：下标从 0 到 3 的 &#34;leet&#34; 和下标从 5 到 8 的 &#34;code&#34; 。只有 1 个字符没有使用（下标为 4），所以我们返回 1 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>s = &#34;sayhelloworld&#34;, dictionary = [&#34;hello&#34;,&#34;world&#34;]
<b>输出：</b>3
<b>解释：</b>将 s 分成两个子字符串：下标从 3 到 7 的 &#34;hello&#34; 和下标从 8 到 12 的 &#34;world&#34; 。下标为 0 ，1 和 2 的字符没有使用，所以我们返回 3 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 50</code></li>
	<li><code>1 &lt;= dictionary.length &lt;= 50</code></li>
	<li><code>1 &lt;= dictionary[i].length &lt;= 50</code></li>
	<li><code>dictionary[i]</code> 和 <code>s</code> 只包含小写英文字母。</li>
	<li><code>dictionary</code> 中的单词互不相同。</li>
</ul>
 
### 思路  

设 $n$ 为 $s$ 的长度。我们可以：

- 直接跳过 $s$ 的最后一个字符，那么问题变成 $s$ 的前 $n-1$ 个字符的子问题。
- 考虑「枚举选哪个」，如果从 $s[j]$ 开始的后缀在 $\textit{dictionary}$ 中，那么问题变成 $s$ 的前 $j-1$ 个字符的子问题。

根据上面的讨论，定义 $\textit{dfs}(i)$ 表示 $s$ 的前 $i$ 个字符的子问题。

- 跳过 $s$ 的最后一个字符，有 $\textit{dfs}(i)=\textit{dfs}(i-1)+1$。
- 考虑「枚举选哪个」，如果从 $s[j]$ 到 $s[i]$ 的子串在 $\textit{dictionary}$ 中，有

$$
\textit{dfs}(i)=\min_{j=0}^{i}\textit{dfs}(j-1)
$$

这两种情况取最小值。

递归边界：$\textit{dfs}(-1)=0$。

答案：$\textit{dfs}(n-1)$。

```go 

```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(L + n^3)$，其中 $n$ 为 $s$ 的长度，$L$ 为 $\textit{dictionary}$ 所有字符串的长度之和。预处理哈希表需要 $\mathcal{O}(L)$ 的时间。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n^2)$，因此时间复杂度为 $\mathcal{O}(n^3)$。所以总的时间复杂度为 $\mathcal{O}(L + n^3)$。
- 空间复杂度：$\mathcal{O}(n+L)$。

## 翻译为递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

做法：

- $\textit{dfs}$ 改成 $f$ 数组；
- 递归改成循环（每个参数都对应一层循环）；
- 递归边界改成 $f$ 数组的初始值。

具体来说，$f[i]$ 的含义和 $\textit{dfs}(i)$ 的含义是一样的，都表示 $s$ 的前 $i$ 个字符的子问题。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 的一致：

- 跳过 $s$ 的最后一个字符，有 $f[i]=f[i-1]+1$。
- 考虑「枚举选哪个」，如果从 $s[j]$ 到 $s[i]$ 的子串在 $\textit{dictionary}$ 中，有

$$
f[i]=\min_{j=0}^{i}f[j-1]
$$

这两种情况取最小值。

但当 $i=0$ 或 $j=0$ 时，等号右边会出现负数下标。或者说，**这种定义方式没有状态能表示递归边界**，即出界的情况。

解决办法：在 $f$ 数组左边添加一个状态来表示 $i=-1$，把原来的 $f[i]$ 改成 $f[i+1]$，$f[j-1]$ 改成 $f[j]$。

相应的递推式为 $f[i+1]=f[i]+1$ 以及 $f[i+1]=\min\limits_{j=0}^{i}f[j]$。

初始值 $f[i]=0$。（翻译自 $\textit{dfs}(-1)=0$。）

答案为 $f[n]$。（翻译自 $\textit{dfs}(n-1)$。）

```go 
func minExtraChar(s string, dictionary []string) int {
	has := map[string]bool{}
	for _, s := range dictionary {
		has[s] = true
	}
	n := len(s)
	f := make([]int, n+1)
	for i := 0; i < n; i++ {
		f[i+1] = f[i] + 1 // 不选
		for j := 0; j <= i; j++ { // 枚举选哪个
			if has[s[j:i+1]] {
				f[i+1] = min(f[i+1], f[j])
			}
		}
	}
	return f[n]
}

func min(a, b int) int { if b < a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(L + n^3)$，其中 $n$ 为 $s$ 的长度，$L$ 为 $\textit{dictionary}$ 所有字符串的长度之和。预处理哈希表需要 $\mathcal{O}(L)$ 的时间。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n^2)$，因此时间复杂度为 $\mathcal{O}(n^3)$。所以总的时间复杂度为 $\mathcal{O}(L + n^3)$。
- 空间复杂度：$\mathcal{O}(n+L)$。