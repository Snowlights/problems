package algorithm

import "math/big"

func _() {
	// 阶乘
	// 偶阶乘
	// https://oeis.org/A000165 Double factorial of even numbers: (2n)!! = 2^n*n!
	// 1, 2, 8, 48, 384, 3840, 46080, 645120, 10321920, 185794560, 3715891200,
	// 81749606400, 1961990553600, 51011754393600, 1428329123020800, 42849873690624000,
	// 1371195958099968000, 46620662575398912000, 1678343852714360832000, 63777066403145711616000
	calcEvenFactorialBig := func(n int) *big.Int {
		return new(big.Int).Lsh(new(big.Int).MulRange(1, int64(n)), uint(n))
	}

	// 奇阶乘
	// https://oeis.org/A001147 Double factorial of odd numbers: (2*n-1)!! = 1*3*5*...*(2*n-1) = A(2*n,n) / 2^n
	// 1, 3, 15, 105, 945, 10395, 135135, 2027025, 34459425, 654729075, 13749310575, 316234143225, 7905853580625, ...
	// Number of ways to choose n disjoint pairs of items from 2*n items
	// Number of perfect matchings in the complete graph K(2n)
	// https://atcoder.jp/contests/abc236/tasks/abc236_d
	// LC1359/双周赛20D 有效的快递序列数目 https://leetcode-cn.com/contest/biweekly-contest-20/problems/count-all-valid-pickup-and-delivery-options/
	// 奇阶乘模 2^64 http://acm.hdu.edu.cn/showproblem.php?pid=6481 https://www.90yang.com/hdu6481-a-math-problem/
	calcOddFactorialBig := func(n int) *big.Int {
		return new(big.Int).Rsh(new(big.Int).MulRange(int64(n+1), int64(2*n)), uint(n))
	}
	_ = []interface{}{calcEvenFactorialBig, calcOddFactorialBig}

	// 组合数

	// 组合数/二项式系数
	// 不取模，仅适用于小范围的 n 和 k
	initComb := func() {
		const mx = 60
		C := [mx + 1][mx + 1]int64{}
		for i := 0; i <= mx; i++ {
			C[i][0], C[i][i] = 1, 1
			for j := 1; j < i; j++ {
				C[i][j] = C[i-1][j-1] + C[i-1][j]
			}
		}
	}

	// O(n) 预处理阶乘及其逆元，O(1) 求组合数
	{
		const mod int64 = 1e9 + 7
		const mx int = 2e6
		// 逆元的定义
		// mod2 * a % mod = 1, mod2 就是a的逆元
		// mod2 = (mod + 1) / a
		// n的阶乘
		F := [mx + 1]int64{1}
		for i := 1; i <= mx; i++ {
			F[i] = F[i-1] * int64(i) % mod
		}
		// 阶乘
		pow := func(x, n int64) int64 {
			//x %= mod
			res := int64(1)
			for ; n > 0; n >>= 1 {
				if n&1 == 1 {
					res = res * x % mod
				}
				x = x * x % mod
			}
			return res
		}
		// n阶乘的逆元
		invF := [...]int64{mx: pow(F[mx], mod-2)}
		for i := mx; i > 0; i-- {
			invF[i-1] = invF[i] * int64(i) % mod
		}
		// A(n, k) (k + 1)... * (n)
		A := func(n, k int) int64 {
			if k < 0 || k > n {
				return 0
			}
			return F[n] * invF[k] % mod
		}
		// C(n, k) 组合：n中取k，有多少种取法
		C := func(n, k int) int64 {
			if k < 0 || k > n {
				return 0
			}
			return F[n] * invF[k] % mod * invF[n-k] % mod
		}
		// P就是排列数.各元素排列的顺序不同被视为是不同的方案.
		// 从n各元素中选k个
		P := func(n, k int) int64 {
			if k < 0 || k > n {
				return 0
			}
			return F[n] * invF[n-k] % mod
		}

		// EXTRA: 卢卡斯定理 https://en.wikipedia.org/wiki/Lucas%27s_theorem
		// https://yangty.blog.luogu.org/lucas-theorem-note
		// C(n,m)%p = C(n%p,m%p) * C(n/p,m/p) % p
		// 注意初始化 F 和 invF 时 mx 取 mod-1
		// 推论：n&m==m 时 C(n,m) 为奇数，否则为偶数
		// https://www.luogu.com.cn/problem/P3807
		// https://www.luogu.com.cn/problem/P7386
		var lucas func(n, k int64) int64
		lucas = func(n, k int64) int64 {
			if k == 0 {
				return 1
			}
			return C(int(n%mod), int(k%mod)) * lucas(n/mod, k/mod) % mod
		}

		// 可重组合 https://en.wikipedia.org/wiki/Combination#Number_of_combinations_with_repetition
		// 方案数 H(n,k)=C(n+k-1,k) https://oeis.org/A059481
		// 相当于把 k 个无区别的球放入 n 个有区别的盒子中，且允许空盒的方案数
		//		隔板法：把 n 个盒子当做 n-1 个隔板，这样相当于总共有 k+n-1个位置，从中选择 k 个位置放球，剩下的位置放隔板。这样就把 k 个球划分成了 n 份，放入对应的盒子中
		// NOTE: 若改成「至多放 k 个球」，则等价于多了一个盒子，用来放「不放入盒子的球」
		// NOTE: mx 要开两倍空间！
		H := func(n, k int) int64 { return C(n+k-1, k) }
		// 也相当于，给出长度和元素范围，求有多少种非降序列
		// 也可以理解成在长度和取值范围-1的格点上走单调路径
		H = func(range_, length int) int64 { return C(range_+length-1, length) }

		// 卡特兰数 Cn = C(2n,n)/(n+1) = C(2n,n)-C(2n,n-1)
		// https://en.wikipedia.org/wiki/Catalan_number
		// https://oeis.org/A000108
		// 从 n=0 开始：1, 1, 2, 5, 14, 42, 132, 429, 1430, 4862, 16796, 58786, 208012, 742900, 2674440, 9694845, 35357670, 129644790
		// 所有在 n×n 格点中不越过对角线的单调路径的个数
		// Number of noncrossing partitions of the n-set (不相交握手问题) LC1259/双周赛13D https://leetcode-cn.com/contest/biweekly-contest-13/problems/handshakes-that-dont-cross/
		// Dyck Path https://mathworld.wolfram.com/DyckPath.html
		// https://www.luogu.com.cn/problem/P1641
		//
		// 将全部偶数提取一个 2，可得 (2n)! = 1*3*5*...*(2n-1) * (2^n) * (n!)
		// 故 C(2*n,n)/(n+1) = (2*n)!/(n!)/(n+1)! = 1*3*5*...*(2n-1)*(2^n)/(n+1)!
		// 又由于 n! 的 2 的因子个数 = n/2 + n/4 + ... + n/2^k <= n-1 当且仅当 n 为 2^k 时取到等号
		// 对比分子分母的 2 的因子个数，可以得出如下结论：
		//     当且仅当 n+1 为 2^k 时，卡特兰数为奇数
		//
		// EXTRA: 高维的情况 https://loj.ac/p/6051
		Catalan := func(n int) int64 { return F[2*n] * invF[n+1] % mod * invF[n] % mod }
		Catalan = func(n int) int64 {
			return new(big.Int).Rem(new(big.Int).Div(new(big.Int).Binomial(int64(2*n), int64(n)), big.NewInt(int64(n+1))), big.NewInt(mod)).Int64()
		}

		// 默慈金数 Motzkin number https://oeis.org/A001006
		// 从 (0,0) 移动到 (n,0) 的网格路径数，每步只能向右移动一格（可以向右上、右下、横向向右），并禁止移动到 y=0 以下的地方
		// M(n) = Sum_{i=0..n/2} C(n,2*i)*Catalan(i)
		// https://en.wikipedia.org/wiki/Motzkin_number
		// 包含生成函数 https://mathworld.wolfram.com/MotzkinNumber.html
		// 生成函数推导 https://zhuanlan.zhihu.com/p/187502941
		// https://blog.csdn.net/acdreamers/article/details/41213667
		// http://acm.hdu.edu.cn/showproblem.php?pid=3723
		Motzkin := func(n int) (res int64) {
			for i := 0; i <= n/2; i++ {
				res = (res + C(n, 2*i)*Catalan(i)) % mod
			}
			return
		}

		// EXTRA: 若仅限定起点为 (0,0)，终点可以是任意 (n,i) https://oeis.org/A005773
		// a(0)=1, a(n) = Sum_{k=0..n-1} M(k)*a(n-k-1)

		// EXTRA: 起点为 (0,i)，终点为 (n,j) https://oeis.org/A081113 Number of paths of length n-1 a king can take from one side of an n X n chessboard to the opposite side
		// a(n) = number of sequences (a_1,a_2,...,a_n) with 1<=a_i<=n for all i and |a_(i+1)-a_(i)|<=1 for 1<=i<=n-1
		// a(n) = Sum_{k=1..n} k*(n-k+1)*M(n-1, k-1) where M() is the Motzkin triangle https://oeis.org/A026300
		// 1, 4, 17, 68, 259, 950, 3387, 11814, 40503, 136946, 457795, 1515926, 4979777, 16246924, 52694573, 170028792, 546148863, 1747255194, 5569898331, 17698806798, 56076828573, 177208108824, 558658899825, 1757365514652

		// 那罗延数 Narayana number (Narayana triangle) https://oeis.org/A001263
		// 从 (0,0) 移动到 (2n,0) 且恰好有 k 个山峰的网格路径数，每步只能向右上或右下移动一格（不能横向），并禁止移动到 x 轴以下的地方
		// N(n,k) = C(n,k)*C(n,k-1)/n
		// https://en.wikipedia.org/wiki/Narayana_number

		// Fuss-Catalan 数、(m-1)-Dyck 路与 m 叉树 https://www.luogu.com.cn/blog/your-alpha1022/solution-p2767

		// 某些组合题可能用到
		pow2 := [mx + 1]int64{1}
		for i := 1; i <= mx; i++ {
			pow2[i] = pow2[i-1] << 1 % mod
		}

		// 组合公式 数学
		// 1*n + 2*(n-1) + 3*(n-2) + ... + (n-1)*2 + n*1 = C(n+2, 3) = n*(n+1)*(n+2) / 6
		// C(n, 0) + C(n, 1) + C(n, 2) + ... + C(n, n) = 2^n
		// C(r, r) + C(r+1,r) + C(r+2,r) + ... + C(n, r) = C(n+1, r+1)

		// 多个数最大公约数
		gcd := func(a, b int) int {
			for a%b != 0 {
				a, b = b, a%b
			}
			return b
		}
		// 多个数最小公倍数
		lcm := func(a, b int) int {
			return a / gcd(a, b) * b
		}

		// 预处理: [2,mx] 范围内的质数
		// 埃筛 埃氏筛 埃拉托斯特尼筛法 Sieve of Eratosthenes
		// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
		// https://oeis.org/A055399 Number of stages of sieve of Eratosthenes needed to identify n as prime or composite
		// https://oeis.org/A230773 Minimum number of steps in an alternate definition of the Sieve of Eratosthenes needed to identify n as prime or composite
		// 质数个数 π(n) https://oeis.org/A000720
		//         π(10^n) https://oeis.org/A006880
		//         4, 25, 168, 1229, 9592, 78498, 664579, 5761455, 50847534, /* 1e9 */
		//         455052511, 4118054813, 37607912018, 346065536839, 3204941750802, 29844570422669, 279238341033925, 2623557157654233, 24739954287740860, 234057667276344607,
		// 思想应用 https://codeforces.com/contest/1646/problem/E
		sieve := func() {
			const mx int = 1e6
			primes := []int{}
			pid := [mx + 1]int{-1, -1}
			for i := 2; i <= mx; i++ {
				if pid[i] == 0 {
					primes = append(primes, i)
					pid[i] = len(primes)
					for j := i * i; j <= mx; j += i {
						pid[j] = -1
					}
				}
			}

			// EXTRA: pi(n), the number of primes <= n https://oeis.org/A000720
			pi := [mx + 1]int{}
			for i := 2; i <= mx; i++ {
				pi[i] = pi[i-1]
				if pid[i] > 0 {
					pi[i]++
				}
			}
		}

		// 分解质因数
		// LPF(n): least prime dividing n (when n > 1); a(1) = 1 https://oeis.org/A020639
		// 有时候数据范围比较大，用 primeFactorsAll 预处理会 MLE，这时候就要用 LPF 了（同样是预处理但是内存占用低）
		// 先预处理出 LPF，然后对要处理的数 v 不断地除 LPF(v) 直到等于 1
		// 		LPF 前缀和 https://oeis.org/A046669 https://oeis.org/A088821 前缀积 https://oeis.org/A072486
		//      - a(n) ~ n^2/(2 log n)
		//		n+LPF(n) https://oeis.org/A061228 the smallest number greater than n which is not coprime to n
		// 		n-LPF(n) https://oeis.org/A046666
		// 			迭代至 0 的次数 https://oeis.org/A175126 相关题目 https://codeforces.com/contest/1076/problem/B
		//		n*LPF(n) https://oeis.org/A285109
		// 		n/LPF(n) https://oeis.org/A032742 即 n 的最大因子 = Max{gcd(n,j); j=n+1..2n-1}
		//
		//		只考虑奇质数 https://oeis.org/A078701
		//
		// GPF(n): greatest prime dividing n, for n >= 2; a(1)=1 https://oeis.org/A006530
		//		GPF(p-1) https://oeis.org/A023503
		//		GPF(p+1) https://oeis.org/A023509
		// 		GPF 前缀和 https://oeis.org/A046670 前缀积 https://oeis.org/A104350
		//		n+GPF(n) https://oeis.org/A070229 the next m>n such that GPF(n)|m
		// 		n-GPF(n) https://oeis.org/A076563
		// 			迭代至 0 的次数 https://oeis.org/A309892
		// 		n*GPF(n) https://oeis.org/A253560
		// 		n/GPF(n) https://oeis.org/A052126
		//      a(1)=1, a(n+1)=a(n)+GPF(a(n)) https://oeis.org/A076271
		//
		// 		n/LPF(n)*GPF(n) https://oeis.org/A130064
		// 		n/GPF(n)*LPF(n) https://oeis.org/A130065
		//
		// https://codeforces.com/problemset/problem/385/C
		// https://codeforces.com/gym/103107/problem/F (另一种做法是欧拉筛）
		lpfAll := func() {
			const mx int = 1e6
			lpf := [mx + 1]int{1: 1}
			for i := 2; i <= mx; i++ {
				if lpf[i] == 0 {
					for j := i; j <= mx; j += i {
						if lpf[j] == 0 { // 去掉这个判断就变成求 GPF，也可以用来（从大到小地）分解质因数
							lpf[j] = i
						}
					}
				}
			}
		}
		{
			// 也可以用欧拉筛求，实际测试下来耗时和上面差不多
			lpf := [mx + 1]int{1: 1}
			primes := []int{} // 可以提前确定空间
			for i := 2; i <= mx; i++ {
				if lpf[i] == 0 {
					lpf[i] = i
					primes = append(primes, i)
				}
				for _, p := range primes {
					if p*i > mx {
						break
					}
					lpf[p*i] = p
					if i%p == 0 {
						break
					}
				}
			}
		}

		_ = []interface{}{A, C, P, H, Catalan, Motzkin, initComb, lucas, lcm, sieve, lpfAll}
	}
}
