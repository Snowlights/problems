### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> ，你可以在一些下标之间遍历。对于两个下标 <code>i</code> 和 <code>j</code>（<code>i != j</code>），当且仅当 <code>gcd(nums[i], nums[j]) &gt; 1</code> 时，我们可以在两个下标之间通行，其中 <code>gcd</code> 是两个数的 <strong>最大公约数</strong> 。</p>

<p>你需要判断 <code>nums</code> 数组中 <strong>任意 </strong>两个满足 <code>i &lt; j</code> 的下标 <code>i</code> 和 <code>j</code> ，是否存在若干次通行可以从 <code>i</code> 遍历到 <code>j</code> 。</p>

<p>如果任意满足条件的下标对都可以遍历，那么返回 <code>true</code> ，否则返回 <code>false</code> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [2,3,6]
<b>输出：</b>true
<b>解释：</b>这个例子中，总共有 3 个下标对：(0, 1) ，(0, 2) 和 (1, 2) 。
从下标 0 到下标 1 ，我们可以遍历 0 -&gt; 2 -&gt; 1 ，我们可以从下标 0 到 2 是因为 gcd(nums[0], nums[2]) = gcd(2, 6) = 2 &gt; 1 ，从下标 2 到 1 是因为 gcd(nums[2], nums[1]) = gcd(6, 3) = 3 &gt; 1 。
从下标 0 到下标 2 ，我们可以直接遍历，因为 gcd(nums[0], nums[2]) = gcd(2, 6) = 2 &gt; 1 。同理，我们也可以从下标 1 到 2 因为 gcd(nums[1], nums[2]) = gcd(3, 6) = 3 &gt; 1 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [3,9,5]
<b>输出：</b>false
<b>解释：</b>我们没法从下标 0 到 2 ，所以返回 false 。
</pre>

<p><strong>示例 3：</strong></p>

<pre><b>输入：</b>nums = [4,3,12,8]
<b>输出：</b>true
<b>解释：</b>总共有 6 个下标对：(0, 1) ，(0, 2) ，(0, 3) ，(1, 2) ，(1, 3) 和 (2, 3) 。所有下标对之间都存在可行的遍历，所以返回 true 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
</ul>
 
### 思路  

对数组 $a$ 进行质因数分解，之后使用并查集将各质数连接起来，最后判断是否连通

```go 

```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n+mx)$， $n$ 为数组长度, $mx$ 为固定值
- 空间复杂度：$\mathcal{O}(n+mx)$。