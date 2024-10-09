package p3500

// http://acm.hdu.edu.cn/showproblem.php?pid=3555
// 3555
//typedef long long ll;
//
//// p3555
//string v3555;
//ll dp3555[30][3];
//
//// 0表示不存在 1 表示上一位是4， 2表示存在
//ll f3555(int i, int pre, bool limit) {
//    if (i == v3555.length()) {
//        if (pre == 2) {
//            return 1;
//        }
//        return 0;
//    }
//
//    if (!limit) {
//        if (dp3555[i][pre] >= 0) {
//            return dp3555[i][pre];
//        }
//    }
//
//    ll res = 0;
//    int up = 9;
//    if (limit) {
//        up = v3555[i] - '0';
//    }
//    for (int j = 0; j <= up; j++)
//        if (pre == 2 || (pre == 1 && j == 9))
//            res += f3555(i + 1, 2, limit && j == up);
//        else if (j == 4)
//            res += f3555(i + 1, 1, limit && j == up);
//        else
//            res += f3555(i + 1, 0, limit && j == up);
//
//    if (!limit) {
//        dp3555[i][pre] = res;
//    }
//    return res;
//}
//
//void P3555() {
//    int T;
//    cin >> T;
//    while (T--) {
//        cin >> v3555;
//        memset(dp3555, -1, sizeof(dp3555));
//        cout << f3555(0, 0, true) << endl;
//    }
//}
//

// http://acm.hdu.edu.cn/showproblem.php?pid=2089
// long long n, m;
//string s;
//long long dp[30][2];
//
//long long f(int i, int pre, int stat, bool limit) {
//    if (i == s.length()) {
//        return 1;
//    }
//
//    if (!limit && dp[i][stat] >= 0) {
//        return dp[i][stat];
//    }
//
//    long long res = 0;
//    int up = 9;
//    if (limit) {
//        up = s[i] - '0';
//    }
//
//    for (int j = 0; j <= up; j++) {
//        if (pre == 6 && j == 2)
//            continue;
//        if (j == 4)
//            continue;
//        res += f(i + 1, j, j == 6, limit && j == up);
//    }
//
//    if (!limit) {
//        dp[i][stat] = res;
//    }
//
//    return res;
//}
//
//
//int main() {
//
//    while (1) {
//        cin >> n >> m;
//        if (n == 0 && m == 0) {
//            break;
//        }
//        s = to_string(n-1);
//        memset(dp, -1, sizeof(dp));
//        long long a = f(0, 0, 0, true);
//        s = to_string(m);
//        memset(dp, -1, sizeof(dp));
//        long long b = f(0, 0, 0, true);
//        cout << b - a << endl;
//    };
//
//    return 0;
//}
