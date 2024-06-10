#### 题目  

<p>给你一个下标从 <strong>0</strong> 开始、由 <code>n</code> 个整数组成的数组 <code>nums</code> 和一个整数 <code>target</code> 。</p>

<p>你的初始位置在下标 <code>0</code> 。在一步操作中，你可以从下标 <code>i</code> 跳跃到任意满足下述条件的下标 <code>j</code> ：</p>

<ul>
	<li><code>0 &lt;= i &lt; j &lt; n</code></li>
	<li><code>-target &lt;= nums[j] - nums[i] &lt;= target</code></li>
</ul>

<p>返回到达下标 <code>n - 1</code> 处所需的 <strong>最大跳跃次数</strong> 。</p>

<p>如果无法到达下标 <code>n - 1</code> ，返回 <code>-1</code> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [1,3,6,4,1,2], target = 2
<strong>输出：</strong>3
<strong>解释：</strong>要想以最大跳跃次数从下标 0 到下标 n - 1 ，可以按下述跳跃序列执行操作：
- 从下标 0 跳跃到下标 1 。 
- 从下标 1 跳跃到下标 3 。 
- 从下标 3 跳跃到下标 5 。 
可以证明，从 0 到 n - 1 的所有方案中，不存在比 3 步更长的跳跃序列。因此，答案是 3 。 </pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [1,3,6,4,1,2], target = 3
<strong>输出：</strong>5
<strong>解释：</strong>要想以最大跳跃次数从下标 0 到下标 n - 1 ，可以按下述跳跃序列执行操作：
- 从下标 0 跳跃到下标 1 。 
- 从下标 1 跳跃到下标 2 。 
- 从下标 2 跳跃到下标 3 。 
- 从下标 3 跳跃到下标 4 。 
- 从下标 4 跳跃到下标 5 。 
可以证明，从 0 到 n - 1 的所有方案中，不存在比 5 步更长的跳跃序列。因此，答案是 5 。 </pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>nums = [1,3,6,4,1,2], target = 0
<strong>输出：</strong>-1
<strong>解释：</strong>可以证明不存在从 0 到 n - 1 的跳跃序列。因此，答案是 -1 。 
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length == n &lt;= 1000</code></li>
	<li><code>-10<sup>9</sup> &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>0 &lt;= target &lt;= 2 * 10<sup>9</sup></code></li>
</ul>
 
#### 思路  

由于只能向左跳，每次跳跃之后都会把问题规模缩小，那么可以定义 $\textit{dfs}(i)$ 表示从 $i$ 跳到 $0$ 的最大跳跃次数。  
用「枚举选哪个」来思考：  
枚举 $j$，如果 $-\textit{target}\le \textit{nums}[i]-\textit{nums}[j] \le \textit{target}$，那么有  

$$
\textit{dfs}(i) = \textit{dfs}(j) + 1
$$

取所有情况的最大值，即为 $\textit{dfs}(i)$。  
如果没有这样的 $j$，那么 $\textit{dfs}(i) = -\infty$。  
递归边界：$\textit{dfs}(0)=0$。  
递归入口：$\textit{dfs}(n-1)$。也就是答案。如果答案是负数就返回 $-1$。  

```go 

```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(n)$。