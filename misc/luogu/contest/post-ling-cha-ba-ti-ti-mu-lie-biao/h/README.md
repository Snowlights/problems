子序列的元素和的异或和

- 定义 $f[i][j]$ 表示从前 $i$ 个数中选出元素和为 $j$ 的方案数的奇偶性。
- 初始值 $f[0][0]=1$，其余为 $0$。
- 状态转移方程为 $f[i+1][j]=f[i][j]\oplus f[i][j-a[i]]$。
- 答案为 $f[n][j]=1$ 的 $j$ 的异或和。
- 第一个维度可以优化掉。