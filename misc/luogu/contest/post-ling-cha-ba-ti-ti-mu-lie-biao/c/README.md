子数组异或和的元素和

- 计算前缀异或和（长为 $n+1$），统计其中第 $k$ 位的 $1$ 的个数 $c$，那么这一位对答案的贡献是 $c\cdot(n+1-c)\cdot 2^k$。